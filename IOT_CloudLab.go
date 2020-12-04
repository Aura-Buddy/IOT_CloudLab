// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iot_cloudlab

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IotCloudlabABI is the input ABI used to generate the binding from.
const IotCloudlabABI = "[{\"inputs\":[],\"name\":\"retrieveFuzzyOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieveHumidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieveOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieveTemperature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"Owner\",\"type\":\"bytes32\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"FuzzyOutput\",\"type\":\"uint256\"}],\"name\":\"storeFuzzyOutput\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Humidity\",\"type\":\"uint256\"}],\"name\":\"storeHumidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Temperature\",\"type\":\"uint256\"}],\"name\":\"storeTemperature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IotCloudlabBin is the compiled bytecode used for deploying new contracts.
var IotCloudlabBin = "0x608060405234801561001057600080fd5b50610245806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063d16aeba91161005b578063d16aeba914610135578063d608f45014610153578063dad1b4e914610171578063ebf765791461019f57610088565b806338b3e81d1461008d578063654cf88c146100bb5780639081d447146100e9578063af0638d514610107575b600080fd5b6100b9600480360360208110156100a357600080fd5b81019080803590602001909291905050506101bd565b005b6100e7600480360360208110156100d157600080fd5b81019080803590602001909291905050506101c7565b005b6100f16101d1565b6040518082815260200191505060405180910390f35b6101336004803603602081101561011d57600080fd5b81019080803590602001909291905050506101db565b005b61013d6101e8565b6040518082815260200191505060405180910390f35b61015b6101f1565b6040518082815260200191505060405180910390f35b61019d6004803603602081101561018757600080fd5b81019080803590602001909291905050506101fb565b005b6101a7610205565b6040518082815260200191505060405180910390f35b8060018190555050565b8060038190555050565b6000600254905090565b600a810260028190555050565b60008054905090565b6000600354905090565b8060008190555050565b600060015490509056fea26469706673582212204570fcdd19cf547d8f34f48f32cd878d48565402bccb7316d7c0ed856d1a201564736f6c63430007040033"

// DeployIotCloudlab deploys a new Ethereum contract, binding an instance of IotCloudlab to it.
func DeployIotCloudlab(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IotCloudlab, error) {
	parsed, err := abi.JSON(strings.NewReader(IotCloudlabABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IotCloudlabBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IotCloudlab{IotCloudlabCaller: IotCloudlabCaller{contract: contract}, IotCloudlabTransactor: IotCloudlabTransactor{contract: contract}, IotCloudlabFilterer: IotCloudlabFilterer{contract: contract}}, nil
}

// IotCloudlab is an auto generated Go binding around an Ethereum contract.
type IotCloudlab struct {
	IotCloudlabCaller     // Read-only binding to the contract
	IotCloudlabTransactor // Write-only binding to the contract
	IotCloudlabFilterer   // Log filterer for contract events
}

// IotCloudlabCaller is an auto generated read-only Go binding around an Ethereum contract.
type IotCloudlabCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IotCloudlabTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IotCloudlabTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IotCloudlabFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IotCloudlabFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IotCloudlabSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IotCloudlabSession struct {
	Contract     *IotCloudlab      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IotCloudlabCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IotCloudlabCallerSession struct {
	Contract *IotCloudlabCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IotCloudlabTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IotCloudlabTransactorSession struct {
	Contract     *IotCloudlabTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IotCloudlabRaw is an auto generated low-level Go binding around an Ethereum contract.
type IotCloudlabRaw struct {
	Contract *IotCloudlab // Generic contract binding to access the raw methods on
}

// IotCloudlabCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IotCloudlabCallerRaw struct {
	Contract *IotCloudlabCaller // Generic read-only contract binding to access the raw methods on
}

// IotCloudlabTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IotCloudlabTransactorRaw struct {
	Contract *IotCloudlabTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIotCloudlab creates a new instance of IotCloudlab, bound to a specific deployed contract.
func NewIotCloudlab(address common.Address, backend bind.ContractBackend) (*IotCloudlab, error) {
	contract, err := bindIotCloudlab(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IotCloudlab{IotCloudlabCaller: IotCloudlabCaller{contract: contract}, IotCloudlabTransactor: IotCloudlabTransactor{contract: contract}, IotCloudlabFilterer: IotCloudlabFilterer{contract: contract}}, nil
}

// NewIotCloudlabCaller creates a new read-only instance of IotCloudlab, bound to a specific deployed contract.
func NewIotCloudlabCaller(address common.Address, caller bind.ContractCaller) (*IotCloudlabCaller, error) {
	contract, err := bindIotCloudlab(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IotCloudlabCaller{contract: contract}, nil
}

// NewIotCloudlabTransactor creates a new write-only instance of IotCloudlab, bound to a specific deployed contract.
func NewIotCloudlabTransactor(address common.Address, transactor bind.ContractTransactor) (*IotCloudlabTransactor, error) {
	contract, err := bindIotCloudlab(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IotCloudlabTransactor{contract: contract}, nil
}

// NewIotCloudlabFilterer creates a new log filterer instance of IotCloudlab, bound to a specific deployed contract.
func NewIotCloudlabFilterer(address common.Address, filterer bind.ContractFilterer) (*IotCloudlabFilterer, error) {
	contract, err := bindIotCloudlab(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IotCloudlabFilterer{contract: contract}, nil
}

// bindIotCloudlab binds a generic wrapper to an already deployed contract.
func bindIotCloudlab(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IotCloudlabABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IotCloudlab *IotCloudlabRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IotCloudlab.Contract.IotCloudlabCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IotCloudlab *IotCloudlabRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IotCloudlab.Contract.IotCloudlabTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IotCloudlab *IotCloudlabRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IotCloudlab.Contract.IotCloudlabTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IotCloudlab *IotCloudlabCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IotCloudlab.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IotCloudlab *IotCloudlabTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IotCloudlab.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IotCloudlab *IotCloudlabTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IotCloudlab.Contract.contract.Transact(opts, method, params...)
}

// RetrieveFuzzyOutput is a free data retrieval call binding the contract method 0x9081d447.
//
// Solidity: function retrieveFuzzyOutput() view returns(uint256)
func (_IotCloudlab *IotCloudlabCaller) RetrieveFuzzyOutput(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IotCloudlab.contract.Call(opts, &out, "retrieveFuzzyOutput")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveFuzzyOutput is a free data retrieval call binding the contract method 0x9081d447.
//
// Solidity: function retrieveFuzzyOutput() view returns(uint256)
func (_IotCloudlab *IotCloudlabSession) RetrieveFuzzyOutput() (*big.Int, error) {
	return _IotCloudlab.Contract.RetrieveFuzzyOutput(&_IotCloudlab.CallOpts)
}

// RetrieveFuzzyOutput is a free data retrieval call binding the contract method 0x9081d447.
//
// Solidity: function retrieveFuzzyOutput() view returns(uint256)
func (_IotCloudlab *IotCloudlabCallerSession) RetrieveFuzzyOutput() (*big.Int, error) {
	return _IotCloudlab.Contract.RetrieveFuzzyOutput(&_IotCloudlab.CallOpts)
}

// RetrieveHumidity is a free data retrieval call binding the contract method 0xebf76579.
//
// Solidity: function retrieveHumidity() view returns(uint256)
func (_IotCloudlab *IotCloudlabCaller) RetrieveHumidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IotCloudlab.contract.Call(opts, &out, "retrieveHumidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveHumidity is a free data retrieval call binding the contract method 0xebf76579.
//
// Solidity: function retrieveHumidity() view returns(uint256)
func (_IotCloudlab *IotCloudlabSession) RetrieveHumidity() (*big.Int, error) {
	return _IotCloudlab.Contract.RetrieveHumidity(&_IotCloudlab.CallOpts)
}

// RetrieveHumidity is a free data retrieval call binding the contract method 0xebf76579.
//
// Solidity: function retrieveHumidity() view returns(uint256)
func (_IotCloudlab *IotCloudlabCallerSession) RetrieveHumidity() (*big.Int, error) {
	return _IotCloudlab.Contract.RetrieveHumidity(&_IotCloudlab.CallOpts)
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xd608f450.
//
// Solidity: function retrieveOwner() view returns(bytes32)
func (_IotCloudlab *IotCloudlabCaller) RetrieveOwner(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IotCloudlab.contract.Call(opts, &out, "retrieveOwner")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RetrieveOwner is a free data retrieval call binding the contract method 0xd608f450.
//
// Solidity: function retrieveOwner() view returns(bytes32)
func (_IotCloudlab *IotCloudlabSession) RetrieveOwner() ([32]byte, error) {
	return _IotCloudlab.Contract.RetrieveOwner(&_IotCloudlab.CallOpts)
}

// RetrieveOwner is a free data retrieval call binding the contract method 0xd608f450.
//
// Solidity: function retrieveOwner() view returns(bytes32)
func (_IotCloudlab *IotCloudlabCallerSession) RetrieveOwner() ([32]byte, error) {
	return _IotCloudlab.Contract.RetrieveOwner(&_IotCloudlab.CallOpts)
}

// RetrieveTemperature is a free data retrieval call binding the contract method 0xd16aeba9.
//
// Solidity: function retrieveTemperature() view returns(uint256)
func (_IotCloudlab *IotCloudlabCaller) RetrieveTemperature(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IotCloudlab.contract.Call(opts, &out, "retrieveTemperature")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveTemperature is a free data retrieval call binding the contract method 0xd16aeba9.
//
// Solidity: function retrieveTemperature() view returns(uint256)
func (_IotCloudlab *IotCloudlabSession) RetrieveTemperature() (*big.Int, error) {
	return _IotCloudlab.Contract.RetrieveTemperature(&_IotCloudlab.CallOpts)
}

// RetrieveTemperature is a free data retrieval call binding the contract method 0xd16aeba9.
//
// Solidity: function retrieveTemperature() view returns(uint256)
func (_IotCloudlab *IotCloudlabCallerSession) RetrieveTemperature() (*big.Int, error) {
	return _IotCloudlab.Contract.RetrieveTemperature(&_IotCloudlab.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 Owner) returns()
func (_IotCloudlab *IotCloudlabTransactor) Store(opts *bind.TransactOpts, Owner [32]byte) (*types.Transaction, error) {
	return _IotCloudlab.contract.Transact(opts, "store", Owner)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 Owner) returns()
func (_IotCloudlab *IotCloudlabSession) Store(Owner [32]byte) (*types.Transaction, error) {
	return _IotCloudlab.Contract.Store(&_IotCloudlab.TransactOpts, Owner)
}

// Store is a paid mutator transaction binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 Owner) returns()
func (_IotCloudlab *IotCloudlabTransactorSession) Store(Owner [32]byte) (*types.Transaction, error) {
	return _IotCloudlab.Contract.Store(&_IotCloudlab.TransactOpts, Owner)
}

// StoreFuzzyOutput is a paid mutator transaction binding the contract method 0xaf0638d5.
//
// Solidity: function storeFuzzyOutput(uint256 FuzzyOutput) returns()
func (_IotCloudlab *IotCloudlabTransactor) StoreFuzzyOutput(opts *bind.TransactOpts, FuzzyOutput *big.Int) (*types.Transaction, error) {
	return _IotCloudlab.contract.Transact(opts, "storeFuzzyOutput", FuzzyOutput)
}

// StoreFuzzyOutput is a paid mutator transaction binding the contract method 0xaf0638d5.
//
// Solidity: function storeFuzzyOutput(uint256 FuzzyOutput) returns()
func (_IotCloudlab *IotCloudlabSession) StoreFuzzyOutput(FuzzyOutput *big.Int) (*types.Transaction, error) {
	return _IotCloudlab.Contract.StoreFuzzyOutput(&_IotCloudlab.TransactOpts, FuzzyOutput)
}

// StoreFuzzyOutput is a paid mutator transaction binding the contract method 0xaf0638d5.
//
// Solidity: function storeFuzzyOutput(uint256 FuzzyOutput) returns()
func (_IotCloudlab *IotCloudlabTransactorSession) StoreFuzzyOutput(FuzzyOutput *big.Int) (*types.Transaction, error) {
	return _IotCloudlab.Contract.StoreFuzzyOutput(&_IotCloudlab.TransactOpts, FuzzyOutput)
}

// StoreHumidity is a paid mutator transaction binding the contract method 0x38b3e81d.
//
// Solidity: function storeHumidity(uint256 Humidity) returns()
func (_IotCloudlab *IotCloudlabTransactor) StoreHumidity(opts *bind.TransactOpts, Humidity *big.Int) (*types.Transaction, error) {
	return _IotCloudlab.contract.Transact(opts, "storeHumidity", Humidity)
}

// StoreHumidity is a paid mutator transaction binding the contract method 0x38b3e81d.
//
// Solidity: function storeHumidity(uint256 Humidity) returns()
func (_IotCloudlab *IotCloudlabSession) StoreHumidity(Humidity *big.Int) (*types.Transaction, error) {
	return _IotCloudlab.Contract.StoreHumidity(&_IotCloudlab.TransactOpts, Humidity)
}

// StoreHumidity is a paid mutator transaction binding the contract method 0x38b3e81d.
//
// Solidity: function storeHumidity(uint256 Humidity) returns()
func (_IotCloudlab *IotCloudlabTransactorSession) StoreHumidity(Humidity *big.Int) (*types.Transaction, error) {
	return _IotCloudlab.Contract.StoreHumidity(&_IotCloudlab.TransactOpts, Humidity)
}

// StoreTemperature is a paid mutator transaction binding the contract method 0xdad1b4e9.
//
// Solidity: function storeTemperature(uint256 Temperature) returns()
func (_IotCloudlab *IotCloudlabTransactor) StoreTemperature(opts *bind.TransactOpts, Temperature *big.Int) (*types.Transaction, error) {
	return _IotCloudlab.contract.Transact(opts, "storeTemperature", Temperature)
}

// StoreTemperature is a paid mutator transaction binding the contract method 0xdad1b4e9.
//
// Solidity: function storeTemperature(uint256 Temperature) returns()
func (_IotCloudlab *IotCloudlabSession) StoreTemperature(Temperature *big.Int) (*types.Transaction, error) {
	return _IotCloudlab.Contract.StoreTemperature(&_IotCloudlab.TransactOpts, Temperature)
}

// StoreTemperature is a paid mutator transaction binding the contract method 0xdad1b4e9.
//
// Solidity: function storeTemperature(uint256 Temperature) returns()
func (_IotCloudlab *IotCloudlabTransactorSession) StoreTemperature(Temperature *big.Int) (*types.Transaction, error) {
	return _IotCloudlab.Contract.StoreTemperature(&_IotCloudlab.TransactOpts, Temperature)
}
