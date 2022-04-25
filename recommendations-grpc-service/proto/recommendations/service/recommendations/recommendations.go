package recommendations

import (
	// calculator "recommendations-grpc-service/pkg/recommendations"
	"os"

	"recommendations-grpc-service/pkg/logger"
	"recommendations-grpc-service/pkg/services/calories"
	"recommendations-grpc-service/pkg/services/proteins"
	"recommendations-grpc-service/pkg/services/water"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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
	connWater, errWater := grpc.Dial(
		waterEndpoint,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if errWater != nil {
		log.Error().
			Str("Service", "water").
			Str("Error", errWater.Error()).
			Msg("Failed to create gRPC Connection with Water Consume Service")
	}
	defer connWater.Close()

	waterClient := water.NewWaterServiceClient(connWater)
	resWater, err := waterClient.SayHello(ctx, &water.Message{
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
	connProteins, errProteins := grpc.Dial(
		proteinsEndpoint,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		log.Error().
			Str("Service", "proteins").
			Str("Error", errProteins.Error()).
			Msg("Failed to create gRPC Connection with Water Consume Service")
	}
	defer connProteins.Close()

	proteinsClient := proteins.NewProteinsServiceClient(connProteins)
	resProteins, err := proteinsClient.SayHello(ctx, &proteins.Message{
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
	connCalories, errCalories := grpc.Dial(
		caloriesEndpoint,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if errCalories != nil {
		log.Error().
			Str("Service", "calories").
			Str("Error", errCalories.Error()).
			Msg("Failed to create gRPC Connection with Water Consume Service")
	}
	defer connCalories.Close()

	caloriesClient := calories.NewCaloriesServiceClient(connCalories)
	resCalories, err := caloriesClient.SayHello(ctx, &calories.Message{
		Necessity: in.Calories,
	})

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
