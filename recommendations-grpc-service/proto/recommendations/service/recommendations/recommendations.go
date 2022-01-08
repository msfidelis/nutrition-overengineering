package recommendations

import (
	// calculator "recommendations-grpc-service/pkg/recommendations"
	"os"

	"recommendations-grpc-service/pkg/logger"
	"recommendations-grpc-service/pkg/services/proteins"
	"recommendations-grpc-service/pkg/services/water"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Response, error) {
	log := logger.Instance()

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

	var connWater *grpc.ClientConn
	connWater, errWater := grpc.Dial(waterEndpoint, grpc.WithInsecure())
	if errWater != nil {
		log.Error().
			Str("Service", "imc").
			Str("Error", errWater.Error()).
			Msg("Failed to create gRPC Connection with Water Consume Service")
	}
	defer connWater.Close()

	waterClient := water.NewWaterServiceClient(connWater)
	resWater, err := waterClient.SayHello(context.Background(), &water.Message{
		Weight: in.Weight,
		Height: in.Height,
	})

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

	proteinsEndpoint := os.Getenv("PROTEINS_SERVICE_ENDPOINT")

	log.Info().
		Str("Service", "proteins").
		Str("PROTEINS_SERVICE_ENDPOINT", proteinsEndpoint).
		Msg("Creating remote connection with gRPC Endpoint for Proteins Consume Service")

	var connProteins *grpc.ClientConn
	connProteins, errProteins := grpc.Dial(proteinsEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Error().
			Str("Service", "proteins").
			Str("Error", errProteins.Error()).
			Msg("Failed to create gRPC Connection with Water Consume Service")
	}
	defer connProteins.Close()

	proteinsClient := proteins.NewProteinsServiceClient(connProteins)
	resProteins, err := proteinsClient.SayHello(context.Background(), &proteins.Message{
		Weight: in.Weight,
	})

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

	caloriesEndpoint := os.Getenv("CALORIES_SERVICE_ENDPOINT")

	log.Info().
		Str("Service", "calories").
		Str("CALORIES_SERVICE_ENDPOINT", caloriesEndpoint).
		Msg("Creating remote connection with gRPC Endpoint for Calories Consume Service")

	var connCalories *grpc.ClientConn
	connCalories, errCalories := grpc.Dial(caloriesEndpoint, grpc.WithInsecure())
	if errCalories != nil {
		log.Error().
			Str("Service", "calories").
			Str("Error", errCalories.Error()).
			Msg("Failed to create gRPC Connection with Water Consume Service")
	}
	defer connCalories.Close()

	return &Response{
		WaterValue:    resWater.Value,
		WaterUnit:     resWater.Unit,
		ProteinsValue: resProteins.Value,
		ProteinsUnit:  resProteins.Unit,
	}, nil
}
