package controller

import (
	"model"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {
	address := c.Param("address")
	balance := service.GetBalance(address)
	nonce := service.GetNonce(address)
	response := model.Account{address, balance, nonce}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": response})
}

func GetBalance(c *gin.Context) {
	address := c.Param("address")
	balance := service.GetBalance(address)
	response := model.Balance{address, balance}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": response})
}
func GetNonce(c *gin.Context) {
	address := c.Param("address")
	nonce := service.GetNonce(address)
	response := model.Nonce{address, nonce}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": response})
}
