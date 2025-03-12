package listeners

import (
	"app/pkg/logger"
	"os"
	"strings"
)

func StartListeners() {
	logInternal := logger.Instance()
	switch strings.ToLower(os.Getenv("MESSAGE_TYPE")) {
	case "sqs":
		StartSQSConsumer()
	// case "KAFKA":
	// 	StartKafkaConsumer()
	default:
		logInternal.Warn().Msg("No listener defined")
	}

}
