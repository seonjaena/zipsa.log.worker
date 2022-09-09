package properties

// RabbitMQ 설정
const (
	RabbitMQHost              = "b-eb4fa4e5-2217-4f02-aa63-4582e063c667.mq.ap-northeast-2.amazonaws.com"
	RabbitMQUsername          = "sky114z"
	RabbitMQPassword          = "na6080su12!@"
	RabbitMQVirtualhost       = "/"
	RabbitMQDeadLogExchange   = "dead-log-exchange"
	RabbitMQDeadLogQueue      = "dead-log-queue"
	RabbitMQDeadLogRoutingkey = "dead-log-queue"
	RabbitMQLogExchange       = "log-exchange"
	RabbitMQLogQueue          = "log-queue"
	RabbitMQLogRoutingkey     = "log-queue"
	RabbitMQPort              = "5671"
	RabbitMQUseSSL            = true
)

// Redis 설정
const (
	RedisHost     = "144.24.83.179"
	RedisPassword = "na6080su12!@"
	RedisPort     = "6379"
)

// 공통 설정
const (
	LogOut   = "stdout"
	LogLevel = "debug"
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
	rabbitmqUseSsl            bool
	logLevel                  string
	logOut                    string
}

func initProperties() *zipsaProp {
	if zipsaPropInstance == nil {
		zipsaPropInstance = &zipsaProp{
			RedisHost,
			RedisPort,
			RedisPassword,
			RabbitMQHost,
			RabbitMQPort,
			RabbitMQVirtualhost,
			RabbitMQUsername,
			RabbitMQPassword,
			RabbitMQDeadLogQueue,
			RabbitMQDeadLogExchange,
			RabbitMQDeadLogRoutingkey,
			RabbitMQLogQueue,
			RabbitMQLogExchange,
			RabbitMQLogRoutingkey,
			RabbitMQUseSSL,
			LogLevel,
			LogOut,
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

func GetRabbitmqUseSsl() bool {
	return zipsaPropInstance.rabbitmqUseSsl
}

func GetLogLevel() string {
	return zipsaPropInstance.logLevel
}

func GetLogOut() string {
	return zipsaPropInstance.logOut
}
