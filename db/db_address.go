package db

import (
	"database/sql"
)

func IsAddressContract(addr string) int {
	db := GetDB()
	if nil == db {
		return -1
	}

	var fAddress string
	var fIsContract int
	rows := db.QueryRow("select * from address where address=" + addr) //获取一行数据
	rows.Scan(&fAddress, &fIsContract)
	if "" == fAddress {
		return -1
	}

	return fIsContract
}

func SetAddressContract(addr string, isContract int) {
	db := GetDB()
	if nil == db {
		return
	}

	var stmt *sql.Stmt
	var err error
	stmt, err = db.Prepare("INSERT INTO address(address, isContract) values(?,?)")
	if err != nil {
		print(err)
		return
	}

	_, err = stmt.Exec(addr, isContract)
	if err != nil {
		print(err)
		return
	}
}
