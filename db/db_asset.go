package db

import (
	"database/sql"
	"log"
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
	rsk := db.QueryRow("select * from AssetDynamicInfo where contract = ? and  dateF = ?", contract, dateS)
	if nil != rsk.Err() {
		return nil, nil, rsk.Err()
	} else {
		var supply float64
		var price float64
		var values = make([]interface{}, 4)
		log.Println(rsk)
		rsk.Scan(values...)

		//supplyF, err := strconv.ParseFloat(supply, 64)
		//if nil == err {
		//	return nil, nil, err
		//}
		//priceF, err := strconv.ParseFloat(price, 64)
		//if nil == err {
		//	return nil, nil, err
		//}
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
