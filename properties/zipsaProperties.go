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
	redisHost               string
	redisPort               string
	redisPassword           string
	redisDB                 int
	redisBufferSize         int
	redisFlushIntervalMS    int
	rabbitmqHost            string
	rabbitmqPort            string
	rabbitmqVirtualhost     string
	rabbitmqUsername        string
	rabbitmqPassword        string
	rabbitmqUseSsl          bool
	rabbitmqDeadLogQueue    string
	rabbitmqDeadLogExchange string
	rabbitmqWaitLogQueue    string
	rabbitmqWaitLogExchange string
	rabbitmqLogQueue        string
	rabbitmqLogExchange     string
	rabbitmqDeadLogTTL      int
	rabbitmqPrefetchCnt     int
	logLevel                string
	logOut                  string
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
			p.MustGetInt("redis.db"),
			p.MustGetInt("redis.buffer-size"),
			p.MustGetInt("redis.flush-interval-ms"),
			p.MustGetString("rabbitmq.host"),
			p.MustGetString("rabbitmq.port"),
			p.MustGetString("rabbitmq.virtualhost"),
			p.MustGetString("rabbitmq.username"),
			p.MustGetString("rabbitmq.password"),
			p.MustGetBool("rabbitmq.use-ssl"),
			p.MustGetString("rabbitmq.dead-log-queue"),
			p.MustGetString("rabbitmq.dead-log-exchange"),
			p.MustGetString("rabbitmq.wait-log-queue"),
			p.MustGetString("rabbitmq.wait-log-exchange"),
			p.MustGetString("rabbitmq.log-queue"),
			p.MustGetString("rabbitmq.log-exchange"),
			p.MustGetInt("rabbitmq.dead-log-ttl"),
			p.MustGetInt("rabbitmq.prefetch-cnt"),
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
	log.Printf("redis.db = %d", GetRedisDB())
	log.Printf("redis.buffer-size = %d", GetRedisBufferSize())
	log.Printf("redis.flush-interval-ms = %d", GetRedisFlushIntervalMS())
	log.Printf("rabbitmq.host = %s", GetRabbitmqHost())
	log.Printf("rabbitmq.port = %s", GetRabbitmqPort())
	log.Printf("rabbitmq.virtualhost = %s", GetRabbitmqVirtualhost())
	log.Printf("rabbitmq.username = %s", GetRabbitmqUsername())
	log.Printf("rabbitmq.password = %s", GetRabbitmqPassword())
	log.Printf("rabbitmq.use-ssl = %t", GetRabbitmqUseSsl())
	log.Printf("rabbitmq.dead-log-queue = %s", GetRabbitmqDeadLogQueue())
	log.Printf("rabbitmq.dead-log-exchange = %s", GetRabbitmqDeadLogExchange())
	log.Printf("rabbitmq.wait-log-queue = %s", GetRabbitmqWaitLogQueue())
	log.Printf("rabbitmq.wait-log-exchange = %s", GetRabbitmqWaitLogExchange())
	log.Printf("rabbitmq.log-queue = %s", GetRabbitmqLogQueue())
	log.Printf("rabbitmq.log-exchange = %s", GetRabbitmqLogExchange())
	log.Printf("rabbitmq.dead-log-ttl = %d", GetRabbitmqDeadLogTTL())
	log.Printf("rabbitmq.prefetch-cnt = %d", GetRabbitmqPrefetchCnt())
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

func GetRedisDB() int {
	return zipsaPropInstance.redisDB
}

func GetRedisBufferSize() int {
	return zipsaPropInstance.redisBufferSize
}

func GetRedisFlushIntervalMS() int {
	return zipsaPropInstance.redisFlushIntervalMS
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

func GetRabbitmqWaitLogQueue() string {
	return zipsaPropInstance.rabbitmqWaitLogQueue
}

func GetRabbitmqWaitLogExchange() string {
	return zipsaPropInstance.rabbitmqWaitLogExchange
}

func GetRabbitmqLogQueue() string {
	return zipsaPropInstance.rabbitmqLogQueue
}

func GetRabbitmqLogExchange() string {
	return zipsaPropInstance.rabbitmqLogExchange
}

func GetRabbitmqUseSsl() bool {
	return zipsaPropInstance.rabbitmqUseSsl
}

func GetRabbitmqDeadLogTTL() int {
	return zipsaPropInstance.rabbitmqDeadLogTTL
}

func GetRabbitmqPrefetchCnt() int {
	return zipsaPropInstance.rabbitmqPrefetchCnt
}

func GetLogLevel() string {
	return zipsaPropInstance.logLevel
}

func GetLogOut() string {
	return zipsaPropInstance.logOut
}
