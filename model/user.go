package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Wallet struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	Address    string    `json:"address"`
	PrivateKey string    `json:"private_key,omitempty"`
	Balance    string    `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateWalletRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}

type WalletResponse struct {
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
}

type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type TransferRequest struct {
	FromAddress string `json:"from_address" binding:"required"`
	PrivateKey  string `json:"private_key" binding:"required"`
	ToAddress   string `json:"to_address" binding:"required"`
	Amount      string `json:"amount" binding:"required"`
	GasLimit    uint64 `json:"gas_limit"`
	GasPrice    string `json:"gas_price"`
}

type TransferResponse struct {
	TxHash      string `json:"tx_hash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Amount      string `json:"amount"`
	GasUsed     uint64 `json:"gas_used,omitempty"`
	Status      string `json:"status"`
	BlockNumber uint64 `json:"block_number,omitempty"`
}

type TransactionStatus struct {
	TxHash        string `json:"tx_hash"`
	Status        string `json:"status"`
	BlockNumber   uint64 `json:"block_number"`
	GasUsed       uint64 `json:"gas_used"`
	Confirmations uint64 `json:"confirmations"`
}
