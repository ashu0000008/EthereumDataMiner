package main

import (
	"EthereumDataMiner/asset"
	"EthereumDataMiner/etherscan"
	"EthereumDataMiner/task"
	"fmt"
	"time"
)

func main() {

	etherscan.ShowGasPrice()
	//asset.GetStakedETHInfo()

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

	taskStakedETH := func() {
		fmt.Println("taskStakedETH start------> ", time.Now())
		asset.GetStakedETHInfo()
		fmt.Println("taskStakedETH end------> ", time.Now())
	}

	task.StartTimerTask(7, taskUsd)
	task.StartTimerTask(8, taskStakedETH)
	task.StartTimerTask(8, taskThreshold)

	task.StartTimerTask(17, taskStakedETH)
	select {} //阻塞主线程停止
}
