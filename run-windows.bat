@ECHO off

SET GOOS=windows
SET GOARCH=amd64

SET rabbitmq_dead_log_exchange=dead-log-exchange
SET rabbitmq_dead_log_queue=dead-log-queue
SET rabbitmq_dead_log_routingkey=dead-log-queue
SET rabbitmq_host=b-eb4fa4e5-2217-4f02-aa63-4582e063c667.mq.ap-northeast-2.amazonaws.com
SET rabbitmq_log_exchange=log-exchange
SET rabbitmq_log_queue=log-queue
SET rabbitmq_log_routingkey=log-queue
SET rabbitmq_password=na6080su12!@
SET rabbitmq_port=5671
SET rabbitmq_use_ssl=true
SET rabbitmq_username=sky114z
SET rabbitmq_virtualhost=/
SET redis_host=144.24.83.179
SET redis_password=na6080su12!@
SET redis_port=	6379
SET log_out=stdout
SET log_level=debug

go run .\main.go