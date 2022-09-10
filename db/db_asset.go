package db

import (
	"database/sql"
	"math/big"
)

func InsertOneDayInfo(dateNow string, contract string, supply big.Float, price big.Float) {
	db := GetDB()
	if nil == db {
		return
	}

	var stmt *sql.Stmt
	var err error
	stmt, err = db.Prepare("INSERT INTO AssetDynamicInfo(dateF, contract, supply, price) values(?,?,?,?)")
	if err != nil {
		print(err)
		return
	}

	_, err = stmt.Exec(dateNow, contract, supply.String(), price.String())
	if err != nil {
		print(err)
		return
	}
}
