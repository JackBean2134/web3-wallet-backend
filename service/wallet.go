package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"web3-wallet-backend/config"
	"web3-wallet-backend/model"
	"web3-wallet-backend/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var cfg *config.Config

func InitService() {
	cfg = config.LoadConfig()
}

func CreateWallet() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate private key: %v", err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address, privateKeyHex, nil
}

func GetBalance(address string) (string, error) {
	if !utils.IsValidAddress(address) {
		return "", fmt.Errorf("invalid Ethereum address format")
	}

	balance, err := utils.GetBalanceFromRPC(cfg.RPCURL, address)
	if err != nil {
		return "", fmt.Errorf("failed to get balance: %v", err)
	}

	return balance, nil
}

func TransferETH(req model.TransferRequest) (*model.TransferResponse, error) {
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()

	fromAddress := common.HexToAddress(req.FromAddress)
	toAddress := common.HexToAddress(req.ToAddress)

	privateKey, err := utils.ParsePrivateKey(req.PrivateKey)
	if err != nil {
		return nil, err
	}

	value, err := utils.EtherToWei(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("invalid amount: %v", err)
	}

	nonce, err := utils.GetNonce(client, fromAddress)
	if err != nil {
		return nil, err
	}

	var gasPrice *big.Int
	if req.GasPrice != "" {
		gasPrice, _ = new(big.Int).SetString(req.GasPrice, 10)
	} else {
		gasPrice, err = utils.GetGasPrice(client)
		if err != nil {
			return nil, err
		}
	}

	var gasLimit uint64
	if req.GasLimit > 0 {
		gasLimit = req.GasLimit
	} else {
		gasLimit, err = utils.EstimateGasLimit(client, fromAddress, toAddress, value)
		if err != nil {
			gasLimit = 21000
		}
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	signedTx, err := utils.SignTransaction(tx, privateKey, chainID)
	if err != nil {
		return nil, err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	response := &model.TransferResponse{
		TxHash: signedTx.Hash().Hex(),
		From:   req.FromAddress,
		To:     req.ToAddress,
		Amount: req.Amount,
		Status: "pending",
	}

	return response, nil
}

func GetTransactionStatus(txHash string) (*model.TransactionStatus, error) {
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()

	hash := common.HexToHash(txHash)

	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, fmt.Errorf("transaction not found or still pending: %v", err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get current block: %v", err)
	}

	status := "failed"
	if receipt.Status == 1 {
		status = "success"
	}

	confirmations := header.Number.Uint64() - receipt.BlockNumber.Uint64()

	return &model.TransactionStatus{
		TxHash:        txHash,
		Status:        status,
		BlockNumber:   receipt.BlockNumber.Uint64(),
		GasUsed:       receipt.GasUsed,
		Confirmations: confirmations,
	}, nil
}
