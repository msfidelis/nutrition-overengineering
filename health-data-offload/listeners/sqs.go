package listeners

import (
	"app/data"
	"app/models"
	"app/pkg/logger"
	"app/processor"
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func StartSQSConsumer() {
	logInternal := logger.Instance()
	queueURL := os.Getenv("SQS_QUEUE_URL")
	if queueURL == "" {
		logInternal.Warn().Msg("SQS_QUEUE_URL not set")
		return
	}

	sess := session.Must(session.NewSession())
	svc := sqs.New(sess)
	for {
		output, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: aws.Int64(1),
			WaitTimeSeconds:     aws.Int64(10),
		})
		if err != nil {
			logInternal.Error().Str("error", err.Error()).Msg("Failed to consume queue")
			time.Sleep(10 * time.Second)
			continue
		}

		for _, message := range output.Messages {
			var msg models.Message

			err = json.Unmarshal([]byte(*message.Body), &msg)
			if err != nil {
				logInternal.Error().Str("error", err.Error()).Msg("Failed parse JSON message")
				continue
			}

			logInternal.Info().Str("id", msg.Id).Msg("Received message")
			logInternal.Info().Str("id", msg.Id).Msg("Processing message")

			report := processor.Execute(msg)
			data.Save(report)

			logInternal.Info().Str("id", msg.Id).Msg("Deleting message from queue")

			_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: message.ReceiptHandle,
			})
			if err != nil {
				logInternal.Error().Str("error", err.Error()).Msg("Failed to delete message from queue")
			}
		}
	}
}
