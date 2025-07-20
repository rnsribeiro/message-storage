// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package message

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
	_ = abi.ConvertType
)

// MessageStorageMetaData contains all meta data concerning the MessageStorage contract.
var MessageStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMessage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"setMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MessageStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageStorageMetaData.ABI instead.
var MessageStorageABI = MessageStorageMetaData.ABI

// MessageStorage is an auto generated Go binding around an Ethereum contract.
type MessageStorage struct {
	MessageStorageCaller     // Read-only binding to the contract
	MessageStorageTransactor // Write-only binding to the contract
	MessageStorageFilterer   // Log filterer for contract events
}

// MessageStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageStorageSession struct {
	Contract     *MessageStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageStorageCallerSession struct {
	Contract *MessageStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MessageStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageStorageTransactorSession struct {
	Contract     *MessageStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MessageStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageStorageRaw struct {
	Contract *MessageStorage // Generic contract binding to access the raw methods on
}

// MessageStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageStorageCallerRaw struct {
	Contract *MessageStorageCaller // Generic read-only contract binding to access the raw methods on
}

// MessageStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageStorageTransactorRaw struct {
	Contract *MessageStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageStorage creates a new instance of MessageStorage, bound to a specific deployed contract.
func NewMessageStorage(address common.Address, backend bind.ContractBackend) (*MessageStorage, error) {
	contract, err := bindMessageStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageStorage{MessageStorageCaller: MessageStorageCaller{contract: contract}, MessageStorageTransactor: MessageStorageTransactor{contract: contract}, MessageStorageFilterer: MessageStorageFilterer{contract: contract}}, nil
}

// NewMessageStorageCaller creates a new read-only instance of MessageStorage, bound to a specific deployed contract.
func NewMessageStorageCaller(address common.Address, caller bind.ContractCaller) (*MessageStorageCaller, error) {
	contract, err := bindMessageStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageStorageCaller{contract: contract}, nil
}

// NewMessageStorageTransactor creates a new write-only instance of MessageStorage, bound to a specific deployed contract.
func NewMessageStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageStorageTransactor, error) {
	contract, err := bindMessageStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageStorageTransactor{contract: contract}, nil
}

// NewMessageStorageFilterer creates a new log filterer instance of MessageStorage, bound to a specific deployed contract.
func NewMessageStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageStorageFilterer, error) {
	contract, err := bindMessageStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageStorageFilterer{contract: contract}, nil
}

// bindMessageStorage binds a generic wrapper to an already deployed contract.
func bindMessageStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageStorage *MessageStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageStorage.Contract.MessageStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageStorage *MessageStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageStorage.Contract.MessageStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageStorage *MessageStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageStorage.Contract.MessageStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageStorage *MessageStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageStorage *MessageStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageStorage *MessageStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageStorage.Contract.contract.Transact(opts, method, params...)
}

// GetMessage is a free data retrieval call binding the contract method 0x86f79edb.
//
// Solidity: function getMessage(uint256 _id) view returns(string)
func (_MessageStorage *MessageStorageCaller) GetMessage(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _MessageStorage.contract.Call(opts, &out, "getMessage", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMessage is a free data retrieval call binding the contract method 0x86f79edb.
//
// Solidity: function getMessage(uint256 _id) view returns(string)
func (_MessageStorage *MessageStorageSession) GetMessage(_id *big.Int) (string, error) {
	return _MessageStorage.Contract.GetMessage(&_MessageStorage.CallOpts, _id)
}

// GetMessage is a free data retrieval call binding the contract method 0x86f79edb.
//
// Solidity: function getMessage(uint256 _id) view returns(string)
func (_MessageStorage *MessageStorageCallerSession) GetMessage(_id *big.Int) (string, error) {
	return _MessageStorage.Contract.GetMessage(&_MessageStorage.CallOpts, _id)
}

// GetTotalMessages is a free data retrieval call binding the contract method 0x147e9108.
//
// Solidity: function getTotalMessages() view returns(uint256)
func (_MessageStorage *MessageStorageCaller) GetTotalMessages(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageStorage.contract.Call(opts, &out, "getTotalMessages")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalMessages is a free data retrieval call binding the contract method 0x147e9108.
//
// Solidity: function getTotalMessages() view returns(uint256)
func (_MessageStorage *MessageStorageSession) GetTotalMessages() (*big.Int, error) {
	return _MessageStorage.Contract.GetTotalMessages(&_MessageStorage.CallOpts)
}

// GetTotalMessages is a free data retrieval call binding the contract method 0x147e9108.
//
// Solidity: function getTotalMessages() view returns(uint256)
func (_MessageStorage *MessageStorageCallerSession) GetTotalMessages() (*big.Int, error) {
	return _MessageStorage.Contract.GetTotalMessages(&_MessageStorage.CallOpts)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string _message) returns(uint256)
func (_MessageStorage *MessageStorageTransactor) SetMessage(opts *bind.TransactOpts, _message string) (*types.Transaction, error) {
	return _MessageStorage.contract.Transact(opts, "setMessage", _message)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string _message) returns(uint256)
func (_MessageStorage *MessageStorageSession) SetMessage(_message string) (*types.Transaction, error) {
	return _MessageStorage.Contract.SetMessage(&_MessageStorage.TransactOpts, _message)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string _message) returns(uint256)
func (_MessageStorage *MessageStorageTransactorSession) SetMessage(_message string) (*types.Transaction, error) {
	return _MessageStorage.Contract.SetMessage(&_MessageStorage.TransactOpts, _message)
}
