package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalanceFromRPC(rpcURL, address string) (string, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get balance: %v", err)
	}

	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))

	return ethValue.String(), nil
}

func IsValidAddress(address string) bool {
	if len(address) != 42 {
		return false
	}
	if address[:2] != "0x" && address[:2] != "0X" {
		return false
	}

	for _, char := range address[2:] {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')) {
			return false
		}
	}

	return true
}

func ParsePrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key format: %v", err)
	}
	return privateKey, nil
}

func GetNonce(client *ethclient.Client, address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, fmt.Errorf("failed to get nonce: %v", err)
	}
	return nonce, nil
}

func EstimateGasLimit(client *ethclient.Client, from, to common.Address, amount *big.Int) (uint64, error) {
	msg := ethereum.CallMsg{
		From:  from,
		To:    &to,
		Value: amount,
	}

	gas, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0, fmt.Errorf("failed to estimate gas: %v", err)
	}

	return gas, nil
}

func GetGasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %v", err)
	}

	return gasPrice, nil
}

func SignTransaction(tx *types.Transaction, privateKey *ecdsa.PrivateKey, chainID *big.Int) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}
	return signedTx, nil
}

func EtherToWei(amount string) (*big.Int, error) {
	etherFloat, ok := new(big.Float).SetString(amount)
	if !ok {
		return nil, fmt.Errorf("invalid amount format")
	}

	weiFloat := new(big.Float).Mul(etherFloat, big.NewFloat(1e18))
	wei := new(big.Int)
	weiFloat.Int(wei)

	return wei, nil
}

func WeiToEther(wei *big.Int) string {
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return ethValue.Text('f', 18)
}
