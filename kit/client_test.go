/**
 @author sunlight
 @date 16:04 2020/12/1
**/

//
package kit

import (
	"eth-kit/test"
	"eth-kit/util/testutil"
	"testing"
)

func TestInitClient(t *testing.T) {
	err := InitEthClient(test.InfuraRpcMainNet, test.InfuraWsMainNet)
	testutil.AssertNil(t, err)
	err = InitEthClient("", "")
	testutil.AssertNil(t, err)
}
