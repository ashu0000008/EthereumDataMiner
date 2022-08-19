package parser

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
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

	//parse transaction
	for _, trans := range blockInfo.Transactions() {
		log.Println(trans)
		parseOneTransaction(blockInfo, trans)
	}

	return
}

func parseOneTransaction(block *types.Block, trans *types.Transaction) *DataTransaction {
	msg, err := trans.AsMessage(types.NewEIP155Signer(trans.ChainId()), block.BaseFee())
	if nil != err {
		return nil
	}

	from := msg.From().Hex()
	to := trans.To().Hex()

	isFromContract := isAddressContract(msg.From())
	isToContract := isAddressContract(*trans.To())

	return &DataTransaction{
		Block:          block.Number().String(),
		TxHash:         trans.Hash().Hex(),
		From:           from,
		To:             to,
		FromIsContract: isFromContract,
		ToIsContract:   isToContract,
	}
}
