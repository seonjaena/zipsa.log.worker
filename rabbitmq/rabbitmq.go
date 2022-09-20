package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"strings"
	"time"
	zp "zipsa.log.worker/properties"
	"zipsa.log.worker/zlog"
)

const (
	DeadLetterExchangeKey string = "x-dead-letter-exchange"
	DeadLetterRoutingKey  string = "x-dead-letter-routing-key"
	MessageTtlKey         string = "x-message-ttl"
	DirectOption          string = "direct"
)

var log = zlog.Instance()
var _conn *amqp.Connection
var _chan *amqp.Channel

func init() {
	GetConn()
	_ = GetChan()
	DeclareExchange()
	DeclareQueue()
	BindQueue()
}

func GetConn() {
	if _conn == nil || _conn.IsClosed() {
		tryConn()
	}
}

func GetChan() *amqp.Channel {
	if _chan == nil {
		createChan()
	}
	return _chan
}

func DeclareQueue() {
	err := _chan.Qos(zp.GetRabbitmqPrefetchCnt(), 0, false)
	if err != nil {
		log.Errorf("set prefetch count failed")
	}
	logArg := amqp.Table{
		DeadLetterExchangeKey: zp.GetRabbitmqWaitLogExchange(),
		DeadLetterRoutingKey:  zp.GetRabbitmqWaitLogQueue(),
	}
	waitLogArg := amqp.Table{
		DeadLetterExchangeKey: zp.GetRabbitmqLogExchange(),
		DeadLetterRoutingKey:  zp.GetRabbitmqLogQueue(),
		MessageTtlKey:         zp.GetRabbitmqDeadLogTTL(),
	}

	_, err = _chan.QueueDeclare(
		zp.GetRabbitmqLogQueue(),
		true,
		false,
		false,
		false,
		logArg,
	)
	if err != nil {
		log.Errorf("declare queue failed: %s", zp.GetRabbitmqLogQueue())
	}
	_, err = _chan.QueueDeclare(
		zp.GetRabbitmqWaitLogQueue(),
		true,
		false,
		false,
		false,
		waitLogArg,
	)
	if err != nil {
		log.Errorf("declare queue failed: %s", zp.GetRabbitmqDeadLogQueue())
	}
	_, err = _chan.QueueDeclare(
		zp.GetRabbitmqDeadLogQueue(),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Errorf("declare queue failed: %s", zp.GetRabbitmqDeadLogQueue())
	}
}

func DeclareExchange() {
	err := _chan.ExchangeDeclare(
		zp.GetRabbitmqLogExchange(),
		DirectOption,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Errorf("declare exchange failed: %s", zp.GetRabbitmqLogExchange())
	}
	err = _chan.ExchangeDeclare(
		zp.GetRabbitmqWaitLogExchange(),
		DirectOption,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Errorf("declare exchange failed: %s", zp.GetRabbitmqDeadLogExchange())
	}
	err = _chan.ExchangeDeclare(
		zp.GetRabbitmqDeadLogExchange(),
		DirectOption,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Errorf("declare exchange failed: %s", zp.GetRabbitmqDeadLogExchange())
	}
}

func BindQueue() {
	err := _chan.QueueBind(
		zp.GetRabbitmqLogQueue(),
		zp.GetRabbitmqLogQueue(),
		zp.GetRabbitmqLogExchange(),
		false,
		nil,
	)
	if err != nil {
		log.Errorf("bind queue failed: %s, %s", zp.GetRabbitmqLogQueue(), zp.GetRabbitmqLogExchange())
	}
	err = _chan.QueueBind(
		zp.GetRabbitmqWaitLogQueue(),
		zp.GetRabbitmqWaitLogQueue(),
		zp.GetRabbitmqWaitLogExchange(),
		false,
		nil,
	)
	if err != nil {
		log.Errorf("bind queue failed: %s, %s", zp.GetRabbitmqLogQueue(), zp.GetRabbitmqLogExchange())
	}
	err = _chan.QueueBind(
		zp.GetRabbitmqDeadLogQueue(),
		zp.GetRabbitmqDeadLogQueue(),
		zp.GetRabbitmqDeadLogExchange(),
		false,
		nil,
	)
	if err != nil {
		log.Errorf("bind queue failed: %s, %s", zp.GetRabbitmqDeadLogQueue(), zp.GetRabbitmqDeadLogExchange())
	}
}

func createChan() {
	for {
		var err error
		log.Infof("rabbitmq channel creating...")
		_chan, err = _conn.Channel()
		if err != nil {
			log.Errorf("create rabbitmq channel failed: %s", err.Error())
		} else {
			go func() {
				log.Errorf("rabbitmq channel closed: %s", <-_chan.NotifyClose(make(chan *amqp.Error)))
				createChan()
			}()

			log.Infof("rabbitmq channel created")

			break
		}
	}
}

func tryConn() {
	for {
		var err error
		log.Infof("rabbitmq connecting...")
		_conn, err = amqp.Dial(makeConnectionURL())
		if err != nil {
			log.Errorf("connect rabbitmq failed: %s", err.Error())
		} else {
			go func() {
				log.Errorf("rabbitmq connection closed: %s", <-_conn.NotifyClose(make(chan *amqp.Error)))
				tryConn()
			}()

			log.Infof("rabbitmq connected")

			break
		}
	}
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

func RetryMsg(d *amqp.Delivery, rejectedErr error) {
	rejectedCnt := getRejectedCnt(d)
	log.Infof("Rejected Count = %d", rejectedCnt)
	if rejectedCnt == 0 || rejectedCnt+1 <= zp.GetRabbitmqRetryCnt() {
		for {
			err := d.Reject(false)
			if err == nil {
				break
			} else {
				time.Sleep(3 * time.Second)
			}
		}
		return
	}
	deadMsg := fmt.Sprintf("%s[^]%s", string(d.Body), rejectedErr)
	for {
		err := _chan.Publish(
			zp.GetRabbitmqDeadLogExchange(),
			zp.GetRabbitmqDeadLogQueue(),
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(deadMsg),
			},
		)
		if err == nil {
			break
		} else {
			time.Sleep(3 * time.Second)
		}
	}
	for {
		err := d.Ack(false)
		if err == nil {
			break
		} else {
			time.Sleep(3 * time.Second)
		}
	}
}

func getRejectedCnt(d *amqp.Delivery) int64 {
	a := d.Headers["x-death"]
	if a == nil {
		return 0
	}

	for _, v := range a.([]interface{}) {
		val := v.(amqp.Table)
		reason, _ := val["reason"].(string)
		if strings.Compare(reason, "rejected") == 0 {
			val2, _ := val["count"].(int64)
			return val2
		}
	}
	return 0
}

func GetChannel() *amqp.Channel {
	return _chan
}
