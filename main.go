package main

import (
	"EthereumDataMiner/asset"
	"EthereumDataMiner/task"
	"fmt"
	"time"
)

func main() {

	taskThreshold := func() {
		fmt.Println("taskThreshold start------> ", time.Now())
		asset.GetThresholdDynamicData()
		fmt.Println("taskThreshold end------> ", time.Now())
	}

	task.StartTimerTask(2, taskThreshold)
	select {} //阻塞主线程停止
}
