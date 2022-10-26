package main

import (
	"EthereumDataMiner/asset"
	"EthereumDataMiner/task"
	"fmt"
	"time"
)

func main() {

	//asset.GetUSDInfo()

	taskThreshold := func() {
		fmt.Println("taskThreshold start------> ", time.Now())
		asset.GetThresholdInfo()
		fmt.Println("taskThreshold end------> ", time.Now())
	}

	taskUsd := func() {
		fmt.Println("taskThreshold start------> ", time.Now())
		asset.GetUSDInfo()
		fmt.Println("taskThreshold end------> ", time.Now())
	}

	task.StartTimerTask(9, taskThreshold)
	task.StartTimerTask(17, taskThreshold)
	task.StartTimerTask(23, taskThreshold)

	task.StartTimerTask(8, taskUsd)
	select {} //阻塞主线程停止
}
