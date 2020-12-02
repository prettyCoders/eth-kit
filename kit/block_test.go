/**
 @author sunlight
 @date 10:33 2020/12/2
**/

//
package kit

import (
	"fmt"
	"github.com/prettyCoders/eth-kit/util/testutil"
	"log"
	"testing"
)

func TestQueryLatestBlockHeight(t *testing.T) {
	InitTestClient()
	height, err := QueryLatestBlockHeight()
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, height)
	fmt.Println(height)
}

func TestQueryBlockByHeight(t *testing.T) {
	InitTestClient()
	block, err := QueryBlockByHeight(9180007)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, block.Hash().Hex(), "0x9654539890470f8dec28d9f7c58ba45d1e8d0df2a1dc2eae33d8252de51b7074")
}

func TestSubNewBlockHeader(t *testing.T) {
	InitTestClient()
	headerChan, sub, err := SubNewBlockHeader()
	testutil.AssertNil(t, err)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headerChan:
			fmt.Println(header.Hash().Hex())
		}
	}
}
