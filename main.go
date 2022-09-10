package main

import (
	"time"
	"zipsa.log.worker/rabbitmq"
	"zipsa.log.worker/zlog"
)

var log = zlog.Instance()

func work() {
	forever := make(chan bool)
	_, err := rabbitmq.GetConn()
	if err != nil {
		log.Errorf("Error Occurred.")
		log.Errorf("error: %s", err.Error())
	}
	go func() {
		for {
			log.Infof("Worker-1!!!")
			time.Sleep(time.Second * 3)
		}
	}()
	<-forever
}

func main() {
	log.Infof("Start!!!")
	work()
	log.Infof("End!!!")
}
