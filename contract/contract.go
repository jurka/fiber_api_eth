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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGroupIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_groupId\",\"type\":\"uint256\"}],\"name\":\"getGroup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_indexId\",\"type\":\"uint256\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ethPriceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdPriceInCents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdCapitalization\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"percentageChange\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 _groupId) view returns(string name, uint256[] indexes)
func (_Contract *ContractCaller) GetGroup(opts *bind.CallOpts, _groupId *big.Int) (struct {
	Name    string
	Indexes []*big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGroup", _groupId)

	outstruct := new(struct {
		Name    string
		Indexes []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Indexes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 _groupId) view returns(string name, uint256[] indexes)
func (_Contract *ContractSession) GetGroup(_groupId *big.Int) (struct {
	Name    string
	Indexes []*big.Int
}, error) {
	return _Contract.Contract.GetGroup(&_Contract.CallOpts, _groupId)
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 _groupId) view returns(string name, uint256[] indexes)
func (_Contract *ContractCallerSession) GetGroup(_groupId *big.Int) (struct {
	Name    string
	Indexes []*big.Int
}, error) {
	return _Contract.Contract.GetGroup(&_Contract.CallOpts, _groupId)
}

// GetGroupIds is a free data retrieval call binding the contract method 0xc02b8b2c.
//
// Solidity: function getGroupIds() view returns(uint256[])
func (_Contract *ContractCaller) GetGroupIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGroupIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetGroupIds is a free data retrieval call binding the contract method 0xc02b8b2c.
//
// Solidity: function getGroupIds() view returns(uint256[])
func (_Contract *ContractSession) GetGroupIds() ([]*big.Int, error) {
	return _Contract.Contract.GetGroupIds(&_Contract.CallOpts)
}

// GetGroupIds is a free data retrieval call binding the contract method 0xc02b8b2c.
//
// Solidity: function getGroupIds() view returns(uint256[])
func (_Contract *ContractCallerSession) GetGroupIds() ([]*big.Int, error) {
	return _Contract.Contract.GetGroupIds(&_Contract.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0x8e7cb6e1.
//
// Solidity: function getIndex(uint256 _indexId) view returns(string name, uint256 ethPriceInWei, uint256 usdPriceInCents, uint256 usdCapitalization, int256 percentageChange)
func (_Contract *ContractCaller) GetIndex(opts *bind.CallOpts, _indexId *big.Int) (struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getIndex", _indexId)

	outstruct := new(struct {
		Name              string
		EthPriceInWei     *big.Int
		UsdPriceInCents   *big.Int
		UsdCapitalization *big.Int
		PercentageChange  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.EthPriceInWei = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UsdPriceInCents = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UsdCapitalization = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PercentageChange = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetIndex is a free data retrieval call binding the contract method 0x8e7cb6e1.
//
// Solidity: function getIndex(uint256 _indexId) view returns(string name, uint256 ethPriceInWei, uint256 usdPriceInCents, uint256 usdCapitalization, int256 percentageChange)
func (_Contract *ContractSession) GetIndex(_indexId *big.Int) (struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}, error) {
	return _Contract.Contract.GetIndex(&_Contract.CallOpts, _indexId)
}

// GetIndex is a free data retrieval call binding the contract method 0x8e7cb6e1.
//
// Solidity: function getIndex(uint256 _indexId) view returns(string name, uint256 ethPriceInWei, uint256 usdPriceInCents, uint256 usdCapitalization, int256 percentageChange)
func (_Contract *ContractCallerSession) GetIndex(_indexId *big.Int) (struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}, error) {
	return _Contract.Contract.GetIndex(&_Contract.CallOpts, _indexId)
}
