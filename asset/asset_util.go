package asset

import (
	"math"
	"math/big"
)

func GetBalance(amount *big.Int, decimal uint8) *big.Float {
	factor := math.Pow10(int(decimal))
	amountF := new(big.Float).SetInt(amount)
	return new(big.Float).Quo(amountF, big.NewFloat(factor))
}
