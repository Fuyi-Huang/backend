package utils

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/fuyi-huang/backend/config"
)

var counter uint32 = 0

func get_counter() uint32 {
	temp := atomic.AddUint32(&counter, 1)

	return temp % 1000
}

func OrderStr() string {
	now := time.Now()
	microsec := int(now.UnixNano() % int64(time.Microsecond))
	uuid := now.Format("20060102150405") + strconv.Itoa(microsec) + strconv.Itoa(config.Mechine_id) + fmt.Sprint(get_counter())
	return uuid
}

func OrderNumber() int64 {
	uuid := time.Now().UnixNano()/1000*10000 + int64(config.Mechine_id)*1000 + int64(get_counter())
	return uuid
}
