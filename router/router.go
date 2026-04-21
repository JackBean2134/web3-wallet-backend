package router

import (
	"web3-wallet-backend/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/wallet/create", controller.CreateWallet)
	r.GET("/wallet/balance", controller.GetBalance)
	r.POST("/wallet/transfer", controller.TransferETH)
	r.GET("/wallet/transaction/status", controller.GetTransactionStatus)

	return r
}
