package etherscan

func ShowGasPrice() {
	client := GetEtherScanClient()
	price, err := client.GasOracle()
	if nil != err {
		return
	}
	println(price.SafeGasPrice)
}

func GetTop100() {

}
