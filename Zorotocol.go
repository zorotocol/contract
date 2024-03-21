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
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"Purchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"purchase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040525f805460ff19169055348015610018575f80fd5b50604051610d0d380380610d0d83398101604081905261003791610133565b5f8054610100600160a81b031916336101000217905560036100598382610216565b5060046100668282610216565b50505f600555506002805460ff191690556102d5565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261009f575f80fd5b81516001600160401b03808211156100b9576100b961007c565b604051601f8301601f19908116603f011681019082821181831017156100e1576100e161007c565b81604052838152602092508660208588010111156100fd575f80fd5b5f91505b8382101561011e5785820183015181830184015290820190610101565b5f602085830101528094505050505092915050565b5f8060408385031215610144575f80fd5b82516001600160401b038082111561015a575f80fd5b61016686838701610090565b9350602085015191508082111561017b575f80fd5b5061018885828601610090565b9150509250929050565b600181811c908216806101a657607f821691505b6020821081036101c457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561021157805f5260205f20601f840160051c810160208510156101ef5750805b601f840160051c820191505b8181101561020e575f81556001016101fb565b50505b505050565b81516001600160401b0381111561022f5761022f61007c565b6102438161023d8454610192565b846101ca565b602080601f831160018114610276575f841561025f5750858301515b5f19600386901b1c1916600185901b1785556102cd565b5f85815260208120601f198616915b828110156102a457888601518255948401946001909101908401610285565b50858210156102c157878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b610a2b806102e25f395ff3fe60806040526004361061009f575f3560e01c8063867904b411610063578063867904b4146101ed5780638da5cb5b1461020c57806395d89b4114610247578063a9059cbb1461025b578063da9581b61461028a578063f2fde38b146102a9576100e0565b806306fdde031461011657806318160ddd1461014057806327e235e314610163578063313ce5671461018e57806370a08231146101b9576100e0565b366100e05760405162461bcd60e51b815260206004820152600a6024820152696e6f207265636569766560b01b60448201526064015b60405180910390fd5b005b60405162461bcd60e51b815260206004820152600b60248201526a6e6f2066616c6c6261636b60a81b60448201526064016100d5565b348015610121575f80fd5b5061012a6102c8565b60405161013791906107a6565b60405180910390f35b34801561014b575f80fd5b5061015560055481565b604051908152602001610137565b34801561016e575f80fd5b5061015561017d3660046107d3565b60016020525f908152604090205481565b348015610199575f80fd5b506002546101a79060ff1681565b60405160ff9091168152602001610137565b3480156101c4575f80fd5b506101556101d33660046107d3565b6001600160a01b03165f9081526001602052604090205490565b3480156101f8575f80fd5b506100de6102073660046107ec565b610354565b348015610217575f80fd5b505f5461022f9061010090046001600160a01b031681565b6040516001600160a01b039091168152602001610137565b348015610252575f80fd5b5061012a6103fe565b348015610266575f80fd5b5061027a6102753660046107ec565b61040b565b6040519015158152602001610137565b348015610295575f80fd5b506101556102a4366004610828565b610420565b3480156102b4575f80fd5b506100de6102c33660046107d3565b610517565b600380546102d5906108dd565b80601f0160208091040260200160405190810160405280929190818152602001828054610301906108dd565b801561034c5780601f106103235761010080835404028352916020019161034c565b820191905f5260205f20905b81548152906001019060200180831161032f57829003601f168201915b505050505081565b5f5461010090046001600160a01b0316331461036e575f80fd5b6001600160a01b0382165f908152600160205260409020546103909082610563565b6001600160a01b0383165f908152600160205260409020556005546103b59082610563565b600555604080516001600160a01b0384168152602081018390527fc65a3f767206d2fdcede0b094a4840e01c0dd0be1888b5ba800346eaa0123c16910160405180910390a15050565b600480546102d5906108dd565b5f610417338484610580565b90505b92915050565b5f805460ff161561046b5760405162461bcd60e51b81526020600482015260156024820152741b9bc81c99595b9d1c985b98de48185b1b1bddd959605a1b60448201526064016100d5565b5f805460ff191660019081178255338252602052604090205461048e9084610715565b335f908152600160205260408120919091556104b6426104b086610e10610730565b90610563565b6005549091506104c69085610715565b6005556040517f6d5507734d4d9e8592119c9e6b3e2bf7eab8bd6252f36e2fa05a5588442479da906104ff903390879085908890610915565b60405180910390a15f805460ff191690559392505050565b5f5461010090046001600160a01b03163314610531575f80fd5b6001600160a01b03811615610560575f8054610100600160a81b0319166101006001600160a01b038416021790555b50565b5f61056e828461095f565b90508281101561041a5761041a610972565b5f805460ff16156105cb5760405162461bcd60e51b81526020600482015260156024820152741b9bc81c99595b9d1c985b98de48185b1b1bddd959605a1b60448201526064016100d5565b5f805460ff191660011790556001600160a01b03831661062d5760405162461bcd60e51b815260206004820152601860248201527f696e76616c69642072656365697665722061646472657373000000000000000060448201526064016100d5565b6001600160a01b0384165f9081526001602052604090205482111561065357505f610705565b6001600160a01b0384165f908152600160205260409020546106759083610715565b6001600160a01b038086165f9081526001602052604080822093909355908516815220546106a39083610563565b6001600160a01b038481165f818152600160209081526040918290209490945580519288168352928201529081018390527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060600160405180910390a15060015b5f805460ff191690559392505050565b5f8282111561072657610726610972565b6104178284610986565b5f825f0361073f57505f61041a565b6107498284610999565b90508161075684836109b0565b1461041a5761041a610972565b5f81518084525f5b818110156107875760208185018101518683018201520161076b565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f6104176020830184610763565b80356001600160a01b03811681146107ce575f80fd5b919050565b5f602082840312156107e3575f80fd5b610417826107b8565b5f80604083850312156107fd575f80fd5b610806836107b8565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b5f8060408385031215610839575f80fd5b82359150602083013567ffffffffffffffff80821115610857575f80fd5b818501915085601f83011261086a575f80fd5b81358181111561087c5761087c610814565b604051601f8201601f19908116603f011681019083821181831017156108a4576108a4610814565b816040528281528860208487010111156108bc575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b600181811c908216806108f157607f821691505b60208210810361090f57634e487b7160e01b5f52602260045260245ffd5b50919050565b60018060a01b0385168152836020820152826040820152608060608201525f6109416080830184610763565b9695505050505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561041a5761041a61094b565b634e487b7160e01b5f52600160045260245ffd5b8181038181111561041a5761041a61094b565b808202811582820484141761041a5761041a61094b565b5f826109ca57634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220e64786a4cb61dee711965b6c196beefbf98505e427a099c643bcf5d32ead198364736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

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
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contract *ContractCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contract *ContractSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractCallerSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address _account, uint256 _hours) returns()
func (_Contract *ContractTransactor) Issue(opts *bind.TransactOpts, _account common.Address, _hours *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "issue", _account, _hours)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address _account, uint256 _hours) returns()
func (_Contract *ContractSession) Issue(_account common.Address, _hours *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Issue(&_Contract.TransactOpts, _account, _hours)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address _account, uint256 _hours) returns()
func (_Contract *ContractTransactorSession) Issue(_account common.Address, _hours *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Issue(&_Contract.TransactOpts, _account, _hours)
}

// Purchase is a paid mutator transaction binding the contract method 0xda9581b6.
//
// Solidity: function purchase(uint256 _hours, string email) returns(uint256)
func (_Contract *ContractTransactor) Purchase(opts *bind.TransactOpts, _hours *big.Int, email string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "purchase", _hours, email)
}

// Purchase is a paid mutator transaction binding the contract method 0xda9581b6.
//
// Solidity: function purchase(uint256 _hours, string email) returns(uint256)
func (_Contract *ContractSession) Purchase(_hours *big.Int, email string) (*types.Transaction, error) {
	return _Contract.Contract.Purchase(&_Contract.TransactOpts, _hours, email)
}

// Purchase is a paid mutator transaction binding the contract method 0xda9581b6.
//
// Solidity: function purchase(uint256 _hours, string email) returns(uint256)
func (_Contract *ContractTransactorSession) Purchase(_hours *big.Int, email string) (*types.Transaction, error) {
	return _Contract.Contract.Purchase(&_Contract.TransactOpts, _hours, email)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _hours) returns(bool)
func (_Contract *ContractTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _hours *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transfer", _to, _hours)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _hours) returns(bool)
func (_Contract *ContractSession) Transfer(_to common.Address, _hours *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, _to, _hours)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _hours) returns(bool)
func (_Contract *ContractTransactorSession) Transfer(_to common.Address, _hours *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, _to, _hours)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, _newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Contract *ContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Contract *ContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Contract.Contract.Fallback(&_Contract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Contract *ContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Contract.Contract.Fallback(&_Contract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractIssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the Contract contract.
type ContractIssueIterator struct {
	Event *ContractIssue // Event containing the contract specifics and raw log

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
func (it *ContractIssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIssue)
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
		it.Event = new(ContractIssue)
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
func (it *ContractIssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIssue represents a Issue event raised by the Contract contract.
type ContractIssue struct {
	Account common.Address
	Hours   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0xc65a3f767206d2fdcede0b094a4840e01c0dd0be1888b5ba800346eaa0123c16.
//
// Solidity: event Issue(address _account, uint256 _hours)
func (_Contract *ContractFilterer) FilterIssue(opts *bind.FilterOpts) (*ContractIssueIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return &ContractIssueIterator{contract: _Contract.contract, event: "Issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0xc65a3f767206d2fdcede0b094a4840e01c0dd0be1888b5ba800346eaa0123c16.
//
// Solidity: event Issue(address _account, uint256 _hours)
func (_Contract *ContractFilterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *ContractIssue) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIssue)
				if err := _Contract.contract.UnpackLog(event, "Issue", log); err != nil {
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

// ParseIssue is a log parse operation binding the contract event 0xc65a3f767206d2fdcede0b094a4840e01c0dd0be1888b5ba800346eaa0123c16.
//
// Solidity: event Issue(address _account, uint256 _hours)
func (_Contract *ContractFilterer) ParseIssue(log types.Log) (*ContractIssue, error) {
	event := new(ContractIssue)
	if err := _Contract.contract.UnpackLog(event, "Issue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPurchaseIterator is returned from FilterPurchase and is used to iterate over the raw logs and unpacked data for Purchase events raised by the Contract contract.
type ContractPurchaseIterator struct {
	Event *ContractPurchase // Event containing the contract specifics and raw log

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
func (it *ContractPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPurchase)
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
		it.Event = new(ContractPurchase)
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
func (it *ContractPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPurchase represents a Purchase event raised by the Contract contract.
type ContractPurchase struct {
	Payer    common.Address
	Hours    *big.Int
	Deadline *big.Int
	Email    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPurchase is a free log retrieval operation binding the contract event 0x6d5507734d4d9e8592119c9e6b3e2bf7eab8bd6252f36e2fa05a5588442479da.
//
// Solidity: event Purchase(address _payer, uint256 _hours, uint256 deadline, string email)
func (_Contract *ContractFilterer) FilterPurchase(opts *bind.FilterOpts) (*ContractPurchaseIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Purchase")
	if err != nil {
		return nil, err
	}
	return &ContractPurchaseIterator{contract: _Contract.contract, event: "Purchase", logs: logs, sub: sub}, nil
}

// WatchPurchase is a free log subscription operation binding the contract event 0x6d5507734d4d9e8592119c9e6b3e2bf7eab8bd6252f36e2fa05a5588442479da.
//
// Solidity: event Purchase(address _payer, uint256 _hours, uint256 deadline, string email)
func (_Contract *ContractFilterer) WatchPurchase(opts *bind.WatchOpts, sink chan<- *ContractPurchase) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Purchase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPurchase)
				if err := _Contract.contract.UnpackLog(event, "Purchase", log); err != nil {
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

// ParsePurchase is a log parse operation binding the contract event 0x6d5507734d4d9e8592119c9e6b3e2bf7eab8bd6252f36e2fa05a5588442479da.
//
// Solidity: event Purchase(address _payer, uint256 _hours, uint256 deadline, string email)
func (_Contract *ContractFilterer) ParsePurchase(log types.Log) (*ContractPurchase, error) {
	event := new(ContractPurchase)
	if err := _Contract.contract.UnpackLog(event, "Purchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

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
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
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
		it.Event = new(ContractTransfer)
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
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address arg0, address arg1, uint256 arg2)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts) (*ContractTransferIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address arg0, address arg1, uint256 arg2)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address arg0, address arg1, uint256 arg2)
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
