package statistics

import (
	"EthereumDataMiner/asset/erc20"
	"EthereumDataMiner/etherscan"
	"EthereumDataMiner/parser"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/context"
)

func GetTop100() {
	client := parser.GetClient()
	if nil == client {
		return
	}

	erc20Threshold, err := erc20.NewErc20(common.HexToAddress(etherscan.CONTRACT_ADDRESS_THRESHOLD), client)
	if nil != err {
		return
	}

	blockNum, err := client.BlockNumber(context.Background())
	if nil != err {
		return
	}
	blockNum = uint64(18233125)
	opts := bind.FilterOpts{Start: blockNum - 2, End: &blockNum, Context: context.Background()}
	iterator, err := erc20Threshold.Erc20Filterer.FilterTransfer(&opts, nil, nil)
	if nil != err {
		return
	}
	println(iterator)
}
