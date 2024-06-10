package main

import (
	"calories-grpc-service/pkg/logger"
	"calories-grpc-service/pkg/tracer"
	"calories-grpc-service/proto/calories/service/calories"
	"context"
	"net"
	"os"

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
		Msg("Listener for calories-grpc-service is created")

	s := calories.Server{}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	calories.RegisterCaloriesServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		logInternal.Error().
			Str("Error", err.Error()).
			Msg("Failed to server gRPC server")
	}

	logInternal.Info().
		Msg("Server calories-grpc-service is enabled")

}
