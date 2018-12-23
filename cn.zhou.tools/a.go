package cn_zhou_tools

import (
	"fmt"
	"time"
)
/*
本案例实现了定时器，定时结束程序的功能
time.Timer()
time.Ticker()

 */
func Break() {

	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()

	for e := range ticker.C {
		 fmt.Println("时间：",e)
		 panic("the  1 minute is arrive to.....")
		//fmt.Println(e,"the 20s is arrive to.....")
	}

}

