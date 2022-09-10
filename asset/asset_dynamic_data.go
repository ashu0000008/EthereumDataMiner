package asset

import (
	"EthereumDataMiner/asset/erc20"
	"EthereumDataMiner/parser"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type CoinDynamicData struct {
	contract string
	date     string
	supply   big.Float
	price    big.Float
}

func GetThresholdDynamicData(contractAddress string) *CoinDynamicData {
	client := parser.GetClient()
	contract, err := erc20.NewErc20(common.HexToAddress(contractAddress), client)
	if nil != err {
		return nil
	}

	decimal, err := contract.Decimals(nil)
	if nil != err {
		return nil
	}

	totalSupply, err := contract.TotalSupply(nil)
	if nil != err {
		return nil
	}

	totalSupplyFloat := GetBalance(totalSupply, decimal)
	println(totalSupplyFloat)

	dataNow := GetDateNow()
	return &CoinDynamicData{contract: contractAddress, date: dataNow, supply: *totalSupplyFloat, price: *big.NewFloat(0)}
}
