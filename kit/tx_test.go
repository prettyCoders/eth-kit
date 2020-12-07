/**
 @author sunlight
 @date 10:43 2020/12/2
**/

//
package kit

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prettyCoders/eth-kit/test"
	"github.com/prettyCoders/eth-kit/util/ethutil"
	"github.com/prettyCoders/eth-kit/util/testutil"
	"log"
	"testing"
)

func TestQueryChainID(t *testing.T) {
	InitTestClient()
	chainID, err := QueryChainID()
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, chainID)
}

func TestQueryTxMessage(t *testing.T) {
	InitTestClient()
	chainID, err := QueryChainID()
	testutil.AssertNil(t, err)
	block, err := QueryBlockByHeight(9180007)
	testutil.AssertNil(t, err)
	for _, tx := range block.Transactions() {
		txMessage, err := QueryTxMessage(chainID, tx)
		testutil.AssertNil(t, err)
		testutil.AssertNotNil(t, txMessage.From())
	}
}

func TestQueryTxReceipt(t *testing.T) {
	InitTestClient()
	block, err := QueryBlockByHeight(9180007)
	testutil.AssertNil(t, err)
	for _, tx := range block.Transactions() {
		txReceipt, err := QueryTxReceipt(tx.Hash().Hex())
		testutil.AssertNil(t, err)
		testutil.AssertNotNil(t, txReceipt.Status)
		testutil.AssertNotNil(t, txReceipt.Logs)
	}
}

func TestQueryTxByHash(t *testing.T) {
	err := InitEthClient(test.InfuraRpcMainNet, test.InfuraWsMainNet)
	if err != nil {
		panic(err)
	}
	tx, isPending, err := QueryTxByHash("0xcfeaab6132d549f2d8d203dd5702f91472e463660657efd174949cd205ddf025")
	fmt.Println(tx.To().Hex())
	testutil.AssertNil(t, err)
	testutil.AssertFalse(t, isPending)
	testutil.AssertNotNil(t, tx)
}

func TestSendTx(t *testing.T) {
	InitTestClient()
	//转ETH
	senderAddress := test.Address
	privateKey := test.PrivateKey
	nonce, err := QueryAddressNonce(senderAddress)
	testutil.AssertNil(t, err)
	gasPrice, err := QuerySuggestGasPrice()
	testutil.AssertNil(t, err)
	value := ethutil.ToWei(0.0001, 18)
	txData := TxData{
		PrivateKeyHex: privateKey,
		Nonce:         nonce,
		Value:         value,
		GasLimit:      21000,
		GasPrice:      gasPrice,
		ToAddress:     "0x52A75088d21ceE6DD8ce7A9ef4448395fDD345e2",
		Data:          nil,
	}
	hash, err := SendTx(&txData)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, hash)
}

func TestBuildRawTransactionHex(t *testing.T) {
	InitTestClient()
	senderAddress := test.Address
	privateKey := test.PrivateKey
	nonce, err := QueryAddressNonce(senderAddress)
	testutil.AssertNil(t, err)
	gasPrice, err := QuerySuggestGasPrice()
	testutil.AssertNil(t, err)
	value := ethutil.ToWei(0.0001, 18)
	txData := TxData{
		PrivateKeyHex: privateKey,
		Nonce:         nonce,
		Value:         value,
		GasLimit:      21000,
		GasPrice:      gasPrice,
		ToAddress:     "0x52A75088d21ceE6DD8ce7A9ef4448395fDD345e2",
		Data:          nil,
	}
	hex, err := BuildRawTransactionHex(&txData)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, hex)
}

func TestSendRawTransaction(t *testing.T) {
	InitTestClient()
	rawTx := "f8690384b2d05e008252089452a75088d21cee6dd8ce7a9ef4448395fdd345e2865af3107a4000802aa09ef32cc0e6197a837ee5dacad9e627237f392d597a5c76aa1feb769b3d8d0536a058d600060e3381fc63c04412c3d68237cb8dc3d13446e2c7f83799f8b678c28d"
	hash, err := SendRawTransaction(rawTx)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, hash)
}

func TestSubContractEvent(t *testing.T) {
	err := InitEthClient(test.InfuraRpcMainNet, test.InfuraWsMainNet)
	if err != nil {
		panic(err)
	}
	var addresses []string
	addresses = append(addresses, "0xdac17f958d2ee523a2206206994597c13d831ec7")
	addresses = append(addresses, "0x514910771af9ca656af840dff83e8264ecf986ca")
	eventChan, sub, err := SubContractEvent(addresses)
	testutil.AssertNil(t, err)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case event := <-eventChan:
			topics := event.Topics
			eventNameHash := topics[0]
			if eventNameHash.Hex() == crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")).Hex() {
				for _, topic := range topics {
					fmt.Println(topic.Hex())
				}
			}
		}
	}
}
