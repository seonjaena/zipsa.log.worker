package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"sync"
	"time"
	zp "zipsa.log.worker/properties"
	"zipsa.log.worker/zlog"
)

type redisLog struct {
	d    *amqp.Delivery
	keys []string
}

type safeBuffer struct {
	buffer        chan redisLog
	innerBuffer   [][]string
	lastDelivery  *amqp.Delivery
	updateLocker  *sync.Mutex
	consumeLocker chan bool
	updateTrigger chan bool
	cleanTrigger  chan bool
}

const (
	Total string = "TOTAL"
	NoDup string = "XDUP"
	OkDup string = "ODUP"
)

var log = zlog.Instance()
var bCtx = context.Background()
var redisClient *redis.Client

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

func (buffer *safeBuffer) Append(body string, d *amqp.Delivery) error {
	_, _ = parseAccessLog(body)
	return nil
}

func parseAccessLog(body string) ([]string, error) {

	curDate := time.Now()

	//월 간 전체 접속자 수 (중복 제거)
	monthTotalNDAccess := fmt.Sprintf("%s^%s^%s", curDate.Format("2006-01"), Total, NoDup)
	//일 간 전체 접속자 수 (중복 제거)
	dayTotalNDAccess := fmt.Sprintf("%s^%s^%s", curDate.Format("2006-01-02"), Total, NoDup)
	//월 간 전체 접속자 수 (중복 허용)
	monthTotalODAccess := fmt.Sprintf("%s^%s^%s", curDate.Format("2006-01"), Total, OkDup)
	//일 간 전체 접속자 수 (중복 허용)
	dayTotalODAccess := fmt.Sprintf("%s^%s^%s", curDate.Format("2006-01-02"), Total, OkDup)
	//월 간 건물 당 접속자 수 (중복 제거)
	monthBuildingNDAccess := fmt.Sprintf("%s^%s^%s", curDate.Format("2006-01"), body, NoDup)
	//일 간 건물 당 접속자 수 (중복 제거)
	dayBuildingNDAccess := fmt.Sprintf("%s^%s^%s", curDate.Format("2006-01-02"), body, NoDup)
	//월 간 건물 당 접속자 수 (중복 허용)
	monthBuildingODAccess := fmt.Sprintf("%s^%s^%s", curDate.Format("2006-01"), body, OkDup)
	//일 간 건물 당 접속자 수 (중복 허용)
	dayBuildingODAccess := fmt.Sprintf("%s^%s^%s", curDate.Format("2006-01-02"), body, OkDup)

	return []string{
		monthTotalNDAccess,
		dayTotalNDAccess,
		monthTotalODAccess,
		dayTotalODAccess,
		monthBuildingNDAccess,
		dayBuildingNDAccess,
		monthBuildingODAccess,
		dayBuildingODAccess,
	}, nil
}
