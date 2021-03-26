package util

import (
	"fmt"
	"time"
)

func TimeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Println("函数执行完成耗时：", tc)
	}
}
