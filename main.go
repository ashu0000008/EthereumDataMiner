package main

import (
	"EthereumDataMiner/task"
	"fmt"
	"time"
)

func main() {

	taskThreshold := func() {
		fmt.Println("taskThreshold start------> ", time.Now())
	}

	task.StartTimerTask(1, taskThreshold)
	select {} //阻塞主线程停止
}
