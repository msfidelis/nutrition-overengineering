package recommendations

import (
	// calculator "recommendations-grpc-service/pkg/recommendations"
	"os"

	"recommendations-grpc-service/pkg/logger"
	"recommendations-grpc-service/pkg/services/calories"
	"recommendations-grpc-service/pkg/services/proteins"
	"recommendations-grpc-service/pkg/services/water"

	"go.opentelemetry.io/otel"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Response, error) {
	log := logger.Instance()
	tr := otel.Tracer("recommendations-grpc-service")

	// Water
	log.Info().
		Float64("Weight", in.Weight).
		Float64("Height", in.Height).
		Msg("Calculating water")

	waterEndpoint := os.Getenv("WATER_SERVICE_ENDPOINT")

	log.Info().
		Str("Service", "water").
		Str("WATER_SERVICE_ENDPOINT", waterEndpoint).
		Msg("Creating remote connection with gRPC Endpoint for Water Consume Service")

	resWater, err := water.Call(ctx, in.Weight, in.Height, tr)

	if err != nil {
		log.Error().
			Str("Service", "water").
			Str("Error", err.Error()).
			Msg("Failed consume water service")
	}

	log.Info().
		Float64("Weight", in.Weight).
		Float64("Height", in.Height).
		Float64("Water", resWater.Value).
		Msg("Water consume Calculated")

	// Proteins

	log.Info().
		Float64("Weight", in.Weight).
		Float64("Height", in.Height).
		Msg("Calculating proteins necessity")

	resProteins, err := proteins.Call(ctx, in.Weight, tr)

	if err != nil {
		log.Error().
			Str("Service", "water").
			Str("Error", err.Error()).
			Msg("Failed consume water service")
	}

	// Calories
	log.Info().
		Float64("Weight", in.Weight).
		Float64("Height", in.Height).
		Msg("Calculating calories necessity")

	resCalories, err := calories.Call(ctx, in.Calories, tr)

	if err != nil {
		log.Error().
			Str("Service", "errCalories").
			Str("Error", err.Error()).
			Msg("Failed consume errCalories service")
	}

	log.Info().
		Str("Service", "calories").
		Float64("Necessity", in.Calories).
		Float64("Maintain", resCalories.Maintain).
		Float64("Loss", resCalories.Loss).
		Float64("Gain", resCalories.Gain).
		Msg("Calories necessity calculated")

	return &Response{
		WaterValue:         resWater.Value,
		WaterUnit:          resWater.Unit,
		ProteinsValue:      resProteins.Value,
		ProteinsUnit:       resProteins.Unit,
		CaloriesToLoss:     resCalories.Loss,
		CaloriesToGain:     resCalories.Gain,
		CaloriesToMaintein: resCalories.Maintain,
	}, nil
}
