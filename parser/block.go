package parser

import (
	"context"
	"log"
	"math/big"
)

func GetBlockInfo(blockNumber uint64) {
	client := getClient()

	blockInfo, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
	if nil != err {
		return
	} else {
		log.Println(blockInfo)
	}

	return
}
