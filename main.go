package main

import (
	"zipsa.log.worker/rabbitmq"
	"zipsa.log.worker/redis"
	"zipsa.log.worker/zlog"
)

var log = zlog.Instance()

func worker() {
	forever := make(chan bool)
	go redis.LogBuffer.FlushData()
	go rabbitmq.ConsumeLog()
	<-forever
}

func main() {
	worker()
}
