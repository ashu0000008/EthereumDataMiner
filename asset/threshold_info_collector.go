package asset

import (
	"EthereumDataMiner/db"
	"EthereumDataMiner/robot"
)

var thresholdContracts = [...]string{
	"0x18084fbA666a33d37592fA2633fD49a74DD93a88", //tbtcv2
	"0x8dAEBADE922dF735c38C80C7eBD708Af50815fAa", //tbtcv1
}

func GetThresholdInfo() {
	for _, contract := range thresholdContracts {
		info := GetThresholdDynamicData(contract)
		if nil == info {
			return
		}
		db.InsertOneDayInfo(info.date, info.contract, info.supply, info.price)
	}

	robot.DispatchThresholdInfo("xxx")
}
