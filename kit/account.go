/**
 @author sunlight
 @date 16:07 2020/12/1
**/

//
package kit

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"eth-kit/util/ethutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"golang.org/x/crypto/sha3"
	"golang.org/x/xerrors"
	"reflect"
	"regexp"
)

//NewMnemonic 生成新的助记词
//bits 表示使用多少位的熵，能被32整除，范围128-256
func NewMnemonic(bits uint16) (string, error) {
	return hdwallet.NewMnemonic(int(bits))
}

//DeriveAccount 派生账户
//mnemonic 助记词
//path 标准bip44路径，例如：m/44'/60'/0'/0/0
func DeriveAccount(mnemonic string, path string) (address string, privateKey string, err error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return "", "", err
	}
	derivationPath := hdwallet.MustParseDerivationPath(path)
	account, err := wallet.Derive(derivationPath, false)
	if err != nil {
		return "", "", err
	}
	privateKey, err = wallet.PrivateKeyHex(account)
	if err != nil {
		return "", "", err
	}
	return account.Address.Hex(), privateKey, nil
}

//QueryEtherBalance 查询地址ETH余额
func QueryEtherBalance(address string) (float64, error) {
	account := common.HexToAddress(address)
	balance, err := RpcClient.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return 0, err
	}
	value, _ := ethutil.ToDecimal(balance, EtherDecimal).Float64()
	return value, nil
}

//QueryEtherPendingBalance 查询地址ETH pending余额（影响账户余额的交易处于pending状态）
func QueryEtherPendingBalance(address string) (float64, error) {
	account := common.HexToAddress(address)
	balance, err := RpcClient.PendingBalanceAt(context.Background(), account)
	if err != nil {
		return 0, err
	}
	value, _ := ethutil.ToDecimal(balance, EtherDecimal).Float64()
	return value, nil
}

//QueryAddressNonce 查询地址当前nonce值
func QueryAddressNonce(address string) (uint64, error) {
	return RpcClient.PendingNonceAt(context.Background(), common.HexToAddress(address))
}

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

//IsContractAddress 判断地址是否是合约地址，若在该地址存储了字节码，该地址是智能合约
func IsContractAddress(address string) (bool, error) {
	bytecode, err := RpcClient.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return false, err
	}
	return len(bytecode) > 0, nil
}

//PrivateKeyToAddress 十六进制私钥转地址
func PrivateKeyToAddress(privateKeyHex string) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", xerrors.Errorf("Failed cast %t to *ecdsa.PublicKey", publicKey)
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address.Hex(), nil
}

//PrivateKeyToPublicKey 私钥推导公钥
func PrivateKeyToPublicKey(privateKeyHex string) (*ecdsa.PublicKey, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, xerrors.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return publicKeyECDSA, nil
}

// PublicKeyBytesToAddress 公钥推导地址
func PublicKeyBytesToAddress(publicKey []byte) common.Address {
	var buf []byte

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey[1:]) // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]

	return common.HexToAddress(hex.EncodeToString(address))
}

//SignMessage 签消息
func SignMessage(message []byte, privateKeyHex string) ([]byte, error) {
	if privateKey, err := crypto.HexToECDSA(privateKeyHex); err != nil {
		return nil, err
	} else {
		hash := crypto.Keccak256Hash(message)
		if signature, err := crypto.Sign(hash.Bytes(), privateKey); err != nil {
			return nil, err
		} else {
			return signature, nil
		}
	}
}

//VerifySignature 验签
func VerifySignature(message []byte, signature []byte, publicKeyECDSA *ecdsa.PublicKey) (bool, error) {
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash := crypto.Keccak256Hash(message)
	if sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature); err != nil {
		return false, err
	} else {
		matches := bytes.Equal(sigPublicKey, publicKeyBytes)
		return matches, nil
	}
}
