/**
 @author sunlight
 @date 16:55 2020/10/30
**/

//
package erc20

import (
	"eth-kit/test"
	"eth-kit/util/testutil"
	"testing"
)

func TestLoadErc20(t *testing.T) {
	erc20, err := LoadErc20(test.DeployedContractAddress)
	testutil.AssertNil(t, err)
	_ = erc20
}
