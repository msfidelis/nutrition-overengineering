package calculator

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/msfidelis/health-api/pkg/logger"
	"github.com/msfidelis/health-api/pkg/services/bmr"
	"github.com/msfidelis/health-api/pkg/services/imc"
	"github.com/msfidelis/health-api/pkg/services/recommendations"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
)

type Request struct {
	Age               int     `json:"age" binding:"required"`
	Weight            float64 `json:"weight" binding:"required"`
	Height            float64 `json:"height" binding:"required"`
	Gender            string  `json:"gender" binding:"required,oneof=M F"`
	ActivityIntensity string  `json:"activity_intensity" binding:"required,oneof=sedentary lightly_active moderately_active very_active extra_active"`
}

type Response struct {
	Status int `json:"status" binding:"required"`
	Imc    struct {
		Result float64 `json:"result"`
		Class  string  `json:"class"`
	} `json:"imc"`
	Basal struct {
		BMR struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"bmr"`
		Necessity struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"necessity"`
	} `json:"basal"`
	HealthInfo struct {
		Age               int     `json:"age"`
		Weight            float64 `json:"weight"`
		Height            float64 `json:"height"`
		Gender            string  `json:"gender"`
		ActivityIntensity string  `json:"activity_intensity"`
	} `json:"health_info"`
	Recomendations struct {
		Protein struct {
			Value int64  `json:"value"`
			Unit  string `json:"unit"`
		} `json:"protein"`
		Water struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"water"`
		Calories struct {
			Maintain struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"maintain_weight"`
			Loss struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"loss_weight"`
			Gain struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"gain_weight"`
		} `json:"calories"`
	} `json:"recomendations"`
}

// Calculator godoc
// @Summary Return calculation for health macros
// @Tags HealthCalculator
// @Produce json
// @Success 200 {object} Response
// @Param message body Request true "Health Information"
// @Router /calculator [post]
func Post(c *gin.Context) {
	var response Response
	var request Request

	// Endpoint Span
	span := trace.SpanFromContext(c.Request.Context())
	span.SetName("Nutrition Calc Service")
	tr := otel.Tracer("health-api")

	log := logger.Instance()

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	span.SetAttributes(
		attribute.String("request.Gender", request.Gender),
		attribute.Float64("request.Weight", request.Weight),
		attribute.Float64("request.Height", request.Height),
		attribute.String("request.ActivityIntensity", request.ActivityIntensity),
		attribute.Int("request.Age", request.Age),
	)

	// BMR
	ctxBMR, spanBMR := tr.Start(c.Request.Context(), "BMR Service Call")

	bmrEndpoint := os.Getenv("BMR_SERVICE_ENDPOINT")

	log.Info().
		Str("Service", "bmr").
		Str("BMR_SERVICE_ENDPOINT", bmrEndpoint).
		Msg("Creating remote connection with gRPC Endpoint for BMR Service")

	spanBMR.SetAttributes(
		attribute.String("Service", "bmr"),
		attribute.String("BMR_SERVICE_ENDPOINT", bmrEndpoint),
	)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(
		bmrEndpoint,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		log.Error().
			Str("Service", "bmr").
			Str("Error", err.Error()).
			Msg("Failed to create gRPC Connection with BMR Service")

		spanBMR.SetAttributes(
			attribute.String("Service", "bmr"),
			attribute.String("Error", err.Error()),
			attribute.String("Message", "Failed to create gRPC Connection with BMR Service"),
		)
	}
	defer conn.Close()

	bmrClient := bmr.NewBMRServiceClient(conn)

	spanBMR.SetAttributes(
		attribute.String("grpc.request.Gender", request.Gender),
		attribute.Float64("grpc.request.Weight", request.Weight),
		attribute.Float64("grpc.request.Height", request.Height),
		attribute.String("grpc.request.ActivityIntensity", request.ActivityIntensity),
	)

	resBMR, err := bmrClient.SayHello(ctxBMR, &bmr.Message{
		Gender:   request.Gender,
		Weight:   request.Weight,
		Height:   request.Height,
		Activity: request.ActivityIntensity,
	})

	spanBMR.SetAttributes(
		attribute.Float64("grpc.response.BMR", resBMR.Bmr),
		attribute.Float64("grpc.response.Necessity", resBMR.Necessity),
		attribute.String("grpc.response.Unity", "kcal"),
	)

	defer spanBMR.End()

	// IMC
	ctxIMC, spanIMC := tr.Start(c.Request.Context(), "IMC Service Call")
	imcEndpoint := os.Getenv("IMC_SERVICE_ENDPOINT")

	log.Info().
		Str("Service", "imc").
		Str("IMC_SERVICE_ENDPOINT", imcEndpoint).
		Msg("Creating remote connection with gRPC Endpoint for IMC Service")

	spanIMC.SetAttributes(
		attribute.String("Service", "imc"),
		attribute.String("IMC_SERVICE_ENDPOINT", imcEndpoint),
	)

	var connIMC *grpc.ClientConn
	connIMC, errIMC := grpc.Dial(
		imcEndpoint,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if errIMC != nil {
		log.Error().
			Str("Service", "imc").
			Str("Error", err.Error()).
			Msg("Failed to create gRPC Connection with IMC Service")

		spanIMC.SetAttributes(
			attribute.String("Service", "imc"),
			attribute.String("Error", err.Error()),
			attribute.String("Message", "Failed to create gRPC Connection with IMC Service"),
		)
	}
	defer connIMC.Close()

	imcClient := imc.NewIMCServiceClient(connIMC)

	spanIMC.SetAttributes(
		attribute.Float64("grpc.request.Weight", request.Weight),
		attribute.Float64("grpc.request.Height", request.Height),
	)

	resIMC, err := imcClient.SayHello(ctxIMC, &imc.Message{
		Weight: request.Weight,
		Height: request.Height,
	})

	spanIMC.SetAttributes(
		attribute.Float64("grpc.response.Imc", resIMC.Imc),
		attribute.String("grpc.response.Class", resIMC.Class),
	)

	defer spanIMC.End()

	// Recommendations
	ctxRecommendations, spanRecommendations := tr.Start(c.Request.Context(), "Recommendations Service Call")
	recommendationsEndpoint := os.Getenv("RECOMMENDATIONS_SERVICE_ENDPOINT")

	log.Info().
		Str("Service", "recommendations").
		Str("RECOMMENDATIONS_SERVICE_ENDPOINT", recommendationsEndpoint).
		Msg("Creating remote connection with gRPC Endpoint for Recommendation Service")

	spanRecommendations.SetAttributes(
		attribute.String("Service", "recommendations"),
		attribute.String("RECOMMENDATIONS_SERVICE_ENDPOINT", recommendationsEndpoint),
	)

	var connRecommendations *grpc.ClientConn
	connRecommendations, errRecommendations := grpc.Dial(
		recommendationsEndpoint,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if errRecommendations != nil {
		log.Error().
			Str("Service", "recommendations").
			Str("Error", err.Error()).
			Msg("Failed to create gRPC Connection with Recommendations Service")

		spanRecommendations.SetAttributes(
			attribute.String("Service", "recommendations"),
			attribute.String("Error", err.Error()),
			attribute.String("Message", "Failed to create gRPC Connection with Recommendations Service"),
		)
	}
	defer connRecommendations.Close()

	recommendationsClient := recommendations.NewRecomendationsServiceClient(connRecommendations)

	spanRecommendations.SetAttributes(
		attribute.String("grpc.request.Gender", request.Gender),
		attribute.Float64("grpc.request.Weight", request.Weight),
		attribute.Float64("grpc.request.Calories", resBMR.Necessity),
	)

	resRecommendations, err := recommendationsClient.SayHello(ctxRecommendations, &recommendations.Message{
		Weight:   request.Weight,
		Height:   request.Height,
		Calories: resBMR.Necessity,
	})

	spanRecommendations.SetAttributes(
		attribute.Int64("grpc.response.Protein", resRecommendations.ProteinsValue),
		attribute.Float64("grpc.response.Water", resRecommendations.WaterValue),
		attribute.Float64("grpc.response.Calories.Maintain", resRecommendations.CaloriesToMaintein),
		attribute.Float64("grpc.response.Calories.Gain", resRecommendations.CaloriesToGain),
		attribute.Float64("grpc.response.Calories.Loss", resRecommendations.CaloriesToLoss),
	)

	defer spanRecommendations.End()

	_, spanResponse := tr.Start(c.Request.Context(), "HTTP Response")

	// BMR Response
	response.Basal.BMR.Value = resBMR.Bmr
	response.Basal.BMR.Unit = "kcal"
	response.Basal.Necessity.Value = resBMR.Necessity
	response.Basal.Necessity.Unit = "kcal"

	// IMC Response
	response.Imc.Result = resIMC.Imc
	response.Imc.Class = resIMC.Class

	response.HealthInfo.Age = request.Age
	response.HealthInfo.Gender = request.Gender
	response.HealthInfo.Weight = request.Weight
	response.HealthInfo.Height = request.Height
	response.HealthInfo.ActivityIntensity = request.ActivityIntensity

	// Recommendations Response
	response.Recomendations.Protein.Value = resRecommendations.ProteinsValue
	response.Recomendations.Protein.Unit = resRecommendations.ProteinsUnit
	response.Recomendations.Water.Value = resRecommendations.WaterValue
	response.Recomendations.Water.Unit = resRecommendations.WaterUnit
	response.Recomendations.Calories.Maintain.Value = resRecommendations.CaloriesToMaintein
	response.Recomendations.Calories.Maintain.Unit = response.Basal.Necessity.Unit
	response.Recomendations.Calories.Gain.Value = resRecommendations.CaloriesToGain
	response.Recomendations.Calories.Gain.Unit = response.Basal.Necessity.Unit
	response.Recomendations.Calories.Loss.Value = resRecommendations.CaloriesToLoss
	response.Recomendations.Calories.Loss.Unit = response.Basal.Necessity.Unit
	response.Status = http.StatusOK

	spanResponse.SetAttributes(
		attribute.Float64("http.response.Basal.BMR.Value", response.Basal.BMR.Value),
		attribute.String("http.response.Basal.BMR.Unit", response.Basal.BMR.Unit),
		attribute.Float64("http.response.Basal.Necessity.Value", response.Basal.Necessity.Value),
		attribute.String("http.response.Basal.Necessity.Unit", response.Basal.Necessity.Unit),

		attribute.Float64("http.response.Imc.Result", response.Imc.Result),
		attribute.String("http.response.Imc.Class", response.Imc.Class),

		attribute.Int("http.response.HealthInfo.Age", request.Age),
		attribute.String("http.response.HealthInfo.Gender", request.Gender),
		attribute.Float64("http.response.HealthInfo.Weight", request.Weight),
		attribute.Float64("http.response.HealthInfo.Height", request.Height),
		attribute.String("http.response.HealthInfo.ActivityIntensity", request.ActivityIntensity),

		attribute.Int64("http.response.Recomendations.Protein.Value", resRecommendations.ProteinsValue),
		attribute.String("http.response.Recomendations.Protein.Unit", resRecommendations.ProteinsUnit),
		attribute.Float64("http.response.Water.Value", resRecommendations.WaterValue),
		attribute.String("http.response.Water.Unit", resRecommendations.WaterUnit),
		attribute.Float64("http.response.Recomendations.Calories.Maintain.Value", resRecommendations.CaloriesToMaintein),
		attribute.String("http.response.Recomendations.Calories.Maintain.Unit", response.Basal.Necessity.Unit),
		attribute.Float64("http.response.Recomendations.Calories.Gain.Value", resRecommendations.CaloriesToGain),
		attribute.String("http.response.Recomendations.Calories.Gain.Unit", response.Basal.Necessity.Unit),
		attribute.Float64("http.response.Recomendations.Calories.Loss.Value", resRecommendations.CaloriesToLoss),
		attribute.String("http.response.Recomendations.Calories.Loss.Unit", response.Basal.Necessity.Unit),
	)

	c.JSON(http.StatusOK, response)
	defer spanResponse.End()
}
