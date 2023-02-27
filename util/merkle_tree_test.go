package util

import (
	"awesomeProject/contract"
	"fmt"
	"github.com/Fueav/merkletree"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"testing"
)

func TestMerkleTree(t *testing.T) {
	var list []merkletree.Content
	//list = append(list, UserInfo{Address: TailorPrefix("0xF6509DB198aCc08D14bBDe21AC9eA952af89eAbB")})
	//list = append(list, UserInfo{Address: TailorPrefix("0x65a22c5Ef3bc163f5201B5e7C9554b2ff8E52512")})
	//list = append(list, UserInfo{Address: TailorPrefix("0xC0806ABF42d45956775590f18a20fF40083455A5")})
	//list = append(list, UserInfo{Address: TailorPrefix("0x82233Dd347DF3e2109eBB8Deb2B3934304516F7b")})
	//list = append(list, UserInfo{Address: TailorPrefix("0x285640d63902B4F638c8A4671dd467b461Bab10e")})
	//list = append(list, UserInfo{Address: TailorPrefix("0xeb1A88d3E2c3006952Fd2Afc5B16E3b646729990")})
	//list = append(list, UserInfo{Address: TailorPrefix("0x978e7ccf82a6c89804220C4FCe31238fd8Fd5d1e")})
	//list = append(list, UserInfo{Address: TailorPrefix("0xd355Ffc771B463eA8ac5F94bCC55924F72ed4092")})
	//list = append(list, UserInfo{Address: TailorPrefix("0x593227907587CFE2Cf5712929643082D0710E0Cc")})
	//list = append(list, UserInfo{Address: TailorPrefix("0xEfdFea8A8bCe5D622611f7B2e17a3EE4b8270147")})
	list = append(list, UserInfo{Address: TailorPrefix("0xAD71563312086f1139BEaEbD8C0e45E247f7E5F1")})
	list = append(list, UserInfo{Address: TailorPrefix("0x62d1003725582817dF4ad219EEa16937802B0A92")})
	list = append(list, UserInfo{Address: TailorPrefix("0xf45FeB68f49b3A24D6a387F110640EA301175208")})
	list = append(list, UserInfo{Address: TailorPrefix("0x24b21F4812b39CDE9D0b6F8D4aC3796f275c8ab6")})
	list = append(list, UserInfo{Address: TailorPrefix("0xD2F9Afd7360dc05eA7f7ef568e2C722Ac654c92F")})
	list = append(list, UserInfo{Address: TailorPrefix("0x7F235614Ece2b136E595Baf8905c83e423f38331")})

	//Create a new Merkle Tree from the list of Content
	tree, err := merkletree.NewTreeWithHashStrategy(list, sha3.NewLegacyKeccak256)
	if err != nil {
		log.Fatal(err)
	}

	//Get the Merkle Root of the tree
	proofRoot := tree.MerkleRoot()

	fmt.Println(common.Bytes2Hex(proofRoot))

	//Verify the entire tree (hashes for each node) is valid
	vt, err := tree.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Verify Tree: ", vt)

	//Verify a specific content in in the tree
	vc, err := tree.VerifyContent(list[0])
	if err != nil {
		log.Fatal(err)
	}

	proofPath, _, err := tree.GetMerklePath(UserInfo{Address: TailorPrefix("0xAD71563312086f1139BEaEbD8C0e45E247f7E5F1")})
	if err != nil {
		return
	}

	fmt.Println("Verify Content: ", vc)

	for i := range tree.Leafs {
		fmt.Println(common.BytesToHash(tree.Leafs[i].Hash).String())
	}

	client, err := ethclient.Dial("https://nd-513-001-835.p2pify.com/16ead84fd5bde769914b64a03ecde444")
	if err != nil {
		fmt.Println(err)
		return
	}

	nftDistribution, err := contract.NewNFTDistribution(common.HexToAddress("0xe92402DF34259a0E2C167E0186aff14E95915f7C"), client)
	if err != nil {
		fmt.Println(nftDistribution)
		return
	}

	var root [32]byte
	copy(root[:], proofRoot[:32])

	path := make([][32]byte, len(proofPath))
	for i := 0; i < len(proofPath); i++ {
		copy(path[i][:], proofPath[i][:32])
	}

	fmt.Println()
	for _, bytes := range proofPath {
		fmt.Println(common.BytesToHash(bytes).String())
	}

	fmt.Println()
	for i := range path {
		fmt.Println(path[i])
	}

	fmt.Println(root)
	fmt.Println(proofRoot)

	tx, err := nftDistribution.VerifyProof(nil, common.HexToAddress("0xAD71563312086f1139BEaEbD8C0e45E247f7E5F1"), root, path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tx)
}
