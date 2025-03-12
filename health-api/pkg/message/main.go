package message

import "os"

func SendMessage(message string) {
	// send message
	switch os.Getenv("MESSAGE_TYPE") {
	case "sqs":
		SendSQSMessage(message)
		return
	case "kafka":
		return
	default:
		return
	}
}
