/**
 @author sunlight
 @date 16:33 2020/10/30
**/

//
package erc20

import (
	"eth-kit/kit"
	"github.com/ethereum/go-ethereum/common"
)

//LoadErc20 加载ERC20合约
func LoadErc20(contractAddress string) (*Erc20, error) {
	instance, err := NewErc20(
		common.HexToAddress(contractAddress),
		kit.RpcClient,
	)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
