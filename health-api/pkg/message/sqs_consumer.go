package message

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/msfidelis/health-api/pkg/logger"
)

func StartSQSConsumer() {
	log := logger.Instance()
	queueURL := os.Getenv("SQS_QUEUE_URL")
	if queueURL == "" {
		log.Warn().Msg("SQS_QUEUE_URL not set")
		return
	}

	sess := session.Must(session.NewSession())
	svc := sqs.New(sess)

	for {
		output, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: aws.Int64(10),
			WaitTimeSeconds:     aws.Int64(10),
		})
		if err != nil {
			log.Error().Msg(err.Error())
			time.Sleep(10 * time.Second)
			continue
		}

		for _, message := range output.Messages {
			log.Info().Msgf("Received message: %s", *message.Body)
			// Process the message here

			_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: message.ReceiptHandle,
			})
			if err != nil {
				log.Error().Msg(err.Error())
			}
		}
	}
}
