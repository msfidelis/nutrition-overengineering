package main

import (
	"context"
	"net"
	"os"
	"recommendations-grpc-service/pkg/logger"
	"recommendations-grpc-service/pkg/tracer"
	"recommendations-grpc-service/proto/recommendations/service/recommendations"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

func main() {

	logInternal := logger.Instance()

	ctx := context.Background()
	cleanup := tracer.InitTracer(ctx)
	defer cleanup()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	lis, err := net.Listen("tcp", ":30000")
	if err != nil {
		log.Error().
			Str("Error", err.Error()).
			Msg("Failed to listen")
	}

	logInternal.Info().
		Msg("Listener for recommendations-grpc-service is created")

	s := recommendations.Server{}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	recommendations.RegisterRecomendationsServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		logInternal.Error().
			Str("Error", err.Error()).
			Msg("Failed to server gRPC server")
	}

	logInternal.Info().
		Msg("Server recommendations-grpc-service is enabled")

}
