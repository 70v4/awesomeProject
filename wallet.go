package main

import (
	"awesomeProject/util"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/rand"
	"time"
)

type WalletTree struct {
	Val   *bind.TransactOpts
	Child []*WalletTree
}

func (s WalletTree) String() string {
	return fmt.Sprintf("{val: %v, leftChild: %v, rightChild: %v}", s.Val)
}

func MultipleTree(dataLst []*bind.TransactOpts) *WalletTree {
	rand.Seed(time.Now().Unix())

	// 剔除第一个账户 原始账户
	data := make([]*bind.TransactOpts, len(dataLst))
	copy(data, dataLst)

	data = append(data[:0], data[1:]...)

	leftWallet := 10

	treeItem := make([]*WalletTree, 0)
	treeRoot := &WalletTree{Val: dataLst[0], Child: make([]*WalletTree, 0)}
	treeItem = append(treeItem, treeRoot)

	for {
		if leftWallet == 0 || len(data) == 0 {
			return treeRoot
		}

		childNum := util.Random(2, 5)
		for i := 0; i < childNum; i++ {
			if len(data) == 0 || len(treeItem) == 0 {
				return treeRoot
			}
			leftWallet--
			index := rand.Intn(len(data))
			child := &WalletTree{Val: data[index], Child: make([]*WalletTree, 0)}
			treeItem[0].Child = append(treeItem[0].Child, child)
			treeItem = append(treeItem, child)
			data = append(data[:index], data[index+1:]...)
		}

		treeItem = append(treeItem[:0], treeItem[1:]...)
	}
}

func PrintTree(s *WalletTree) {
	if s.Val == nil {
		return
	}
	fmt.Println("root node: ", s.Val.From)
	if len(s.Child) == 0 {
		return
	}
	for i := range s.Child {
		fmt.Print("child node: ", s.Child[i].Val.From, "-")
	}
	fmt.Println()
	if len(s.Child) != 0 {
		for i := range s.Child {
			PrintTree(s.Child[i])
		}
	}
}
