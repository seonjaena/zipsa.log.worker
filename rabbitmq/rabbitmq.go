package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	zp "zipsa.log.worker/properties"
)

func GetConn() (*amqp.Connection, error) {
	conn, err := amqp.Dial(makeConnectionURL())
	return conn, err
}

func makeConnectionURL() string {
	var mqProtocol string
	if zp.GetRabbitmqUseSsl() {
		mqProtocol = "amqps"
	} else {
		mqProtocol = "amqp"
	}
	return fmt.Sprintf("%s://%s:%s@%s/%s", mqProtocol, zp.GetRabbitmqUsername(), zp.GetRabbitmqPassword(), zp.GetRabbitmqHost(), zp.GetRabbitmqVirtualhost())
}
