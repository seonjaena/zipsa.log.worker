package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"strings"
	"sync"
	"time"
	"unicode"
	zp "zipsa.log.worker/properties"
	"zipsa.log.worker/zlog"
)

type redisLog struct {
	d    *amqp.Delivery
	keys []string
	body string
}

type logBuffer struct {
	buffer          chan redisLog
	innerKeyBuffer  [][]string
	innerDataBuffer []string
	lastBuffer      *amqp.Delivery
	updateLocker    *sync.Mutex
	consumeLocker   chan bool
	updateTrigger   chan bool
}

const (
	Total string = "TOTAL"
	NoDup string = "XDUP"
	OkDup string = "ODUP"
)

var log = zlog.Instance()
var bCtx = context.Background()
var redisClient *redis.Client
var LogBuffer logBuffer

func init() {
	LogBuffer = logBuffer{
		buffer:          make(chan redisLog),
		innerKeyBuffer:  [][]string{},
		innerDataBuffer: []string{},
		lastBuffer:      nil,
		updateLocker:    &sync.Mutex{},
		consumeLocker:   make(chan bool),
		updateTrigger:   make(chan bool),
	}
	GetConn()
}

func GetConn() {
	if redisClient == nil {
		tryConn()
	}
}

func tryConn() {
	for {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", zp.GetRedisHost(), zp.GetRedisPort()),
			Password: zp.GetRedisPassword(),
			DB:       zp.GetRedisDB(),
		})

		result := redisClient.Ping(bCtx)
		message, err := result.Result()
		if err != nil {
			log.Errorf("connect redis failed, message=%s, error=%s", message, err.Error())
		} else {
			log.Infof("redis connected")
			break
		}
	}
}

func (logBuffer *logBuffer) Append(data string, delivery *amqp.Delivery) error {
	keys, body, err := parseAccessLog(data)
	if err != nil {
		log.Errorf("Error Occurred")
	}
	logBuffer.buffer <- redisLog{d: delivery, keys: keys, body: body}
	<-logBuffer.consumeLocker
	return nil
}

func (logBuffer *logBuffer) FlushData() {
	for {
		select {
		case message := <-logBuffer.buffer:
			logBuffer.updateLocker.Lock()
			logBuffer.innerKeyBuffer = append(logBuffer.innerKeyBuffer, message.keys)
			logBuffer.innerDataBuffer = append(logBuffer.innerDataBuffer, message.body)
			logBuffer.lastBuffer = message.d
			if len(logBuffer.innerKeyBuffer) >= zp.GetRedisBufferSize() && logBuffer.lastBuffer != nil {
				logBuffer.updateToRedis()
				for {
					err := logBuffer.lastBuffer.Ack(true)
					if err == nil {
						break
					} else {
						time.Sleep(3 * time.Second)
					}
				}
				logBuffer.innerDataBuffer = nil
			}
			logBuffer.consumeLocker <- true
			logBuffer.updateLocker.Unlock()
		case <-time.After(time.Millisecond * time.Duration(zp.GetRedisFlushIntervalMS())):
			logBuffer.updateLocker.Lock()
			if len(logBuffer.innerKeyBuffer) > 0 && logBuffer.lastBuffer != nil {
				logBuffer.updateToRedis()
				for {
					err := logBuffer.lastBuffer.Ack(true)
					if err == nil {
						break
					} else {
						time.Sleep(3 * time.Second)
					}
				}
				logBuffer.innerKeyBuffer = nil
				logBuffer.innerDataBuffer = nil
			}
			logBuffer.updateLocker.Unlock()
		}
	}
}

func parseAccessLog(body string) ([]string, string, error) {

	messages := strings.Split(body, "^")
	date := messages[0]
	userNo := messages[1]
	buildingNo := messages[2]

	isProperFormat := checkAccessLogFormat(date, userNo, buildingNo)

	if !isProperFormat {
		log.Errorf("not good format")
	}

	monthlyDate := date[0:8]
	dailyDate := date

	var monthTotalNDAccess string
	var dayTotalNDAccess string
	var monthTotalODAccess string
	var dayTotalODAccess string
	var monthBuildingNDAccess string
	var dayBuildingNDAccess string
	var monthBuildingODAccess string
	var dayBuildingODAccess string

	//월 간 전체 접속자 수 (중복 제거)
	monthTotalNDAccess = fmt.Sprintf("%s^%s^%s", monthlyDate, Total, NoDup)
	//일 간 전체 접속자 수 (중복 제거)
	dayTotalNDAccess = fmt.Sprintf("%s^%s^%s", dailyDate, Total, NoDup)
	//월 간 전체 접속자 수 (중복 허용)
	monthTotalODAccess = fmt.Sprintf("%s^%s^%s", monthlyDate, Total, OkDup)
	//일 간 전체 접속자 수 (중복 허용)
	dayTotalODAccess = fmt.Sprintf("%s^%s^%s", dailyDate, Total, OkDup)
	if strings.Compare(strings.Trim(buildingNo, " "), "") != 0 {
		//월 간 건물 당 접속자 수 (중복 제거)
		monthBuildingNDAccess = fmt.Sprintf("%s^%s^%s", monthlyDate, buildingNo, NoDup)
		//일 간 건물 당 접속자 수 (중복 제거)
		dayBuildingNDAccess = fmt.Sprintf("%s^%s^%s", dailyDate, buildingNo, NoDup)
		//월 간 건물 당 접속자 수 (중복 허용)
		monthBuildingODAccess = fmt.Sprintf("%s^%s^%s", monthlyDate, buildingNo, OkDup)
		//일 간 건물 당 접속자 수 (중복 허용)
		dayBuildingODAccess = fmt.Sprintf("%s^%s^%s", dailyDate, buildingNo, OkDup)
	}

	return []string{
		monthTotalNDAccess,
		dayTotalNDAccess,
		monthTotalODAccess,
		dayTotalODAccess,
		monthBuildingNDAccess,
		dayBuildingNDAccess,
		monthBuildingODAccess,
		dayBuildingODAccess,
	}, userNo, nil
}

func checkAccessLogFormat(date string, userNo string, buildingNo string) bool {

	if len(date) != 10 {
		return false
	}

	if strings.Count(date, "-") != 2 {
		return false
	}

	for _, c := range strings.ReplaceAll(date, "-", "") {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	for _, c := range userNo {
		if !unicode.IsDigit(c) {
			log.Errorf("UserIDX is not proper format, userNo = %s", userNo)
			return false
		}
	}

	for _, c := range buildingNo {
		if !unicode.IsDigit(c) {
			log.Errorf("BuildingIDX is not proper format, buildingNo = %s", buildingNo)
			return false
		}
	}

	return true
}

func (logBuffer *logBuffer) updateToRedis() bool {
	pipe := redisClient.TxPipeline()
	for i := 0; i < len(logBuffer.innerKeyBuffer); i++ {
		pipe.PFAdd(bCtx, logBuffer.innerKeyBuffer[i][0], logBuffer.innerDataBuffer[i])
		pipe.PFAdd(bCtx, logBuffer.innerKeyBuffer[i][1], logBuffer.innerDataBuffer[i])
		pipe.Incr(bCtx, logBuffer.innerKeyBuffer[i][2])
		pipe.Incr(bCtx, logBuffer.innerKeyBuffer[i][3])
		for j := 4; j <= 7; j++ {
			if strings.Compare(logBuffer.innerKeyBuffer[i][j], "") == 0 {
				return false
			}
		}
		pipe.PFAdd(bCtx, logBuffer.innerKeyBuffer[i][4], logBuffer.innerDataBuffer[i])
		pipe.PFAdd(bCtx, logBuffer.innerKeyBuffer[i][5], logBuffer.innerDataBuffer[i])
		pipe.Incr(bCtx, logBuffer.innerKeyBuffer[i][6])
		pipe.Incr(bCtx, logBuffer.innerKeyBuffer[i][7])
	}
	_, err := pipe.Exec(bCtx)
	if err != nil {
		log.Errorf("Failed to update to redis, error = %s", err)
		_ = pipe.Discard()
		return false
	} else {
		log.Infof("Success to update to redis")
		_ = pipe.Close()
		return true
	}
}
