package main

import (
	"app/listeners"
	"app/pkg/logger"
	"app/pkg/tracer"
	"context"
	"net/http"
	"os"
	"strconv"
	"sync"

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

	numWorkersStr := os.Getenv("WORKERS")
	if numWorkersStr == "" {
		numWorkersStr = "3"
	}

	numWorkers, err := strconv.ParseInt(numWorkersStr, 10, 32)
	if err != nil {
		logInternal.Fatal().Err(err).Str("WORKERS", os.Getenv("WORKERS")).Msg("Failed to paser WORKERS variable")
	}

	var wg sync.WaitGroup

	for i := 0; i < int(numWorkers); i++ {
		wg.Add(1)
		logInternal.Info().Int("Worker", i).Msg("Starting worker")
		go listeners.StartListeners()
	}

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
