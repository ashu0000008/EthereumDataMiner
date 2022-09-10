package db

import (
	"database/sql"
	"math/big"
)

func InsertOneDayInfo2(dateNow string, contract string, supply big.Float, price big.Float) {
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

func findInfo(db *sql.DB, dateS string, contract string) (*big.Float, *big.Float, error) {
	//queryString := fmt.Sprintf("SELECT * from AssetDynamicInfo where contract=%s;", contract)
	rsk := db.QueryRow("select supply,price from AssetSeeker.AssetDynamicInfo where contract = ? and  dateF = ?", contract, dateS)
	if nil != rsk.Err() {
		return nil, nil, rsk.Err()
	} else {
		var supply float64
		var price float64
		err := rsk.Scan(&supply, &price)
		if nil != err {
			return nil, nil, err
		}
		return big.NewFloat(supply), big.NewFloat(price), nil
	}
}

func InsertOneDayInfo(dateNow string, contract string, supply big.Float, price big.Float) {
	db := GetDB()
	if nil == db {
		return
	}

	_, _, err := findInfo(db, dateNow, contract)
	if nil == err {
		//update data
		return
	}

	var stmt *sql.Stmt
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
