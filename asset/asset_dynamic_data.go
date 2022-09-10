package asset

import (
	"EthereumDataMiner/asset/tbtcv2"
	"EthereumDataMiner/db"
	"EthereumDataMiner/parser"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const thresholdContract = "0x18084fbA666a33d37592fA2633fD49a74DD93a88"

func GetThresholdDynamicData() {
	client := parser.GetClient()
	contract, err := tbtcv2.NewContract(common.HexToAddress(thresholdContract), client)
	if nil != err {
		return
	}

	decimal, err := contract.Decimals(nil)
	if nil != err {
		return
	}

	totalSupply, err := contract.TotalSupply(nil)
	if nil != err {
		return
	}

	totalSupplyFloat := GetBalance(totalSupply, decimal)
	println(totalSupplyFloat)

	dataNow := GetDateNow()
	db.InsertOneDayInfo(dataNow, thresholdContract, *totalSupplyFloat, *big.NewFloat(0))
}
