package main

import (
	"awesomeProject/config"
	"awesomeProject/contract"
	"awesomeProject/global"
	"awesomeProject/trade"
	"awesomeProject/util"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	logger "github.com/ipfs/go-log"
	"math/big"
)

func main() {
	logger.SetLogLevel("*", "INFO")
	global.Viper = config.InitViper()

	_ESSENCE_TYPEHASH := crypto.Keccak256Hash([]byte("mint(address to,uint256 profileId,uint256 essenceId,uint256 nonce,uint256 deadline)"))

	wallets := trade.NewPuppetWallet()
	fmt.Println("------Address: ", wallets.PuppetWallets[0].From)

	ccmw, err := contract.NewType(common.HexToAddress("0xAf53c3101c4b57A3d48100832ab8d1732b58C64C"), wallets.BscClient)
	if err != nil {
		fmt.Println("1", err)
		panic(err)
	}

	vote, err := contract.NewVote(common.HexToAddress("0x2723522702093601e6360CAe665518C4f63e9dA6"), wallets.BscClient)
	if err != nil {
		fmt.Println("2", err)
		panic(err)
	}

	nonce, err := ccmw.GetNonce(nil, big.NewInt(89), wallets.PuppetWallets[0].From, big.NewInt(30))
	if err != nil {
		fmt.Println("3", err)
		panic(err)
	}
	fmt.Println("nonce : ", nonce)

	height, err := wallets.BscClient.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("3.5", err)
		panic(err)
	}
	block, err := wallets.BscClient.BlockByNumber(context.Background(), big.NewInt(int64(height)))
	if err != nil {
		fmt.Println("3.5", err)
		panic(err)
	}
	fmt.Println("time: ", height, "++", block.Time(), "++", int64(block.Time()))
	futureTime := block.Time()

	hashed := crypto.Keccak256Hash(
		_ESSENCE_TYPEHASH.Bytes(),
		common.LeftPadBytes(wallets.PuppetWallets[0].From.Bytes(), 32),
		common.LeftPadBytes(big.NewInt(89).Bytes(), 32),
		common.LeftPadBytes(big.NewInt(30).Bytes(), 32),
		common.LeftPadBytes(nonce.Bytes(), 32),
		common.LeftPadBytes(big.NewInt(int64(futureTime)).Bytes(), 32),
	)

	secondHashed := crypto.Keccak256Hash(
		[]byte("\x19\x01"),
		//[]byte("0xe073ecf470948495cc09cecc3753d6d130af47a09e38cb401fad6d688c00fa30"),
		common.HexToHash("0xe073ecf470948495cc09cecc3753d6d130af47a09e38cb401fad6d688c00fa30").Bytes(),
		hashed.Bytes(),
	)

	priKeyString, err := util.DerivePrivateKeyWithNumber(config.Cfg.PuppetWallet.Mnemonic, 0)
	if err != nil {
		fmt.Println("6", err)
		panic(err)
	}

	priKey, err := crypto.HexToECDSA(priKeyString[2:])
	if err != nil {
		fmt.Println("7", err)
		panic(err)
	}

	signature, err := crypto.Sign(secondHashed.Bytes(), priKey)
	if err != nil {
		fmt.Println("8", err)
		panic(err)
	}
	fmt.Println("sign: ", signature)
	r := signature[:32]
	s := signature[32:64]
	v := signature[len(signature)-1:]
	fmt.Println("sign1: ", r, len(r))
	fmt.Println("sign2: ", s)
	fmt.Println("sign3: ", v)

	var vrsPacked []byte
	vrsPacked = append(vrsPacked, common.LeftPadBytes(v, 32)...)
	//vrsPacked = append(vrsPacked, common.LeftPadBytes(big.NewInt(28).Bytes(), 32)...)
	vrsPacked = append(vrsPacked, r...)
	vrsPacked = append(vrsPacked, s...)
	vrsPacked = append(vrsPacked, common.LeftPadBytes(big.NewInt(int64(futureTime)).Bytes(), 32)...)

	fmt.Println("vrsPacked", vrsPacked)

	abi, err := contract.VoteMetaData.GetAbi()
	if err != nil {
		fmt.Println("9", err)
		return
	}

	input, err := abi.Pack("collect", contract.DataTypesCollectParams{
		Collector: wallets.PuppetWallets[0].From,
		ProfileId: big.NewInt(89),
		EssenceId: big.NewInt(30),
	}, vrsPacked, []byte{})

	fmt.Println("input:", common.Bytes2Hex(input))

	ecrecover, err := crypto.Ecrecover(secondHashed.Bytes(), signature)
	if err != nil {
		fmt.Println("10", err)
		return
	}

	fmt.Println("publickey", common.BytesToAddress(ecrecover))

	collect, err := vote.Collect(wallets.PuppetWallets[0], contract.DataTypesCollectParams{
		Collector: wallets.PuppetWallets[0].From,
		ProfileId: big.NewInt(89),
		EssenceId: big.NewInt(30),
	}, vrsPacked, []byte{})
	if err != nil {
		fmt.Println("9", err)
		return
	}
	fmt.Println(collect.Hash().String())
}

type EIP712Signature struct {
	v        []byte
	r        []byte
	s        []byte
	deadline *big.Int
}
