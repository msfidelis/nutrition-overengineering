package message

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/msfidelis/health-api/pkg/logger"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func SendSQSMessage(ctx context.Context, message string) {
	log := logger.Instance()
	span := trace.SpanFromContext(ctx)
	span.SetName("SQS Send Message")
	defer span.End()

	queueURL := os.Getenv("SQS_QUEUE_URL")
	if queueURL == "" {
		log.Warn().Msg("SQS_QUEUE_URL not set")
		return
	}

	span.SetAttributes(
		attribute.String("sqs.url", queueURL),
	)

	sess := session.Must(session.NewSession())
	svc := sqs.New(sess)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(queueURL),
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
}
