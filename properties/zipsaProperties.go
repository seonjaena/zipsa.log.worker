package properties

import (
	"os"
)

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
	rabbitmqUseSsl            string
	logLevel                  string
	logOut                    string
}

func initProperties() *zipsaProp {
	if zipsaPropInstance == nil {
		zipsaPropInstance = &zipsaProp{
			os.Getenv("redis_host"),
			os.Getenv("redis_port"),
			os.Getenv("redis_password"),
			os.Getenv("rabbitmq_host"),
			os.Getenv("rabbitmq_port"),
			os.Getenv("rabbitmq_virtualhost"),
			os.Getenv("rabbitmq_username"),
			os.Getenv("rabbitmq_password"),
			os.Getenv("rabbitmq_dead_log_queue"),
			os.Getenv("rabbitmq_dead_log_exchange"),
			os.Getenv("rabbitmq_dead_log_routingkey"),
			os.Getenv("rabbitmq_log_queue"),
			os.Getenv("rabbitmq_log_exchange"),
			os.Getenv("rabbitmq_log_routingkey"),
			os.Getenv("rabbitmq_use_ssl"),
			os.Getenv("log_level"),
			os.Getenv("log_out"),
		}
	}

	return zipsaPropInstance
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

func GetRabbitmqUseSsl() string {
	return zipsaPropInstance.rabbitmqUseSsl
}

func GetLogLevel() string {
	return zipsaPropInstance.logLevel
}

func GetLogOut() string {
	return zipsaPropInstance.logOut
}
