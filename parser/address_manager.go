package parser

import (
	"EthereumDataMiner/db"
	"context"
	"github.com/ethereum/go-ethereum/common"
)

func isAddressContract(addr common.Address) int {
	isContract := db.IsAddressContract(addr.Hex())
	if 0 == isContract || 1 == isContract {
		return isContract
	}

	client := GetClient()
	code, err := client.CodeAt(context.Background(), addr, nil)
	if nil == err {
		var result = 0
		if nil == code || len(code) == 0 {
			result = 0
		} else {
			result = 1
		}

		//记录数据库
		db.SetAddressContract(addr.Hex(), result)

		return result
	} else {
		return 0
	}
}
