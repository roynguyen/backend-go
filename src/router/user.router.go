package router

import (
	"controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	v1 := router.Group("/eth")
	{
		v1.GET("/:address/balance", controller.GetBalance)
		v1.GET("/:address/nonce", controller.GetNonce)
		v1.GET("/:address", controller.GetAccount)
	}
	router.Run()
}
