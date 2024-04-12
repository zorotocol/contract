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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"NewDeadline\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"Purchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDeadline\",\"type\":\"uint256\"}],\"name\":\"incDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"purchase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newState\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040525f805460ff19169055348015610018575f80fd5b506040516111a13803806111a183398101604081905261003791610143565b5f8054610100600160a81b031916336101000217905560036100598382610226565b5060046100668282610226565b50505f600555506002805460ff19908116909155600780549091169055426006556102e5565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126100af575f80fd5b81516001600160401b03808211156100c9576100c961008c565b604051601f8301601f19908116603f011681019082821181831017156100f1576100f161008c565b816040528381526020925086602085880101111561010d575f80fd5b5f91505b8382101561012e5785820183015181830184015290820190610111565b5f602085830101528094505050505092915050565b5f8060408385031215610154575f80fd5b82516001600160401b038082111561016a575f80fd5b610176868387016100a0565b9350602085015191508082111561018b575f80fd5b50610198858286016100a0565b9150509250929050565b600181811c908216806101b657607f821691505b6020821081036101d457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561022157805f5260205f20601f840160051c810160208510156101ff5750805b601f840160051c820191505b8181101561021e575f815560010161020b565b50505b505050565b81516001600160401b0381111561023f5761023f61008c565b6102538161024d84546101a2565b846101da565b602080601f831160018114610286575f841561026f5750858301515b5f19600386901b1c1916600185901b1785556102dd565b5f85815260208120601f198616915b828110156102b457888601518255948401946001909101908401610295565b50858210156102d157878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b610eaf806102f25f395ff3fe608060405260043610610117575f3560e01c8063700efb281161009f578063a9059cbb11610063578063a9059cbb146103a1578063bedb86fb146103c0578063da9581b6146103df578063dd62ed3e146103fe578063f2fde38b1461044257610158565b8063700efb28146102e057806370a08231146102ff578063867904b4146103335780638da5cb5b1461035257806395d89b411461038d57610158565b806327e235e3116100e657806327e235e314610229578063313ce567146102545780635c6581651461027f5780635f8d96de146102b55780636805b84b146102c957610158565b806306fdde031461018e578063095ea7b3146101b857806318160ddd146101e757806323b872dd1461020a57610158565b366101585760405162461bcd60e51b815260206004820152600a6024820152696e6f207265636569766560b01b60448201526064015b60405180910390fd5b005b60405162461bcd60e51b815260206004820152600b60248201526a6e6f2066616c6c6261636b60a81b604482015260640161014d565b348015610199575f80fd5b506101a2610461565b6040516101af9190610b54565b60405180910390f35b3480156101c3575f80fd5b506101d76101d2366004610b81565b6104ed565b60405190151581526020016101af565b3480156101f2575f80fd5b506101fc60055481565b6040519081526020016101af565b348015610215575f80fd5b506101d7610224366004610ba9565b610583565b348015610234575f80fd5b506101fc610243366004610be2565b60016020525f908152604090205481565b34801561025f575f80fd5b5060025461026d9060ff1681565b60405160ff90911681526020016101af565b34801561028a575f80fd5b506101fc610299366004610bfb565b600860209081525f928352604080842090915290825290205481565b3480156102c0575f80fd5b506006546101fc565b3480156102d4575f80fd5b5060075460ff166101d7565b3480156102eb575f80fd5b506101566102fa366004610c2c565b6106b6565b34801561030a575f80fd5b506101fc610319366004610be2565b6001600160a01b03165f9081526001602052604090205490565b34801561033e575f80fd5b5061015661034d366004610b81565b610726565b34801561035d575f80fd5b505f546103759061010090046001600160a01b031681565b6040516001600160a01b0390911681526020016101af565b348015610398575f80fd5b506101a26107d0565b3480156103ac575f80fd5b506101d76103bb366004610b81565b6107dd565b3480156103cb575f80fd5b506101566103da366004610c43565b6107f2565b3480156103ea575f80fd5b506101fc6103f9366004610c7d565b610863565b348015610409575f80fd5b506101fc610418366004610bfb565b6001600160a01b039182165f90815260086020908152604080832093909416825291909152205490565b34801561044d575f80fd5b5061015661045c366004610be2565b6109b3565b6003805461046e90610d32565b80601f016020809104026020016040519081016040528092919081815260200182805461049a90610d32565b80156104e55780601f106104bc576101008083540402835291602001916104e5565b820191905f5260205f20905b8154815290600101906020018083116104c857829003601f168201915b505050505081565b5f805460ff16156105105760405162461bcd60e51b815260040161014d90610d6a565b5f805460ff19166001178155338082526008602090815260408084206001600160a01b03881680865290835293819020869055518581527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a350505f805460ff1916905550600190565b5f805460ff16156105a65760405162461bcd60e51b815260040161014d90610d6a565b5f805460ff191660011781556001600160a01b03851681526008602090815260408083203384529091529020546105dd90836109ff565b6001600160a01b0385165f8181526008602090815260408083203384528252808320949094559181526001909152205461061790836109ff565b6001600160a01b038086165f9081526001602052604080822093909355908516815220546106459083610a1a565b6001600160a01b038481165f818152600160209081526040918290209490945580519288168352928201529081018390527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060600160405180910390a15060015f805460ff191690559392505050565b5f5461010090046001600160a01b031633146106d0575f80fd5b80600654036106dd575f80fd5b6006546106ea9082610a1a565b60068190556040519081527f41f916eb2a7065889bf3a2125806adaa5ca5a1c9be36f179340901abbf666d23906020015b60405180910390a150565b5f5461010090046001600160a01b03163314610740575f80fd5b6001600160a01b0382165f908152600160205260409020546107629082610a1a565b6001600160a01b0383165f908152600160205260409020556005546107879082610a1a565b600555604080516001600160a01b0384168152602081018390527fc65a3f767206d2fdcede0b094a4840e01c0dd0be1888b5ba800346eaa0123c16910160405180910390a15050565b6004805461046e90610d32565b5f6107e9338484610a37565b90505b92915050565b5f5461010090046001600160a01b0316331461080c575f80fd5b60075481151560ff909116151503610822575f80fd5b6007805460ff19168215159081179091556040519081527f0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd29060200161071b565b6007545f9060ff16156108a15760405162461bcd60e51b81526020600482015260066024820152651c185d5cd95960d21b604482015260640161014d565b5f5460ff16156108c35760405162461bcd60e51b815260040161014d90610d6a565b5f805460ff19166001908117825533825260205260409020546108e690846109ff565b335f9081526001602052604081209190915561090e4261090886610e10610ade565b90610a1a565b905080600654116109555760405162461bcd60e51b8152602060048201526011602482015270191958591b1a5b9948195e18d959591959607a1b604482015260640161014d565b60055461096290856109ff565b6005556040517f6d5507734d4d9e8592119c9e6b3e2bf7eab8bd6252f36e2fa05a5588442479da9061099b903390879085908890610d99565b60405180910390a15f805460ff191690559392505050565b5f5461010090046001600160a01b031633146109cd575f80fd5b6001600160a01b038116156109fc575f8054610100600160a81b0319166101006001600160a01b038416021790555b50565b5f82821115610a1057610a10610dcf565b6107e98284610df7565b5f610a258284610e0a565b9050828110156107ec576107ec610dcf565b5f805460ff1615610a5a5760405162461bcd60e51b815260040161014d90610d6a565b5f805460ff191660011790556001600160a01b038316610abc5760405162461bcd60e51b815260206004820152601860248201527f696e76616c696420726563656976657220616464726573730000000000000000604482015260640161014d565b6001600160a01b0384165f9081526001602052604090205461061790836109ff565b5f825f03610aed57505f6107ec565b610af78284610e1d565b905081610b048483610e34565b146107ec576107ec610dcf565b5f81518084525f5b81811015610b3557602081850181015186830182015201610b19565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f6107e96020830184610b11565b80356001600160a01b0381168114610b7c575f80fd5b919050565b5f8060408385031215610b92575f80fd5b610b9b83610b66565b946020939093013593505050565b5f805f60608486031215610bbb575f80fd5b610bc484610b66565b9250610bd260208501610b66565b9150604084013590509250925092565b5f60208284031215610bf2575f80fd5b6107e982610b66565b5f8060408385031215610c0c575f80fd5b610c1583610b66565b9150610c2360208401610b66565b90509250929050565b5f60208284031215610c3c575f80fd5b5035919050565b5f60208284031215610c53575f80fd5b81358015158114610c62575f80fd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b5f8060408385031215610c8e575f80fd5b82359150602083013567ffffffffffffffff80821115610cac575f80fd5b818501915085601f830112610cbf575f80fd5b813581811115610cd157610cd1610c69565b604051601f8201601f19908116603f01168101908382118183101715610cf957610cf9610c69565b81604052828152886020848701011115610d11575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b600181811c90821680610d4657607f821691505b602082108103610d6457634e487b7160e01b5f52602260045260245ffd5b50919050565b6020808252601590820152741b9bc81c99595b9d1c985b98de48185b1b1bddd959605a1b604082015260600190565b60018060a01b0385168152836020820152826040820152608060608201525f610dc56080830184610b11565b9695505050505050565b634e487b7160e01b5f52600160045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b818103818111156107ec576107ec610de3565b808201808211156107ec576107ec610de3565b80820281158282048414176107ec576107ec610de3565b5f82610e4e57634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220a3cd9c4673ef14f6300ac11fea1484db893868b1fee947d485ecdee4db4c585964736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address spender, address delegate) view returns(uint256)
func (_Contract *ContractCaller) Allowance(opts *bind.CallOpts, spender common.Address, delegate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowance", spender, delegate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address spender, address delegate) view returns(uint256)
func (_Contract *ContractSession) Allowance(spender common.Address, delegate common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, spender, delegate)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address spender, address delegate) view returns(uint256)
func (_Contract *ContractCallerSession) Allowance(spender common.Address, delegate common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, spender, delegate)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Contract *ContractCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Contract *ContractSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowed(&_Contract.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Contract *ContractCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowed(&_Contract.CallOpts, arg0, arg1)
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

// GetDeadline is a free data retrieval call binding the contract method 0x5f8d96de.
//
// Solidity: function getDeadline() view returns(uint256)
func (_Contract *ContractCaller) GetDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDeadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeadline is a free data retrieval call binding the contract method 0x5f8d96de.
//
// Solidity: function getDeadline() view returns(uint256)
func (_Contract *ContractSession) GetDeadline() (*big.Int, error) {
	return _Contract.Contract.GetDeadline(&_Contract.CallOpts)
}

// GetDeadline is a free data retrieval call binding the contract method 0x5f8d96de.
//
// Solidity: function getDeadline() view returns(uint256)
func (_Contract *ContractCallerSession) GetDeadline() (*big.Int, error) {
	return _Contract.Contract.GetDeadline(&_Contract.CallOpts)
}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_Contract *ContractCaller) GetPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_Contract *ContractSession) GetPaused() (bool, error) {
	return _Contract.Contract.GetPaused(&_Contract.CallOpts)
}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_Contract *ContractCallerSession) GetPaused() (bool, error) {
	return _Contract.Contract.GetPaused(&_Contract.CallOpts)
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

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 amount) returns(bool)
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, delegate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", delegate, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 amount) returns(bool)
func (_Contract *ContractSession) Approve(delegate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, delegate, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 amount) returns(bool)
func (_Contract *ContractTransactorSession) Approve(delegate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, delegate, amount)
}

// IncDeadline is a paid mutator transaction binding the contract method 0x700efb28.
//
// Solidity: function incDeadline(uint256 newDeadline) returns()
func (_Contract *ContractTransactor) IncDeadline(opts *bind.TransactOpts, newDeadline *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "incDeadline", newDeadline)
}

// IncDeadline is a paid mutator transaction binding the contract method 0x700efb28.
//
// Solidity: function incDeadline(uint256 newDeadline) returns()
func (_Contract *ContractSession) IncDeadline(newDeadline *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncDeadline(&_Contract.TransactOpts, newDeadline)
}

// IncDeadline is a paid mutator transaction binding the contract method 0x700efb28.
//
// Solidity: function incDeadline(uint256 newDeadline) returns()
func (_Contract *ContractTransactorSession) IncDeadline(newDeadline *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncDeadline(&_Contract.TransactOpts, newDeadline)
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

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool newState) returns()
func (_Contract *ContractTransactor) SetPause(opts *bind.TransactOpts, newState bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPause", newState)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool newState) returns()
func (_Contract *ContractSession) SetPause(newState bool) (*types.Transaction, error) {
	return _Contract.Contract.SetPause(&_Contract.TransactOpts, newState)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool newState) returns()
func (_Contract *ContractTransactorSession) SetPause(newState bool) (*types.Transaction, error) {
	return _Contract.Contract.SetPause(&_Contract.TransactOpts, newState)
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

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address spender, address buyer, uint256 numTokens) returns(bool)
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, spender common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", spender, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address spender, address buyer, uint256 numTokens) returns(bool)
func (_Contract *ContractSession) TransferFrom(spender common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, spender, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address spender, address buyer, uint256 numTokens) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom(spender common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, spender, buyer, numTokens)
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

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

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
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
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
		it.Event = new(ContractApproval)
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
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// ContractNewDeadlineIterator is returned from FilterNewDeadline and is used to iterate over the raw logs and unpacked data for NewDeadline events raised by the Contract contract.
type ContractNewDeadlineIterator struct {
	Event *ContractNewDeadline // Event containing the contract specifics and raw log

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
func (it *ContractNewDeadlineIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewDeadline)
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
		it.Event = new(ContractNewDeadline)
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
func (it *ContractNewDeadlineIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewDeadlineIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewDeadline represents a NewDeadline event raised by the Contract contract.
type ContractNewDeadline struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewDeadline is a free log retrieval operation binding the contract event 0x41f916eb2a7065889bf3a2125806adaa5ca5a1c9be36f179340901abbf666d23.
//
// Solidity: event NewDeadline(uint256 arg0)
func (_Contract *ContractFilterer) FilterNewDeadline(opts *bind.FilterOpts) (*ContractNewDeadlineIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewDeadline")
	if err != nil {
		return nil, err
	}
	return &ContractNewDeadlineIterator{contract: _Contract.contract, event: "NewDeadline", logs: logs, sub: sub}, nil
}

// WatchNewDeadline is a free log subscription operation binding the contract event 0x41f916eb2a7065889bf3a2125806adaa5ca5a1c9be36f179340901abbf666d23.
//
// Solidity: event NewDeadline(uint256 arg0)
func (_Contract *ContractFilterer) WatchNewDeadline(opts *bind.WatchOpts, sink chan<- *ContractNewDeadline) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewDeadline")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewDeadline)
				if err := _Contract.contract.UnpackLog(event, "NewDeadline", log); err != nil {
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

// ParseNewDeadline is a log parse operation binding the contract event 0x41f916eb2a7065889bf3a2125806adaa5ca5a1c9be36f179340901abbf666d23.
//
// Solidity: event NewDeadline(uint256 arg0)
func (_Contract *ContractFilterer) ParseNewDeadline(log types.Log) (*ContractNewDeadline, error) {
	event := new(ContractNewDeadline)
	if err := _Contract.contract.UnpackLog(event, "NewDeadline", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Contract contract.
type ContractPausedIterator struct {
	Event *ContractPaused // Event containing the contract specifics and raw log

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
func (it *ContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPaused)
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
		it.Event = new(ContractPaused)
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
func (it *ContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPaused represents a Paused event raised by the Contract contract.
type ContractPaused struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool arg0)
func (_Contract *ContractFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractPausedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractPausedIterator{contract: _Contract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool arg0)
func (_Contract *ContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractPaused) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPaused)
				if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool arg0)
func (_Contract *ContractFilterer) ParsePaused(log types.Log) (*ContractPaused, error) {
	event := new(ContractPaused)
	if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
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
