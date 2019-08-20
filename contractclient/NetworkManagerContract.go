// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lition

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LitionABI is the input ABI used to generate the binding from.
const LitionABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"notary_block\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"miners\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"blocks_mined\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"user_gas\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"largest_tx\",\"type\":\"uint32\"}],\"name\":\"get_signature_hash_from_notary\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"}],\"name\":\"get_signatures_count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodesCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"r\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"e\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"r\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"e\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"get_signatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"i\",\"type\":\"uint16\"}],\"name\":\"getNodeList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"r\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"e\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"r\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"e\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"updateNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"store_signature\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nodeName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"publickey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"enode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"print\",\"type\":\"event\"}]"

// Lition is an auto generated Go binding around an Ethereum contract.
type Lition struct {
	LitionCaller     // Read-only binding to the contract
	LitionTransactor // Write-only binding to the contract
	LitionFilterer   // Log filterer for contract events
}

// LitionCaller is an auto generated read-only Go binding around an Ethereum contract.
type LitionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LitionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LitionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LitionSession struct {
	Contract     *Lition           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LitionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LitionCallerSession struct {
	Contract *LitionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LitionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LitionTransactorSession struct {
	Contract     *LitionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LitionRaw is an auto generated low-level Go binding around an Ethereum contract.
type LitionRaw struct {
	Contract *Lition // Generic contract binding to access the raw methods on
}

// LitionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LitionCallerRaw struct {
	Contract *LitionCaller // Generic read-only contract binding to access the raw methods on
}

// LitionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LitionTransactorRaw struct {
	Contract *LitionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLition creates a new instance of Lition, bound to a specific deployed contract.
func NewLition(address common.Address, backend bind.ContractBackend) (*Lition, error) {
	contract, err := bindLition(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lition{LitionCaller: LitionCaller{contract: contract}, LitionTransactor: LitionTransactor{contract: contract}, LitionFilterer: LitionFilterer{contract: contract}}, nil
}

// NewLitionCaller creates a new read-only instance of Lition, bound to a specific deployed contract.
func NewLitionCaller(address common.Address, caller bind.ContractCaller) (*LitionCaller, error) {
	contract, err := bindLition(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LitionCaller{contract: contract}, nil
}

// NewLitionTransactor creates a new write-only instance of Lition, bound to a specific deployed contract.
func NewLitionTransactor(address common.Address, transactor bind.ContractTransactor) (*LitionTransactor, error) {
	contract, err := bindLition(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LitionTransactor{contract: contract}, nil
}

// NewLitionFilterer creates a new log filterer instance of Lition, bound to a specific deployed contract.
func NewLitionFilterer(address common.Address, filterer bind.ContractFilterer) (*LitionFilterer, error) {
	contract, err := bindLition(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LitionFilterer{contract: contract}, nil
}

// bindLition binds a generic wrapper to an already deployed contract.
func bindLition(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LitionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lition *LitionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lition.Contract.LitionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lition *LitionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lition.Contract.LitionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lition *LitionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lition.Contract.LitionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lition *LitionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lition.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lition *LitionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lition.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lition *LitionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lition.Contract.contract.Transact(opts, method, params...)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x7f11a8ed.
//
// Solidity: function getNodeDetails(uint16 _index) constant returns(string n, string r, string p, string ip, string e, uint256 i)
func (_Lition *LitionCaller) GetNodeDetails(opts *bind.CallOpts, _index uint16) (struct {
	N  string
	R  string
	P  string
	Ip string
	E  string
	I  *big.Int
}, error) {
	ret := new(struct {
		N  string
		R  string
		P  string
		Ip string
		E  string
		I  *big.Int
	})
	out := ret
	err := _Lition.contract.Call(opts, out, "getNodeDetails", _index)
	return *ret, err
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x7f11a8ed.
//
// Solidity: function getNodeDetails(uint16 _index) constant returns(string n, string r, string p, string ip, string e, uint256 i)
func (_Lition *LitionSession) GetNodeDetails(_index uint16) (struct {
	N  string
	R  string
	P  string
	Ip string
	E  string
	I  *big.Int
}, error) {
	return _Lition.Contract.GetNodeDetails(&_Lition.CallOpts, _index)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x7f11a8ed.
//
// Solidity: function getNodeDetails(uint16 _index) constant returns(string n, string r, string p, string ip, string e, uint256 i)
func (_Lition *LitionCallerSession) GetNodeDetails(_index uint16) (struct {
	N  string
	R  string
	P  string
	Ip string
	E  string
	I  *big.Int
}, error) {
	return _Lition.Contract.GetNodeDetails(&_Lition.CallOpts, _index)
}

// GetNodeList is a free data retrieval call binding the contract method 0xdeb043c6.
//
// Solidity: function getNodeList(uint16 i) constant returns(string n, string r, string p, string ip, string e)
func (_Lition *LitionCaller) GetNodeList(opts *bind.CallOpts, i uint16) (struct {
	N  string
	R  string
	P  string
	Ip string
	E  string
}, error) {
	ret := new(struct {
		N  string
		R  string
		P  string
		Ip string
		E  string
	})
	out := ret
	err := _Lition.contract.Call(opts, out, "getNodeList", i)
	return *ret, err
}

// GetNodeList is a free data retrieval call binding the contract method 0xdeb043c6.
//
// Solidity: function getNodeList(uint16 i) constant returns(string n, string r, string p, string ip, string e)
func (_Lition *LitionSession) GetNodeList(i uint16) (struct {
	N  string
	R  string
	P  string
	Ip string
	E  string
}, error) {
	return _Lition.Contract.GetNodeList(&_Lition.CallOpts, i)
}

// GetNodeList is a free data retrieval call binding the contract method 0xdeb043c6.
//
// Solidity: function getNodeList(uint16 i) constant returns(string n, string r, string p, string ip, string e)
func (_Lition *LitionCallerSession) GetNodeList(i uint16) (struct {
	N  string
	R  string
	P  string
	Ip string
	E  string
}, error) {
	return _Lition.Contract.GetNodeList(&_Lition.CallOpts, i)
}

// GetNodesCounter is a free data retrieval call binding the contract method 0x6168d293.
//
// Solidity: function getNodesCounter() constant returns(uint256)
func (_Lition *LitionCaller) GetNodesCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lition.contract.Call(opts, out, "getNodesCounter")
	return *ret0, err
}

// GetNodesCounter is a free data retrieval call binding the contract method 0x6168d293.
//
// Solidity: function getNodesCounter() constant returns(uint256)
func (_Lition *LitionSession) GetNodesCounter() (*big.Int, error) {
	return _Lition.Contract.GetNodesCounter(&_Lition.CallOpts)
}

// GetNodesCounter is a free data retrieval call binding the contract method 0x6168d293.
//
// Solidity: function getNodesCounter() constant returns(uint256)
func (_Lition *LitionCallerSession) GetNodesCounter() (*big.Int, error) {
	return _Lition.Contract.GetNodesCounter(&_Lition.CallOpts)
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0x2aad5f6e.
//
// Solidity: function get_signature_hash_from_notary(uint256 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx) constant returns(bytes32)
func (_Lition *LitionCaller) GetSignatureHashFromNotary(opts *bind.CallOpts, notary_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Lition.contract.Call(opts, out, "get_signature_hash_from_notary", notary_block, miners, blocks_mined, users, user_gas, largest_tx)
	return *ret0, err
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0x2aad5f6e.
//
// Solidity: function get_signature_hash_from_notary(uint256 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx) constant returns(bytes32)
func (_Lition *LitionSession) GetSignatureHashFromNotary(notary_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32) ([32]byte, error) {
	return _Lition.Contract.GetSignatureHashFromNotary(&_Lition.CallOpts, notary_block, miners, blocks_mined, users, user_gas, largest_tx)
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0x2aad5f6e.
//
// Solidity: function get_signature_hash_from_notary(uint256 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint32[] user_gas, uint32 largest_tx) constant returns(bytes32)
func (_Lition *LitionCallerSession) GetSignatureHashFromNotary(notary_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32) ([32]byte, error) {
	return _Lition.Contract.GetSignatureHashFromNotary(&_Lition.CallOpts, notary_block, miners, blocks_mined, users, user_gas, largest_tx)
}

// GetSignatures is a free data retrieval call binding the contract method 0xa05b9f85.
//
// Solidity: function get_signatures(uint256 block_no, uint256 index) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_Lition *LitionCaller) GetSignatures(opts *bind.CallOpts, block_no *big.Int, index *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	ret := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	out := ret
	err := _Lition.contract.Call(opts, out, "get_signatures", block_no, index)
	return *ret, err
}

// GetSignatures is a free data retrieval call binding the contract method 0xa05b9f85.
//
// Solidity: function get_signatures(uint256 block_no, uint256 index) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_Lition *LitionSession) GetSignatures(block_no *big.Int, index *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Lition.Contract.GetSignatures(&_Lition.CallOpts, block_no, index)
}

// GetSignatures is a free data retrieval call binding the contract method 0xa05b9f85.
//
// Solidity: function get_signatures(uint256 block_no, uint256 index) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_Lition *LitionCallerSession) GetSignatures(block_no *big.Int, index *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Lition.Contract.GetSignatures(&_Lition.CallOpts, block_no, index)
}

// GetSignaturesCount is a free data retrieval call binding the contract method 0x559c270e.
//
// Solidity: function get_signatures_count(uint256 block_no) constant returns(uint256)
func (_Lition *LitionCaller) GetSignaturesCount(opts *bind.CallOpts, block_no *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lition.contract.Call(opts, out, "get_signatures_count", block_no)
	return *ret0, err
}

// GetSignaturesCount is a free data retrieval call binding the contract method 0x559c270e.
//
// Solidity: function get_signatures_count(uint256 block_no) constant returns(uint256)
func (_Lition *LitionSession) GetSignaturesCount(block_no *big.Int) (*big.Int, error) {
	return _Lition.Contract.GetSignaturesCount(&_Lition.CallOpts, block_no)
}

// GetSignaturesCount is a free data retrieval call binding the contract method 0x559c270e.
//
// Solidity: function get_signatures_count(uint256 block_no) constant returns(uint256)
func (_Lition *LitionCallerSession) GetSignaturesCount(block_no *big.Int) (*big.Int, error) {
	return _Lition.Contract.GetSignaturesCount(&_Lition.CallOpts, block_no)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x82cb1a2a.
//
// Solidity: function registerNode(string n, string r, string p, string e, string ip) returns()
func (_Lition *LitionTransactor) RegisterNode(opts *bind.TransactOpts, n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _Lition.contract.Transact(opts, "registerNode", n, r, p, e, ip)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x82cb1a2a.
//
// Solidity: function registerNode(string n, string r, string p, string e, string ip) returns()
func (_Lition *LitionSession) RegisterNode(n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _Lition.Contract.RegisterNode(&_Lition.TransactOpts, n, r, p, e, ip)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x82cb1a2a.
//
// Solidity: function registerNode(string n, string r, string p, string e, string ip) returns()
func (_Lition *LitionTransactorSession) RegisterNode(n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _Lition.Contract.RegisterNode(&_Lition.TransactOpts, n, r, p, e, ip)
}

// StoreSignature is a paid mutator transaction binding the contract method 0xf202c84d.
//
// Solidity: function store_signature(uint256 block_no, uint8 v, bytes32 r, bytes32 s) returns()
func (_Lition *LitionTransactor) StoreSignature(opts *bind.TransactOpts, block_no *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Lition.contract.Transact(opts, "store_signature", block_no, v, r, s)
}

// StoreSignature is a paid mutator transaction binding the contract method 0xf202c84d.
//
// Solidity: function store_signature(uint256 block_no, uint8 v, bytes32 r, bytes32 s) returns()
func (_Lition *LitionSession) StoreSignature(block_no *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Lition.Contract.StoreSignature(&_Lition.TransactOpts, block_no, v, r, s)
}

// StoreSignature is a paid mutator transaction binding the contract method 0xf202c84d.
//
// Solidity: function store_signature(uint256 block_no, uint8 v, bytes32 r, bytes32 s) returns()
func (_Lition *LitionTransactorSession) StoreSignature(block_no *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Lition.Contract.StoreSignature(&_Lition.TransactOpts, block_no, v, r, s)
}

// UpdateNode is a paid mutator transaction binding the contract method 0xe1d33203.
//
// Solidity: function updateNode(string n, string r, string p, string e, string ip) returns()
func (_Lition *LitionTransactor) UpdateNode(opts *bind.TransactOpts, n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _Lition.contract.Transact(opts, "updateNode", n, r, p, e, ip)
}

// UpdateNode is a paid mutator transaction binding the contract method 0xe1d33203.
//
// Solidity: function updateNode(string n, string r, string p, string e, string ip) returns()
func (_Lition *LitionSession) UpdateNode(n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _Lition.Contract.UpdateNode(&_Lition.TransactOpts, n, r, p, e, ip)
}

// UpdateNode is a paid mutator transaction binding the contract method 0xe1d33203.
//
// Solidity: function updateNode(string n, string r, string p, string e, string ip) returns()
func (_Lition *LitionTransactorSession) UpdateNode(n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _Lition.Contract.UpdateNode(&_Lition.TransactOpts, n, r, p, e, ip)
}

// LitionPrintIterator is returned from FilterPrint and is used to iterate over the raw logs and unpacked data for Print events raised by the Lition contract.
type LitionPrintIterator struct {
	Event *LitionPrint // Event containing the contract specifics and raw log

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
func (it *LitionPrintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionPrint)
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
		it.Event = new(LitionPrint)
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
func (it *LitionPrintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionPrintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionPrint represents a Print event raised by the Lition contract.
type LitionPrint struct {
	NodeName  string
	Role      string
	Publickey string
	Enode     string
	Ip        string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrint is a free log retrieval operation binding the contract event 0x8f48d31c5e32025ea0c67fbf4573ae86f4b46e5bde075c4dca076b5d293ce408.
//
// Solidity: event print(string nodeName, string role, string publickey, string enode, string ip)
func (_Lition *LitionFilterer) FilterPrint(opts *bind.FilterOpts) (*LitionPrintIterator, error) {

	logs, sub, err := _Lition.contract.FilterLogs(opts, "print")
	if err != nil {
		return nil, err
	}
	return &LitionPrintIterator{contract: _Lition.contract, event: "print", logs: logs, sub: sub}, nil
}

// WatchPrint is a free log subscription operation binding the contract event 0x8f48d31c5e32025ea0c67fbf4573ae86f4b46e5bde075c4dca076b5d293ce408.
//
// Solidity: event print(string nodeName, string role, string publickey, string enode, string ip)
func (_Lition *LitionFilterer) WatchPrint(opts *bind.WatchOpts, sink chan<- *LitionPrint) (event.Subscription, error) {

	logs, sub, err := _Lition.contract.WatchLogs(opts, "print")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionPrint)
				if err := _Lition.contract.UnpackLog(event, "print", log); err != nil {
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

// ParsePrint is a log parse operation binding the contract event 0x8f48d31c5e32025ea0c67fbf4573ae86f4b46e5bde075c4dca076b5d293ce408.
//
// Solidity: event print(string nodeName, string role, string publickey, string enode, string ip)
func (_Lition *LitionFilterer) ParsePrint(log types.Log) (*LitionPrint, error) {
	event := new(LitionPrint)
	if err := _Lition.contract.UnpackLog(event, "print", log); err != nil {
		return nil, err
	}
	return event, nil
}
