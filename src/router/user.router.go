package router

import (
	"controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() { //*mux.Router {
	/*router := mux.NewRouter()
	router.HandleFunc("/auth", authentication.CreateTokenEndpoint).Methods("POST")
	router.HandleFunc("/eth/{address}/balance", authentication.ValidateMiddleware(controller.GetBalance)).Methods("GET")
	router.HandleFunc("/eth/{address}/nonce", controller.GetNonce).Methods("GET")
	router.HandleFunc("/eth/{address}", controller.GetAccount).Methods("GET")
	return router*/
	router := gin.Default()
	v1 := router.Group("/eth")
	{
		v1.GET("/:address/balance", controller.GetBalance)
		v1.GET("/:address/nonce", controller.GetNonce)
		v1.GET("/:address", controller.GetAccount)
	}
	router.Run()
}
