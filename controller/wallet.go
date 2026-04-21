package controller

import (
	"net/http"

	"web3-wallet-backend/model"
	"web3-wallet-backend/service"

	"github.com/gin-gonic/gin"
)

func CreateWallet(c *gin.Context) {
	var req model.CreateWalletRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	address, privateKey, err := service.CreateWallet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := model.WalletResponse{
		Address:    address,
		PrivateKey: privateKey,
	}

	c.JSON(http.StatusOK, response)
}

func GetBalance(c *gin.Context) {
	address := c.Query("address")

	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "address parameter is required"})
		return
	}

	balance, err := service.GetBalance(address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := model.BalanceResponse{
		Address: address,
		Balance: balance,
	}

	c.JSON(http.StatusOK, response)
}
