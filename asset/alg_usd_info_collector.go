package asset

import (
	"EthereumDataMiner/db"
	"EthereumDataMiner/robot"
	"fmt"
	"math/big"
)

var algUsdContracts = [...]string{
	"0xdAC17F958D2ee523a2206206994597C13D831ec7", //usdt
	"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", //usdc
	"0x4Fabb145d64652a948d72533023f6E7A623C7C53", //busd
}

func GetAlgUSDInfo() {
	msg := ""
	total := big.NewFloat(0)

	for i, contract := range algUsdContracts {

		info := GetThresholdDynamicData(contract)
		if nil == info {
			return
		}

		if i == 0 {
			msg = fmt.Sprintf("%s         %s", "usdt", info.supply.String())
		} else if i == 1 {
			msg = fmt.Sprintf("%s\n%s         %s", msg, "usdc", info.supply.String())
		} else if i == 2 {
			msg = fmt.Sprintf("%s\n%s         %s", msg, "busd", info.supply.String())
		}

		db.InsertOneDayInfo(info.date, info.contract, info.supply, info.price)
		total.Add(total, &info.supply)
	}

	msg = fmt.Sprintf("%s\n%s    %s", msg, "usd total", total.String())
	robot.DispatchThresholdInfo(msg)
}
