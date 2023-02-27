// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DataTypesCollectParams is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCollectParams struct {
	Collector common.Address
	ProfileId *big.Int
	EssenceId *big.Int
}

// DataTypesCreateProfileParams is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreateProfileParams struct {
	To       common.Address
	Handle   string
	Avatar   string
	Metadata string
	Operator common.Address
}

// DataTypesEIP712Signature is an auto generated low-level Go binding around an user-defined struct.
type DataTypesEIP712Signature struct {
	V        uint8
	R        [32]byte
	S        [32]byte
	Deadline *big.Int
}

// DataTypesRegisterEssenceParams is an auto generated low-level Go binding around an user-defined struct.
type DataTypesRegisterEssenceParams struct {
	ProfileId        *big.Int
	Name             string
	Symbol           string
	EssenceTokenURI  string
	EssenceMw        common.Address
	Transferable     bool
	DeployAtRegister bool
}

// DataTypesSubscribeParams is an auto generated low-level Go binding around an user-defined struct.
type DataTypesSubscribeParams struct {
	ProfileIds []*big.Int
}

// VoteMetaData contains all meta data concerning the Vote contract.
var VoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"preData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"postData\",\"type\":\"bytes\"}],\"name\":\"CollectEssence\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"CreateProfile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"essenceNFT\",\"type\":\"address\"}],\"name\":\"DeployEssenceNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subscribeNFT\",\"type\":\"address\"}],\"name\":\"DeploySubscribeNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"essenceTokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"essenceMw\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"prepareReturnData\",\"type\":\"bytes\"}],\"name\":\"RegisterEssence\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newAvatar\",\"type\":\"string\"}],\"name\":\"SetAvatar\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"prepareReturnData\",\"type\":\"bytes\"}],\"name\":\"SetEssenceData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMetadata\",\"type\":\"string\"}],\"name\":\"SetMetadata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDescriptor\",\"type\":\"address\"}],\"name\":\"SetNFTDescriptor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"preOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SetNamespaceOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prevApproved\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"SetOperatorApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"SetPrimaryProfile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"prepareReturnData\",\"type\":\"bytes\"}],\"name\":\"SetSubscribeData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"preDatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"postDatas\",\"type\":\"bytes[]\"}],\"name\":\"Subscribe\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ENGINE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ESSENCE_BEACON\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBSCRIBE_BEACON\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.CollectParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"preData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"postData\",\"type\":\"bytes\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.CollectParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"preData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"postData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"collectWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.CreateProfileParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"preData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"postData\",\"type\":\"bytes\"}],\"name\":\"createProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getAvatar\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"name\":\"getEssenceMw\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"name\":\"getEssenceNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"name\":\"getEssenceNFTTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getHandleByProfileId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNFTDescriptor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNamespaceOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getPrimaryProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getProfileIdByHandle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getSubscribeMw\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getSubscribeNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getSubscribeNFTTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"toPause\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"essenceTokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"essenceMw\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"deployAtRegister\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.RegisterEssenceParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"registerEssence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"essenceTokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"essenceMw\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"deployAtRegister\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.RegisterEssenceParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"registerEssenceWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"}],\"name\":\"setAvatar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"setAvatarWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setEssenceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"setEssenceDataWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"setMetadataWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"descriptor\",\"type\":\"address\"}],\"name\":\"setNFTDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setNamespaceOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setOperatorApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"setPrimaryProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"setPrimaryProfileWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setSubscribeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mw\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"setSubscribeDataWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataTypes.SubscribeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"preDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"postDatas\",\"type\":\"bytes[]\"}],\"name\":\"subscribe\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataTypes.SubscribeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"preDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"postDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"subscribeWithSig\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBurned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// VoteABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteMetaData.ABI instead.
var VoteABI = VoteMetaData.ABI

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Vote *VoteCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Vote *VoteSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Vote.Contract.DOMAINSEPARATOR(&_Vote.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Vote *VoteCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Vote.Contract.DOMAINSEPARATOR(&_Vote.CallOpts)
}

// ENGINE is a free data retrieval call binding the contract method 0x4785e8d4.
//
// Solidity: function ENGINE() view returns(address)
func (_Vote *VoteCaller) ENGINE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "ENGINE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ENGINE is a free data retrieval call binding the contract method 0x4785e8d4.
//
// Solidity: function ENGINE() view returns(address)
func (_Vote *VoteSession) ENGINE() (common.Address, error) {
	return _Vote.Contract.ENGINE(&_Vote.CallOpts)
}

// ENGINE is a free data retrieval call binding the contract method 0x4785e8d4.
//
// Solidity: function ENGINE() view returns(address)
func (_Vote *VoteCallerSession) ENGINE() (common.Address, error) {
	return _Vote.Contract.ENGINE(&_Vote.CallOpts)
}

// ESSENCEBEACON is a free data retrieval call binding the contract method 0x1ec7d106.
//
// Solidity: function ESSENCE_BEACON() view returns(address)
func (_Vote *VoteCaller) ESSENCEBEACON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "ESSENCE_BEACON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ESSENCEBEACON is a free data retrieval call binding the contract method 0x1ec7d106.
//
// Solidity: function ESSENCE_BEACON() view returns(address)
func (_Vote *VoteSession) ESSENCEBEACON() (common.Address, error) {
	return _Vote.Contract.ESSENCEBEACON(&_Vote.CallOpts)
}

// ESSENCEBEACON is a free data retrieval call binding the contract method 0x1ec7d106.
//
// Solidity: function ESSENCE_BEACON() view returns(address)
func (_Vote *VoteCallerSession) ESSENCEBEACON() (common.Address, error) {
	return _Vote.Contract.ESSENCEBEACON(&_Vote.CallOpts)
}

// SUBSCRIBEBEACON is a free data retrieval call binding the contract method 0x5b996c0f.
//
// Solidity: function SUBSCRIBE_BEACON() view returns(address)
func (_Vote *VoteCaller) SUBSCRIBEBEACON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "SUBSCRIBE_BEACON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SUBSCRIBEBEACON is a free data retrieval call binding the contract method 0x5b996c0f.
//
// Solidity: function SUBSCRIBE_BEACON() view returns(address)
func (_Vote *VoteSession) SUBSCRIBEBEACON() (common.Address, error) {
	return _Vote.Contract.SUBSCRIBEBEACON(&_Vote.CallOpts)
}

// SUBSCRIBEBEACON is a free data retrieval call binding the contract method 0x5b996c0f.
//
// Solidity: function SUBSCRIBE_BEACON() view returns(address)
func (_Vote *VoteCallerSession) SUBSCRIBEBEACON() (common.Address, error) {
	return _Vote.Contract.SUBSCRIBEBEACON(&_Vote.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vote *VoteCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vote *VoteSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Vote.Contract.BalanceOf(&_Vote.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Vote *VoteCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Vote.Contract.BalanceOf(&_Vote.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Vote *VoteCaller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Vote *VoteSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Vote.Contract.GetApproved(&_Vote.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Vote *VoteCallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Vote.Contract.GetApproved(&_Vote.CallOpts, arg0)
}

// GetAvatar is a free data retrieval call binding the contract method 0x1328ec9b.
//
// Solidity: function getAvatar(uint256 profileId) view returns(string)
func (_Vote *VoteCaller) GetAvatar(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getAvatar", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAvatar is a free data retrieval call binding the contract method 0x1328ec9b.
//
// Solidity: function getAvatar(uint256 profileId) view returns(string)
func (_Vote *VoteSession) GetAvatar(profileId *big.Int) (string, error) {
	return _Vote.Contract.GetAvatar(&_Vote.CallOpts, profileId)
}

// GetAvatar is a free data retrieval call binding the contract method 0x1328ec9b.
//
// Solidity: function getAvatar(uint256 profileId) view returns(string)
func (_Vote *VoteCallerSession) GetAvatar(profileId *big.Int) (string, error) {
	return _Vote.Contract.GetAvatar(&_Vote.CallOpts, profileId)
}

// GetEssenceMw is a free data retrieval call binding the contract method 0x5e316de4.
//
// Solidity: function getEssenceMw(uint256 profileId, uint256 essenceId) view returns(address)
func (_Vote *VoteCaller) GetEssenceMw(opts *bind.CallOpts, profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getEssenceMw", profileId, essenceId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEssenceMw is a free data retrieval call binding the contract method 0x5e316de4.
//
// Solidity: function getEssenceMw(uint256 profileId, uint256 essenceId) view returns(address)
func (_Vote *VoteSession) GetEssenceMw(profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	return _Vote.Contract.GetEssenceMw(&_Vote.CallOpts, profileId, essenceId)
}

// GetEssenceMw is a free data retrieval call binding the contract method 0x5e316de4.
//
// Solidity: function getEssenceMw(uint256 profileId, uint256 essenceId) view returns(address)
func (_Vote *VoteCallerSession) GetEssenceMw(profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	return _Vote.Contract.GetEssenceMw(&_Vote.CallOpts, profileId, essenceId)
}

// GetEssenceNFT is a free data retrieval call binding the contract method 0xa92dc7e8.
//
// Solidity: function getEssenceNFT(uint256 profileId, uint256 essenceId) view returns(address)
func (_Vote *VoteCaller) GetEssenceNFT(opts *bind.CallOpts, profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getEssenceNFT", profileId, essenceId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEssenceNFT is a free data retrieval call binding the contract method 0xa92dc7e8.
//
// Solidity: function getEssenceNFT(uint256 profileId, uint256 essenceId) view returns(address)
func (_Vote *VoteSession) GetEssenceNFT(profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	return _Vote.Contract.GetEssenceNFT(&_Vote.CallOpts, profileId, essenceId)
}

// GetEssenceNFT is a free data retrieval call binding the contract method 0xa92dc7e8.
//
// Solidity: function getEssenceNFT(uint256 profileId, uint256 essenceId) view returns(address)
func (_Vote *VoteCallerSession) GetEssenceNFT(profileId *big.Int, essenceId *big.Int) (common.Address, error) {
	return _Vote.Contract.GetEssenceNFT(&_Vote.CallOpts, profileId, essenceId)
}

// GetEssenceNFTTokenURI is a free data retrieval call binding the contract method 0x206f8e86.
//
// Solidity: function getEssenceNFTTokenURI(uint256 profileId, uint256 essenceId) view returns(string)
func (_Vote *VoteCaller) GetEssenceNFTTokenURI(opts *bind.CallOpts, profileId *big.Int, essenceId *big.Int) (string, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getEssenceNFTTokenURI", profileId, essenceId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEssenceNFTTokenURI is a free data retrieval call binding the contract method 0x206f8e86.
//
// Solidity: function getEssenceNFTTokenURI(uint256 profileId, uint256 essenceId) view returns(string)
func (_Vote *VoteSession) GetEssenceNFTTokenURI(profileId *big.Int, essenceId *big.Int) (string, error) {
	return _Vote.Contract.GetEssenceNFTTokenURI(&_Vote.CallOpts, profileId, essenceId)
}

// GetEssenceNFTTokenURI is a free data retrieval call binding the contract method 0x206f8e86.
//
// Solidity: function getEssenceNFTTokenURI(uint256 profileId, uint256 essenceId) view returns(string)
func (_Vote *VoteCallerSession) GetEssenceNFTTokenURI(profileId *big.Int, essenceId *big.Int) (string, error) {
	return _Vote.Contract.GetEssenceNFTTokenURI(&_Vote.CallOpts, profileId, essenceId)
}

// GetHandleByProfileId is a free data retrieval call binding the contract method 0x2307c987.
//
// Solidity: function getHandleByProfileId(uint256 profileId) view returns(string)
func (_Vote *VoteCaller) GetHandleByProfileId(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getHandleByProfileId", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandleByProfileId is a free data retrieval call binding the contract method 0x2307c987.
//
// Solidity: function getHandleByProfileId(uint256 profileId) view returns(string)
func (_Vote *VoteSession) GetHandleByProfileId(profileId *big.Int) (string, error) {
	return _Vote.Contract.GetHandleByProfileId(&_Vote.CallOpts, profileId)
}

// GetHandleByProfileId is a free data retrieval call binding the contract method 0x2307c987.
//
// Solidity: function getHandleByProfileId(uint256 profileId) view returns(string)
func (_Vote *VoteCallerSession) GetHandleByProfileId(profileId *big.Int) (string, error) {
	return _Vote.Contract.GetHandleByProfileId(&_Vote.CallOpts, profileId)
}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 profileId) view returns(string)
func (_Vote *VoteCaller) GetMetadata(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getMetadata", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 profileId) view returns(string)
func (_Vote *VoteSession) GetMetadata(profileId *big.Int) (string, error) {
	return _Vote.Contract.GetMetadata(&_Vote.CallOpts, profileId)
}

// GetMetadata is a free data retrieval call binding the contract method 0xa574cea4.
//
// Solidity: function getMetadata(uint256 profileId) view returns(string)
func (_Vote *VoteCallerSession) GetMetadata(profileId *big.Int) (string, error) {
	return _Vote.Contract.GetMetadata(&_Vote.CallOpts, profileId)
}

// GetNFTDescriptor is a free data retrieval call binding the contract method 0x17c4d5d2.
//
// Solidity: function getNFTDescriptor() view returns(address)
func (_Vote *VoteCaller) GetNFTDescriptor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getNFTDescriptor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNFTDescriptor is a free data retrieval call binding the contract method 0x17c4d5d2.
//
// Solidity: function getNFTDescriptor() view returns(address)
func (_Vote *VoteSession) GetNFTDescriptor() (common.Address, error) {
	return _Vote.Contract.GetNFTDescriptor(&_Vote.CallOpts)
}

// GetNFTDescriptor is a free data retrieval call binding the contract method 0x17c4d5d2.
//
// Solidity: function getNFTDescriptor() view returns(address)
func (_Vote *VoteCallerSession) GetNFTDescriptor() (common.Address, error) {
	return _Vote.Contract.GetNFTDescriptor(&_Vote.CallOpts)
}

// GetNamespaceOwner is a free data retrieval call binding the contract method 0x42cfa815.
//
// Solidity: function getNamespaceOwner() view returns(address)
func (_Vote *VoteCaller) GetNamespaceOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getNamespaceOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNamespaceOwner is a free data retrieval call binding the contract method 0x42cfa815.
//
// Solidity: function getNamespaceOwner() view returns(address)
func (_Vote *VoteSession) GetNamespaceOwner() (common.Address, error) {
	return _Vote.Contract.GetNamespaceOwner(&_Vote.CallOpts)
}

// GetNamespaceOwner is a free data retrieval call binding the contract method 0x42cfa815.
//
// Solidity: function getNamespaceOwner() view returns(address)
func (_Vote *VoteCallerSession) GetNamespaceOwner() (common.Address, error) {
	return _Vote.Contract.GetNamespaceOwner(&_Vote.CallOpts)
}

// GetOperatorApproval is a free data retrieval call binding the contract method 0xf0ab9600.
//
// Solidity: function getOperatorApproval(uint256 profileId, address operator) view returns(bool)
func (_Vote *VoteCaller) GetOperatorApproval(opts *bind.CallOpts, profileId *big.Int, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getOperatorApproval", profileId, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetOperatorApproval is a free data retrieval call binding the contract method 0xf0ab9600.
//
// Solidity: function getOperatorApproval(uint256 profileId, address operator) view returns(bool)
func (_Vote *VoteSession) GetOperatorApproval(profileId *big.Int, operator common.Address) (bool, error) {
	return _Vote.Contract.GetOperatorApproval(&_Vote.CallOpts, profileId, operator)
}

// GetOperatorApproval is a free data retrieval call binding the contract method 0xf0ab9600.
//
// Solidity: function getOperatorApproval(uint256 profileId, address operator) view returns(bool)
func (_Vote *VoteCallerSession) GetOperatorApproval(profileId *big.Int, operator common.Address) (bool, error) {
	return _Vote.Contract.GetOperatorApproval(&_Vote.CallOpts, profileId, operator)
}

// GetPrimaryProfile is a free data retrieval call binding the contract method 0xdfa52f07.
//
// Solidity: function getPrimaryProfile(address user) view returns(uint256)
func (_Vote *VoteCaller) GetPrimaryProfile(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getPrimaryProfile", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryProfile is a free data retrieval call binding the contract method 0xdfa52f07.
//
// Solidity: function getPrimaryProfile(address user) view returns(uint256)
func (_Vote *VoteSession) GetPrimaryProfile(user common.Address) (*big.Int, error) {
	return _Vote.Contract.GetPrimaryProfile(&_Vote.CallOpts, user)
}

// GetPrimaryProfile is a free data retrieval call binding the contract method 0xdfa52f07.
//
// Solidity: function getPrimaryProfile(address user) view returns(uint256)
func (_Vote *VoteCallerSession) GetPrimaryProfile(user common.Address) (*big.Int, error) {
	return _Vote.Contract.GetPrimaryProfile(&_Vote.CallOpts, user)
}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_Vote *VoteCaller) GetProfileIdByHandle(opts *bind.CallOpts, handle string) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getProfileIdByHandle", handle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_Vote *VoteSession) GetProfileIdByHandle(handle string) (*big.Int, error) {
	return _Vote.Contract.GetProfileIdByHandle(&_Vote.CallOpts, handle)
}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_Vote *VoteCallerSession) GetProfileIdByHandle(handle string) (*big.Int, error) {
	return _Vote.Contract.GetProfileIdByHandle(&_Vote.CallOpts, handle)
}

// GetSubscribeMw is a free data retrieval call binding the contract method 0x56bfd88c.
//
// Solidity: function getSubscribeMw(uint256 profileId) view returns(address)
func (_Vote *VoteCaller) GetSubscribeMw(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getSubscribeMw", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubscribeMw is a free data retrieval call binding the contract method 0x56bfd88c.
//
// Solidity: function getSubscribeMw(uint256 profileId) view returns(address)
func (_Vote *VoteSession) GetSubscribeMw(profileId *big.Int) (common.Address, error) {
	return _Vote.Contract.GetSubscribeMw(&_Vote.CallOpts, profileId)
}

// GetSubscribeMw is a free data retrieval call binding the contract method 0x56bfd88c.
//
// Solidity: function getSubscribeMw(uint256 profileId) view returns(address)
func (_Vote *VoteCallerSession) GetSubscribeMw(profileId *big.Int) (common.Address, error) {
	return _Vote.Contract.GetSubscribeMw(&_Vote.CallOpts, profileId)
}

// GetSubscribeNFT is a free data retrieval call binding the contract method 0xf4954aef.
//
// Solidity: function getSubscribeNFT(uint256 profileId) view returns(address)
func (_Vote *VoteCaller) GetSubscribeNFT(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getSubscribeNFT", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubscribeNFT is a free data retrieval call binding the contract method 0xf4954aef.
//
// Solidity: function getSubscribeNFT(uint256 profileId) view returns(address)
func (_Vote *VoteSession) GetSubscribeNFT(profileId *big.Int) (common.Address, error) {
	return _Vote.Contract.GetSubscribeNFT(&_Vote.CallOpts, profileId)
}

// GetSubscribeNFT is a free data retrieval call binding the contract method 0xf4954aef.
//
// Solidity: function getSubscribeNFT(uint256 profileId) view returns(address)
func (_Vote *VoteCallerSession) GetSubscribeNFT(profileId *big.Int) (common.Address, error) {
	return _Vote.Contract.GetSubscribeNFT(&_Vote.CallOpts, profileId)
}

// GetSubscribeNFTTokenURI is a free data retrieval call binding the contract method 0x303c14f8.
//
// Solidity: function getSubscribeNFTTokenURI(uint256 profileId) view returns(string)
func (_Vote *VoteCaller) GetSubscribeNFTTokenURI(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getSubscribeNFTTokenURI", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSubscribeNFTTokenURI is a free data retrieval call binding the contract method 0x303c14f8.
//
// Solidity: function getSubscribeNFTTokenURI(uint256 profileId) view returns(string)
func (_Vote *VoteSession) GetSubscribeNFTTokenURI(profileId *big.Int) (string, error) {
	return _Vote.Contract.GetSubscribeNFTTokenURI(&_Vote.CallOpts, profileId)
}

// GetSubscribeNFTTokenURI is a free data retrieval call binding the contract method 0x303c14f8.
//
// Solidity: function getSubscribeNFTTokenURI(uint256 profileId) view returns(string)
func (_Vote *VoteCallerSession) GetSubscribeNFTTokenURI(profileId *big.Int) (string, error) {
	return _Vote.Contract.GetSubscribeNFTTokenURI(&_Vote.CallOpts, profileId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Vote *VoteCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Vote *VoteSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vote.Contract.IsApprovedForAll(&_Vote.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Vote *VoteCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vote.Contract.IsApprovedForAll(&_Vote.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vote *VoteCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vote *VoteSession) Name() (string, error) {
	return _Vote.Contract.Name(&_Vote.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vote *VoteCallerSession) Name() (string, error) {
	return _Vote.Contract.Name(&_Vote.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Vote *VoteCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Vote *VoteSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Vote.Contract.Nonces(&_Vote.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Vote *VoteCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Vote.Contract.Nonces(&_Vote.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Vote *VoteCaller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Vote *VoteSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Vote.Contract.OwnerOf(&_Vote.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Vote *VoteCallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Vote.Contract.OwnerOf(&_Vote.CallOpts, id)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vote *VoteCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vote *VoteSession) Paused() (bool, error) {
	return _Vote.Contract.Paused(&_Vote.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vote *VoteCallerSession) Paused() (bool, error) {
	return _Vote.Contract.Paused(&_Vote.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vote *VoteCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vote *VoteSession) ProxiableUUID() ([32]byte, error) {
	return _Vote.Contract.ProxiableUUID(&_Vote.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vote *VoteCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Vote.Contract.ProxiableUUID(&_Vote.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vote *VoteCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vote *VoteSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vote.Contract.SupportsInterface(&_Vote.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vote *VoteCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vote.Contract.SupportsInterface(&_Vote.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vote *VoteCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vote *VoteSession) Symbol() (string, error) {
	return _Vote.Contract.Symbol(&_Vote.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vote *VoteCallerSession) Symbol() (string, error) {
	return _Vote.Contract.Symbol(&_Vote.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Vote *VoteCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Vote *VoteSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Vote.Contract.TokenURI(&_Vote.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Vote *VoteCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Vote.Contract.TokenURI(&_Vote.CallOpts, tokenId)
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_Vote *VoteCaller) TotalBurned(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "totalBurned")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_Vote *VoteSession) TotalBurned() (*big.Int, error) {
	return _Vote.Contract.TotalBurned(&_Vote.CallOpts)
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_Vote *VoteCallerSession) TotalBurned() (*big.Int, error) {
	return _Vote.Contract.TotalBurned(&_Vote.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Vote *VoteCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Vote *VoteSession) TotalMinted() (*big.Int, error) {
	return _Vote.Contract.TotalMinted(&_Vote.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Vote *VoteCallerSession) TotalMinted() (*big.Int, error) {
	return _Vote.Contract.TotalMinted(&_Vote.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vote *VoteCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vote *VoteSession) TotalSupply() (*big.Int, error) {
	return _Vote.Contract.TotalSupply(&_Vote.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vote *VoteCallerSession) TotalSupply() (*big.Int, error) {
	return _Vote.Contract.TotalSupply(&_Vote.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Vote *VoteCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Vote *VoteSession) Version() (*big.Int, error) {
	return _Vote.Contract.Version(&_Vote.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Vote *VoteCallerSession) Version() (*big.Int, error) {
	return _Vote.Contract.Version(&_Vote.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Vote *VoteTransactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "approve", spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Vote *VoteSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Approve(&_Vote.TransactOpts, spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Vote *VoteTransactorSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Approve(&_Vote.TransactOpts, spender, id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Vote *VoteTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Vote *VoteSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Burn(&_Vote.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Vote *VoteTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Burn(&_Vote.TransactOpts, tokenId)
}

// Collect is a paid mutator transaction binding the contract method 0xbe10bc61.
//
// Solidity: function collect((address,uint256,uint256) params, bytes preData, bytes postData) returns(uint256 tokenId)
func (_Vote *VoteTransactor) Collect(opts *bind.TransactOpts, params DataTypesCollectParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "collect", params, preData, postData)
}

// Collect is a paid mutator transaction binding the contract method 0xbe10bc61.
//
// Solidity: function collect((address,uint256,uint256) params, bytes preData, bytes postData) returns(uint256 tokenId)
func (_Vote *VoteSession) Collect(params DataTypesCollectParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _Vote.Contract.Collect(&_Vote.TransactOpts, params, preData, postData)
}

// Collect is a paid mutator transaction binding the contract method 0xbe10bc61.
//
// Solidity: function collect((address,uint256,uint256) params, bytes preData, bytes postData) returns(uint256 tokenId)
func (_Vote *VoteTransactorSession) Collect(params DataTypesCollectParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _Vote.Contract.Collect(&_Vote.TransactOpts, params, preData, postData)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0x54036644.
//
// Solidity: function collectWithSig((address,uint256,uint256) params, bytes preData, bytes postData, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_Vote *VoteTransactor) CollectWithSig(opts *bind.TransactOpts, params DataTypesCollectParams, preData []byte, postData []byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "collectWithSig", params, preData, postData, sender, sig)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0x54036644.
//
// Solidity: function collectWithSig((address,uint256,uint256) params, bytes preData, bytes postData, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_Vote *VoteSession) CollectWithSig(params DataTypesCollectParams, preData []byte, postData []byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.CollectWithSig(&_Vote.TransactOpts, params, preData, postData, sender, sig)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0x54036644.
//
// Solidity: function collectWithSig((address,uint256,uint256) params, bytes preData, bytes postData, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_Vote *VoteTransactorSession) CollectWithSig(params DataTypesCollectParams, preData []byte, postData []byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.CollectWithSig(&_Vote.TransactOpts, params, preData, postData, sender, sig)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x166fad6f.
//
// Solidity: function createProfile((address,string,string,string,address) params, bytes preData, bytes postData) payable returns(uint256 tokenID)
func (_Vote *VoteTransactor) CreateProfile(opts *bind.TransactOpts, params DataTypesCreateProfileParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "createProfile", params, preData, postData)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x166fad6f.
//
// Solidity: function createProfile((address,string,string,string,address) params, bytes preData, bytes postData) payable returns(uint256 tokenID)
func (_Vote *VoteSession) CreateProfile(params DataTypesCreateProfileParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _Vote.Contract.CreateProfile(&_Vote.TransactOpts, params, preData, postData)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x166fad6f.
//
// Solidity: function createProfile((address,string,string,string,address) params, bytes preData, bytes postData) payable returns(uint256 tokenID)
func (_Vote *VoteTransactorSession) CreateProfile(params DataTypesCreateProfileParams, preData []byte, postData []byte) (*types.Transaction, error) {
	return _Vote.Contract.CreateProfile(&_Vote.TransactOpts, params, preData, postData)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _owner, string name, string symbol) returns()
func (_Vote *VoteTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "initialize", _owner, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _owner, string name, string symbol) returns()
func (_Vote *VoteSession) Initialize(_owner common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Vote.Contract.Initialize(&_Vote.TransactOpts, _owner, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x90657147.
//
// Solidity: function initialize(address _owner, string name, string symbol) returns()
func (_Vote *VoteTransactorSession) Initialize(_owner common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Vote.Contract.Initialize(&_Vote.TransactOpts, _owner, name, symbol)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool toPause) returns()
func (_Vote *VoteTransactor) Pause(opts *bind.TransactOpts, toPause bool) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "pause", toPause)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool toPause) returns()
func (_Vote *VoteSession) Pause(toPause bool) (*types.Transaction, error) {
	return _Vote.Contract.Pause(&_Vote.TransactOpts, toPause)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool toPause) returns()
func (_Vote *VoteTransactorSession) Pause(toPause bool) (*types.Transaction, error) {
	return _Vote.Contract.Pause(&_Vote.TransactOpts, toPause)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactor) Permit(opts *bind.TransactOpts, spender common.Address, tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "permit", spender, tokenId, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteSession) Permit(spender common.Address, tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.Permit(&_Vote.TransactOpts, spender, tokenId, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactorSession) Permit(spender common.Address, tokenId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.Permit(&_Vote.TransactOpts, spender, tokenId, sig)
}

// RegisterEssence is a paid mutator transaction binding the contract method 0x71fd0a98.
//
// Solidity: function registerEssence((uint256,string,string,string,address,bool,bool) params, bytes initData) returns(uint256)
func (_Vote *VoteTransactor) RegisterEssence(opts *bind.TransactOpts, params DataTypesRegisterEssenceParams, initData []byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "registerEssence", params, initData)
}

// RegisterEssence is a paid mutator transaction binding the contract method 0x71fd0a98.
//
// Solidity: function registerEssence((uint256,string,string,string,address,bool,bool) params, bytes initData) returns(uint256)
func (_Vote *VoteSession) RegisterEssence(params DataTypesRegisterEssenceParams, initData []byte) (*types.Transaction, error) {
	return _Vote.Contract.RegisterEssence(&_Vote.TransactOpts, params, initData)
}

// RegisterEssence is a paid mutator transaction binding the contract method 0x71fd0a98.
//
// Solidity: function registerEssence((uint256,string,string,string,address,bool,bool) params, bytes initData) returns(uint256)
func (_Vote *VoteTransactorSession) RegisterEssence(params DataTypesRegisterEssenceParams, initData []byte) (*types.Transaction, error) {
	return _Vote.Contract.RegisterEssence(&_Vote.TransactOpts, params, initData)
}

// RegisterEssenceWithSig is a paid mutator transaction binding the contract method 0xe5fee479.
//
// Solidity: function registerEssenceWithSig((uint256,string,string,string,address,bool,bool) params, bytes initData, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_Vote *VoteTransactor) RegisterEssenceWithSig(opts *bind.TransactOpts, params DataTypesRegisterEssenceParams, initData []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "registerEssenceWithSig", params, initData, sig)
}

// RegisterEssenceWithSig is a paid mutator transaction binding the contract method 0xe5fee479.
//
// Solidity: function registerEssenceWithSig((uint256,string,string,string,address,bool,bool) params, bytes initData, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_Vote *VoteSession) RegisterEssenceWithSig(params DataTypesRegisterEssenceParams, initData []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.RegisterEssenceWithSig(&_Vote.TransactOpts, params, initData, sig)
}

// RegisterEssenceWithSig is a paid mutator transaction binding the contract method 0xe5fee479.
//
// Solidity: function registerEssenceWithSig((uint256,string,string,string,address,bool,bool) params, bytes initData, (uint8,bytes32,bytes32,uint256) sig) returns(uint256 tokenId)
func (_Vote *VoteTransactorSession) RegisterEssenceWithSig(params DataTypesRegisterEssenceParams, initData []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.RegisterEssenceWithSig(&_Vote.TransactOpts, params, initData, sig)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Vote *VoteTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Vote *VoteSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.SafeTransferFrom(&_Vote.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Vote *VoteTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.SafeTransferFrom(&_Vote.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Vote *VoteTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Vote *VoteSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Vote.Contract.SafeTransferFrom0(&_Vote.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Vote *VoteTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Vote.Contract.SafeTransferFrom0(&_Vote.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Vote *VoteTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Vote *VoteSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Vote.Contract.SetApprovalForAll(&_Vote.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Vote *VoteTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Vote.Contract.SetApprovalForAll(&_Vote.TransactOpts, operator, approved)
}

// SetAvatar is a paid mutator transaction binding the contract method 0x87beeb2b.
//
// Solidity: function setAvatar(uint256 profileId, string avatar) returns()
func (_Vote *VoteTransactor) SetAvatar(opts *bind.TransactOpts, profileId *big.Int, avatar string) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setAvatar", profileId, avatar)
}

// SetAvatar is a paid mutator transaction binding the contract method 0x87beeb2b.
//
// Solidity: function setAvatar(uint256 profileId, string avatar) returns()
func (_Vote *VoteSession) SetAvatar(profileId *big.Int, avatar string) (*types.Transaction, error) {
	return _Vote.Contract.SetAvatar(&_Vote.TransactOpts, profileId, avatar)
}

// SetAvatar is a paid mutator transaction binding the contract method 0x87beeb2b.
//
// Solidity: function setAvatar(uint256 profileId, string avatar) returns()
func (_Vote *VoteTransactorSession) SetAvatar(profileId *big.Int, avatar string) (*types.Transaction, error) {
	return _Vote.Contract.SetAvatar(&_Vote.TransactOpts, profileId, avatar)
}

// SetAvatarWithSig is a paid mutator transaction binding the contract method 0x6463788d.
//
// Solidity: function setAvatarWithSig(uint256 profileId, string avatar, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactor) SetAvatarWithSig(opts *bind.TransactOpts, profileId *big.Int, avatar string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setAvatarWithSig", profileId, avatar, sig)
}

// SetAvatarWithSig is a paid mutator transaction binding the contract method 0x6463788d.
//
// Solidity: function setAvatarWithSig(uint256 profileId, string avatar, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteSession) SetAvatarWithSig(profileId *big.Int, avatar string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SetAvatarWithSig(&_Vote.TransactOpts, profileId, avatar, sig)
}

// SetAvatarWithSig is a paid mutator transaction binding the contract method 0x6463788d.
//
// Solidity: function setAvatarWithSig(uint256 profileId, string avatar, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactorSession) SetAvatarWithSig(profileId *big.Int, avatar string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SetAvatarWithSig(&_Vote.TransactOpts, profileId, avatar, sig)
}

// SetEssenceData is a paid mutator transaction binding the contract method 0x0eee16f4.
//
// Solidity: function setEssenceData(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data) returns()
func (_Vote *VoteTransactor) SetEssenceData(opts *bind.TransactOpts, profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setEssenceData", profileId, essenceId, uri, mw, data)
}

// SetEssenceData is a paid mutator transaction binding the contract method 0x0eee16f4.
//
// Solidity: function setEssenceData(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data) returns()
func (_Vote *VoteSession) SetEssenceData(profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _Vote.Contract.SetEssenceData(&_Vote.TransactOpts, profileId, essenceId, uri, mw, data)
}

// SetEssenceData is a paid mutator transaction binding the contract method 0x0eee16f4.
//
// Solidity: function setEssenceData(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data) returns()
func (_Vote *VoteTransactorSession) SetEssenceData(profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _Vote.Contract.SetEssenceData(&_Vote.TransactOpts, profileId, essenceId, uri, mw, data)
}

// SetEssenceDataWithSig is a paid mutator transaction binding the contract method 0x5ea97079.
//
// Solidity: function setEssenceDataWithSig(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactor) SetEssenceDataWithSig(opts *bind.TransactOpts, profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setEssenceDataWithSig", profileId, essenceId, uri, mw, data, sig)
}

// SetEssenceDataWithSig is a paid mutator transaction binding the contract method 0x5ea97079.
//
// Solidity: function setEssenceDataWithSig(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteSession) SetEssenceDataWithSig(profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SetEssenceDataWithSig(&_Vote.TransactOpts, profileId, essenceId, uri, mw, data, sig)
}

// SetEssenceDataWithSig is a paid mutator transaction binding the contract method 0x5ea97079.
//
// Solidity: function setEssenceDataWithSig(uint256 profileId, uint256 essenceId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactorSession) SetEssenceDataWithSig(profileId *big.Int, essenceId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SetEssenceDataWithSig(&_Vote.TransactOpts, profileId, essenceId, uri, mw, data, sig)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x593aa283.
//
// Solidity: function setMetadata(uint256 profileId, string metadata) returns()
func (_Vote *VoteTransactor) SetMetadata(opts *bind.TransactOpts, profileId *big.Int, metadata string) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setMetadata", profileId, metadata)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x593aa283.
//
// Solidity: function setMetadata(uint256 profileId, string metadata) returns()
func (_Vote *VoteSession) SetMetadata(profileId *big.Int, metadata string) (*types.Transaction, error) {
	return _Vote.Contract.SetMetadata(&_Vote.TransactOpts, profileId, metadata)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x593aa283.
//
// Solidity: function setMetadata(uint256 profileId, string metadata) returns()
func (_Vote *VoteTransactorSession) SetMetadata(profileId *big.Int, metadata string) (*types.Transaction, error) {
	return _Vote.Contract.SetMetadata(&_Vote.TransactOpts, profileId, metadata)
}

// SetMetadataWithSig is a paid mutator transaction binding the contract method 0xd0e56c38.
//
// Solidity: function setMetadataWithSig(uint256 profileId, string metadata, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactor) SetMetadataWithSig(opts *bind.TransactOpts, profileId *big.Int, metadata string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setMetadataWithSig", profileId, metadata, sig)
}

// SetMetadataWithSig is a paid mutator transaction binding the contract method 0xd0e56c38.
//
// Solidity: function setMetadataWithSig(uint256 profileId, string metadata, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteSession) SetMetadataWithSig(profileId *big.Int, metadata string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SetMetadataWithSig(&_Vote.TransactOpts, profileId, metadata, sig)
}

// SetMetadataWithSig is a paid mutator transaction binding the contract method 0xd0e56c38.
//
// Solidity: function setMetadataWithSig(uint256 profileId, string metadata, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactorSession) SetMetadataWithSig(profileId *big.Int, metadata string, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SetMetadataWithSig(&_Vote.TransactOpts, profileId, metadata, sig)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address descriptor) returns()
func (_Vote *VoteTransactor) SetNFTDescriptor(opts *bind.TransactOpts, descriptor common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setNFTDescriptor", descriptor)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address descriptor) returns()
func (_Vote *VoteSession) SetNFTDescriptor(descriptor common.Address) (*types.Transaction, error) {
	return _Vote.Contract.SetNFTDescriptor(&_Vote.TransactOpts, descriptor)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address descriptor) returns()
func (_Vote *VoteTransactorSession) SetNFTDescriptor(descriptor common.Address) (*types.Transaction, error) {
	return _Vote.Contract.SetNFTDescriptor(&_Vote.TransactOpts, descriptor)
}

// SetNamespaceOwner is a paid mutator transaction binding the contract method 0xf7727c9f.
//
// Solidity: function setNamespaceOwner(address owner) returns()
func (_Vote *VoteTransactor) SetNamespaceOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setNamespaceOwner", owner)
}

// SetNamespaceOwner is a paid mutator transaction binding the contract method 0xf7727c9f.
//
// Solidity: function setNamespaceOwner(address owner) returns()
func (_Vote *VoteSession) SetNamespaceOwner(owner common.Address) (*types.Transaction, error) {
	return _Vote.Contract.SetNamespaceOwner(&_Vote.TransactOpts, owner)
}

// SetNamespaceOwner is a paid mutator transaction binding the contract method 0xf7727c9f.
//
// Solidity: function setNamespaceOwner(address owner) returns()
func (_Vote *VoteTransactorSession) SetNamespaceOwner(owner common.Address) (*types.Transaction, error) {
	return _Vote.Contract.SetNamespaceOwner(&_Vote.TransactOpts, owner)
}

// SetOperatorApproval is a paid mutator transaction binding the contract method 0xafa83e68.
//
// Solidity: function setOperatorApproval(uint256 profileId, address operator, bool approved) returns()
func (_Vote *VoteTransactor) SetOperatorApproval(opts *bind.TransactOpts, profileId *big.Int, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setOperatorApproval", profileId, operator, approved)
}

// SetOperatorApproval is a paid mutator transaction binding the contract method 0xafa83e68.
//
// Solidity: function setOperatorApproval(uint256 profileId, address operator, bool approved) returns()
func (_Vote *VoteSession) SetOperatorApproval(profileId *big.Int, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Vote.Contract.SetOperatorApproval(&_Vote.TransactOpts, profileId, operator, approved)
}

// SetOperatorApproval is a paid mutator transaction binding the contract method 0xafa83e68.
//
// Solidity: function setOperatorApproval(uint256 profileId, address operator, bool approved) returns()
func (_Vote *VoteTransactorSession) SetOperatorApproval(profileId *big.Int, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Vote.Contract.SetOperatorApproval(&_Vote.TransactOpts, profileId, operator, approved)
}

// SetPrimaryProfile is a paid mutator transaction binding the contract method 0x7c08d9fd.
//
// Solidity: function setPrimaryProfile(uint256 profileId) returns()
func (_Vote *VoteTransactor) SetPrimaryProfile(opts *bind.TransactOpts, profileId *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setPrimaryProfile", profileId)
}

// SetPrimaryProfile is a paid mutator transaction binding the contract method 0x7c08d9fd.
//
// Solidity: function setPrimaryProfile(uint256 profileId) returns()
func (_Vote *VoteSession) SetPrimaryProfile(profileId *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.SetPrimaryProfile(&_Vote.TransactOpts, profileId)
}

// SetPrimaryProfile is a paid mutator transaction binding the contract method 0x7c08d9fd.
//
// Solidity: function setPrimaryProfile(uint256 profileId) returns()
func (_Vote *VoteTransactorSession) SetPrimaryProfile(profileId *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.SetPrimaryProfile(&_Vote.TransactOpts, profileId)
}

// SetPrimaryProfileWithSig is a paid mutator transaction binding the contract method 0xa95c1cbf.
//
// Solidity: function setPrimaryProfileWithSig(uint256 profileId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactor) SetPrimaryProfileWithSig(opts *bind.TransactOpts, profileId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setPrimaryProfileWithSig", profileId, sig)
}

// SetPrimaryProfileWithSig is a paid mutator transaction binding the contract method 0xa95c1cbf.
//
// Solidity: function setPrimaryProfileWithSig(uint256 profileId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteSession) SetPrimaryProfileWithSig(profileId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SetPrimaryProfileWithSig(&_Vote.TransactOpts, profileId, sig)
}

// SetPrimaryProfileWithSig is a paid mutator transaction binding the contract method 0xa95c1cbf.
//
// Solidity: function setPrimaryProfileWithSig(uint256 profileId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactorSession) SetPrimaryProfileWithSig(profileId *big.Int, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SetPrimaryProfileWithSig(&_Vote.TransactOpts, profileId, sig)
}

// SetSubscribeData is a paid mutator transaction binding the contract method 0x8d54db1b.
//
// Solidity: function setSubscribeData(uint256 profileId, string uri, address mw, bytes data) returns()
func (_Vote *VoteTransactor) SetSubscribeData(opts *bind.TransactOpts, profileId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setSubscribeData", profileId, uri, mw, data)
}

// SetSubscribeData is a paid mutator transaction binding the contract method 0x8d54db1b.
//
// Solidity: function setSubscribeData(uint256 profileId, string uri, address mw, bytes data) returns()
func (_Vote *VoteSession) SetSubscribeData(profileId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _Vote.Contract.SetSubscribeData(&_Vote.TransactOpts, profileId, uri, mw, data)
}

// SetSubscribeData is a paid mutator transaction binding the contract method 0x8d54db1b.
//
// Solidity: function setSubscribeData(uint256 profileId, string uri, address mw, bytes data) returns()
func (_Vote *VoteTransactorSession) SetSubscribeData(profileId *big.Int, uri string, mw common.Address, data []byte) (*types.Transaction, error) {
	return _Vote.Contract.SetSubscribeData(&_Vote.TransactOpts, profileId, uri, mw, data)
}

// SetSubscribeDataWithSig is a paid mutator transaction binding the contract method 0x2229f28d.
//
// Solidity: function setSubscribeDataWithSig(uint256 profileId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactor) SetSubscribeDataWithSig(opts *bind.TransactOpts, profileId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "setSubscribeDataWithSig", profileId, uri, mw, data, sig)
}

// SetSubscribeDataWithSig is a paid mutator transaction binding the contract method 0x2229f28d.
//
// Solidity: function setSubscribeDataWithSig(uint256 profileId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteSession) SetSubscribeDataWithSig(profileId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SetSubscribeDataWithSig(&_Vote.TransactOpts, profileId, uri, mw, data, sig)
}

// SetSubscribeDataWithSig is a paid mutator transaction binding the contract method 0x2229f28d.
//
// Solidity: function setSubscribeDataWithSig(uint256 profileId, string uri, address mw, bytes data, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_Vote *VoteTransactorSession) SetSubscribeDataWithSig(profileId *big.Int, uri string, mw common.Address, data []byte, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SetSubscribeDataWithSig(&_Vote.TransactOpts, profileId, uri, mw, data, sig)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4940e2d1.
//
// Solidity: function subscribe((uint256[]) params, bytes[] preDatas, bytes[] postDatas) returns(uint256[])
func (_Vote *VoteTransactor) Subscribe(opts *bind.TransactOpts, params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "subscribe", params, preDatas, postDatas)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4940e2d1.
//
// Solidity: function subscribe((uint256[]) params, bytes[] preDatas, bytes[] postDatas) returns(uint256[])
func (_Vote *VoteSession) Subscribe(params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte) (*types.Transaction, error) {
	return _Vote.Contract.Subscribe(&_Vote.TransactOpts, params, preDatas, postDatas)
}

// Subscribe is a paid mutator transaction binding the contract method 0x4940e2d1.
//
// Solidity: function subscribe((uint256[]) params, bytes[] preDatas, bytes[] postDatas) returns(uint256[])
func (_Vote *VoteTransactorSession) Subscribe(params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte) (*types.Transaction, error) {
	return _Vote.Contract.Subscribe(&_Vote.TransactOpts, params, preDatas, postDatas)
}

// SubscribeWithSig is a paid mutator transaction binding the contract method 0x026103e9.
//
// Solidity: function subscribeWithSig((uint256[]) params, bytes[] preDatas, bytes[] postDatas, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256[])
func (_Vote *VoteTransactor) SubscribeWithSig(opts *bind.TransactOpts, params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "subscribeWithSig", params, preDatas, postDatas, sender, sig)
}

// SubscribeWithSig is a paid mutator transaction binding the contract method 0x026103e9.
//
// Solidity: function subscribeWithSig((uint256[]) params, bytes[] preDatas, bytes[] postDatas, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256[])
func (_Vote *VoteSession) SubscribeWithSig(params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SubscribeWithSig(&_Vote.TransactOpts, params, preDatas, postDatas, sender, sig)
}

// SubscribeWithSig is a paid mutator transaction binding the contract method 0x026103e9.
//
// Solidity: function subscribeWithSig((uint256[]) params, bytes[] preDatas, bytes[] postDatas, address sender, (uint8,bytes32,bytes32,uint256) sig) returns(uint256[])
func (_Vote *VoteTransactorSession) SubscribeWithSig(params DataTypesSubscribeParams, preDatas [][]byte, postDatas [][]byte, sender common.Address, sig DataTypesEIP712Signature) (*types.Transaction, error) {
	return _Vote.Contract.SubscribeWithSig(&_Vote.TransactOpts, params, preDatas, postDatas, sender, sig)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Vote *VoteTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Vote *VoteSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.TransferFrom(&_Vote.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Vote *VoteTransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.TransferFrom(&_Vote.TransactOpts, from, to, id)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vote *VoteTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vote *VoteSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Vote.Contract.UpgradeTo(&_Vote.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vote *VoteTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Vote.Contract.UpgradeTo(&_Vote.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vote *VoteTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vote *VoteSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vote.Contract.UpgradeToAndCall(&_Vote.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vote *VoteTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vote.Contract.UpgradeToAndCall(&_Vote.TransactOpts, newImplementation, data)
}

// VoteAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Vote contract.
type VoteAdminChangedIterator struct {
	Event *VoteAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteAdminChanged represents a AdminChanged event raised by the Vote contract.
type VoteAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vote *VoteFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VoteAdminChangedIterator, error) {

	logs, sub, err := _Vote.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VoteAdminChangedIterator{contract: _Vote.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vote *VoteFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VoteAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Vote.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteAdminChanged)
				if err := _Vote.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vote *VoteFilterer) ParseAdminChanged(log types.Log) (*VoteAdminChanged, error) {
	event := new(VoteAdminChanged)
	if err := _Vote.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vote contract.
type VoteApprovalIterator struct {
	Event *VoteApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteApproval represents a Approval event raised by the Vote contract.
type VoteApproval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Vote *VoteFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*VoteApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &VoteApprovalIterator{contract: _Vote.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Vote *VoteFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VoteApproval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteApproval)
				if err := _Vote.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Vote *VoteFilterer) ParseApproval(log types.Log) (*VoteApproval, error) {
	event := new(VoteApproval)
	if err := _Vote.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Vote contract.
type VoteApprovalForAllIterator struct {
	Event *VoteApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteApprovalForAll represents a ApprovalForAll event raised by the Vote contract.
type VoteApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Vote *VoteFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*VoteApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &VoteApprovalForAllIterator{contract: _Vote.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Vote *VoteFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *VoteApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteApprovalForAll)
				if err := _Vote.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Vote *VoteFilterer) ParseApprovalForAll(log types.Log) (*VoteApprovalForAll, error) {
	event := new(VoteApprovalForAll)
	if err := _Vote.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Vote contract.
type VoteBeaconUpgradedIterator struct {
	Event *VoteBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteBeaconUpgraded represents a BeaconUpgraded event raised by the Vote contract.
type VoteBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vote *VoteFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VoteBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VoteBeaconUpgradedIterator{contract: _Vote.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vote *VoteFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VoteBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteBeaconUpgraded)
				if err := _Vote.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vote *VoteFilterer) ParseBeaconUpgraded(log types.Log) (*VoteBeaconUpgraded, error) {
	event := new(VoteBeaconUpgraded)
	if err := _Vote.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteCollectEssenceIterator is returned from FilterCollectEssence and is used to iterate over the raw logs and unpacked data for CollectEssence events raised by the Vote contract.
type VoteCollectEssenceIterator struct {
	Event *VoteCollectEssence // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteCollectEssenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteCollectEssence)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteCollectEssence)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteCollectEssenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteCollectEssenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteCollectEssence represents a CollectEssence event raised by the Vote contract.
type VoteCollectEssence struct {
	Collector common.Address
	ProfileId *big.Int
	EssenceId *big.Int
	TokenId   *big.Int
	PreData   []byte
	PostData  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectEssence is a free log retrieval operation binding the contract event 0xd4665c34529353804143794023e073d96230b57acf608ae514de33d10a09354f.
//
// Solidity: event CollectEssence(address indexed collector, uint256 indexed profileId, uint256 indexed essenceId, uint256 tokenId, bytes preData, bytes postData)
func (_Vote *VoteFilterer) FilterCollectEssence(opts *bind.FilterOpts, collector []common.Address, profileId []*big.Int, essenceId []*big.Int) (*VoteCollectEssenceIterator, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "CollectEssence", collectorRule, profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteCollectEssenceIterator{contract: _Vote.contract, event: "CollectEssence", logs: logs, sub: sub}, nil
}

// WatchCollectEssence is a free log subscription operation binding the contract event 0xd4665c34529353804143794023e073d96230b57acf608ae514de33d10a09354f.
//
// Solidity: event CollectEssence(address indexed collector, uint256 indexed profileId, uint256 indexed essenceId, uint256 tokenId, bytes preData, bytes postData)
func (_Vote *VoteFilterer) WatchCollectEssence(opts *bind.WatchOpts, sink chan<- *VoteCollectEssence, collector []common.Address, profileId []*big.Int, essenceId []*big.Int) (event.Subscription, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "CollectEssence", collectorRule, profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteCollectEssence)
				if err := _Vote.contract.UnpackLog(event, "CollectEssence", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectEssence is a log parse operation binding the contract event 0xd4665c34529353804143794023e073d96230b57acf608ae514de33d10a09354f.
//
// Solidity: event CollectEssence(address indexed collector, uint256 indexed profileId, uint256 indexed essenceId, uint256 tokenId, bytes preData, bytes postData)
func (_Vote *VoteFilterer) ParseCollectEssence(log types.Log) (*VoteCollectEssence, error) {
	event := new(VoteCollectEssence)
	if err := _Vote.contract.UnpackLog(event, "CollectEssence", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteCreateProfileIterator is returned from FilterCreateProfile and is used to iterate over the raw logs and unpacked data for CreateProfile events raised by the Vote contract.
type VoteCreateProfileIterator struct {
	Event *VoteCreateProfile // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteCreateProfileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteCreateProfile)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteCreateProfile)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteCreateProfileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteCreateProfileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteCreateProfile represents a CreateProfile event raised by the Vote contract.
type VoteCreateProfile struct {
	To        common.Address
	ProfileId *big.Int
	Handle    string
	Avatar    string
	Metadata  string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateProfile is a free log retrieval operation binding the contract event 0x0d0246441d2185882e392b0a85e33212f2e1f668cf78c11c1808d421b2549fa6.
//
// Solidity: event CreateProfile(address indexed to, uint256 indexed profileId, string handle, string avatar, string metadata)
func (_Vote *VoteFilterer) FilterCreateProfile(opts *bind.FilterOpts, to []common.Address, profileId []*big.Int) (*VoteCreateProfileIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "CreateProfile", toRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteCreateProfileIterator{contract: _Vote.contract, event: "CreateProfile", logs: logs, sub: sub}, nil
}

// WatchCreateProfile is a free log subscription operation binding the contract event 0x0d0246441d2185882e392b0a85e33212f2e1f668cf78c11c1808d421b2549fa6.
//
// Solidity: event CreateProfile(address indexed to, uint256 indexed profileId, string handle, string avatar, string metadata)
func (_Vote *VoteFilterer) WatchCreateProfile(opts *bind.WatchOpts, sink chan<- *VoteCreateProfile, to []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "CreateProfile", toRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteCreateProfile)
				if err := _Vote.contract.UnpackLog(event, "CreateProfile", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateProfile is a log parse operation binding the contract event 0x0d0246441d2185882e392b0a85e33212f2e1f668cf78c11c1808d421b2549fa6.
//
// Solidity: event CreateProfile(address indexed to, uint256 indexed profileId, string handle, string avatar, string metadata)
func (_Vote *VoteFilterer) ParseCreateProfile(log types.Log) (*VoteCreateProfile, error) {
	event := new(VoteCreateProfile)
	if err := _Vote.contract.UnpackLog(event, "CreateProfile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteDeployEssenceNFTIterator is returned from FilterDeployEssenceNFT and is used to iterate over the raw logs and unpacked data for DeployEssenceNFT events raised by the Vote contract.
type VoteDeployEssenceNFTIterator struct {
	Event *VoteDeployEssenceNFT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteDeployEssenceNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteDeployEssenceNFT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteDeployEssenceNFT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteDeployEssenceNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteDeployEssenceNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteDeployEssenceNFT represents a DeployEssenceNFT event raised by the Vote contract.
type VoteDeployEssenceNFT struct {
	ProfileId  *big.Int
	EssenceId  *big.Int
	EssenceNFT common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeployEssenceNFT is a free log retrieval operation binding the contract event 0x9582a6b88daf141f38fa06c06d493aa8a09546d1b508289c833a9d48498ff189.
//
// Solidity: event DeployEssenceNFT(uint256 indexed profileId, uint256 indexed essenceId, address indexed essenceNFT)
func (_Vote *VoteFilterer) FilterDeployEssenceNFT(opts *bind.FilterOpts, profileId []*big.Int, essenceId []*big.Int, essenceNFT []common.Address) (*VoteDeployEssenceNFTIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}
	var essenceNFTRule []interface{}
	for _, essenceNFTItem := range essenceNFT {
		essenceNFTRule = append(essenceNFTRule, essenceNFTItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "DeployEssenceNFT", profileIdRule, essenceIdRule, essenceNFTRule)
	if err != nil {
		return nil, err
	}
	return &VoteDeployEssenceNFTIterator{contract: _Vote.contract, event: "DeployEssenceNFT", logs: logs, sub: sub}, nil
}

// WatchDeployEssenceNFT is a free log subscription operation binding the contract event 0x9582a6b88daf141f38fa06c06d493aa8a09546d1b508289c833a9d48498ff189.
//
// Solidity: event DeployEssenceNFT(uint256 indexed profileId, uint256 indexed essenceId, address indexed essenceNFT)
func (_Vote *VoteFilterer) WatchDeployEssenceNFT(opts *bind.WatchOpts, sink chan<- *VoteDeployEssenceNFT, profileId []*big.Int, essenceId []*big.Int, essenceNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}
	var essenceNFTRule []interface{}
	for _, essenceNFTItem := range essenceNFT {
		essenceNFTRule = append(essenceNFTRule, essenceNFTItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "DeployEssenceNFT", profileIdRule, essenceIdRule, essenceNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteDeployEssenceNFT)
				if err := _Vote.contract.UnpackLog(event, "DeployEssenceNFT", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeployEssenceNFT is a log parse operation binding the contract event 0x9582a6b88daf141f38fa06c06d493aa8a09546d1b508289c833a9d48498ff189.
//
// Solidity: event DeployEssenceNFT(uint256 indexed profileId, uint256 indexed essenceId, address indexed essenceNFT)
func (_Vote *VoteFilterer) ParseDeployEssenceNFT(log types.Log) (*VoteDeployEssenceNFT, error) {
	event := new(VoteDeployEssenceNFT)
	if err := _Vote.contract.UnpackLog(event, "DeployEssenceNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteDeploySubscribeNFTIterator is returned from FilterDeploySubscribeNFT and is used to iterate over the raw logs and unpacked data for DeploySubscribeNFT events raised by the Vote contract.
type VoteDeploySubscribeNFTIterator struct {
	Event *VoteDeploySubscribeNFT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteDeploySubscribeNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteDeploySubscribeNFT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteDeploySubscribeNFT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteDeploySubscribeNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteDeploySubscribeNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteDeploySubscribeNFT represents a DeploySubscribeNFT event raised by the Vote contract.
type VoteDeploySubscribeNFT struct {
	ProfileId    *big.Int
	SubscribeNFT common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeploySubscribeNFT is a free log retrieval operation binding the contract event 0x8fb25da3d7c3a4880851ad04d6fca5a6f663ca562d932462e1444544321de175.
//
// Solidity: event DeploySubscribeNFT(uint256 indexed profileId, address indexed subscribeNFT)
func (_Vote *VoteFilterer) FilterDeploySubscribeNFT(opts *bind.FilterOpts, profileId []*big.Int, subscribeNFT []common.Address) (*VoteDeploySubscribeNFTIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var subscribeNFTRule []interface{}
	for _, subscribeNFTItem := range subscribeNFT {
		subscribeNFTRule = append(subscribeNFTRule, subscribeNFTItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "DeploySubscribeNFT", profileIdRule, subscribeNFTRule)
	if err != nil {
		return nil, err
	}
	return &VoteDeploySubscribeNFTIterator{contract: _Vote.contract, event: "DeploySubscribeNFT", logs: logs, sub: sub}, nil
}

// WatchDeploySubscribeNFT is a free log subscription operation binding the contract event 0x8fb25da3d7c3a4880851ad04d6fca5a6f663ca562d932462e1444544321de175.
//
// Solidity: event DeploySubscribeNFT(uint256 indexed profileId, address indexed subscribeNFT)
func (_Vote *VoteFilterer) WatchDeploySubscribeNFT(opts *bind.WatchOpts, sink chan<- *VoteDeploySubscribeNFT, profileId []*big.Int, subscribeNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var subscribeNFTRule []interface{}
	for _, subscribeNFTItem := range subscribeNFT {
		subscribeNFTRule = append(subscribeNFTRule, subscribeNFTItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "DeploySubscribeNFT", profileIdRule, subscribeNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteDeploySubscribeNFT)
				if err := _Vote.contract.UnpackLog(event, "DeploySubscribeNFT", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeploySubscribeNFT is a log parse operation binding the contract event 0x8fb25da3d7c3a4880851ad04d6fca5a6f663ca562d932462e1444544321de175.
//
// Solidity: event DeploySubscribeNFT(uint256 indexed profileId, address indexed subscribeNFT)
func (_Vote *VoteFilterer) ParseDeploySubscribeNFT(log types.Log) (*VoteDeploySubscribeNFT, error) {
	event := new(VoteDeploySubscribeNFT)
	if err := _Vote.contract.UnpackLog(event, "DeploySubscribeNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the Vote contract.
type VoteInitializeIterator struct {
	Event *VoteInitialize // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteInitialize)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteInitialize)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteInitialize represents a Initialize event raised by the Vote contract.
type VoteInitialize struct {
	Owner  common.Address
	Name   string
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x5b98d40fc5d761472c5e9e4eeda87b2e7d3d58f72ed64216d9350fe5bc1b5bf1.
//
// Solidity: event Initialize(address indexed owner, string name, string symbol)
func (_Vote *VoteFilterer) FilterInitialize(opts *bind.FilterOpts, owner []common.Address) (*VoteInitializeIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "Initialize", ownerRule)
	if err != nil {
		return nil, err
	}
	return &VoteInitializeIterator{contract: _Vote.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x5b98d40fc5d761472c5e9e4eeda87b2e7d3d58f72ed64216d9350fe5bc1b5bf1.
//
// Solidity: event Initialize(address indexed owner, string name, string symbol)
func (_Vote *VoteFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *VoteInitialize, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "Initialize", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteInitialize)
				if err := _Vote.contract.UnpackLog(event, "Initialize", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialize is a log parse operation binding the contract event 0x5b98d40fc5d761472c5e9e4eeda87b2e7d3d58f72ed64216d9350fe5bc1b5bf1.
//
// Solidity: event Initialize(address indexed owner, string name, string symbol)
func (_Vote *VoteFilterer) ParseInitialize(log types.Log) (*VoteInitialize, error) {
	event := new(VoteInitialize)
	if err := _Vote.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Vote contract.
type VoteInitializedIterator struct {
	Event *VoteInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteInitialized represents a Initialized event raised by the Vote contract.
type VoteInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vote *VoteFilterer) FilterInitialized(opts *bind.FilterOpts) (*VoteInitializedIterator, error) {

	logs, sub, err := _Vote.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VoteInitializedIterator{contract: _Vote.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vote *VoteFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VoteInitialized) (event.Subscription, error) {

	logs, sub, err := _Vote.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteInitialized)
				if err := _Vote.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vote *VoteFilterer) ParseInitialized(log types.Log) (*VoteInitialized, error) {
	event := new(VoteInitialized)
	if err := _Vote.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Vote contract.
type VotePausedIterator struct {
	Event *VotePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VotePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VotePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VotePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotePaused represents a Paused event raised by the Vote contract.
type VotePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vote *VoteFilterer) FilterPaused(opts *bind.FilterOpts) (*VotePausedIterator, error) {

	logs, sub, err := _Vote.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VotePausedIterator{contract: _Vote.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vote *VoteFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VotePaused) (event.Subscription, error) {

	logs, sub, err := _Vote.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotePaused)
				if err := _Vote.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vote *VoteFilterer) ParsePaused(log types.Log) (*VotePaused, error) {
	event := new(VotePaused)
	if err := _Vote.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteRegisterEssenceIterator is returned from FilterRegisterEssence and is used to iterate over the raw logs and unpacked data for RegisterEssence events raised by the Vote contract.
type VoteRegisterEssenceIterator struct {
	Event *VoteRegisterEssence // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteRegisterEssenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteRegisterEssence)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteRegisterEssence)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteRegisterEssenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteRegisterEssenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteRegisterEssence represents a RegisterEssence event raised by the Vote contract.
type VoteRegisterEssence struct {
	ProfileId         *big.Int
	EssenceId         *big.Int
	Name              string
	Symbol            string
	EssenceTokenURI   string
	EssenceMw         common.Address
	PrepareReturnData []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRegisterEssence is a free log retrieval operation binding the contract event 0x7f8f05cb8699d58aa2b69fa7a1b1ab0db87ab881493e58a6842cec13b2982011.
//
// Solidity: event RegisterEssence(uint256 indexed profileId, uint256 indexed essenceId, string name, string symbol, string essenceTokenURI, address essenceMw, bytes prepareReturnData)
func (_Vote *VoteFilterer) FilterRegisterEssence(opts *bind.FilterOpts, profileId []*big.Int, essenceId []*big.Int) (*VoteRegisterEssenceIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "RegisterEssence", profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteRegisterEssenceIterator{contract: _Vote.contract, event: "RegisterEssence", logs: logs, sub: sub}, nil
}

// WatchRegisterEssence is a free log subscription operation binding the contract event 0x7f8f05cb8699d58aa2b69fa7a1b1ab0db87ab881493e58a6842cec13b2982011.
//
// Solidity: event RegisterEssence(uint256 indexed profileId, uint256 indexed essenceId, string name, string symbol, string essenceTokenURI, address essenceMw, bytes prepareReturnData)
func (_Vote *VoteFilterer) WatchRegisterEssence(opts *bind.WatchOpts, sink chan<- *VoteRegisterEssence, profileId []*big.Int, essenceId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "RegisterEssence", profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteRegisterEssence)
				if err := _Vote.contract.UnpackLog(event, "RegisterEssence", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegisterEssence is a log parse operation binding the contract event 0x7f8f05cb8699d58aa2b69fa7a1b1ab0db87ab881493e58a6842cec13b2982011.
//
// Solidity: event RegisterEssence(uint256 indexed profileId, uint256 indexed essenceId, string name, string symbol, string essenceTokenURI, address essenceMw, bytes prepareReturnData)
func (_Vote *VoteFilterer) ParseRegisterEssence(log types.Log) (*VoteRegisterEssence, error) {
	event := new(VoteRegisterEssence)
	if err := _Vote.contract.UnpackLog(event, "RegisterEssence", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteSetAvatarIterator is returned from FilterSetAvatar and is used to iterate over the raw logs and unpacked data for SetAvatar events raised by the Vote contract.
type VoteSetAvatarIterator struct {
	Event *VoteSetAvatar // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteSetAvatarIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteSetAvatar)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteSetAvatar)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteSetAvatarIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteSetAvatarIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteSetAvatar represents a SetAvatar event raised by the Vote contract.
type VoteSetAvatar struct {
	ProfileId *big.Int
	NewAvatar string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetAvatar is a free log retrieval operation binding the contract event 0x60f52aedb1db78381295534bb468e00f2962aea6c3bbb89eb141a27599ddac0e.
//
// Solidity: event SetAvatar(uint256 indexed profileId, string newAvatar)
func (_Vote *VoteFilterer) FilterSetAvatar(opts *bind.FilterOpts, profileId []*big.Int) (*VoteSetAvatarIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "SetAvatar", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteSetAvatarIterator{contract: _Vote.contract, event: "SetAvatar", logs: logs, sub: sub}, nil
}

// WatchSetAvatar is a free log subscription operation binding the contract event 0x60f52aedb1db78381295534bb468e00f2962aea6c3bbb89eb141a27599ddac0e.
//
// Solidity: event SetAvatar(uint256 indexed profileId, string newAvatar)
func (_Vote *VoteFilterer) WatchSetAvatar(opts *bind.WatchOpts, sink chan<- *VoteSetAvatar, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "SetAvatar", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteSetAvatar)
				if err := _Vote.contract.UnpackLog(event, "SetAvatar", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAvatar is a log parse operation binding the contract event 0x60f52aedb1db78381295534bb468e00f2962aea6c3bbb89eb141a27599ddac0e.
//
// Solidity: event SetAvatar(uint256 indexed profileId, string newAvatar)
func (_Vote *VoteFilterer) ParseSetAvatar(log types.Log) (*VoteSetAvatar, error) {
	event := new(VoteSetAvatar)
	if err := _Vote.contract.UnpackLog(event, "SetAvatar", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteSetEssenceDataIterator is returned from FilterSetEssenceData and is used to iterate over the raw logs and unpacked data for SetEssenceData events raised by the Vote contract.
type VoteSetEssenceDataIterator struct {
	Event *VoteSetEssenceData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteSetEssenceDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteSetEssenceData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteSetEssenceData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteSetEssenceDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteSetEssenceDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteSetEssenceData represents a SetEssenceData event raised by the Vote contract.
type VoteSetEssenceData struct {
	ProfileId         *big.Int
	EssenceId         *big.Int
	TokenURI          string
	Mw                common.Address
	PrepareReturnData []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetEssenceData is a free log retrieval operation binding the contract event 0x1644f271ba3edff28cd2563431920a960fcfebffd068680e682b2de30ca31694.
//
// Solidity: event SetEssenceData(uint256 indexed profileId, uint256 indexed essenceId, string tokenURI, address mw, bytes prepareReturnData)
func (_Vote *VoteFilterer) FilterSetEssenceData(opts *bind.FilterOpts, profileId []*big.Int, essenceId []*big.Int) (*VoteSetEssenceDataIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "SetEssenceData", profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteSetEssenceDataIterator{contract: _Vote.contract, event: "SetEssenceData", logs: logs, sub: sub}, nil
}

// WatchSetEssenceData is a free log subscription operation binding the contract event 0x1644f271ba3edff28cd2563431920a960fcfebffd068680e682b2de30ca31694.
//
// Solidity: event SetEssenceData(uint256 indexed profileId, uint256 indexed essenceId, string tokenURI, address mw, bytes prepareReturnData)
func (_Vote *VoteFilterer) WatchSetEssenceData(opts *bind.WatchOpts, sink chan<- *VoteSetEssenceData, profileId []*big.Int, essenceId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "SetEssenceData", profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteSetEssenceData)
				if err := _Vote.contract.UnpackLog(event, "SetEssenceData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetEssenceData is a log parse operation binding the contract event 0x1644f271ba3edff28cd2563431920a960fcfebffd068680e682b2de30ca31694.
//
// Solidity: event SetEssenceData(uint256 indexed profileId, uint256 indexed essenceId, string tokenURI, address mw, bytes prepareReturnData)
func (_Vote *VoteFilterer) ParseSetEssenceData(log types.Log) (*VoteSetEssenceData, error) {
	event := new(VoteSetEssenceData)
	if err := _Vote.contract.UnpackLog(event, "SetEssenceData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteSetMetadataIterator is returned from FilterSetMetadata and is used to iterate over the raw logs and unpacked data for SetMetadata events raised by the Vote contract.
type VoteSetMetadataIterator struct {
	Event *VoteSetMetadata // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteSetMetadataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteSetMetadata)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteSetMetadata)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteSetMetadataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteSetMetadataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteSetMetadata represents a SetMetadata event raised by the Vote contract.
type VoteSetMetadata struct {
	ProfileId   *big.Int
	NewMetadata string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMetadata is a free log retrieval operation binding the contract event 0xcf080687acfb4df18bbb9959fd7284daa7103f9d1c9164e70999fdedac2a3030.
//
// Solidity: event SetMetadata(uint256 indexed profileId, string newMetadata)
func (_Vote *VoteFilterer) FilterSetMetadata(opts *bind.FilterOpts, profileId []*big.Int) (*VoteSetMetadataIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "SetMetadata", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteSetMetadataIterator{contract: _Vote.contract, event: "SetMetadata", logs: logs, sub: sub}, nil
}

// WatchSetMetadata is a free log subscription operation binding the contract event 0xcf080687acfb4df18bbb9959fd7284daa7103f9d1c9164e70999fdedac2a3030.
//
// Solidity: event SetMetadata(uint256 indexed profileId, string newMetadata)
func (_Vote *VoteFilterer) WatchSetMetadata(opts *bind.WatchOpts, sink chan<- *VoteSetMetadata, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "SetMetadata", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteSetMetadata)
				if err := _Vote.contract.UnpackLog(event, "SetMetadata", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMetadata is a log parse operation binding the contract event 0xcf080687acfb4df18bbb9959fd7284daa7103f9d1c9164e70999fdedac2a3030.
//
// Solidity: event SetMetadata(uint256 indexed profileId, string newMetadata)
func (_Vote *VoteFilterer) ParseSetMetadata(log types.Log) (*VoteSetMetadata, error) {
	event := new(VoteSetMetadata)
	if err := _Vote.contract.UnpackLog(event, "SetMetadata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteSetNFTDescriptorIterator is returned from FilterSetNFTDescriptor and is used to iterate over the raw logs and unpacked data for SetNFTDescriptor events raised by the Vote contract.
type VoteSetNFTDescriptorIterator struct {
	Event *VoteSetNFTDescriptor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteSetNFTDescriptorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteSetNFTDescriptor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteSetNFTDescriptor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteSetNFTDescriptorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteSetNFTDescriptorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteSetNFTDescriptor represents a SetNFTDescriptor event raised by the Vote contract.
type VoteSetNFTDescriptor struct {
	NewDescriptor common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetNFTDescriptor is a free log retrieval operation binding the contract event 0xb72b36bb50852d5d293221b9d1357433e4df37b01e636d3587c2bab7e8531ead.
//
// Solidity: event SetNFTDescriptor(address indexed newDescriptor)
func (_Vote *VoteFilterer) FilterSetNFTDescriptor(opts *bind.FilterOpts, newDescriptor []common.Address) (*VoteSetNFTDescriptorIterator, error) {

	var newDescriptorRule []interface{}
	for _, newDescriptorItem := range newDescriptor {
		newDescriptorRule = append(newDescriptorRule, newDescriptorItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "SetNFTDescriptor", newDescriptorRule)
	if err != nil {
		return nil, err
	}
	return &VoteSetNFTDescriptorIterator{contract: _Vote.contract, event: "SetNFTDescriptor", logs: logs, sub: sub}, nil
}

// WatchSetNFTDescriptor is a free log subscription operation binding the contract event 0xb72b36bb50852d5d293221b9d1357433e4df37b01e636d3587c2bab7e8531ead.
//
// Solidity: event SetNFTDescriptor(address indexed newDescriptor)
func (_Vote *VoteFilterer) WatchSetNFTDescriptor(opts *bind.WatchOpts, sink chan<- *VoteSetNFTDescriptor, newDescriptor []common.Address) (event.Subscription, error) {

	var newDescriptorRule []interface{}
	for _, newDescriptorItem := range newDescriptor {
		newDescriptorRule = append(newDescriptorRule, newDescriptorItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "SetNFTDescriptor", newDescriptorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteSetNFTDescriptor)
				if err := _Vote.contract.UnpackLog(event, "SetNFTDescriptor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetNFTDescriptor is a log parse operation binding the contract event 0xb72b36bb50852d5d293221b9d1357433e4df37b01e636d3587c2bab7e8531ead.
//
// Solidity: event SetNFTDescriptor(address indexed newDescriptor)
func (_Vote *VoteFilterer) ParseSetNFTDescriptor(log types.Log) (*VoteSetNFTDescriptor, error) {
	event := new(VoteSetNFTDescriptor)
	if err := _Vote.contract.UnpackLog(event, "SetNFTDescriptor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteSetNamespaceOwnerIterator is returned from FilterSetNamespaceOwner and is used to iterate over the raw logs and unpacked data for SetNamespaceOwner events raised by the Vote contract.
type VoteSetNamespaceOwnerIterator struct {
	Event *VoteSetNamespaceOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteSetNamespaceOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteSetNamespaceOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteSetNamespaceOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteSetNamespaceOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteSetNamespaceOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteSetNamespaceOwner represents a SetNamespaceOwner event raised by the Vote contract.
type VoteSetNamespaceOwner struct {
	PreOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetNamespaceOwner is a free log retrieval operation binding the contract event 0x46e05b7612d0b60892e73d72895fb0752deda3fa69984aca10c3cc72889ef7df.
//
// Solidity: event SetNamespaceOwner(address indexed preOwner, address indexed newOwner)
func (_Vote *VoteFilterer) FilterSetNamespaceOwner(opts *bind.FilterOpts, preOwner []common.Address, newOwner []common.Address) (*VoteSetNamespaceOwnerIterator, error) {

	var preOwnerRule []interface{}
	for _, preOwnerItem := range preOwner {
		preOwnerRule = append(preOwnerRule, preOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "SetNamespaceOwner", preOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VoteSetNamespaceOwnerIterator{contract: _Vote.contract, event: "SetNamespaceOwner", logs: logs, sub: sub}, nil
}

// WatchSetNamespaceOwner is a free log subscription operation binding the contract event 0x46e05b7612d0b60892e73d72895fb0752deda3fa69984aca10c3cc72889ef7df.
//
// Solidity: event SetNamespaceOwner(address indexed preOwner, address indexed newOwner)
func (_Vote *VoteFilterer) WatchSetNamespaceOwner(opts *bind.WatchOpts, sink chan<- *VoteSetNamespaceOwner, preOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var preOwnerRule []interface{}
	for _, preOwnerItem := range preOwner {
		preOwnerRule = append(preOwnerRule, preOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "SetNamespaceOwner", preOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteSetNamespaceOwner)
				if err := _Vote.contract.UnpackLog(event, "SetNamespaceOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetNamespaceOwner is a log parse operation binding the contract event 0x46e05b7612d0b60892e73d72895fb0752deda3fa69984aca10c3cc72889ef7df.
//
// Solidity: event SetNamespaceOwner(address indexed preOwner, address indexed newOwner)
func (_Vote *VoteFilterer) ParseSetNamespaceOwner(log types.Log) (*VoteSetNamespaceOwner, error) {
	event := new(VoteSetNamespaceOwner)
	if err := _Vote.contract.UnpackLog(event, "SetNamespaceOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteSetOperatorApprovalIterator is returned from FilterSetOperatorApproval and is used to iterate over the raw logs and unpacked data for SetOperatorApproval events raised by the Vote contract.
type VoteSetOperatorApprovalIterator struct {
	Event *VoteSetOperatorApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteSetOperatorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteSetOperatorApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteSetOperatorApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteSetOperatorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteSetOperatorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteSetOperatorApproval represents a SetOperatorApproval event raised by the Vote contract.
type VoteSetOperatorApproval struct {
	ProfileId    *big.Int
	Operator     common.Address
	PrevApproved bool
	Approved     bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetOperatorApproval is a free log retrieval operation binding the contract event 0x470544bc4035b5a73e4bb39628742f459c69af3a00b7568beff4bcb76aa51b6e.
//
// Solidity: event SetOperatorApproval(uint256 indexed profileId, address indexed operator, bool prevApproved, bool approved)
func (_Vote *VoteFilterer) FilterSetOperatorApproval(opts *bind.FilterOpts, profileId []*big.Int, operator []common.Address) (*VoteSetOperatorApprovalIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "SetOperatorApproval", profileIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &VoteSetOperatorApprovalIterator{contract: _Vote.contract, event: "SetOperatorApproval", logs: logs, sub: sub}, nil
}

// WatchSetOperatorApproval is a free log subscription operation binding the contract event 0x470544bc4035b5a73e4bb39628742f459c69af3a00b7568beff4bcb76aa51b6e.
//
// Solidity: event SetOperatorApproval(uint256 indexed profileId, address indexed operator, bool prevApproved, bool approved)
func (_Vote *VoteFilterer) WatchSetOperatorApproval(opts *bind.WatchOpts, sink chan<- *VoteSetOperatorApproval, profileId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "SetOperatorApproval", profileIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteSetOperatorApproval)
				if err := _Vote.contract.UnpackLog(event, "SetOperatorApproval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetOperatorApproval is a log parse operation binding the contract event 0x470544bc4035b5a73e4bb39628742f459c69af3a00b7568beff4bcb76aa51b6e.
//
// Solidity: event SetOperatorApproval(uint256 indexed profileId, address indexed operator, bool prevApproved, bool approved)
func (_Vote *VoteFilterer) ParseSetOperatorApproval(log types.Log) (*VoteSetOperatorApproval, error) {
	event := new(VoteSetOperatorApproval)
	if err := _Vote.contract.UnpackLog(event, "SetOperatorApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteSetPrimaryProfileIterator is returned from FilterSetPrimaryProfile and is used to iterate over the raw logs and unpacked data for SetPrimaryProfile events raised by the Vote contract.
type VoteSetPrimaryProfileIterator struct {
	Event *VoteSetPrimaryProfile // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteSetPrimaryProfileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteSetPrimaryProfile)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteSetPrimaryProfile)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteSetPrimaryProfileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteSetPrimaryProfileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteSetPrimaryProfile represents a SetPrimaryProfile event raised by the Vote contract.
type VoteSetPrimaryProfile struct {
	User      common.Address
	ProfileId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetPrimaryProfile is a free log retrieval operation binding the contract event 0xdf187b01e55e0f335a75956346e6d4417e6ea27aed3bc1bf2106d9d60bdc50d0.
//
// Solidity: event SetPrimaryProfile(address indexed user, uint256 indexed profileId)
func (_Vote *VoteFilterer) FilterSetPrimaryProfile(opts *bind.FilterOpts, user []common.Address, profileId []*big.Int) (*VoteSetPrimaryProfileIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "SetPrimaryProfile", userRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteSetPrimaryProfileIterator{contract: _Vote.contract, event: "SetPrimaryProfile", logs: logs, sub: sub}, nil
}

// WatchSetPrimaryProfile is a free log subscription operation binding the contract event 0xdf187b01e55e0f335a75956346e6d4417e6ea27aed3bc1bf2106d9d60bdc50d0.
//
// Solidity: event SetPrimaryProfile(address indexed user, uint256 indexed profileId)
func (_Vote *VoteFilterer) WatchSetPrimaryProfile(opts *bind.WatchOpts, sink chan<- *VoteSetPrimaryProfile, user []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "SetPrimaryProfile", userRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteSetPrimaryProfile)
				if err := _Vote.contract.UnpackLog(event, "SetPrimaryProfile", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPrimaryProfile is a log parse operation binding the contract event 0xdf187b01e55e0f335a75956346e6d4417e6ea27aed3bc1bf2106d9d60bdc50d0.
//
// Solidity: event SetPrimaryProfile(address indexed user, uint256 indexed profileId)
func (_Vote *VoteFilterer) ParseSetPrimaryProfile(log types.Log) (*VoteSetPrimaryProfile, error) {
	event := new(VoteSetPrimaryProfile)
	if err := _Vote.contract.UnpackLog(event, "SetPrimaryProfile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteSetSubscribeDataIterator is returned from FilterSetSubscribeData and is used to iterate over the raw logs and unpacked data for SetSubscribeData events raised by the Vote contract.
type VoteSetSubscribeDataIterator struct {
	Event *VoteSetSubscribeData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteSetSubscribeDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteSetSubscribeData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteSetSubscribeData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteSetSubscribeDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteSetSubscribeDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteSetSubscribeData represents a SetSubscribeData event raised by the Vote contract.
type VoteSetSubscribeData struct {
	ProfileId         *big.Int
	TokenURI          string
	Mw                common.Address
	PrepareReturnData []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetSubscribeData is a free log retrieval operation binding the contract event 0x56e2b6a9becdb5836c1184738ad1cddd0096556aa7ada94493a4a4feb117bd86.
//
// Solidity: event SetSubscribeData(uint256 indexed profileId, string tokenURI, address mw, bytes prepareReturnData)
func (_Vote *VoteFilterer) FilterSetSubscribeData(opts *bind.FilterOpts, profileId []*big.Int) (*VoteSetSubscribeDataIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "SetSubscribeData", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteSetSubscribeDataIterator{contract: _Vote.contract, event: "SetSubscribeData", logs: logs, sub: sub}, nil
}

// WatchSetSubscribeData is a free log subscription operation binding the contract event 0x56e2b6a9becdb5836c1184738ad1cddd0096556aa7ada94493a4a4feb117bd86.
//
// Solidity: event SetSubscribeData(uint256 indexed profileId, string tokenURI, address mw, bytes prepareReturnData)
func (_Vote *VoteFilterer) WatchSetSubscribeData(opts *bind.WatchOpts, sink chan<- *VoteSetSubscribeData, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "SetSubscribeData", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteSetSubscribeData)
				if err := _Vote.contract.UnpackLog(event, "SetSubscribeData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSubscribeData is a log parse operation binding the contract event 0x56e2b6a9becdb5836c1184738ad1cddd0096556aa7ada94493a4a4feb117bd86.
//
// Solidity: event SetSubscribeData(uint256 indexed profileId, string tokenURI, address mw, bytes prepareReturnData)
func (_Vote *VoteFilterer) ParseSetSubscribeData(log types.Log) (*VoteSetSubscribeData, error) {
	event := new(VoteSetSubscribeData)
	if err := _Vote.contract.UnpackLog(event, "SetSubscribeData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteSubscribeIterator is returned from FilterSubscribe and is used to iterate over the raw logs and unpacked data for Subscribe events raised by the Vote contract.
type VoteSubscribeIterator struct {
	Event *VoteSubscribe // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteSubscribeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteSubscribe)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteSubscribe)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteSubscribeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteSubscribeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteSubscribe represents a Subscribe event raised by the Vote contract.
type VoteSubscribe struct {
	Sender     common.Address
	ProfileIds []*big.Int
	PreDatas   [][]byte
	PostDatas  [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubscribe is a free log retrieval operation binding the contract event 0xcda1d04cbb59c6063db6682ab282a88dc856e841cd4d5562471f2f8bbcbd6d6b.
//
// Solidity: event Subscribe(address indexed sender, uint256[] profileIds, bytes[] preDatas, bytes[] postDatas)
func (_Vote *VoteFilterer) FilterSubscribe(opts *bind.FilterOpts, sender []common.Address) (*VoteSubscribeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "Subscribe", senderRule)
	if err != nil {
		return nil, err
	}
	return &VoteSubscribeIterator{contract: _Vote.contract, event: "Subscribe", logs: logs, sub: sub}, nil
}

// WatchSubscribe is a free log subscription operation binding the contract event 0xcda1d04cbb59c6063db6682ab282a88dc856e841cd4d5562471f2f8bbcbd6d6b.
//
// Solidity: event Subscribe(address indexed sender, uint256[] profileIds, bytes[] preDatas, bytes[] postDatas)
func (_Vote *VoteFilterer) WatchSubscribe(opts *bind.WatchOpts, sink chan<- *VoteSubscribe, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "Subscribe", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteSubscribe)
				if err := _Vote.contract.UnpackLog(event, "Subscribe", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscribe is a log parse operation binding the contract event 0xcda1d04cbb59c6063db6682ab282a88dc856e841cd4d5562471f2f8bbcbd6d6b.
//
// Solidity: event Subscribe(address indexed sender, uint256[] profileIds, bytes[] preDatas, bytes[] postDatas)
func (_Vote *VoteFilterer) ParseSubscribe(log types.Log) (*VoteSubscribe, error) {
	event := new(VoteSubscribe)
	if err := _Vote.contract.UnpackLog(event, "Subscribe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vote contract.
type VoteTransferIterator struct {
	Event *VoteTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteTransfer represents a Transfer event raised by the Vote contract.
type VoteTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Vote *VoteFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*VoteTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &VoteTransferIterator{contract: _Vote.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Vote *VoteFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VoteTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteTransfer)
				if err := _Vote.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Vote *VoteFilterer) ParseTransfer(log types.Log) (*VoteTransfer, error) {
	event := new(VoteTransfer)
	if err := _Vote.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Vote contract.
type VoteUnpausedIterator struct {
	Event *VoteUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteUnpaused represents a Unpaused event raised by the Vote contract.
type VoteUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vote *VoteFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VoteUnpausedIterator, error) {

	logs, sub, err := _Vote.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VoteUnpausedIterator{contract: _Vote.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vote *VoteFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VoteUnpaused) (event.Subscription, error) {

	logs, sub, err := _Vote.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteUnpaused)
				if err := _Vote.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vote *VoteFilterer) ParseUnpaused(log types.Log) (*VoteUnpaused, error) {
	event := new(VoteUnpaused)
	if err := _Vote.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Vote contract.
type VoteUpgradedIterator struct {
	Event *VoteUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoteUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VoteUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VoteUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteUpgraded represents a Upgraded event raised by the Vote contract.
type VoteUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vote *VoteFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VoteUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VoteUpgradedIterator{contract: _Vote.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vote *VoteFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VoteUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteUpgraded)
				if err := _Vote.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vote *VoteFilterer) ParseUpgraded(log types.Log) (*VoteUpgraded, error) {
	event := new(VoteUpgraded)
	if err := _Vote.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
