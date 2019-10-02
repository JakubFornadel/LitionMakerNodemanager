// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package litionScClient

import (
	"math/big"
	"strings"

	ethereum "gitlab.com/lition/lition"
	"gitlab.com/lition/lition/accounts/abi"
	"gitlab.com/lition/lition/accounts/abi/bind"
	"gitlab.com/lition/lition/common"
	"gitlab.com/lition/lition/core/types"
	"gitlab.com/lition/lition/event"
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

// LitionScClientABI is the input ABI used to generate the binding from.
const LitionScClientABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"}],\"name\":\"get_signatures_count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodesCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"r\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"e\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"r\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"e\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"registerNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"get_signatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"notary_block\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"miners\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"blocks_mined\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"user_gas\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"largest_tx\",\"type\":\"uint64\"}],\"name\":\"get_signature_hash_from_notary\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"i\",\"type\":\"uint16\"}],\"name\":\"getNodeList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"r\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"e\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"r\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"e\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"updateNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"store_signature\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nodeName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"publickey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"enode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"print\",\"type\":\"event\"}]"

// LitionScClient is an auto generated Go binding around an Ethereum contract.
type LitionScClient struct {
	LitionScClientCaller     // Read-only binding to the contract
	LitionScClientTransactor // Write-only binding to the contract
	LitionScClientFilterer   // Log filterer for contract events
}

// LitionScClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type LitionScClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionScClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LitionScClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionScClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LitionScClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LitionScClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LitionScClientSession struct {
	Contract     *LitionScClient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LitionScClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LitionScClientCallerSession struct {
	Contract *LitionScClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LitionScClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LitionScClientTransactorSession struct {
	Contract     *LitionScClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LitionScClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type LitionScClientRaw struct {
	Contract *LitionScClient // Generic contract binding to access the raw methods on
}

// LitionScClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LitionScClientCallerRaw struct {
	Contract *LitionScClientCaller // Generic read-only contract binding to access the raw methods on
}

// LitionScClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LitionScClientTransactorRaw struct {
	Contract *LitionScClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLitionScClient creates a new instance of LitionScClient, bound to a specific deployed contract.
func NewLitionScClient(address common.Address, backend bind.ContractBackend) (*LitionScClient, error) {
	contract, err := bindLitionScClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LitionScClient{LitionScClientCaller: LitionScClientCaller{contract: contract}, LitionScClientTransactor: LitionScClientTransactor{contract: contract}, LitionScClientFilterer: LitionScClientFilterer{contract: contract}}, nil
}

// NewLitionScClientCaller creates a new read-only instance of LitionScClient, bound to a specific deployed contract.
func NewLitionScClientCaller(address common.Address, caller bind.ContractCaller) (*LitionScClientCaller, error) {
	contract, err := bindLitionScClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LitionScClientCaller{contract: contract}, nil
}

// NewLitionScClientTransactor creates a new write-only instance of LitionScClient, bound to a specific deployed contract.
func NewLitionScClientTransactor(address common.Address, transactor bind.ContractTransactor) (*LitionScClientTransactor, error) {
	contract, err := bindLitionScClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LitionScClientTransactor{contract: contract}, nil
}

// NewLitionScClientFilterer creates a new log filterer instance of LitionScClient, bound to a specific deployed contract.
func NewLitionScClientFilterer(address common.Address, filterer bind.ContractFilterer) (*LitionScClientFilterer, error) {
	contract, err := bindLitionScClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LitionScClientFilterer{contract: contract}, nil
}

// bindLitionScClient binds a generic wrapper to an already deployed contract.
func bindLitionScClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LitionScClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LitionScClient *LitionScClientRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LitionScClient.Contract.LitionScClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LitionScClient *LitionScClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LitionScClient.Contract.LitionScClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LitionScClient *LitionScClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LitionScClient.Contract.LitionScClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LitionScClient *LitionScClientCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LitionScClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LitionScClient *LitionScClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LitionScClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LitionScClient *LitionScClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LitionScClient.Contract.contract.Transact(opts, method, params...)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x7f11a8ed.
//
// Solidity: function getNodeDetails(uint16 _index) constant returns(string n, string r, string p, string ip, string e, uint256 i)
func (_LitionScClient *LitionScClientCaller) GetNodeDetails(opts *bind.CallOpts, _index uint16) (struct {
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
	err := _LitionScClient.contract.Call(opts, out, "getNodeDetails", _index)
	return *ret, err
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x7f11a8ed.
//
// Solidity: function getNodeDetails(uint16 _index) constant returns(string n, string r, string p, string ip, string e, uint256 i)
func (_LitionScClient *LitionScClientSession) GetNodeDetails(_index uint16) (struct {
	N  string
	R  string
	P  string
	Ip string
	E  string
	I  *big.Int
}, error) {
	return _LitionScClient.Contract.GetNodeDetails(&_LitionScClient.CallOpts, _index)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x7f11a8ed.
//
// Solidity: function getNodeDetails(uint16 _index) constant returns(string n, string r, string p, string ip, string e, uint256 i)
func (_LitionScClient *LitionScClientCallerSession) GetNodeDetails(_index uint16) (struct {
	N  string
	R  string
	P  string
	Ip string
	E  string
	I  *big.Int
}, error) {
	return _LitionScClient.Contract.GetNodeDetails(&_LitionScClient.CallOpts, _index)
}

// GetNodeList is a free data retrieval call binding the contract method 0xdeb043c6.
//
// Solidity: function getNodeList(uint16 i) constant returns(string n, string r, string p, string ip, string e)
func (_LitionScClient *LitionScClientCaller) GetNodeList(opts *bind.CallOpts, i uint16) (struct {
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
	err := _LitionScClient.contract.Call(opts, out, "getNodeList", i)
	return *ret, err
}

// GetNodeList is a free data retrieval call binding the contract method 0xdeb043c6.
//
// Solidity: function getNodeList(uint16 i) constant returns(string n, string r, string p, string ip, string e)
func (_LitionScClient *LitionScClientSession) GetNodeList(i uint16) (struct {
	N  string
	R  string
	P  string
	Ip string
	E  string
}, error) {
	return _LitionScClient.Contract.GetNodeList(&_LitionScClient.CallOpts, i)
}

// GetNodeList is a free data retrieval call binding the contract method 0xdeb043c6.
//
// Solidity: function getNodeList(uint16 i) constant returns(string n, string r, string p, string ip, string e)
func (_LitionScClient *LitionScClientCallerSession) GetNodeList(i uint16) (struct {
	N  string
	R  string
	P  string
	Ip string
	E  string
}, error) {
	return _LitionScClient.Contract.GetNodeList(&_LitionScClient.CallOpts, i)
}

// GetNodesCounter is a free data retrieval call binding the contract method 0x6168d293.
//
// Solidity: function getNodesCounter() constant returns(uint256)
func (_LitionScClient *LitionScClientCaller) GetNodesCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "getNodesCounter")
	return *ret0, err
}

// GetNodesCounter is a free data retrieval call binding the contract method 0x6168d293.
//
// Solidity: function getNodesCounter() constant returns(uint256)
func (_LitionScClient *LitionScClientSession) GetNodesCounter() (*big.Int, error) {
	return _LitionScClient.Contract.GetNodesCounter(&_LitionScClient.CallOpts)
}

// GetNodesCounter is a free data retrieval call binding the contract method 0x6168d293.
//
// Solidity: function getNodesCounter() constant returns(uint256)
func (_LitionScClient *LitionScClientCallerSession) GetNodesCounter() (*big.Int, error) {
	return _LitionScClient.Contract.GetNodesCounter(&_LitionScClient.CallOpts)
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0xa0c795ea.
//
// Solidity: function get_signature_hash_from_notary(uint256 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint64[] user_gas, uint64 largest_tx) constant returns(bytes32)
func (_LitionScClient *LitionScClientCaller) GetSignatureHashFromNotary(opts *bind.CallOpts, notary_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint64, largest_tx uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "get_signature_hash_from_notary", notary_block, miners, blocks_mined, users, user_gas, largest_tx)
	return *ret0, err
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0xa0c795ea.
//
// Solidity: function get_signature_hash_from_notary(uint256 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint64[] user_gas, uint64 largest_tx) constant returns(bytes32)
func (_LitionScClient *LitionScClientSession) GetSignatureHashFromNotary(notary_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint64, largest_tx uint64) ([32]byte, error) {
	return _LitionScClient.Contract.GetSignatureHashFromNotary(&_LitionScClient.CallOpts, notary_block, miners, blocks_mined, users, user_gas, largest_tx)
}

// GetSignatureHashFromNotary is a free data retrieval call binding the contract method 0xa0c795ea.
//
// Solidity: function get_signature_hash_from_notary(uint256 notary_block, address[] miners, uint32[] blocks_mined, address[] users, uint64[] user_gas, uint64 largest_tx) constant returns(bytes32)
func (_LitionScClient *LitionScClientCallerSession) GetSignatureHashFromNotary(notary_block *big.Int, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint64, largest_tx uint64) ([32]byte, error) {
	return _LitionScClient.Contract.GetSignatureHashFromNotary(&_LitionScClient.CallOpts, notary_block, miners, blocks_mined, users, user_gas, largest_tx)
}

// GetSignatures is a free data retrieval call binding the contract method 0xa05b9f85.
//
// Solidity: function get_signatures(uint256 block_no, uint256 index) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_LitionScClient *LitionScClientCaller) GetSignatures(opts *bind.CallOpts, block_no *big.Int, index *big.Int) (struct {
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
	err := _LitionScClient.contract.Call(opts, out, "get_signatures", block_no, index)
	return *ret, err
}

// GetSignatures is a free data retrieval call binding the contract method 0xa05b9f85.
//
// Solidity: function get_signatures(uint256 block_no, uint256 index) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_LitionScClient *LitionScClientSession) GetSignatures(block_no *big.Int, index *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _LitionScClient.Contract.GetSignatures(&_LitionScClient.CallOpts, block_no, index)
}

// GetSignatures is a free data retrieval call binding the contract method 0xa05b9f85.
//
// Solidity: function get_signatures(uint256 block_no, uint256 index) constant returns(uint8 v, bytes32 r, bytes32 s)
func (_LitionScClient *LitionScClientCallerSession) GetSignatures(block_no *big.Int, index *big.Int) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _LitionScClient.Contract.GetSignatures(&_LitionScClient.CallOpts, block_no, index)
}

// GetSignaturesCount is a free data retrieval call binding the contract method 0x559c270e.
//
// Solidity: function get_signatures_count(uint256 block_no) constant returns(uint256)
func (_LitionScClient *LitionScClientCaller) GetSignaturesCount(opts *bind.CallOpts, block_no *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LitionScClient.contract.Call(opts, out, "get_signatures_count", block_no)
	return *ret0, err
}

// GetSignaturesCount is a free data retrieval call binding the contract method 0x559c270e.
//
// Solidity: function get_signatures_count(uint256 block_no) constant returns(uint256)
func (_LitionScClient *LitionScClientSession) GetSignaturesCount(block_no *big.Int) (*big.Int, error) {
	return _LitionScClient.Contract.GetSignaturesCount(&_LitionScClient.CallOpts, block_no)
}

// GetSignaturesCount is a free data retrieval call binding the contract method 0x559c270e.
//
// Solidity: function get_signatures_count(uint256 block_no) constant returns(uint256)
func (_LitionScClient *LitionScClientCallerSession) GetSignaturesCount(block_no *big.Int) (*big.Int, error) {
	return _LitionScClient.Contract.GetSignaturesCount(&_LitionScClient.CallOpts, block_no)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x82cb1a2a.
//
// Solidity: function registerNode(string n, string r, string p, string e, string ip) returns()
func (_LitionScClient *LitionScClientTransactor) RegisterNode(opts *bind.TransactOpts, n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "registerNode", n, r, p, e, ip)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x82cb1a2a.
//
// Solidity: function registerNode(string n, string r, string p, string e, string ip) returns()
func (_LitionScClient *LitionScClientSession) RegisterNode(n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _LitionScClient.Contract.RegisterNode(&_LitionScClient.TransactOpts, n, r, p, e, ip)
}

// RegisterNode is a paid mutator transaction binding the contract method 0x82cb1a2a.
//
// Solidity: function registerNode(string n, string r, string p, string e, string ip) returns()
func (_LitionScClient *LitionScClientTransactorSession) RegisterNode(n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _LitionScClient.Contract.RegisterNode(&_LitionScClient.TransactOpts, n, r, p, e, ip)
}

// StoreSignature is a paid mutator transaction binding the contract method 0xf202c84d.
//
// Solidity: function store_signature(uint256 block_no, uint8 v, bytes32 r, bytes32 s) returns()
func (_LitionScClient *LitionScClientTransactor) StoreSignature(opts *bind.TransactOpts, block_no *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "store_signature", block_no, v, r, s)
}

// StoreSignature is a paid mutator transaction binding the contract method 0xf202c84d.
//
// Solidity: function store_signature(uint256 block_no, uint8 v, bytes32 r, bytes32 s) returns()
func (_LitionScClient *LitionScClientSession) StoreSignature(block_no *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LitionScClient.Contract.StoreSignature(&_LitionScClient.TransactOpts, block_no, v, r, s)
}

// StoreSignature is a paid mutator transaction binding the contract method 0xf202c84d.
//
// Solidity: function store_signature(uint256 block_no, uint8 v, bytes32 r, bytes32 s) returns()
func (_LitionScClient *LitionScClientTransactorSession) StoreSignature(block_no *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LitionScClient.Contract.StoreSignature(&_LitionScClient.TransactOpts, block_no, v, r, s)
}

// UpdateNode is a paid mutator transaction binding the contract method 0xe1d33203.
//
// Solidity: function updateNode(string n, string r, string p, string e, string ip) returns()
func (_LitionScClient *LitionScClientTransactor) UpdateNode(opts *bind.TransactOpts, n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _LitionScClient.contract.Transact(opts, "updateNode", n, r, p, e, ip)
}

// UpdateNode is a paid mutator transaction binding the contract method 0xe1d33203.
//
// Solidity: function updateNode(string n, string r, string p, string e, string ip) returns()
func (_LitionScClient *LitionScClientSession) UpdateNode(n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _LitionScClient.Contract.UpdateNode(&_LitionScClient.TransactOpts, n, r, p, e, ip)
}

// UpdateNode is a paid mutator transaction binding the contract method 0xe1d33203.
//
// Solidity: function updateNode(string n, string r, string p, string e, string ip) returns()
func (_LitionScClient *LitionScClientTransactorSession) UpdateNode(n string, r string, p string, e string, ip string) (*types.Transaction, error) {
	return _LitionScClient.Contract.UpdateNode(&_LitionScClient.TransactOpts, n, r, p, e, ip)
}

// LitionScClientPrintIterator is returned from FilterPrint and is used to iterate over the raw logs and unpacked data for Print events raised by the LitionScClient contract.
type LitionScClientPrintIterator struct {
	Event *LitionScClientPrint // Event containing the contract specifics and raw log

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
func (it *LitionScClientPrintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LitionScClientPrint)
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
		it.Event = new(LitionScClientPrint)
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
func (it *LitionScClientPrintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LitionScClientPrintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LitionScClientPrint represents a Print event raised by the LitionScClient contract.
type LitionScClientPrint struct {
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
func (_LitionScClient *LitionScClientFilterer) FilterPrint(opts *bind.FilterOpts) (*LitionScClientPrintIterator, error) {

	logs, sub, err := _LitionScClient.contract.FilterLogs(opts, "print")
	if err != nil {
		return nil, err
	}
	return &LitionScClientPrintIterator{contract: _LitionScClient.contract, event: "print", logs: logs, sub: sub}, nil
}

// WatchPrint is a free log subscription operation binding the contract event 0x8f48d31c5e32025ea0c67fbf4573ae86f4b46e5bde075c4dca076b5d293ce408.
//
// Solidity: event print(string nodeName, string role, string publickey, string enode, string ip)
func (_LitionScClient *LitionScClientFilterer) WatchPrint(opts *bind.WatchOpts, sink chan<- *LitionScClientPrint) (event.Subscription, error) {

	logs, sub, err := _LitionScClient.contract.WatchLogs(opts, "print")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LitionScClientPrint)
				if err := _LitionScClient.contract.UnpackLog(event, "print", log); err != nil {
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
func (_LitionScClient *LitionScClientFilterer) ParsePrint(log types.Log) (*LitionScClientPrint, error) {
	event := new(LitionScClientPrint)
	if err := _LitionScClient.contract.UnpackLog(event, "print", log); err != nil {
		return nil, err
	}
	return event, nil
}
