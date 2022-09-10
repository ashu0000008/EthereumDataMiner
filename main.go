package main

import (
	"EthereumDataMiner/asset"
)

func main() {

	asset.GetThresholdInfo()

	//taskThreshold := func() {
	//	fmt.Println("taskThreshold start------> ", time.Now())
	//	asset.GetThresholdInfo()
	//	fmt.Println("taskThreshold end------> ", time.Now())
	//}
	//
	//task.StartTimerTask(10, taskThreshold)
	//select {} //阻塞主线程停止
}
