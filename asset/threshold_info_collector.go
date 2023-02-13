package asset

import (
	"EthereumDataMiner/db"
	"EthereumDataMiner/robot"
	"fmt"
)

var thresholdContracts = [...]string{
	"0x18084fbA666a33d37592fA2633fD49a74DD93a88", //tbtcv2
	"0x8dAEBADE922dF735c38C80C7eBD708Af50815fAa", //tbtcv1
	"0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599", //wbtc
	"0xEB4C2781e4ebA804CE9a9803C67d0893436bB27D", //renbtc
}

func GetThresholdInfo() {
	msg := ""

	for i, contract := range thresholdContracts {
		info := GetThresholdDynamicData(contract)
		if nil == info {
			return
		}

		if i == 0 {
			msg = fmt.Sprintf("%s    %s", "tbtcv2", info.supply.String())
		} else if i == 1 {
			msg = fmt.Sprintf("%s\n%s    %s", msg, "tbtcv1", info.supply.String())
		} else if i == 2 {
			msg = fmt.Sprintf("%s\n%s      %s", msg, "wbtc", info.supply.String())
		} else if i == 3 {
			msg = fmt.Sprintf("%s\n%s      %s", msg, "renbtc", info.supply.String())
		}

		db.InsertOneDayInfo(info.date, info.contract, info.supply, info.price)
	}

	robot.DispatchThresholdInfo(msg)
}
