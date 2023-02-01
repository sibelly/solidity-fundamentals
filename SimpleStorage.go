// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addedNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"storedNumber\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"favoriteNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_favoriteNumber\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156100115760006000fd5b50610017565b61017b806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100465760003560e01c80632e64cec11461004c578063471f7cdf1461006a5780636057361d1461008857610046565b60006000fd5b6100546100b7565b6040518082815260200191505060405180910390f35b6100726100c9565b6040518082815260200191505060405180910390f35b6100b56004803603602081101561009f5760006000fd5b81019080803590602001909291905050506100d2565b005b600060006000505490506100c6565b90565b60006000505481565b806000600050547fe304654f140653267f4d09df4968c896d7d6fd0b8e23c45f1db26a86cf930001600060005054840133604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a38060006000508190909055505b5056fea2646970667358221220262d78206a2ae3c0be22c61e70445a0bbf6da1a0dcf27a5c7fadeef04d55e99564736f6c63430007030033",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainMetaData.Bin instead.
var MainBin = MainMetaData.Bin

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// FavoriteNumber is a free data retrieval call binding the contract method 0x471f7cdf.
//
// Solidity: function favoriteNumber() view returns(uint256)
func (_Main *MainCaller) FavoriteNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "favoriteNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FavoriteNumber is a free data retrieval call binding the contract method 0x471f7cdf.
//
// Solidity: function favoriteNumber() view returns(uint256)
func (_Main *MainSession) FavoriteNumber() (*big.Int, error) {
	return _Main.Contract.FavoriteNumber(&_Main.CallOpts)
}

// FavoriteNumber is a free data retrieval call binding the contract method 0x471f7cdf.
//
// Solidity: function favoriteNumber() view returns(uint256)
func (_Main *MainCallerSession) FavoriteNumber() (*big.Int, error) {
	return _Main.Contract.FavoriteNumber(&_Main.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Main *MainCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Main *MainSession) Retrieve() (*big.Int, error) {
	return _Main.Contract.Retrieve(&_Main.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Main *MainCallerSession) Retrieve() (*big.Int, error) {
	return _Main.Contract.Retrieve(&_Main.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_Main *MainTransactor) Store(opts *bind.TransactOpts, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "store", _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_Main *MainSession) Store(_favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Store(&_Main.TransactOpts, _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_Main *MainTransactorSession) Store(_favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Store(&_Main.TransactOpts, _favoriteNumber)
}

// MainStoredNumberIterator is returned from FilterStoredNumber and is used to iterate over the raw logs and unpacked data for StoredNumber events raised by the Main contract.
type MainStoredNumberIterator struct {
	Event *MainStoredNumber // Event containing the contract specifics and raw log

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
func (it *MainStoredNumberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainStoredNumber)
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
		it.Event = new(MainStoredNumber)
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
func (it *MainStoredNumberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainStoredNumberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainStoredNumber represents a StoredNumber event raised by the Main contract.
type MainStoredNumber struct {
	OldNumber   *big.Int
	NewNumber   *big.Int
	AddedNumber *big.Int
	Sender      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStoredNumber is a free log retrieval operation binding the contract event 0xe304654f140653267f4d09df4968c896d7d6fd0b8e23c45f1db26a86cf930001.
//
// Solidity: event storedNumber(uint256 indexed oldNumber, uint256 indexed newNumber, uint256 addedNumber, address sender)
func (_Main *MainFilterer) FilterStoredNumber(opts *bind.FilterOpts, oldNumber []*big.Int, newNumber []*big.Int) (*MainStoredNumberIterator, error) {

	var oldNumberRule []interface{}
	for _, oldNumberItem := range oldNumber {
		oldNumberRule = append(oldNumberRule, oldNumberItem)
	}
	var newNumberRule []interface{}
	for _, newNumberItem := range newNumber {
		newNumberRule = append(newNumberRule, newNumberItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "storedNumber", oldNumberRule, newNumberRule)
	if err != nil {
		return nil, err
	}
	return &MainStoredNumberIterator{contract: _Main.contract, event: "storedNumber", logs: logs, sub: sub}, nil
}

// WatchStoredNumber is a free log subscription operation binding the contract event 0xe304654f140653267f4d09df4968c896d7d6fd0b8e23c45f1db26a86cf930001.
//
// Solidity: event storedNumber(uint256 indexed oldNumber, uint256 indexed newNumber, uint256 addedNumber, address sender)
func (_Main *MainFilterer) WatchStoredNumber(opts *bind.WatchOpts, sink chan<- *MainStoredNumber, oldNumber []*big.Int, newNumber []*big.Int) (event.Subscription, error) {

	var oldNumberRule []interface{}
	for _, oldNumberItem := range oldNumber {
		oldNumberRule = append(oldNumberRule, oldNumberItem)
	}
	var newNumberRule []interface{}
	for _, newNumberItem := range newNumber {
		newNumberRule = append(newNumberRule, newNumberItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "storedNumber", oldNumberRule, newNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainStoredNumber)
				if err := _Main.contract.UnpackLog(event, "storedNumber", log); err != nil {
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

// ParseStoredNumber is a log parse operation binding the contract event 0xe304654f140653267f4d09df4968c896d7d6fd0b8e23c45f1db26a86cf930001.
//
// Solidity: event storedNumber(uint256 indexed oldNumber, uint256 indexed newNumber, uint256 addedNumber, address sender)
func (_Main *MainFilterer) ParseStoredNumber(log types.Log) (*MainStoredNumber, error) {
	event := new(MainStoredNumber)
	if err := _Main.contract.UnpackLog(event, "storedNumber", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
