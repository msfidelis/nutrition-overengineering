package main

import (
	"app/listeners"
	"app/pkg/logger"
	"app/pkg/tracer"
	"context"
	"net/http"

	"github.com/rs/zerolog"
)

func main() {
	logInternal := logger.Instance()

	ctx := context.Background()
	cleanup := tracer.InitTracer(ctx)
	defer cleanup()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	go startHTTPHealthCheckServer()

	logInternal.Info().Msg("Starting Consumer")

	listeners.StartListeners()
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
