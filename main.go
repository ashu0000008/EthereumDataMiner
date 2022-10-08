package main

import (
	"EthereumDataMiner/asset"
	"EthereumDataMiner/task"
	"fmt"
	"time"
)

func main() {

	//asset.GetThresholdInfo()

	taskThreshold := func() {
		fmt.Println("taskThreshold start------> ", time.Now())
		asset.GetThresholdInfo()
		fmt.Println("taskThreshold end------> ", time.Now())
	}

	task.StartTimerTask(9, taskThreshold)
	task.StartTimerTask(17, taskThreshold)
	task.StartTimerTask(23, taskThreshold)
	select {} //阻塞主线程停止
}
