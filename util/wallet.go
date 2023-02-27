package util

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/shopspring/decimal"
	"github.com/tyler-smith/go-bip39"
	"math/big"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ValidateAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func NewMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return mnemonic, err
	}

	return mnemonic, nil
}

func DerivePrivateKeyWithNumber(mnemonic string, number int) (string, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return "", err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + strconv.Itoa(number))
	account, err := wallet.Derive(path, false)
	if err != nil {
		return "", err
	}

	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(crypto.FromECDSA(privateKey)), nil
}

func DerivePrivateKey(mnemonic string) (string, error) {
	return DerivePrivateKeyWithNumber(mnemonic, 0)
}

func GenerateAddress(private string) (string, error) {
	privateKey, err := crypto.HexToECDSA(private[2:])
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {

		return "", err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).String()

	return address, err
}

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max - min)
	return r + min
}

func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

func TailorPrefix(address string) string {
	if !strings.HasPrefix(address, "0x") {
		return address
	}
	return address[2:]
}
