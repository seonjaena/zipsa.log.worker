package properties

import "github.com/magiconair/properties"

func init() {

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
		p := properties.MustLoadFiles([]string{
			"conf/rabbitmq.properties",
			"conf/redis.properties",
			"conf/common.properties",
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
			p.MustGetString("rabbitmq.use-ssl"),
			p.MustGetString("log.level"),
			p.MustGetString("log.out"),
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
