package calculator

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/msfidelis/health-api/pkg/logger"
	"github.com/msfidelis/health-api/pkg/services/bmr"
	"github.com/msfidelis/health-api/pkg/services/imc"
	"github.com/msfidelis/health-api/pkg/services/recommendations"
	"google.golang.org/grpc"
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

	log := logger.Instance()

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// BMR
	bmrEndpoint := os.Getenv("BMR_SERVICE_ENDPOINT")

	log.Info().
		Str("Service", "bmr").
		Str("BMR_SERVICE_ENDPOINT", bmrEndpoint).
		Msg("Creating remote connection with gRPC Endpoint for BMR Service")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(bmrEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Error().
			Str("Service", "bmr").
			Str("Error", err.Error()).
			Msg("Failed to create gRPC Connection with BMR Service")
	}
	defer conn.Close()

	bmrClient := bmr.NewBMRServiceClient(conn)
	resBMR, err := bmrClient.SayHello(context.Background(), &bmr.Message{
		Gender:   request.Gender,
		Weight:   request.Weight,
		Height:   request.Height,
		Activity: request.ActivityIntensity,
	})

	// IMC
	imcEndpoint := os.Getenv("IMC_SERVICE_ENDPOINT")

	log.Info().
		Str("Service", "imc").
		Str("IMC_SERVICE_ENDPOINT", imcEndpoint).
		Msg("Creating remote connection with gRPC Endpoint for IMC Service")

	var connIMC *grpc.ClientConn
	connIMC, errIMC := grpc.Dial(imcEndpoint, grpc.WithInsecure())
	if errIMC != nil {
		log.Error().
			Str("Service", "imc").
			Str("Error", err.Error()).
			Msg("Failed to create gRPC Connection with IMC Service")
	}
	defer connIMC.Close()

	imcClient := imc.NewIMCServiceClient(connIMC)
	resIMC, err := imcClient.SayHello(context.Background(), &imc.Message{
		Weight: request.Weight,
		Height: request.Height,
	})

	// Recommendations
	recommendationsEndpoint := os.Getenv("RECOMMENDATIONS_SERVICE_ENDPOINT")

	log.Info().
		Str("Service", "recommendations").
		Str("RECOMMENDATIONS_SERVICE_ENDPOINT", recommendationsEndpoint).
		Msg("Creating remote connection with gRPC Endpoint for Recommendation Service")

	var connRecommendations *grpc.ClientConn
	connRecommendations, errRecommendations := grpc.Dial(recommendationsEndpoint, grpc.WithInsecure())
	if errRecommendations != nil {
		log.Error().
			Str("Service", "recommendations").
			Str("Error", err.Error()).
			Msg("Failed to create gRPC Connection with Recommendations Service")
	}
	defer connRecommendations.Close()

	recommendationsClient := recommendations.NewRecomendationsServiceClient(connRecommendations)
	resRecommendations, err := recommendationsClient.SayHello(context.Background(), &recommendations.Message{
		Weight:   request.Weight,
		Height:   request.Height,
		Calories: resBMR.Necessity,
	})

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

	c.JSON(http.StatusOK, response)
}
