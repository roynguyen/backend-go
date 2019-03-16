package model

import "math/big"

type Account struct {
	Address string   `json:"address,omitempty"`
	Balance *big.Int `json:"balance,omitempty"`
	Nonce   *big.Int `json:"nonce,omitempty"`
}

type Balance struct {
	Address string   `json:"address,omitempty"`
	Balance *big.Int `json:"balance,omitempty"`
}

type Nonce struct {
	Address string   `json:"address,omitempty"`
	Nonce   *big.Int `json:"nonce,omitempty"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}
