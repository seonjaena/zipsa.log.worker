package main

import (
	zp "zipsa.log.worker/properties"
	"zipsa.log.worker/rabbitmq"
	"zipsa.log.worker/redis"
	"zipsa.log.worker/zlog"
)

var log = zlog.Instance()

func worker() {
	forever := make(chan bool)
	go redis.LogBuffer.FlushData()
	go consumeLog()
	<-forever
}

func main() {
	log.Infof("Start!!!")
	worker()
	log.Infof("End!!!")
}

func consumeLog() {
	channel := rabbitmq.GetChannel()
	msg, err := channel.Consume(
		zp.GetRabbitmqLogQueue(),
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Errorf("consume message failed")
	}
	for d := range msg {
		redis.LogBuffer.Append(string(d.Body), &d)
	}
}
