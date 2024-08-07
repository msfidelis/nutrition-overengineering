package main

import (
	"database-store-service/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	log := logger.Instance()
	log.Info().Msg("Starting Database Store Service")

	wg := sync.WaitGroup{}
	wg.Add(10)

	// Healthcheck Probe :8080
	go startHTTPHealthCheckServer()

	// Nats Client
	nc, err := nats.Connect(os.Getenv("NATS_URI"))
	defer nc.Close()

	if err != nil {
		log.Error().
			Str("Error", err.Error()).
			Str("NATS_URI", os.Getenv(("NATS_URI"))).
			Msg("Error to connect to Nats")
	}
	log.Info().
		Str("NATS_URI", os.Getenv("NATS_URI")).
		Msg("Connected to NATS")

	// Create JetStream Context
	js, err := nc.JetStream()

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "orders",
		Subjects: []string{"ORDERS.*"},
	})

	if err != nil {
		log.Error().
			Str("Error", err.Error()).
			Str("NATS_URI", os.Getenv(("NATS_URI"))).
			Msg("Failed to add Stream")
	}

	if err != nil {
		log.Error().
			Str("Error", err.Error()).
			Str("NATS_URI", os.Getenv(("NATS_URI"))).
			Msg("Failed to add a durable consumer")
	}

	js.QueueSubscribe("ORDERS.*", "store", func(m *nats.Msg) {
		log.Info().
			Msg(fmt.Sprintf("Received a person: %+v\n", string(m.Data)))
		m.Ack()
	}, nats.ManualAck())

	// Wait for a message to come in
	wg.Wait()

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
