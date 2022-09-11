package main

import (
	"zipsa.log.worker/rabbitmq"
	_ "zipsa.log.worker/redis"
	"zipsa.log.worker/zlog"
)

var log = zlog.Instance()

func worker() {
	forever := make(chan bool)
	go rabbitmq.ConsumeLog()
	<-forever
}

func main() {
	log.Infof("Start!!!")
	worker()
	log.Infof("End!!!")
}
