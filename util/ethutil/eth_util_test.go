/**
 @author sunlight
 @date 18:16 2020/11/13
**/

//
package ethutil

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

func TestToWei(t *testing.T) {
	t.Parallel()
	amount := decimal.NewFromFloat(0.02)
	got := ToWei(amount, 38)
	expected := new(big.Int)
	expected.SetString("2000000000000000000000000000000000000", 10)
	if got.Cmp(expected) != 0 {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestToDecimal(t *testing.T) {
	t.Parallel()
	weiAmount := big.NewInt(0)
	weiAmount.SetString("2000000000000000000000000000000000000", 10)
	ethAmount := ToDecimal(weiAmount, 38)
	f64, _ := ethAmount.Float64()
	expected := 0.02
	if f64 != expected {
		t.Errorf("%v does not equal expected %v", ethAmount, expected)
	}
}

func TestCalcGasLimit(t *testing.T) {
	t.Parallel()
	gasPrice := big.NewInt(0)
	gasPrice.SetString("2000000000", 10)
	gasLimit := uint64(21000)
	expected := big.NewInt(0)
	expected.SetString("42000000000000", 10)
	gasCost := CalcGasCost(gasLimit, gasPrice)
	if gasCost.Cmp(expected) != 0 {
		t.Errorf("expected %s, got %s", gasCost, expected)
	}
}

func TestSigRSV(t *testing.T) {
	t.Parallel()

	sig := "0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301"
	r, s, v := SigRSV(sig)
	expectedR := "789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c6"
	expectedS := "2621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde023"
	expectedV := uint8(28)
	if hexutil.Encode(r[:])[2:] != expectedR {
		t.FailNow()
	}
	if hexutil.Encode(s[:])[2:] != expectedS {
		t.FailNow()
	}
	if v != expectedV {
		t.FailNow()
	}

	fmt.Println(ToWei(0.0001, 18).String())
}
