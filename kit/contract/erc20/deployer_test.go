/**
 @author sunlight
 @date 15:10 2020/10/30
**/

//
package erc20

import (
	"eth-kit/kit"
	"eth-kit/test"
	"eth-kit/util/testutil"
	"fmt"
	"math/big"
	"testing"
)

func InitTestClient() {
	err := kit.InitEthClient(test.InfuraRpcRopsten, test.InfuraWsRopsten)
	if err != nil {
		panic(err)
	}
}

func TestDeployErc20(t *testing.T) {
	InitTestClient()
	privateKeyHex := test.PrivateKey
	fromAddress, err := kit.PrivateKeyToAddress(privateKeyHex)
	testutil.AssertNil(t, err)
	nonce, err := kit.QueryAddressNonce(fromAddress)
	testutil.AssertNil(t, err)
	gasLimit := uint64(1800000)
	gasPrice, err := kit.QuerySuggestGasPrice()
	testutil.AssertNil(t, err)
	erc20DeployData := DeployData{
		DeployBaseData: DeployBaseData{
			PrivateKey: privateKeyHex,
			Nonce:      nonce,
			GasLimit:   gasLimit,
			GasPrice:   gasPrice,
		},
		InitialSupply: big.NewInt(1000000),
		TokenName:     "Filecoin6Month",
		TokenSymbol:   "FIL6",
		TokenDecimals: 8,
	}
	address, tx, instance, err := Deploy(&erc20DeployData)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, address)
	testutil.AssertNotNil(t, tx)
	testutil.AssertNotNil(t, instance)
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
}
