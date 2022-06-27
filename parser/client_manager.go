package parser

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

const EthNode = "https://eth-mainnet.alchemyapi.io/v2/tZXSQKLFyzp2n5p5COu4zO0EEOYQ96sU"

var mEthClient *ethclient.Client = nil

func getClient() *ethclient.Client {
	if nil != mEthClient {
		return mEthClient
	}

	client, err := ethclient.Dial(EthNode)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	mEthClient = client
	return client
}
