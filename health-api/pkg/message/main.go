package message

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/trace"
)

func SendMessage(ctx context.Context, message string) {

	span := trace.SpanFromContext(ctx)
	span.SetName("Define Message Type")
	defer span.End()

	// send message
	switch os.Getenv("MESSAGE_TYPE") {
	case "sqs":
		SendSQSMessage(trace.ContextWithSpan(ctx, span), message)
		return
	case "kafka":
		return
	default:
		return
	}
}
