package main

import (
	"github.com/RyaoChengfeng/wzj-checkin/util"
	"sync"
)

func main() {
	go util.StartCron()
	//util.CheckInTask()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
