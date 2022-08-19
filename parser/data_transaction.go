package parser

type DataTransaction struct {
	Block  string `json:"block"`
	TxHash string `json:"txHash"`
	From   string `json:"from"`
	To     string `json:"to"`

	FromIsContract int `json:"fromIsContract"`
	ToIsContract   int `json:"toIsContract"`
}

type DataAddress struct {
	Address    string `json:"from"`
	IsContract int    `json:"isContract"`
}
