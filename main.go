package main

import (
	mq "zipsa.log.worker/rabbitmq"
	"zipsa.log.worker/redis"
	"zipsa.log.worker/zlog"
)

var log = zlog.Instance()

func worker() {
	forever := make(chan bool)
	go redis.LogBuffer.FlushData()
	go mq.ConsumeLog()
	<-forever
}

func main() {
	log.Infof("Start!!!")
	worker()
	log.Infof("End!!!")
}
