package main

import (
	"bmr-grpc-service/pkg/logger"
	"bmr-grpc-service/pkg/tracer"
	"bmr-grpc-service/proto/bmr/service/bmr"
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {

	ctx := context.Background()
	cleanup := tracer.InitTracer(ctx)
	defer cleanup()

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

	// Healthcheck Probe :8080
	go startHTTPHealthCheckServer()

	logInternal.Info().
		Msg("Listener for bmr-grpc-service is created")

	s := bmr.Server{}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     15 * time.Second,
			MaxConnectionAge:      30 * time.Second,
			MaxConnectionAgeGrace: 5 * time.Second,
			Time:                  5 * time.Second,
			Timeout:               1 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.MaxRecvMsgSize(4<<20), // 4MB
		grpc.MaxSendMsgSize(4<<20), // 4MB
		grpc.InitialWindowSize(1<<20),
		grpc.InitialConnWindowSize(1<<20),
	)

	bmr.RegisterBMRServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		logInternal.Error().
			Str("Error", err.Error()).
			Msg("Failed to server gRPC server")
	}

	logInternal.Info().
		Msg("Server bmr-grpc-service is enabled")

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
