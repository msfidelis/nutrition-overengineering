package main

import (
	"database-store-service/pkg/logger"
	"database-store-service/pkg/envs"
	"database-store-service/pkg/migration"
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	log := logger.Instance()
	log.Info().Msg("Database Store Service Starting")

	log.Info()
		.Msg("Starting Database Migrations")

	migration.Migrate()

	wg := sync.WaitGroup{}
	wg.Add(10)

	// Nats Client
	natsUri := envs.Getenv("NATS_URI", "nats://0.0.0.0:4222")

	log.Info(). 
		Str("Nats URI", natsUri).
		Msg("Trying to connect to NATS")	

	nc, err := nats.Connect(natsUri)
	defer nc.Close()

	if err != nil {
		log.Error().
			Str("Error", err.Error()).
			Msg("Error to connect to Nats")
	}

	// Create JetStream Context
	js, err := nc.JetStream()

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "reports",
		Subjects: []string{"REPORTS.*"},
	})

	if err != nil {
		log.Error().
			Str("Error", err.Error()).
			Msg("Failed to add Stream")
	}

	if err != nil {
		log.Error().
			Str("Error", err.Error()).
			Msg("Failed to add a durable consumer")
	}

	js.QueueSubscribe("REPORTS.*", "store", func(m *nats.Msg) {
		log.Info().
			Msg(fmt.Sprintf("Received a health report: %+v\n", string(m.Data)))
		m.Ack()
	}, nats.ManualAck())

	// Wait for a message to come in
	wg.Wait()

}
