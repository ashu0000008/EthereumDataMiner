package etherscan

import "github.com/nanmu42/etherscan-api"

var _client *etherscan.Client = nil

func GetEtherScanClient() *etherscan.Client {
	if nil != _client {
		return _client
	}

	_client := etherscan.New(etherscan.Mainnet, KEY)
	return _client
}
