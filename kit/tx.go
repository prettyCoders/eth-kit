/**
 @author sunlight
 @date 10:28 2020/12/2
**/

//
package kit

import (
	"bytes"
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type TxData struct {
	PrivateKeyHex string
	ToAddress     string
	Nonce         uint64
	Value         *big.Int
	GasLimit      uint64
	GasPrice      *big.Int
	Data          []byte
}

//QueryChainID 查询链ID
func QueryChainID() (*big.Int, error) {
	if chainID, err := RpcClient.ChainID(context.Background()); err != nil {
		return nil, err
	} else {
		return chainID, nil
	}
}

//QueryTxMessage 查询交易数据，BlockByNumber里面的交易列表没有from
func QueryTxMessage(chainID *big.Int, tx *types.Transaction) (*types.Message, error) {
	if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil); err != nil {
		return nil, err
	} else {
		return &msg, nil
	}
}

//QueryTxReceipt 查询交易收据，可用于验证交易是否成功
func QueryTxReceipt(hash string) (*types.Receipt, error) {
	return RpcClient.TransactionReceipt(context.Background(), common.HexToHash(hash))
}

//QueryTxByHash 根据hash查询交易
func QueryTxByHash(hash string) (tx *types.Transaction, isPending bool, err error) {
	return RpcClient.TransactionByHash(context.Background(), common.HexToHash(hash))
}

//SendTx 发送交易
func SendTx(txData *TxData) (string, error) {
	tx := buildTx(txData)
	err := signAndBroadcast(txData.PrivateKeyHex, tx)
	if err != nil {
		return "", err
	} else {
		return tx.Hash().Hex(), nil
	}
}

//buildTx 构建交易
func buildTx(txData *TxData) *types.Transaction {
	return types.NewTransaction(
		txData.Nonce,
		common.HexToAddress(txData.ToAddress),
		txData.Value,
		txData.GasLimit,
		txData.GasPrice,
		txData.Data,
	)
}

//signTransaction 签名交易
func signTransaction(privateKeyHex string, tx *types.Transaction) (*types.Transaction, error) {
	//将十六进制私钥转换成ECDSA私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	chainID, err := RpcClient.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	return types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
}

//signAndBroadcast 签名并广播
func signAndBroadcast(privateKeyHex string, tx *types.Transaction) error {
	signedTx, err := signTransaction(privateKeyHex, tx)
	if err != nil {
		return err
	}
	return RpcClient.SendTransaction(context.Background(), signedTx)
}

//BuildRawTransactionHex 构建原始交易，可以直接广播的十六进制字符串
func BuildRawTransactionHex(txData *TxData) (string, error) {
	tx := buildTx(txData)
	signedTx, err := signTransaction(txData.PrivateKeyHex, tx)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	err = signedTx.EncodeRLP(&buffer)
	if err != nil {
		return "", err
	}
	rawTxHex := hex.EncodeToString(buffer.Bytes())
	return rawTxHex, nil
}

//SendRawTransaction 广播原始交易
func SendRawTransaction(rawTx string) (string, error) {
	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		return "", err
	}
	tx := new(types.Transaction)
	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		return "", err
	}
	err = RpcClient.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", err
	} else {
		return tx.Hash().Hex(), nil
	}
}

//SubContractEvent 订阅合约事件，暂时此方法只订阅所有事件，过滤待测试
func SubContractEvent(contractAddress []string) (chan types.Log, ethereum.Subscription, error) {
	var addresses []common.Address
	for _, address := range contractAddress {
		addresses = append(addresses, common.HexToAddress(address))
	}
	query := ethereum.FilterQuery{
		Addresses: addresses,
	}
	logs := make(chan types.Log)

	sub, err := WsClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return nil, nil, err
	}
	return logs, sub, nil
}
