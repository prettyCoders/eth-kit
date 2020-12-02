/**
 @author sunlight
 @date 18:09 2020/10/30
**/

//
package contract

import (
	"github.com/prettyCoders/eth-kit/kit"
	"github.com/prettyCoders/eth-kit/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

const Erc20ContractAddress = "0x30F04061CBbD7adc5C4F43049D4D9Af8A574f8ad"

func InitTestClient() {
	err := kit.InitEthClient(test.InfuraRpcRopsten, test.InfuraWsRopsten)
	if err != nil {
		panic(err)
	}
}

//TestGetByteCode 读取已部署的智能合约的字节码
func TestGetByteCode(t *testing.T) {
	InitTestClient()
	byteCode, err := QueryByteCode(Erc20ContractAddress)
	assert.Nil(t, err)
	assert.NotNil(t, byteCode)
}
