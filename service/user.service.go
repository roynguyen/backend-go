package service

import (
	"math/big"

	"github.com/roynguyen/backend-go/config"

	Web3 "github.com/regcostajr/go-web3"
	Providers "github.com/regcostajr/go-web3/providers"
)

var provider = Providers.NewHTTPProvider(
	config.DefaultConfig.URL,
	config.DefaultConfig.TIMEOUT,
	config.DefaultConfig.HTTPS,
)
var web3 = Web3.NewWeb3(provider)

func GetBalance(address string) *big.Int {
	res, err := web3.Eth.GetBalance(address, "latest")
	if err != nil {
		return big.NewInt(0)
	}
	return res
}
func GetNonce(address string) *big.Int {
	res, err := web3.Eth.GetTransactionCount(address, "latest")
	if err != nil {
		return big.NewInt(0)
	}
	return res
}
