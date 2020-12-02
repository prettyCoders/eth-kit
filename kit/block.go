/**
 @author sunlight
 @date 10:28 2020/12/2
**/

//
package kit

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

//QueryLatestBlockHeight 查询最新区块高度
func QueryLatestBlockHeight() (uint64, error) {
	header, err := RpcClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

//QueryBlockByHeight 根据区块高度查询区块
func QueryBlockByHeight(height uint64) (*types.Block, error) {
	return RpcClient.BlockByNumber(context.Background(), big.NewInt(int64(height)))
}

//SubNewBlockHeader 订阅新区块头
func SubNewBlockHeader() (chan *types.Header, ethereum.Subscription, error) {
	//创建通道接收新区块的header
	headers := make(chan *types.Header)
	//订阅
	sub, err := WsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		return nil, nil, err
	}
	return headers, sub, nil
}
