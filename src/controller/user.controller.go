package controller

import (
	"encoding/json"
	"model"
	"net/http"
	"service"

	"github.com/gorilla/mux"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	address := params["address"]
	balance := service.GetBalance(address)
	response := model.Balance{address, balance}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func GetNonce(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	address := params["address"]
	nonce := service.GetNonce(address)
	response := model.Nonce{address, nonce}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	address := params["address"]
	balance := service.GetBalance(address)
	nonce := service.GetNonce(address)
	response := model.Account{address, balance, nonce}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
