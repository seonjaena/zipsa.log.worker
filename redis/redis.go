package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"strings"
	"unicode"
	zp "zipsa.log.worker/properties"
	"zipsa.log.worker/zlog"
)

type redisLog struct {
	d    *amqp.Delivery
	keys []string
	body string
}

const (
	Total string = "TOTAL"
	NoDup string = "XDUP"
	OkDup string = "ODUP"
)

var log = zlog.Instance()
var bCtx = context.Background()
var redisClient *redis.Client
var LogStruct *redisLog

func init() {
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

func (redisLog *redisLog) ProcessAccessLog(data string, delivery *amqp.Delivery) {
	keys, body, err := parseAccessLog(data)
	if err != nil {
		log.Errorf("Error Occurred")
	}
	redisLog.d = delivery
	redisLog.keys = keys
	redisLog.body = body
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

	//월 간 전체 접속자 수 (중복 제거)
	monthTotalNDAccess := fmt.Sprintf("%s^%s^%s", monthlyDate, Total, NoDup)
	//일 간 전체 접속자 수 (중복 제거)
	dayTotalNDAccess := fmt.Sprintf("%s^%s^%s", dailyDate, Total, NoDup)
	//월 간 전체 접속자 수 (중복 허용)
	monthTotalODAccess := fmt.Sprintf("%s^%s^%s", monthlyDate, Total, OkDup)
	//일 간 전체 접속자 수 (중복 허용)
	dayTotalODAccess := fmt.Sprintf("%s^%s^%s", dailyDate, Total, OkDup)
	//월 간 건물 당 접속자 수 (중복 제거)
	monthBuildingNDAccess := fmt.Sprintf("%s^%s^%s", monthlyDate, buildingNo, NoDup)
	//일 간 건물 당 접속자 수 (중복 제거)
	dayBuildingNDAccess := fmt.Sprintf("%s^%s^%s", dailyDate, buildingNo, NoDup)
	//월 간 건물 당 접속자 수 (중복 허용)
	monthBuildingODAccess := fmt.Sprintf("%s^%s^%s", monthlyDate, buildingNo, OkDup)
	//일 간 건물 당 접속자 수 (중복 허용)
	dayBuildingODAccess := fmt.Sprintf("%s^%s^%s", dailyDate, buildingNo, OkDup)

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

	if len(date) == 10 {
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

func (redisLog *redisLog) updateToRedis(keys []string, data string) bool {
	pipe := redisClient.TxPipeline()
	pipe.PFAdd(bCtx, keys[0], data)
	pipe.PFAdd(bCtx, keys[1], data)
	pipe.Incr(bCtx, keys[2])
	pipe.Incr(bCtx, keys[3])
	for i := 4; i <= 5; i++ {
		if keys[i] != "" {
			pipe.PFAdd(bCtx, keys[i], data)
		}
	}

	for i := 6; i <= 7; i++ {
		if keys[i] != "" {
			pipe.Incr(bCtx, keys[i])
		}
	}
	_, err := pipe.Exec(bCtx)
	if err != nil {
		log.Errorf("Failed to update to redis, error = %s", err)
		_ = pipe.Discard()
		return false
	} else {
		_ = pipe.Close()
		return true
	}
}
