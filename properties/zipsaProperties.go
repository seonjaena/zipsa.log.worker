package properties

import (
	"fmt"
	"github.com/magiconair/properties"
	"log"
)

var Profile string

func init() {
	initProperties()
}

var zipsaPropInstance *zipsaProp

type zipsaProp struct {
	redisHost                 string
	redisPort                 string
	redisPassword             string
	rabbitmqHost              string
	rabbitmqPort              string
	rabbitmqVirtualhost       string
	rabbitmqUsername          string
	rabbitmqPassword          string
	rabbitmqDeadLogQueue      string
	rabbitmqDeadLogExchange   string
	rabbitmqDeadLogRoutingkey string
	rabbitmqLogQueue          string
	rabbitmqLogExchange       string
	rabbitmqLogRoutingkey     string
	rabbitmqUseSsl            bool
	logLevel                  string
	logOut                    string
}

func initProperties() *zipsaProp {

	if zipsaPropInstance == nil {
		p := properties.MustLoadFiles([]string{
			fmt.Sprintf("conf/%s/log.properties", Profile),
			fmt.Sprintf("conf/%s/rabbitmq.properties", Profile),
			fmt.Sprintf("conf/%s/redis.properties", Profile),
		}, properties.UTF8, false)

		zipsaPropInstance = &zipsaProp{
			p.MustGetString("redis.host"),
			p.MustGetString("redis.port"),
			p.MustGetString("redis.password"),
			p.MustGetString("rabbitmq.host"),
			p.MustGetString("rabbitmq.port"),
			p.MustGetString("rabbitmq.virtualhost"),
			p.MustGetString("rabbitmq.username"),
			p.MustGetString("rabbitmq.password"),
			p.MustGetString("rabbitmq.dead-log-queue"),
			p.MustGetString("rabbitmq.dead-log-exchange"),
			p.MustGetString("rabbitmq.dead-log-routingkey"),
			p.MustGetString("rabbitmq.log-queue"),
			p.MustGetString("rabbitmq.log-exchange"),
			p.MustGetString("rabbitmq.log-routingkey"),
			p.MustGetBool("rabbitmq.use-ssl"),
			p.MustGetString("log.level"),
			p.MustGetString("log.out"),
		}
	}

	printProperties()

	return zipsaPropInstance
}

func printProperties() {
	log.Printf("Profile = %s", Profile)
	log.Printf("redis.host = %s", GetRedisHost())
	log.Printf("redis.port = %s", GetRedisPort())
	log.Printf("redis.password = %s", GetRedisPassword())
	log.Printf("rabbitmq.host = %s", GetRabbitmqHost())
	log.Printf("rabbitmq.port = %s", GetRabbitmqPort())
	log.Printf("rabbitmq.virtualhost = %s", GetRabbitmqVirtualhost())
	log.Printf("rabbitmq.username = %s", GetRabbitmqUsername())
	log.Printf("rabbitmq.password = %s", GetRabbitmqPassword())
	log.Printf("rabbitmq.dead-log-queue = %s", GetRabbitmqDeadLogQueue())
	log.Printf("rabbitmq.dead-log-exchange = %s", GetRabbitmqDeadLogExchange())
	log.Printf("rabbitmq.dead-log-routingkey = %s", GetRabbitmqDeadLogRoutingkey())
	log.Printf("rabbitmq.log-queue = %s", GetRabbitmqLogQueue())
	log.Printf("rabbitmq.log-exchange = %s", GetRabbitmqLogExchange())
	log.Printf("rabbitmq.log-routingkey = %s", GetRabbitmqLogRoutingkey())
	log.Printf("rabbitmq.use-ssl = %t", GetRabbitmqUseSsl())
	log.Printf("log.level = %s", GetLogLevel())
	log.Printf("log.out = %s", GetLogOut())
}

func GetRedisHost() string {
	return zipsaPropInstance.redisHost
}

func GetRedisPort() string {
	return zipsaPropInstance.redisPort
}

func GetRedisPassword() string {
	return zipsaPropInstance.redisPassword
}

func GetRabbitmqHost() string {
	return zipsaPropInstance.rabbitmqHost
}

func GetRabbitmqPort() string {
	return zipsaPropInstance.rabbitmqPort
}

func GetRabbitmqVirtualhost() string {
	return zipsaPropInstance.rabbitmqVirtualhost
}

func GetRabbitmqUsername() string {
	return zipsaPropInstance.rabbitmqUsername
}

func GetRabbitmqPassword() string {
	return zipsaPropInstance.rabbitmqPassword
}

func GetRabbitmqDeadLogQueue() string {
	return zipsaPropInstance.rabbitmqDeadLogQueue
}

func GetRabbitmqDeadLogExchange() string {
	return zipsaPropInstance.rabbitmqDeadLogExchange
}

func GetRabbitmqDeadLogRoutingkey() string {
	return zipsaPropInstance.rabbitmqDeadLogRoutingkey
}

func GetRabbitmqLogQueue() string {
	return zipsaPropInstance.rabbitmqLogQueue
}

func GetRabbitmqLogExchange() string {
	return zipsaPropInstance.rabbitmqLogExchange
}

func GetRabbitmqLogRoutingkey() string {
	return zipsaPropInstance.rabbitmqLogRoutingkey
}

func GetRabbitmqUseSsl() bool {
	return zipsaPropInstance.rabbitmqUseSsl
}

func GetLogLevel() string {
	return zipsaPropInstance.logLevel
}

func GetLogOut() string {
	return zipsaPropInstance.logOut
}
