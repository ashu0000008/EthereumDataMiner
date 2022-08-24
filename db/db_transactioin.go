package db

import (
	"EthereumDataMiner/model"
	"database/sql"
)

func SaveOneTransaction(trans model.DataTransaction) {
	db := GetDB()
	if nil == db {
		return
	}

	var stmt *sql.Stmt
	var err error
	stmt, err = db.Prepare("INSERT INTO transaction(txhash, block, from, to) values(?,?,?,?)")
	if err != nil {
		print(err)
		return
	}

	_, err = stmt.Exec(trans.TxHash, trans.Block, trans.From, trans.To)
	if err != nil {
		print(err)
		return
	}
}
