package data

import (
	"app/models"
	"app/pkg/logger"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func SaveDynamoDB(report models.Report) {
	logInternal := logger.Instance()
	logInternal.Info().Str("Id", report.Id).Msg("Saving to DynamoDB")
	logInternal.Info().Str("Id", report.Id).Msg("Successfully saved item to DynamoDB")

	// Save report to DynamoDB
	tableName := os.Getenv("DYNAMODB_TABLE")
	if tableName == "" {
		logInternal.Warn().Msg("DYNAMODB_TABLE not set")
		return
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	av, err := dynamodbattribute.MarshalMap(report)
	// av, err := dynamodbattribute.MarshalMap(report)
	if err != nil {
		logInternal.Error().Str("Id", report.Id).Str("error", err.Error()).Msg("Failed to marshal report")
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		logInternal.Error().Str("Id", report.Id).Str("error", err.Error()).Msg("Failed to save item to DynamoDB")
	}

}
