/**
 @author sunlight
 @date 16:43 2020/10/30
**/

//
package erc20

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prettyCoders/eth-kit/util/ethutil"
	"math/big"
)

//操作合约基础数据
type OperateBaseData struct {
	PrivateKey string
	EtherValue float64
	Nonce      uint64
	GasLimit   uint64
	GasPrice   *big.Int
}

//转账数据
type TransferData struct {
	*OperateBaseData
	ToAddress string
	Value     float64
}

//BalanceOf 查询余额
func BalanceOf(address string, erc20 *Erc20) (float64, error) {
	value, err := erc20.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return 0, err
	}
	decimal, err := erc20.Decimals(nil)
	if err != nil {
		return 0, err
	}
	balance, _ := ethutil.ToDecimal(value, int(decimal)).Float64()
	return balance, nil
}

//Transfer 转账
func Transfer(transferData *TransferData, erc20 *Erc20) (string, error) {
	privateKey, err := crypto.HexToECDSA(transferData.PrivateKey)
	if err != nil {
		return "", err
	}
	etherValue := ethutil.ToWei(transferData.EtherValue, 18)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(transferData.Nonce))
	auth.Value = etherValue
	auth.GasLimit = transferData.GasLimit
	auth.GasPrice = transferData.GasPrice
	toAddress := common.HexToAddress(transferData.ToAddress)
	decimal, err := erc20.Decimals(nil)
	if err != nil {
		return "", err
	}
	value := ethutil.ToWei(transferData.Value, int(decimal))
	tx, err := erc20.Transfer(auth, toAddress, value)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}
