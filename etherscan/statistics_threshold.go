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
	client := GetEtherScanClient()
	//bn, err := client.BlockNumber(0, "")
	//if nil != err {
	//	return
	//}

	logs, err := client.GetLogs(18233123, 18233125, CONTRACT_ADDRESS_THRESHOLD, "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	if nil != err {
		return
	}
	println(logs)
}
