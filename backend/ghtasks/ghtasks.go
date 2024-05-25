// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ghtasks

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

// GhtasksMetaData contains all meta data concerning the Ghtasks contract.
var GhtasksMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"don\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"donId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"donSecretsSlotId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"donSecretsVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleOracleFulfillment\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"response\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"newOrg\",\"inputs\":[{\"name\":\"initalOwners\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"orgId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"request\",\"inputs\":[{\"name\":\"orgId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"args\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDonSecret\",\"inputs\":[{\"name\":\"slotId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOrgOwners\",\"inputs\":[{\"name\":\"orgId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOrgRepo\",\"inputs\":[{\"name\":\"orgId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"repoOwner\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"repoName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSource\",\"inputs\":[{\"name\":\"sourceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"source\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sources\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"subscriptionId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalOrgs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"NewOrg\",\"inputs\":[{\"name\":\"orgId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"initalOwners\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Request\",\"inputs\":[{\"name\":\"orgId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sourceId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"args\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"requestId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestFulfilled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestSent\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDonSecret\",\"inputs\":[{\"name\":\"slotId\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetOrgOwners\",\"inputs\":[{\"name\":\"orgId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owners\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"bool[]\",\"indexed\":false,\"internalType\":\"bool[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetOrgRepo\",\"inputs\":[{\"name\":\"orgId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"repoOwner\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"repoName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetSource\",\"inputs\":[{\"name\":\"sourceId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"source\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyArgs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptySource\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoInlineSecrets\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRouterCanFulfill\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// GhtasksABI is the input ABI used to generate the binding from.
// Deprecated: Use GhtasksMetaData.ABI instead.
var GhtasksABI = GhtasksMetaData.ABI

// Ghtasks is an auto generated Go binding around an Ethereum contract.
type Ghtasks struct {
	GhtasksCaller     // Read-only binding to the contract
	GhtasksTransactor // Write-only binding to the contract
	GhtasksFilterer   // Log filterer for contract events
}

// GhtasksCaller is an auto generated read-only Go binding around an Ethereum contract.
type GhtasksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GhtasksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GhtasksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GhtasksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GhtasksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GhtasksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GhtasksSession struct {
	Contract     *Ghtasks          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GhtasksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GhtasksCallerSession struct {
	Contract *GhtasksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GhtasksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GhtasksTransactorSession struct {
	Contract     *GhtasksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GhtasksRaw is an auto generated low-level Go binding around an Ethereum contract.
type GhtasksRaw struct {
	Contract *Ghtasks // Generic contract binding to access the raw methods on
}

// GhtasksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GhtasksCallerRaw struct {
	Contract *GhtasksCaller // Generic read-only contract binding to access the raw methods on
}

// GhtasksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GhtasksTransactorRaw struct {
	Contract *GhtasksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGhtasks creates a new instance of Ghtasks, bound to a specific deployed contract.
func NewGhtasks(address common.Address, backend bind.ContractBackend) (*Ghtasks, error) {
	contract, err := bindGhtasks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ghtasks{GhtasksCaller: GhtasksCaller{contract: contract}, GhtasksTransactor: GhtasksTransactor{contract: contract}, GhtasksFilterer: GhtasksFilterer{contract: contract}}, nil
}

// NewGhtasksCaller creates a new read-only instance of Ghtasks, bound to a specific deployed contract.
func NewGhtasksCaller(address common.Address, caller bind.ContractCaller) (*GhtasksCaller, error) {
	contract, err := bindGhtasks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GhtasksCaller{contract: contract}, nil
}

// NewGhtasksTransactor creates a new write-only instance of Ghtasks, bound to a specific deployed contract.
func NewGhtasksTransactor(address common.Address, transactor bind.ContractTransactor) (*GhtasksTransactor, error) {
	contract, err := bindGhtasks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GhtasksTransactor{contract: contract}, nil
}

// NewGhtasksFilterer creates a new log filterer instance of Ghtasks, bound to a specific deployed contract.
func NewGhtasksFilterer(address common.Address, filterer bind.ContractFilterer) (*GhtasksFilterer, error) {
	contract, err := bindGhtasks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GhtasksFilterer{contract: contract}, nil
}

// bindGhtasks binds a generic wrapper to an already deployed contract.
func bindGhtasks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GhtasksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ghtasks *GhtasksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ghtasks.Contract.GhtasksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ghtasks *GhtasksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ghtasks.Contract.GhtasksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ghtasks *GhtasksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ghtasks.Contract.GhtasksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ghtasks *GhtasksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ghtasks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ghtasks *GhtasksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ghtasks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ghtasks *GhtasksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ghtasks.Contract.contract.Transact(opts, method, params...)
}

// DonId is a free data retrieval call binding the contract method 0x8dbe7b9d.
//
// Solidity: function donId() view returns(bytes32)
func (_Ghtasks *GhtasksCaller) DonId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ghtasks.contract.Call(opts, &out, "donId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DonId is a free data retrieval call binding the contract method 0x8dbe7b9d.
//
// Solidity: function donId() view returns(bytes32)
func (_Ghtasks *GhtasksSession) DonId() ([32]byte, error) {
	return _Ghtasks.Contract.DonId(&_Ghtasks.CallOpts)
}

// DonId is a free data retrieval call binding the contract method 0x8dbe7b9d.
//
// Solidity: function donId() view returns(bytes32)
func (_Ghtasks *GhtasksCallerSession) DonId() ([32]byte, error) {
	return _Ghtasks.Contract.DonId(&_Ghtasks.CallOpts)
}

// DonSecretsSlotId is a free data retrieval call binding the contract method 0xd7257cf5.
//
// Solidity: function donSecretsSlotId() view returns(uint8)
func (_Ghtasks *GhtasksCaller) DonSecretsSlotId(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ghtasks.contract.Call(opts, &out, "donSecretsSlotId")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DonSecretsSlotId is a free data retrieval call binding the contract method 0xd7257cf5.
//
// Solidity: function donSecretsSlotId() view returns(uint8)
func (_Ghtasks *GhtasksSession) DonSecretsSlotId() (uint8, error) {
	return _Ghtasks.Contract.DonSecretsSlotId(&_Ghtasks.CallOpts)
}

// DonSecretsSlotId is a free data retrieval call binding the contract method 0xd7257cf5.
//
// Solidity: function donSecretsSlotId() view returns(uint8)
func (_Ghtasks *GhtasksCallerSession) DonSecretsSlotId() (uint8, error) {
	return _Ghtasks.Contract.DonSecretsSlotId(&_Ghtasks.CallOpts)
}

// DonSecretsVersion is a free data retrieval call binding the contract method 0x6421288a.
//
// Solidity: function donSecretsVersion() view returns(uint64)
func (_Ghtasks *GhtasksCaller) DonSecretsVersion(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Ghtasks.contract.Call(opts, &out, "donSecretsVersion")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DonSecretsVersion is a free data retrieval call binding the contract method 0x6421288a.
//
// Solidity: function donSecretsVersion() view returns(uint64)
func (_Ghtasks *GhtasksSession) DonSecretsVersion() (uint64, error) {
	return _Ghtasks.Contract.DonSecretsVersion(&_Ghtasks.CallOpts)
}

// DonSecretsVersion is a free data retrieval call binding the contract method 0x6421288a.
//
// Solidity: function donSecretsVersion() view returns(uint64)
func (_Ghtasks *GhtasksCallerSession) DonSecretsVersion() (uint64, error) {
	return _Ghtasks.Contract.DonSecretsVersion(&_Ghtasks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ghtasks *GhtasksCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ghtasks.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ghtasks *GhtasksSession) Owner() (common.Address, error) {
	return _Ghtasks.Contract.Owner(&_Ghtasks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ghtasks *GhtasksCallerSession) Owner() (common.Address, error) {
	return _Ghtasks.Contract.Owner(&_Ghtasks.CallOpts)
}

// Sources is a free data retrieval call binding the contract method 0xa3663d37.
//
// Solidity: function sources(uint256 ) view returns(string)
func (_Ghtasks *GhtasksCaller) Sources(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Ghtasks.contract.Call(opts, &out, "sources", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Sources is a free data retrieval call binding the contract method 0xa3663d37.
//
// Solidity: function sources(uint256 ) view returns(string)
func (_Ghtasks *GhtasksSession) Sources(arg0 *big.Int) (string, error) {
	return _Ghtasks.Contract.Sources(&_Ghtasks.CallOpts, arg0)
}

// Sources is a free data retrieval call binding the contract method 0xa3663d37.
//
// Solidity: function sources(uint256 ) view returns(string)
func (_Ghtasks *GhtasksCallerSession) Sources(arg0 *big.Int) (string, error) {
	return _Ghtasks.Contract.Sources(&_Ghtasks.CallOpts, arg0)
}

// SubscriptionId is a free data retrieval call binding the contract method 0x09c1ba2e.
//
// Solidity: function subscriptionId() view returns(uint64)
func (_Ghtasks *GhtasksCaller) SubscriptionId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Ghtasks.contract.Call(opts, &out, "subscriptionId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SubscriptionId is a free data retrieval call binding the contract method 0x09c1ba2e.
//
// Solidity: function subscriptionId() view returns(uint64)
func (_Ghtasks *GhtasksSession) SubscriptionId() (uint64, error) {
	return _Ghtasks.Contract.SubscriptionId(&_Ghtasks.CallOpts)
}

// SubscriptionId is a free data retrieval call binding the contract method 0x09c1ba2e.
//
// Solidity: function subscriptionId() view returns(uint64)
func (_Ghtasks *GhtasksCallerSession) SubscriptionId() (uint64, error) {
	return _Ghtasks.Contract.SubscriptionId(&_Ghtasks.CallOpts)
}

// TotalOrgs is a free data retrieval call binding the contract method 0x94428654.
//
// Solidity: function totalOrgs() view returns(uint256)
func (_Ghtasks *GhtasksCaller) TotalOrgs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ghtasks.contract.Call(opts, &out, "totalOrgs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOrgs is a free data retrieval call binding the contract method 0x94428654.
//
// Solidity: function totalOrgs() view returns(uint256)
func (_Ghtasks *GhtasksSession) TotalOrgs() (*big.Int, error) {
	return _Ghtasks.Contract.TotalOrgs(&_Ghtasks.CallOpts)
}

// TotalOrgs is a free data retrieval call binding the contract method 0x94428654.
//
// Solidity: function totalOrgs() view returns(uint256)
func (_Ghtasks *GhtasksCallerSession) TotalOrgs() (*big.Int, error) {
	return _Ghtasks.Contract.TotalOrgs(&_Ghtasks.CallOpts)
}

// HandleOracleFulfillment is a paid mutator transaction binding the contract method 0x0ca76175.
//
// Solidity: function handleOracleFulfillment(bytes32 requestId, bytes response, bytes err) returns()
func (_Ghtasks *GhtasksTransactor) HandleOracleFulfillment(opts *bind.TransactOpts, requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _Ghtasks.contract.Transact(opts, "handleOracleFulfillment", requestId, response, err)
}

// HandleOracleFulfillment is a paid mutator transaction binding the contract method 0x0ca76175.
//
// Solidity: function handleOracleFulfillment(bytes32 requestId, bytes response, bytes err) returns()
func (_Ghtasks *GhtasksSession) HandleOracleFulfillment(requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _Ghtasks.Contract.HandleOracleFulfillment(&_Ghtasks.TransactOpts, requestId, response, err)
}

// HandleOracleFulfillment is a paid mutator transaction binding the contract method 0x0ca76175.
//
// Solidity: function handleOracleFulfillment(bytes32 requestId, bytes response, bytes err) returns()
func (_Ghtasks *GhtasksTransactorSession) HandleOracleFulfillment(requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _Ghtasks.Contract.HandleOracleFulfillment(&_Ghtasks.TransactOpts, requestId, response, err)
}

// NewOrg is a paid mutator transaction binding the contract method 0x1586dea8.
//
// Solidity: function newOrg(address[] initalOwners) returns(uint256 orgId)
func (_Ghtasks *GhtasksTransactor) NewOrg(opts *bind.TransactOpts, initalOwners []common.Address) (*types.Transaction, error) {
	return _Ghtasks.contract.Transact(opts, "newOrg", initalOwners)
}

// NewOrg is a paid mutator transaction binding the contract method 0x1586dea8.
//
// Solidity: function newOrg(address[] initalOwners) returns(uint256 orgId)
func (_Ghtasks *GhtasksSession) NewOrg(initalOwners []common.Address) (*types.Transaction, error) {
	return _Ghtasks.Contract.NewOrg(&_Ghtasks.TransactOpts, initalOwners)
}

// NewOrg is a paid mutator transaction binding the contract method 0x1586dea8.
//
// Solidity: function newOrg(address[] initalOwners) returns(uint256 orgId)
func (_Ghtasks *GhtasksTransactorSession) NewOrg(initalOwners []common.Address) (*types.Transaction, error) {
	return _Ghtasks.Contract.NewOrg(&_Ghtasks.TransactOpts, initalOwners)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ghtasks *GhtasksTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ghtasks.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ghtasks *GhtasksSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ghtasks.Contract.RenounceOwnership(&_Ghtasks.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ghtasks *GhtasksTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ghtasks.Contract.RenounceOwnership(&_Ghtasks.TransactOpts)
}

// Request is a paid mutator transaction binding the contract method 0x03594344.
//
// Solidity: function request(uint256 orgId, uint256 sourceId, string[] args, uint32 gasLimit) returns(bytes32 requestId)
func (_Ghtasks *GhtasksTransactor) Request(opts *bind.TransactOpts, orgId *big.Int, sourceId *big.Int, args []string, gasLimit uint32) (*types.Transaction, error) {
	return _Ghtasks.contract.Transact(opts, "request", orgId, sourceId, args, gasLimit)
}

// Request is a paid mutator transaction binding the contract method 0x03594344.
//
// Solidity: function request(uint256 orgId, uint256 sourceId, string[] args, uint32 gasLimit) returns(bytes32 requestId)
func (_Ghtasks *GhtasksSession) Request(orgId *big.Int, sourceId *big.Int, args []string, gasLimit uint32) (*types.Transaction, error) {
	return _Ghtasks.Contract.Request(&_Ghtasks.TransactOpts, orgId, sourceId, args, gasLimit)
}

// Request is a paid mutator transaction binding the contract method 0x03594344.
//
// Solidity: function request(uint256 orgId, uint256 sourceId, string[] args, uint32 gasLimit) returns(bytes32 requestId)
func (_Ghtasks *GhtasksTransactorSession) Request(orgId *big.Int, sourceId *big.Int, args []string, gasLimit uint32) (*types.Transaction, error) {
	return _Ghtasks.Contract.Request(&_Ghtasks.TransactOpts, orgId, sourceId, args, gasLimit)
}

// SetDonSecret is a paid mutator transaction binding the contract method 0x2831246b.
//
// Solidity: function setDonSecret(uint8 slotId, uint64 version) returns()
func (_Ghtasks *GhtasksTransactor) SetDonSecret(opts *bind.TransactOpts, slotId uint8, version uint64) (*types.Transaction, error) {
	return _Ghtasks.contract.Transact(opts, "setDonSecret", slotId, version)
}

// SetDonSecret is a paid mutator transaction binding the contract method 0x2831246b.
//
// Solidity: function setDonSecret(uint8 slotId, uint64 version) returns()
func (_Ghtasks *GhtasksSession) SetDonSecret(slotId uint8, version uint64) (*types.Transaction, error) {
	return _Ghtasks.Contract.SetDonSecret(&_Ghtasks.TransactOpts, slotId, version)
}

// SetDonSecret is a paid mutator transaction binding the contract method 0x2831246b.
//
// Solidity: function setDonSecret(uint8 slotId, uint64 version) returns()
func (_Ghtasks *GhtasksTransactorSession) SetDonSecret(slotId uint8, version uint64) (*types.Transaction, error) {
	return _Ghtasks.Contract.SetDonSecret(&_Ghtasks.TransactOpts, slotId, version)
}

// SetOrgOwners is a paid mutator transaction binding the contract method 0x28adc4d4.
//
// Solidity: function setOrgOwners(uint256 orgId, address[] owners, bool[] values) returns()
func (_Ghtasks *GhtasksTransactor) SetOrgOwners(opts *bind.TransactOpts, orgId *big.Int, owners []common.Address, values []bool) (*types.Transaction, error) {
	return _Ghtasks.contract.Transact(opts, "setOrgOwners", orgId, owners, values)
}

// SetOrgOwners is a paid mutator transaction binding the contract method 0x28adc4d4.
//
// Solidity: function setOrgOwners(uint256 orgId, address[] owners, bool[] values) returns()
func (_Ghtasks *GhtasksSession) SetOrgOwners(orgId *big.Int, owners []common.Address, values []bool) (*types.Transaction, error) {
	return _Ghtasks.Contract.SetOrgOwners(&_Ghtasks.TransactOpts, orgId, owners, values)
}

// SetOrgOwners is a paid mutator transaction binding the contract method 0x28adc4d4.
//
// Solidity: function setOrgOwners(uint256 orgId, address[] owners, bool[] values) returns()
func (_Ghtasks *GhtasksTransactorSession) SetOrgOwners(orgId *big.Int, owners []common.Address, values []bool) (*types.Transaction, error) {
	return _Ghtasks.Contract.SetOrgOwners(&_Ghtasks.TransactOpts, orgId, owners, values)
}

// SetOrgRepo is a paid mutator transaction binding the contract method 0x4a7bf594.
//
// Solidity: function setOrgRepo(uint256 orgId, string repoOwner, string repoName, bool value) returns()
func (_Ghtasks *GhtasksTransactor) SetOrgRepo(opts *bind.TransactOpts, orgId *big.Int, repoOwner string, repoName string, value bool) (*types.Transaction, error) {
	return _Ghtasks.contract.Transact(opts, "setOrgRepo", orgId, repoOwner, repoName, value)
}

// SetOrgRepo is a paid mutator transaction binding the contract method 0x4a7bf594.
//
// Solidity: function setOrgRepo(uint256 orgId, string repoOwner, string repoName, bool value) returns()
func (_Ghtasks *GhtasksSession) SetOrgRepo(orgId *big.Int, repoOwner string, repoName string, value bool) (*types.Transaction, error) {
	return _Ghtasks.Contract.SetOrgRepo(&_Ghtasks.TransactOpts, orgId, repoOwner, repoName, value)
}

// SetOrgRepo is a paid mutator transaction binding the contract method 0x4a7bf594.
//
// Solidity: function setOrgRepo(uint256 orgId, string repoOwner, string repoName, bool value) returns()
func (_Ghtasks *GhtasksTransactorSession) SetOrgRepo(orgId *big.Int, repoOwner string, repoName string, value bool) (*types.Transaction, error) {
	return _Ghtasks.Contract.SetOrgRepo(&_Ghtasks.TransactOpts, orgId, repoOwner, repoName, value)
}

// SetSource is a paid mutator transaction binding the contract method 0xa71e5da5.
//
// Solidity: function setSource(uint256 sourceId, string source) returns()
func (_Ghtasks *GhtasksTransactor) SetSource(opts *bind.TransactOpts, sourceId *big.Int, source string) (*types.Transaction, error) {
	return _Ghtasks.contract.Transact(opts, "setSource", sourceId, source)
}

// SetSource is a paid mutator transaction binding the contract method 0xa71e5da5.
//
// Solidity: function setSource(uint256 sourceId, string source) returns()
func (_Ghtasks *GhtasksSession) SetSource(sourceId *big.Int, source string) (*types.Transaction, error) {
	return _Ghtasks.Contract.SetSource(&_Ghtasks.TransactOpts, sourceId, source)
}

// SetSource is a paid mutator transaction binding the contract method 0xa71e5da5.
//
// Solidity: function setSource(uint256 sourceId, string source) returns()
func (_Ghtasks *GhtasksTransactorSession) SetSource(sourceId *big.Int, source string) (*types.Transaction, error) {
	return _Ghtasks.Contract.SetSource(&_Ghtasks.TransactOpts, sourceId, source)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ghtasks *GhtasksTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ghtasks.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ghtasks *GhtasksSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ghtasks.Contract.TransferOwnership(&_Ghtasks.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ghtasks *GhtasksTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ghtasks.Contract.TransferOwnership(&_Ghtasks.TransactOpts, newOwner)
}

// GhtasksNewOrgIterator is returned from FilterNewOrg and is used to iterate over the raw logs and unpacked data for NewOrg events raised by the Ghtasks contract.
type GhtasksNewOrgIterator struct {
	Event *GhtasksNewOrg // Event containing the contract specifics and raw log

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
func (it *GhtasksNewOrgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GhtasksNewOrg)
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
		it.Event = new(GhtasksNewOrg)
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
func (it *GhtasksNewOrgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GhtasksNewOrgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GhtasksNewOrg represents a NewOrg event raised by the Ghtasks contract.
type GhtasksNewOrg struct {
	OrgId        *big.Int
	InitalOwners []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewOrg is a free log retrieval operation binding the contract event 0xd5ecb412197d745517901e124cfe396e52736fb4bb02fcadc68325c989c216d6.
//
// Solidity: event NewOrg(uint256 indexed orgId, address[] initalOwners)
func (_Ghtasks *GhtasksFilterer) FilterNewOrg(opts *bind.FilterOpts, orgId []*big.Int) (*GhtasksNewOrgIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Ghtasks.contract.FilterLogs(opts, "NewOrg", orgIdRule)
	if err != nil {
		return nil, err
	}
	return &GhtasksNewOrgIterator{contract: _Ghtasks.contract, event: "NewOrg", logs: logs, sub: sub}, nil
}

// WatchNewOrg is a free log subscription operation binding the contract event 0xd5ecb412197d745517901e124cfe396e52736fb4bb02fcadc68325c989c216d6.
//
// Solidity: event NewOrg(uint256 indexed orgId, address[] initalOwners)
func (_Ghtasks *GhtasksFilterer) WatchNewOrg(opts *bind.WatchOpts, sink chan<- *GhtasksNewOrg, orgId []*big.Int) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Ghtasks.contract.WatchLogs(opts, "NewOrg", orgIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GhtasksNewOrg)
				if err := _Ghtasks.contract.UnpackLog(event, "NewOrg", log); err != nil {
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

// ParseNewOrg is a log parse operation binding the contract event 0xd5ecb412197d745517901e124cfe396e52736fb4bb02fcadc68325c989c216d6.
//
// Solidity: event NewOrg(uint256 indexed orgId, address[] initalOwners)
func (_Ghtasks *GhtasksFilterer) ParseNewOrg(log types.Log) (*GhtasksNewOrg, error) {
	event := new(GhtasksNewOrg)
	if err := _Ghtasks.contract.UnpackLog(event, "NewOrg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GhtasksOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ghtasks contract.
type GhtasksOwnershipTransferredIterator struct {
	Event *GhtasksOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GhtasksOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GhtasksOwnershipTransferred)
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
		it.Event = new(GhtasksOwnershipTransferred)
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
func (it *GhtasksOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GhtasksOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GhtasksOwnershipTransferred represents a OwnershipTransferred event raised by the Ghtasks contract.
type GhtasksOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ghtasks *GhtasksFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GhtasksOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ghtasks.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GhtasksOwnershipTransferredIterator{contract: _Ghtasks.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ghtasks *GhtasksFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GhtasksOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ghtasks.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GhtasksOwnershipTransferred)
				if err := _Ghtasks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ghtasks *GhtasksFilterer) ParseOwnershipTransferred(log types.Log) (*GhtasksOwnershipTransferred, error) {
	event := new(GhtasksOwnershipTransferred)
	if err := _Ghtasks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GhtasksRequestIterator is returned from FilterRequest and is used to iterate over the raw logs and unpacked data for Request events raised by the Ghtasks contract.
type GhtasksRequestIterator struct {
	Event *GhtasksRequest // Event containing the contract specifics and raw log

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
func (it *GhtasksRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GhtasksRequest)
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
		it.Event = new(GhtasksRequest)
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
func (it *GhtasksRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GhtasksRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GhtasksRequest represents a Request event raised by the Ghtasks contract.
type GhtasksRequest struct {
	OrgId     *big.Int
	SourceId  *big.Int
	Args      []string
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequest is a free log retrieval operation binding the contract event 0xe867da8ec9dfeab798dcac1d653a6349004340f8ed3d9a70022717ee7dafee39.
//
// Solidity: event Request(uint256 indexed orgId, uint256 indexed sourceId, string[] args, bytes32 requestId)
func (_Ghtasks *GhtasksFilterer) FilterRequest(opts *bind.FilterOpts, orgId []*big.Int, sourceId []*big.Int) (*GhtasksRequestIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _Ghtasks.contract.FilterLogs(opts, "Request", orgIdRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &GhtasksRequestIterator{contract: _Ghtasks.contract, event: "Request", logs: logs, sub: sub}, nil
}

// WatchRequest is a free log subscription operation binding the contract event 0xe867da8ec9dfeab798dcac1d653a6349004340f8ed3d9a70022717ee7dafee39.
//
// Solidity: event Request(uint256 indexed orgId, uint256 indexed sourceId, string[] args, bytes32 requestId)
func (_Ghtasks *GhtasksFilterer) WatchRequest(opts *bind.WatchOpts, sink chan<- *GhtasksRequest, orgId []*big.Int, sourceId []*big.Int) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _Ghtasks.contract.WatchLogs(opts, "Request", orgIdRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GhtasksRequest)
				if err := _Ghtasks.contract.UnpackLog(event, "Request", log); err != nil {
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

// ParseRequest is a log parse operation binding the contract event 0xe867da8ec9dfeab798dcac1d653a6349004340f8ed3d9a70022717ee7dafee39.
//
// Solidity: event Request(uint256 indexed orgId, uint256 indexed sourceId, string[] args, bytes32 requestId)
func (_Ghtasks *GhtasksFilterer) ParseRequest(log types.Log) (*GhtasksRequest, error) {
	event := new(GhtasksRequest)
	if err := _Ghtasks.contract.UnpackLog(event, "Request", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GhtasksRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the Ghtasks contract.
type GhtasksRequestFulfilledIterator struct {
	Event *GhtasksRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *GhtasksRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GhtasksRequestFulfilled)
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
		it.Event = new(GhtasksRequestFulfilled)
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
func (it *GhtasksRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GhtasksRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GhtasksRequestFulfilled represents a RequestFulfilled event raised by the Ghtasks contract.
type GhtasksRequestFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 indexed id)
func (_Ghtasks *GhtasksFilterer) FilterRequestFulfilled(opts *bind.FilterOpts, id [][32]byte) (*GhtasksRequestFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Ghtasks.contract.FilterLogs(opts, "RequestFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &GhtasksRequestFulfilledIterator{contract: _Ghtasks.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 indexed id)
func (_Ghtasks *GhtasksFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *GhtasksRequestFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Ghtasks.contract.WatchLogs(opts, "RequestFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GhtasksRequestFulfilled)
				if err := _Ghtasks.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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

// ParseRequestFulfilled is a log parse operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 indexed id)
func (_Ghtasks *GhtasksFilterer) ParseRequestFulfilled(log types.Log) (*GhtasksRequestFulfilled, error) {
	event := new(GhtasksRequestFulfilled)
	if err := _Ghtasks.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GhtasksRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the Ghtasks contract.
type GhtasksRequestSentIterator struct {
	Event *GhtasksRequestSent // Event containing the contract specifics and raw log

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
func (it *GhtasksRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GhtasksRequestSent)
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
		it.Event = new(GhtasksRequestSent)
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
func (it *GhtasksRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GhtasksRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GhtasksRequestSent represents a RequestSent event raised by the Ghtasks contract.
type GhtasksRequestSent struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 indexed id)
func (_Ghtasks *GhtasksFilterer) FilterRequestSent(opts *bind.FilterOpts, id [][32]byte) (*GhtasksRequestSentIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Ghtasks.contract.FilterLogs(opts, "RequestSent", idRule)
	if err != nil {
		return nil, err
	}
	return &GhtasksRequestSentIterator{contract: _Ghtasks.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 indexed id)
func (_Ghtasks *GhtasksFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *GhtasksRequestSent, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Ghtasks.contract.WatchLogs(opts, "RequestSent", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GhtasksRequestSent)
				if err := _Ghtasks.contract.UnpackLog(event, "RequestSent", log); err != nil {
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

// ParseRequestSent is a log parse operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 indexed id)
func (_Ghtasks *GhtasksFilterer) ParseRequestSent(log types.Log) (*GhtasksRequestSent, error) {
	event := new(GhtasksRequestSent)
	if err := _Ghtasks.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GhtasksSetDonSecretIterator is returned from FilterSetDonSecret and is used to iterate over the raw logs and unpacked data for SetDonSecret events raised by the Ghtasks contract.
type GhtasksSetDonSecretIterator struct {
	Event *GhtasksSetDonSecret // Event containing the contract specifics and raw log

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
func (it *GhtasksSetDonSecretIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GhtasksSetDonSecret)
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
		it.Event = new(GhtasksSetDonSecret)
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
func (it *GhtasksSetDonSecretIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GhtasksSetDonSecretIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GhtasksSetDonSecret represents a SetDonSecret event raised by the Ghtasks contract.
type GhtasksSetDonSecret struct {
	SlotId  uint8
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetDonSecret is a free log retrieval operation binding the contract event 0x55b42049afa565e6cc53a77f450246e3c6132a01c9adac325eb4793dd815db65.
//
// Solidity: event SetDonSecret(uint8 slotId, uint64 version)
func (_Ghtasks *GhtasksFilterer) FilterSetDonSecret(opts *bind.FilterOpts) (*GhtasksSetDonSecretIterator, error) {

	logs, sub, err := _Ghtasks.contract.FilterLogs(opts, "SetDonSecret")
	if err != nil {
		return nil, err
	}
	return &GhtasksSetDonSecretIterator{contract: _Ghtasks.contract, event: "SetDonSecret", logs: logs, sub: sub}, nil
}

// WatchSetDonSecret is a free log subscription operation binding the contract event 0x55b42049afa565e6cc53a77f450246e3c6132a01c9adac325eb4793dd815db65.
//
// Solidity: event SetDonSecret(uint8 slotId, uint64 version)
func (_Ghtasks *GhtasksFilterer) WatchSetDonSecret(opts *bind.WatchOpts, sink chan<- *GhtasksSetDonSecret) (event.Subscription, error) {

	logs, sub, err := _Ghtasks.contract.WatchLogs(opts, "SetDonSecret")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GhtasksSetDonSecret)
				if err := _Ghtasks.contract.UnpackLog(event, "SetDonSecret", log); err != nil {
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

// ParseSetDonSecret is a log parse operation binding the contract event 0x55b42049afa565e6cc53a77f450246e3c6132a01c9adac325eb4793dd815db65.
//
// Solidity: event SetDonSecret(uint8 slotId, uint64 version)
func (_Ghtasks *GhtasksFilterer) ParseSetDonSecret(log types.Log) (*GhtasksSetDonSecret, error) {
	event := new(GhtasksSetDonSecret)
	if err := _Ghtasks.contract.UnpackLog(event, "SetDonSecret", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GhtasksSetOrgOwnersIterator is returned from FilterSetOrgOwners and is used to iterate over the raw logs and unpacked data for SetOrgOwners events raised by the Ghtasks contract.
type GhtasksSetOrgOwnersIterator struct {
	Event *GhtasksSetOrgOwners // Event containing the contract specifics and raw log

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
func (it *GhtasksSetOrgOwnersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GhtasksSetOrgOwners)
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
		it.Event = new(GhtasksSetOrgOwners)
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
func (it *GhtasksSetOrgOwnersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GhtasksSetOrgOwnersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GhtasksSetOrgOwners represents a SetOrgOwners event raised by the Ghtasks contract.
type GhtasksSetOrgOwners struct {
	OrgId  *big.Int
	Owners []common.Address
	Values []bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOrgOwners is a free log retrieval operation binding the contract event 0x1eb5b0d601211b7d94c584fc45e21075929a092fd458f04e587a730bfcb9f227.
//
// Solidity: event SetOrgOwners(uint256 indexed orgId, address[] owners, bool[] values)
func (_Ghtasks *GhtasksFilterer) FilterSetOrgOwners(opts *bind.FilterOpts, orgId []*big.Int) (*GhtasksSetOrgOwnersIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Ghtasks.contract.FilterLogs(opts, "SetOrgOwners", orgIdRule)
	if err != nil {
		return nil, err
	}
	return &GhtasksSetOrgOwnersIterator{contract: _Ghtasks.contract, event: "SetOrgOwners", logs: logs, sub: sub}, nil
}

// WatchSetOrgOwners is a free log subscription operation binding the contract event 0x1eb5b0d601211b7d94c584fc45e21075929a092fd458f04e587a730bfcb9f227.
//
// Solidity: event SetOrgOwners(uint256 indexed orgId, address[] owners, bool[] values)
func (_Ghtasks *GhtasksFilterer) WatchSetOrgOwners(opts *bind.WatchOpts, sink chan<- *GhtasksSetOrgOwners, orgId []*big.Int) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Ghtasks.contract.WatchLogs(opts, "SetOrgOwners", orgIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GhtasksSetOrgOwners)
				if err := _Ghtasks.contract.UnpackLog(event, "SetOrgOwners", log); err != nil {
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

// ParseSetOrgOwners is a log parse operation binding the contract event 0x1eb5b0d601211b7d94c584fc45e21075929a092fd458f04e587a730bfcb9f227.
//
// Solidity: event SetOrgOwners(uint256 indexed orgId, address[] owners, bool[] values)
func (_Ghtasks *GhtasksFilterer) ParseSetOrgOwners(log types.Log) (*GhtasksSetOrgOwners, error) {
	event := new(GhtasksSetOrgOwners)
	if err := _Ghtasks.contract.UnpackLog(event, "SetOrgOwners", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GhtasksSetOrgRepoIterator is returned from FilterSetOrgRepo and is used to iterate over the raw logs and unpacked data for SetOrgRepo events raised by the Ghtasks contract.
type GhtasksSetOrgRepoIterator struct {
	Event *GhtasksSetOrgRepo // Event containing the contract specifics and raw log

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
func (it *GhtasksSetOrgRepoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GhtasksSetOrgRepo)
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
		it.Event = new(GhtasksSetOrgRepo)
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
func (it *GhtasksSetOrgRepoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GhtasksSetOrgRepoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GhtasksSetOrgRepo represents a SetOrgRepo event raised by the Ghtasks contract.
type GhtasksSetOrgRepo struct {
	OrgId     *big.Int
	RepoOwner string
	RepoName  string
	Value     bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetOrgRepo is a free log retrieval operation binding the contract event 0xd8c4dc50539ee60cb1507e9fff0b6aa0239a0598f5be13b2f80c45d13fb0542a.
//
// Solidity: event SetOrgRepo(uint256 indexed orgId, string repoOwner, string repoName, bool value)
func (_Ghtasks *GhtasksFilterer) FilterSetOrgRepo(opts *bind.FilterOpts, orgId []*big.Int) (*GhtasksSetOrgRepoIterator, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Ghtasks.contract.FilterLogs(opts, "SetOrgRepo", orgIdRule)
	if err != nil {
		return nil, err
	}
	return &GhtasksSetOrgRepoIterator{contract: _Ghtasks.contract, event: "SetOrgRepo", logs: logs, sub: sub}, nil
}

// WatchSetOrgRepo is a free log subscription operation binding the contract event 0xd8c4dc50539ee60cb1507e9fff0b6aa0239a0598f5be13b2f80c45d13fb0542a.
//
// Solidity: event SetOrgRepo(uint256 indexed orgId, string repoOwner, string repoName, bool value)
func (_Ghtasks *GhtasksFilterer) WatchSetOrgRepo(opts *bind.WatchOpts, sink chan<- *GhtasksSetOrgRepo, orgId []*big.Int) (event.Subscription, error) {

	var orgIdRule []interface{}
	for _, orgIdItem := range orgId {
		orgIdRule = append(orgIdRule, orgIdItem)
	}

	logs, sub, err := _Ghtasks.contract.WatchLogs(opts, "SetOrgRepo", orgIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GhtasksSetOrgRepo)
				if err := _Ghtasks.contract.UnpackLog(event, "SetOrgRepo", log); err != nil {
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

// ParseSetOrgRepo is a log parse operation binding the contract event 0xd8c4dc50539ee60cb1507e9fff0b6aa0239a0598f5be13b2f80c45d13fb0542a.
//
// Solidity: event SetOrgRepo(uint256 indexed orgId, string repoOwner, string repoName, bool value)
func (_Ghtasks *GhtasksFilterer) ParseSetOrgRepo(log types.Log) (*GhtasksSetOrgRepo, error) {
	event := new(GhtasksSetOrgRepo)
	if err := _Ghtasks.contract.UnpackLog(event, "SetOrgRepo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GhtasksSetSourceIterator is returned from FilterSetSource and is used to iterate over the raw logs and unpacked data for SetSource events raised by the Ghtasks contract.
type GhtasksSetSourceIterator struct {
	Event *GhtasksSetSource // Event containing the contract specifics and raw log

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
func (it *GhtasksSetSourceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GhtasksSetSource)
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
		it.Event = new(GhtasksSetSource)
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
func (it *GhtasksSetSourceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GhtasksSetSourceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GhtasksSetSource represents a SetSource event raised by the Ghtasks contract.
type GhtasksSetSource struct {
	SourceId *big.Int
	Source   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetSource is a free log retrieval operation binding the contract event 0xd3fcbeb202b4d93dfc8583da25b2ab144876b4dfd528b966dcd0e833c1e9262d.
//
// Solidity: event SetSource(uint256 indexed sourceId, string source)
func (_Ghtasks *GhtasksFilterer) FilterSetSource(opts *bind.FilterOpts, sourceId []*big.Int) (*GhtasksSetSourceIterator, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _Ghtasks.contract.FilterLogs(opts, "SetSource", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &GhtasksSetSourceIterator{contract: _Ghtasks.contract, event: "SetSource", logs: logs, sub: sub}, nil
}

// WatchSetSource is a free log subscription operation binding the contract event 0xd3fcbeb202b4d93dfc8583da25b2ab144876b4dfd528b966dcd0e833c1e9262d.
//
// Solidity: event SetSource(uint256 indexed sourceId, string source)
func (_Ghtasks *GhtasksFilterer) WatchSetSource(opts *bind.WatchOpts, sink chan<- *GhtasksSetSource, sourceId []*big.Int) (event.Subscription, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _Ghtasks.contract.WatchLogs(opts, "SetSource", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GhtasksSetSource)
				if err := _Ghtasks.contract.UnpackLog(event, "SetSource", log); err != nil {
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

// ParseSetSource is a log parse operation binding the contract event 0xd3fcbeb202b4d93dfc8583da25b2ab144876b4dfd528b966dcd0e833c1e9262d.
//
// Solidity: event SetSource(uint256 indexed sourceId, string source)
func (_Ghtasks *GhtasksFilterer) ParseSetSource(log types.Log) (*GhtasksSetSource, error) {
	event := new(GhtasksSetSource)
	if err := _Ghtasks.contract.UnpackLog(event, "SetSource", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
