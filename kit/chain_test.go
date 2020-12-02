/**
 @author sunlight
 @date 10:51 2020/12/2
**/

//
package kit

import (
	"github.com/prettyCoders/eth-kit/util/testutil"
	"testing"
)

func TestQuerySuggestGasPrice(t *testing.T) {
	InitTestClient()
	gasPrice, err := QuerySuggestGasPrice()
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, gasPrice)
}
