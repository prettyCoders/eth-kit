/**
 @author sunlight
 @date 14:16 2020/10/30
**/

//
package erc20

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prettyCoders/eth-kit/kit"
	"math/big"
)

type DeployBaseData struct {
	PrivateKey string
	Nonce      uint64
	GasLimit   uint64
	GasPrice   *big.Int
}

type DeployData struct {
	DeployBaseData
	InitialSupply *big.Int
	TokenName     string
	TokenSymbol   string
	TokenDecimals uint8
}

func Deploy(deployData *DeployData) (*common.Address, *types.Transaction, *Erc20, error) {
	privateKey, err := crypto.HexToECDSA(deployData.PrivateKey)
	if err != nil {
		return nil, nil, nil, err
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(deployData.Nonce))
	auth.Value = big.NewInt(0)          // in wei
	auth.GasLimit = deployData.GasLimit // in units
	auth.GasPrice = deployData.GasPrice

	address, tx, instance, err := DeployErc20(
		auth,
		kit.RpcClient,
		deployData.InitialSupply,
		deployData.TokenName,
		deployData.TokenSymbol,
		deployData.TokenDecimals,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	return &address, tx, instance, nil
}
