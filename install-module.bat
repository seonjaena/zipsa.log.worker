@ECHO off

go mod init zipsa.log.worker

go get -u github.com/streadway/amqp
go get -u github.com/sirupsen/logrus
go get -u github.com/google/martian
go get -u github.com/magiconair/properties