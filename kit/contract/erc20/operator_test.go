/**
 @author sunlight
 @date 18:54 2020/12/1
**/

//
package erc20

import (
	"eth-kit/kit"
	"eth-kit/test"
	"eth-kit/util/testutil"
	"testing"
)

func TestBalanceOf(t *testing.T) {
	InitTestClient()
	token, err := LoadErc20(test.DeployedContractAddress)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, token)
	balance, err := BalanceOf(test.Address, token)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, balance)
}

func TestTransfer(t *testing.T) {
	InitTestClient()
	token, err := LoadErc20(test.DeployedContractAddress)
	testutil.AssertNil(t, err)
	senderAddress := test.Address
	privateKey := test.PrivateKey
	nonce, err := kit.QueryAddressNonce(senderAddress)
	testutil.AssertNil(t, err)
	gasPrice, err := kit.QuerySuggestGasPrice()
	testutil.AssertNil(t, err)
	operateBaseData := OperateBaseData{
		Nonce:      nonce,
		EtherValue: 0,
		PrivateKey: privateKey,
		GasLimit:   60000,
		GasPrice:   gasPrice,
	}
	transferData := TransferData{OperateBaseData: &operateBaseData, Value: 100, ToAddress: "0x52A75088d21ceE6DD8ce7A9ef4448395fDD345e2"}
	hash, err := Transfer(&transferData, token)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, hash)
}
