package main

import (
	"calories-grpc-service/pkg/logger"
	"calories-grpc-service/proto/calories/service/calories"
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

func main() {

	logInternal := logger.Instance()

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

	grpcServer := grpc.NewServer()

	calories.RegisterCaloriesServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		logInternal.Error().
			Str("Error", err.Error()).
			Msg("Failed to server gRPC server")
	}

	logInternal.Info().
		Msg("Server calories-grpc-service is enabled")

}
