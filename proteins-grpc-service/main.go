package main

import (
	"context"
	"net"
	"os"
	"proteins-grpc-service/pkg/logger"
	"proteins-grpc-service/pkg/tracer"
	"proteins-grpc-service/proto/proteins/service/proteins"

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
		Msg("Listener for proteins-grpc-service is created")

	s := proteins.Server{}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	proteins.RegisterProteinsServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		logInternal.Error().
			Str("Error", err.Error()).
			Msg("Failed to server gRPC server")
	}

	logInternal.Info().
		Msg("Server proteins-grpc-service is enabled")

}
