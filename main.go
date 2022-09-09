package main

import (
	"context"
	b64 "encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"zipsa.log.worker/zlog"
)

var log = zlog.Instance()

func HandleRequest(ctx context.Context, event events.RabbitMQEvent) {
	messages := event.MessagesByQueue
	for _, message := range messages {
		for _, data := range message {
			decodeString, _ := b64.StdEncoding.DecodeString(data.Data)
			log.Printf("data = %s", decodeString)
		}
	}
}

func main() {
	lambda.Start(HandleRequest)
}
