package asset

import (
	"EthereumDataMiner/db"
)

const tbtcv2Contract = "0x18084fbA666a33d37592fA2633fD49a74DD93a88"

func GetThresholdInfo() {
	info := GetThresholdDynamicData(tbtcv2Contract)
	if nil == info {
		return
	}
	db.InsertOneDayInfo(info.date, info.contract, info.supply, info.price)
}
