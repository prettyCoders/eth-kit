/**
 @author sunlight
 @date 18:07 2020/10/30
**/

//
package contract

import (
	"context"
	"eth-kit/kit"
	"github.com/ethereum/go-ethereum/common"
)

//QueryByteCode 获取合约字节码
func QueryByteCode(contractAddress string) ([]byte, error) {
	return kit.RpcClient.CodeAt(
		context.Background(),
		common.HexToAddress(contractAddress),
		nil, // nil is latest block
	)
}
