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

// TypeMetaData contains all meta data concerning the Type contract.
var TypeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"namespace\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"CollectPermissionMwSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"postProcess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"preProcess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"essenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setEssenceMwData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TypeABI is the input ABI used to generate the binding from.
// Deprecated: Use TypeMetaData.ABI instead.
var TypeABI = TypeMetaData.ABI

// Type is an auto generated Go binding around an Ethereum contract.
type Type struct {
	TypeCaller     // Read-only binding to the contract
	TypeTransactor // Write-only binding to the contract
	TypeFilterer   // Log filterer for contract events
}

// TypeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TypeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TypeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TypeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TypeSession struct {
	Contract     *Type             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TypeCallerSession struct {
	Contract *TypeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TypeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TypeTransactorSession struct {
	Contract     *TypeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TypeRaw struct {
	Contract *Type // Generic contract binding to access the raw methods on
}

// TypeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TypeCallerRaw struct {
	Contract *TypeCaller // Generic read-only contract binding to access the raw methods on
}

// TypeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TypeTransactorRaw struct {
	Contract *TypeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewType creates a new instance of Type, bound to a specific deployed contract.
func NewType(address common.Address, backend bind.ContractBackend) (*Type, error) {
	contract, err := bindType(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Type{TypeCaller: TypeCaller{contract: contract}, TypeTransactor: TypeTransactor{contract: contract}, TypeFilterer: TypeFilterer{contract: contract}}, nil
}

// NewTypeCaller creates a new read-only instance of Type, bound to a specific deployed contract.
func NewTypeCaller(address common.Address, caller bind.ContractCaller) (*TypeCaller, error) {
	contract, err := bindType(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypeCaller{contract: contract}, nil
}

// NewTypeTransactor creates a new write-only instance of Type, bound to a specific deployed contract.
func NewTypeTransactor(address common.Address, transactor bind.ContractTransactor) (*TypeTransactor, error) {
	contract, err := bindType(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypeTransactor{contract: contract}, nil
}

// NewTypeFilterer creates a new log filterer instance of Type, bound to a specific deployed contract.
func NewTypeFilterer(address common.Address, filterer bind.ContractFilterer) (*TypeFilterer, error) {
	contract, err := bindType(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypeFilterer{contract: contract}, nil
}

// bindType binds a generic wrapper to an already deployed contract.
func bindType(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Type *TypeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Type.Contract.TypeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Type *TypeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Type.Contract.TypeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Type *TypeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Type.Contract.TypeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Type *TypeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Type.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Type *TypeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Type.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Type *TypeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Type.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Type *TypeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Type.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Type *TypeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Type.Contract.DOMAINSEPARATOR(&_Type.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Type *TypeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Type.Contract.DOMAINSEPARATOR(&_Type.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x48d7b156.
//
// Solidity: function getNonce(uint256 profileId, address collector, uint256 essenceId) view returns(uint256)
func (_Type *TypeCaller) GetNonce(opts *bind.CallOpts, profileId *big.Int, collector common.Address, essenceId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Type.contract.Call(opts, &out, "getNonce", profileId, collector, essenceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x48d7b156.
//
// Solidity: function getNonce(uint256 profileId, address collector, uint256 essenceId) view returns(uint256)
func (_Type *TypeSession) GetNonce(profileId *big.Int, collector common.Address, essenceId *big.Int) (*big.Int, error) {
	return _Type.Contract.GetNonce(&_Type.CallOpts, profileId, collector, essenceId)
}

// GetNonce is a free data retrieval call binding the contract method 0x48d7b156.
//
// Solidity: function getNonce(uint256 profileId, address collector, uint256 essenceId) view returns(uint256)
func (_Type *TypeCallerSession) GetNonce(profileId *big.Int, collector common.Address, essenceId *big.Int) (*big.Int, error) {
	return _Type.Contract.GetNonce(&_Type.CallOpts, profileId, collector, essenceId)
}

// PostProcess is a paid mutator transaction binding the contract method 0x2eae9262.
//
// Solidity: function postProcess(uint256 , uint256 , address , address , bytes ) returns()
func (_Type *TypeTransactor) PostProcess(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _Type.contract.Transact(opts, "postProcess", arg0, arg1, arg2, arg3, arg4)
}

// PostProcess is a paid mutator transaction binding the contract method 0x2eae9262.
//
// Solidity: function postProcess(uint256 , uint256 , address , address , bytes ) returns()
func (_Type *TypeSession) PostProcess(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _Type.Contract.PostProcess(&_Type.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// PostProcess is a paid mutator transaction binding the contract method 0x2eae9262.
//
// Solidity: function postProcess(uint256 , uint256 , address , address , bytes ) returns()
func (_Type *TypeTransactorSession) PostProcess(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _Type.Contract.PostProcess(&_Type.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// PreProcess is a paid mutator transaction binding the contract method 0xf2dd081b.
//
// Solidity: function preProcess(uint256 profileId, uint256 essenceId, address collector, address , bytes data) returns()
func (_Type *TypeTransactor) PreProcess(opts *bind.TransactOpts, profileId *big.Int, essenceId *big.Int, collector common.Address, arg3 common.Address, data []byte) (*types.Transaction, error) {
	return _Type.contract.Transact(opts, "preProcess", profileId, essenceId, collector, arg3, data)
}

// PreProcess is a paid mutator transaction binding the contract method 0xf2dd081b.
//
// Solidity: function preProcess(uint256 profileId, uint256 essenceId, address collector, address , bytes data) returns()
func (_Type *TypeSession) PreProcess(profileId *big.Int, essenceId *big.Int, collector common.Address, arg3 common.Address, data []byte) (*types.Transaction, error) {
	return _Type.Contract.PreProcess(&_Type.TransactOpts, profileId, essenceId, collector, arg3, data)
}

// PreProcess is a paid mutator transaction binding the contract method 0xf2dd081b.
//
// Solidity: function preProcess(uint256 profileId, uint256 essenceId, address collector, address , bytes data) returns()
func (_Type *TypeTransactorSession) PreProcess(profileId *big.Int, essenceId *big.Int, collector common.Address, arg3 common.Address, data []byte) (*types.Transaction, error) {
	return _Type.Contract.PreProcess(&_Type.TransactOpts, profileId, essenceId, collector, arg3, data)
}

// SetEssenceMwData is a paid mutator transaction binding the contract method 0xa58d4f0e.
//
// Solidity: function setEssenceMwData(uint256 profileId, uint256 essenceId, bytes data) returns(bytes)
func (_Type *TypeTransactor) SetEssenceMwData(opts *bind.TransactOpts, profileId *big.Int, essenceId *big.Int, data []byte) (*types.Transaction, error) {
	return _Type.contract.Transact(opts, "setEssenceMwData", profileId, essenceId, data)
}

// SetEssenceMwData is a paid mutator transaction binding the contract method 0xa58d4f0e.
//
// Solidity: function setEssenceMwData(uint256 profileId, uint256 essenceId, bytes data) returns(bytes)
func (_Type *TypeSession) SetEssenceMwData(profileId *big.Int, essenceId *big.Int, data []byte) (*types.Transaction, error) {
	return _Type.Contract.SetEssenceMwData(&_Type.TransactOpts, profileId, essenceId, data)
}

// SetEssenceMwData is a paid mutator transaction binding the contract method 0xa58d4f0e.
//
// Solidity: function setEssenceMwData(uint256 profileId, uint256 essenceId, bytes data) returns(bytes)
func (_Type *TypeTransactorSession) SetEssenceMwData(profileId *big.Int, essenceId *big.Int, data []byte) (*types.Transaction, error) {
	return _Type.Contract.SetEssenceMwData(&_Type.TransactOpts, profileId, essenceId, data)
}

// TypeCollectPermissionMwSetIterator is returned from FilterCollectPermissionMwSet and is used to iterate over the raw logs and unpacked data for CollectPermissionMwSet events raised by the Type contract.
type TypeCollectPermissionMwSetIterator struct {
	Event *TypeCollectPermissionMwSet // Event containing the contract specifics and raw log

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
func (it *TypeCollectPermissionMwSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TypeCollectPermissionMwSet)
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
		it.Event = new(TypeCollectPermissionMwSet)
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
func (it *TypeCollectPermissionMwSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TypeCollectPermissionMwSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TypeCollectPermissionMwSet represents a CollectPermissionMwSet event raised by the Type contract.
type TypeCollectPermissionMwSet struct {
	Namespace common.Address
	ProfileId *big.Int
	EssenceId *big.Int
	Signer    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectPermissionMwSet is a free log retrieval operation binding the contract event 0x275238907a05de8c78cfde06762f5ee0b05711cd4310c59d90200e07c31b8017.
//
// Solidity: event CollectPermissionMwSet(address indexed namespace, uint256 indexed profileId, uint256 indexed essenceId, address signer)
func (_Type *TypeFilterer) FilterCollectPermissionMwSet(opts *bind.FilterOpts, namespace []common.Address, profileId []*big.Int, essenceId []*big.Int) (*TypeCollectPermissionMwSetIterator, error) {

	var namespaceRule []interface{}
	for _, namespaceItem := range namespace {
		namespaceRule = append(namespaceRule, namespaceItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _Type.contract.FilterLogs(opts, "CollectPermissionMwSet", namespaceRule, profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return &TypeCollectPermissionMwSetIterator{contract: _Type.contract, event: "CollectPermissionMwSet", logs: logs, sub: sub}, nil
}

// WatchCollectPermissionMwSet is a free log subscription operation binding the contract event 0x275238907a05de8c78cfde06762f5ee0b05711cd4310c59d90200e07c31b8017.
//
// Solidity: event CollectPermissionMwSet(address indexed namespace, uint256 indexed profileId, uint256 indexed essenceId, address signer)
func (_Type *TypeFilterer) WatchCollectPermissionMwSet(opts *bind.WatchOpts, sink chan<- *TypeCollectPermissionMwSet, namespace []common.Address, profileId []*big.Int, essenceId []*big.Int) (event.Subscription, error) {

	var namespaceRule []interface{}
	for _, namespaceItem := range namespace {
		namespaceRule = append(namespaceRule, namespaceItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var essenceIdRule []interface{}
	for _, essenceIdItem := range essenceId {
		essenceIdRule = append(essenceIdRule, essenceIdItem)
	}

	logs, sub, err := _Type.contract.WatchLogs(opts, "CollectPermissionMwSet", namespaceRule, profileIdRule, essenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TypeCollectPermissionMwSet)
				if err := _Type.contract.UnpackLog(event, "CollectPermissionMwSet", log); err != nil {
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

// ParseCollectPermissionMwSet is a log parse operation binding the contract event 0x275238907a05de8c78cfde06762f5ee0b05711cd4310c59d90200e07c31b8017.
//
// Solidity: event CollectPermissionMwSet(address indexed namespace, uint256 indexed profileId, uint256 indexed essenceId, address signer)
func (_Type *TypeFilterer) ParseCollectPermissionMwSet(log types.Log) (*TypeCollectPermissionMwSet, error) {
	event := new(TypeCollectPermissionMwSet)
	if err := _Type.contract.UnpackLog(event, "CollectPermissionMwSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
