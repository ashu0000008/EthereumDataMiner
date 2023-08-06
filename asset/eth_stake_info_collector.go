package asset

import (
	"EthereumDataMiner/db"
	"fmt"
)

var someETHContracts = [...]string{
	"0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84", //stETH
	"0xbe9895146f7af43049ca1c1ae358b0541ea49704", //CBETH
	"0xae78736Cd615f374D3085123A210448E74Fc6393", //rETH
	"0x5E8422345238F34275888049021821E8E08CAa1f", //fraxETH
	"0x0100546F2cD4C9D97f798fFC9755E47865FF7Ee6", //alETH
	"0xFe2e637202056d30016725477c5da089Ab0A043A", //sETH2
	"0xDFe66B14D37C77F4E9b180cEb433d1b164f0281D", //StakeHound
	"0xE95A203B1a91a908F9B9CE46459d101078c2c3cb", //ANKRETH
	"0xc3D088842DcF02C13699F936BB83DFBBc6f721Ab", //vETH
}

func GetStakedETHInfo() {
	msg := ""

	for i, contract := range someETHContracts {
		info := GetThresholdDynamicData(contract)
		if nil == info {
			return
		}

		if i == 0 {
			msg = fmt.Sprintf("%s    %s", "stETH", info.supply.String())
		} else if i == 1 {
			msg = fmt.Sprintf("%s\n%s    %s", msg, "CBETH", info.supply.String())
		} else if i == 2 {
			msg = fmt.Sprintf("%s\n%s      %s", msg, "rETH", info.supply.String())
		} else if i == 3 {
			msg = fmt.Sprintf("%s\n%s      %s", msg, "fraxETH", info.supply.String())
		} else if i == 4 {
			msg = fmt.Sprintf("%s\n%s      %s", msg, "alETH", info.supply.String())
		} else if i == 5 {
			msg = fmt.Sprintf("%s\n%s      %s", msg, "sETH2", info.supply.String())
		} else if i == 6 {
			msg = fmt.Sprintf("%s\n%s      %s", msg, "StakeHound", info.supply.String())
		} else if i == 7 {
			msg = fmt.Sprintf("%s\n%s      %s", msg, "ANKRETH", info.supply.String())
		} else if i == 8 {
			msg = fmt.Sprintf("%s\n%s      %s", msg, "vETH", info.supply.String())
		}

		db.InsertOneDayInfo(info.date, info.contract, info.supply, info.price)
	}

	//robot.DispatchThresholdInfo(msg)
}
