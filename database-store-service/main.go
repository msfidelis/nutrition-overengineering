package main

import (
	"database-store-service/pkg/logger"
	"fmt"
	"os"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	log := logger.Instance()
	log.Info().Msg("Hello World")

	wg := sync.WaitGroup{}
	wg.Add(10)

	// Nats Client
	nc, err := nats.Connect(os.Getenv("NATS_URI"))
	defer nc.Close()

	if err != nil {
		log.Error().
			Str("Error", err.Error()).
			Msg("Error to connect to Nats")
	}

	// Create JetStream Context
	js, err := nc.JetStream()

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "orders",
		Subjects: []string{"ORDERS.*"},
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

	js.QueueSubscribe("ORDERS.*", "store", func(m *nats.Msg) {
		log.Info().
			Msg(fmt.Sprintf("Received a person: %+v\n", string(m.Data)))
		m.Ack()
	}, nats.ManualAck())

	// Wait for a message to come in
	wg.Wait()

}
