package data

import (
	"app/models"
	"os"
)

func Save(r models.Report) {
	switch os.Getenv("DATABASE_TYPE") {
	case "dynamodb":
		SaveDynamoDB(r)
		return
	case "s3":
		return
	case "postgres":
		return
	default:
		return
	}
}
