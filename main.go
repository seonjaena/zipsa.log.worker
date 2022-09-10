package main

import (
	"zipsa.log.worker/rabbitmq"
	"zipsa.log.worker/zlog"
)

var log = zlog.Instance()

func worker() {
	forever := make(chan bool)
	rabbitmq.ConsumeLog()
	<-forever
}

func main() {
	log.Infof("Start!!!")
	worker()
	log.Infof("End!!!")
}
