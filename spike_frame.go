package main

import (
	"awesomeProject/request"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"sync"
)

func TestBatchWithdrawNFT(reqNumstart int, reqNumend int, contractAddress string) {
	var wg sync.WaitGroup
	for i := reqNumstart; i < reqNumend; i++ {
		wg.Add(1)
		go func(orderId int) {
			defer wg.Done()
			req := request.BatchWithdrawalNFTService{OrderId: uuid.New().String(), TokenID: int64(orderId), ToAddress: "0xE88a42c47928818E5775fb3cd076792353bE938b", Cb: "111", ContractAddress: contractAddress}
			Post(req, "http://localhost:3000/tx-api/v1/hotWallet/withdrawNFT")
		}(i)
	}
	wg.Wait()
}

func TestBatchWithdrawToken(reqNum int, contractAddress string) {
	var wg sync.WaitGroup
	for i := 0; i < reqNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req := request.BatchWithdrawalTokenService{OrderId: uuid.New().String(), Amount: "1.189", ToAddress: "0xE88a42c47928818E5775fb3cd076792353bE938b", Cb: "111", ContractAddress: contractAddress}
			Post(req, "http://localhost:3000/tx-api/v1/hotWallet/withdrawToken")
		}()
	}
	wg.Wait()
}

func TestBatchMint(reqNum int) {
	var wg sync.WaitGroup

	for i := 0; i < reqNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req := request.BatchMintNFTService{OrderId: uuid.New().String(), TokenURI: "http://oasis-server.oss-cn-shanghai.aliyuncs.com/json", Cb: "111", ContractAddress: 0}
			//Post(req, "http://119.13.91.31:3000/tx-api/v1/hotWallet/mint")
			Post(req, "http://localhost:3000/tx-api/v1/hotWallet/mint")
		}()
	}
	wg.Wait()
}

func Post(body interface{}, url string) {
	client := resty.New()
	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(url)
	if err != nil {
		fmt.Println(err)
	}
}
