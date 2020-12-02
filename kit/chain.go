/**
 @author sunlight
 @date 18:29 2020/12/1
**/

//
package kit

import (
	"context"
	"math/big"
)

func QuerySuggestGasPrice() (*big.Int, error) {
	return RpcClient.SuggestGasPrice(context.Background())
}
