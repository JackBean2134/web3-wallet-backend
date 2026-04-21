package main

import (
	"log"
	"web3-wallet-backend/router"
	"web3-wallet-backend/service"
)

func main() {
	service.InitService()

	r := router.SetupRouter()

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
