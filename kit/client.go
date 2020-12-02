/**
 @author sunlight
 @date 11:54 2020/12/1
**/

//
package kit

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	RpcClient *ethclient.Client //rpc客户端连接
	WsClient  *ethclient.Client //websocket客户端连接
)

//InitEthClient 初始化eth客户端，由于websocket连接用的少，不会严格要求调用者必须提供参数，只会在控制台打印警告信息
func InitEthClient(rpcRawUrl string, wsRawUrl string) error {
	if rpcRawUrl != "" {
		if client, err := ethclient.Dial(rpcRawUrl); err != nil {
			return err
		} else {
			RpcClient = client
		}
	} else {
		fmt.Println("some function maybe panic because rpc raw url is empty")
	}

	if wsRawUrl != "" {
		if client, err := ethclient.Dial(wsRawUrl); err != nil {
			return err
		} else {
			WsClient = client
		}
	} else {
		fmt.Println("some function maybe panic because websocket raw url is empty")
	}
	return nil
}
