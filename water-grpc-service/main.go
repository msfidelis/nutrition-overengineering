package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"water-grpc-service/pkg/logger"
	"water-grpc-service/pkg/tracer"
	"water-grpc-service/proto/water/service/water"

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

	// Healthcheck Probe :8080
	go startHTTPHealthCheckServer()

	logInternal.Info().
		Msg("Listener for water-grpc-service is created")

	s := water.Server{}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	water.RegisterWaterServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		logInternal.Error().
			Str("Error", err.Error()).
			Msg("Failed to server gRPC server")
	}

	logInternal.Info().
		Msg("Server water-grpc-service is enabled")

}

func startHTTPHealthCheckServer() {
	logInternal := logger.Instance()
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	logInternal.Info().Msg("Starting HTTP healthcheck server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logInternal.Fatal().Err(err).Msg("Failed to start HTTP healthcheck server")
	}
}
