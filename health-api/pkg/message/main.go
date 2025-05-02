package message

import (
	"context"
	"os"
)

func SendMessage(ctx context.Context, message string) {

	// send message
	switch os.Getenv("MESSAGE_TYPE") {
	case "sqs":
		SendSQSMessage(ctx, message)
		return
	case "kafka":
		return
	default:
		return
	}
}
