/**
 @author sunlight
 @date 16:48 2020/12/1
**/

//
package kit

import (
	"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prettyCoders/eth-kit/test"
	"github.com/prettyCoders/eth-kit/util/testutil"
	"testing"
)

func TestNewMnemonic(t *testing.T) {
	mnemonic, err := NewMnemonic(256)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, mnemonic)
}

func TestDeriveAccount(t *testing.T) {
	address, privateKey, err := DeriveAccount(test.Mnemonic, test.DerivationPath)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, address, test.Address)
	testutil.AssertEqual(t, privateKey, test.PrivateKey)
}

func InitTestClient() {
	err := InitEthClient(test.InfuraRpcRopsten, test.InfuraWsRopsten)
	if err != nil {
		panic(err)
	}
}

func TestQueryEtherBalance(t *testing.T) {
	InitTestClient()
	balance, err := QueryEtherBalance(test.Address)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, balance)
}

func TestQueryEtherPendingBalance(t *testing.T) {
	InitTestClient()
	balance, err := QueryEtherPendingBalance(test.Address)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, balance)
}

func TestQueryAddressNonce(t *testing.T) {
	InitTestClient()
	nonce, err := QueryAddressNonce(test.Address)
	testutil.AssertNil(t, err)
	testutil.AssertNotNil(t, nonce)
}

func TestIsValidAddress(t *testing.T) {
	t.Parallel()
	validAddress := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	invalidAddress := "0xabc"
	invalidAddress2 := "323b5d4c32345ced77393b3530b1eed0f346429d"
	{
		got := IsValidAddress(validAddress)
		expected := true

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}

	{
		got := IsValidAddress(invalidAddress)
		expected := false

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}

	{
		got := IsValidAddress(invalidAddress2)
		expected := false

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}
}

func TestIsZeroAddress(t *testing.T) {
	t.Parallel()
	validAddress := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	zeroAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")

	{
		isZeroAddress := IsZeroAddress(validAddress)

		if isZeroAddress {
			t.Error("Expected to be false")
		}
	}

	{
		isZeroAddress := IsZeroAddress(zeroAddress)

		if !isZeroAddress {
			t.Error("Expected to be true")
		}
	}

	{
		isZeroAddress := IsZeroAddress(validAddress.Hex())

		if isZeroAddress {
			t.Error("Expected to be false")
		}
	}

	{
		isZeroAddress := IsZeroAddress(zeroAddress.Hex())

		if !isZeroAddress {
			t.Error("Expected to be true")
		}
	}
}

func TestIsContractAddress(t *testing.T) {
	InitTestClient()
	b, err := IsContractAddress(test.Address)
	testutil.AssertNil(t, err)
	testutil.AssertFalse(t, b)
}

func TestPrivateKeyToAddress(t *testing.T) {
	address, err := PrivateKeyToAddress(test.PrivateKey)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, address, test.Address)
}

func TestPrivateKeyToPublicKey(t *testing.T) {
	publicKey, err := PrivateKeyToPublicKey(test.PrivateKey)
	publicKeyBytes := hexutil.Encode(crypto.FromECDSAPub(publicKey))
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, publicKeyBytes, "0x044049a17765de66afdf539c91752a716c7acfa1280982d995245fd5138ef0a7374becfebfa0bdc92731dc8c44bc87f26f92671348484605672a9c821a6e0d5692")
}

func TestPublicKeyBytesToAddress(t *testing.T) {
	t.Parallel()
	{
		publicKeyBytes, _ := hex.DecodeString("049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05")
		got := PublicKeyBytesToAddress(publicKeyBytes).Hex()
		expected := "0x96216849c49358B10257cb55b28eA603c874b05E"

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}
}

func TestSignVerifyMessage(t *testing.T) {
	//sign
	message := []byte("hello")
	privateKeyHex := test.PrivateKey
	signature, err := SignMessage(message, privateKeyHex)
	testutil.AssertNil(t, err)
	signatureHex := hexutil.Encode(signature)
	testutil.AssertEqual(t, signatureHex, "0x41e2b1bc65c1e185b54cdb115501eba16a49f38bd9f982dd99daa89828897af77b67354ebf93fdb8c7c4e53af0f470fb98d7a14dcb314da48a145e0605870fea01")

	//verify
	publicKeyECDSA, err := PrivateKeyToPublicKey(privateKeyHex)
	testutil.AssertNil(t, err)
	verified, err := VerifySignature(message, signature, publicKeyECDSA)
	testutil.AssertNil(t, err)
	testutil.AssertTrue(t, verified)

	hash := crypto.Keccak256Hash(message)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	//SigToPub 签名验证
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	testutil.AssertNil(t, err)
	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches := bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	testutil.AssertTrue(t, matches)

	//go-ethereum 自带的签名验证方法
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery ID
	verified = crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	testutil.AssertTrue(t, verified)
}
