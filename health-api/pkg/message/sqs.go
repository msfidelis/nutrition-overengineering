package message

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/msfidelis/health-api/pkg/logger"
)

func SendSQSMessage(message string) {
	log := logger.Instance()
	queueURL := os.Getenv("SQS_QUEUE_URL")
	if queueURL == "" {
		log.Warn().Msg("SQS_QUEUE_URL not set")
		return
	}

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
