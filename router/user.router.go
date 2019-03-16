package router

import (
	"github.com/roynguyen/backend-go/controller"
	authentication "github.com/roynguyen/backend-go/middleware"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/auth", authentication.CreateTokenEndpoint).Methods("POST")
	router.HandleFunc("/eth/{address}/balance", authentication.ValidateMiddleware(controller.GetBalance)).Methods("GET")
	router.HandleFunc("/eth/{address}/nonce", controller.GetNonce).Methods("GET")
	router.HandleFunc("/eth/{address}", controller.GetAccount).Methods("GET")
	return router
}
