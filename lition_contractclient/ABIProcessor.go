package litioncontractclient

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const litionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"_deposit_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chains\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"},{\"name\":\"last_notary\",\"type\":\"uint256\"},{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"total_vesting\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"vesting\",\"type\":\"uint256\"}],\"name\":\"vest_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"notary_block\",\"type\":\"uint256\"},{\"name\":\"miners\",\"type\":\"address[]\"},{\"name\":\"blocks_mined\",\"type\":\"uint256[]\"},{\"name\":\"users\",\"type\":\"address[]\"},{\"name\":\"user_gas\",\"type\":\"uint256[]\"},{\"name\":\"largest_tx\",\"type\":\"uint256\"},{\"components\":[{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"notary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"has_deposited\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"info\",\"type\":\"string\"},{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"vesting\",\"type\":\"uint256\"},{\"name\":\"init_endpoint\",\"type\":\"string\"}],\"name\":\"register_chain\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"has_vested\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"deposit_in_chain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"next_id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"description\",\"type\":\"string\"}],\"name\":\"NewChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"NewChainEndpoint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"depositer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"datetime\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"chain_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"depositer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"datetime\",\"type\":\"uint256\"}],\"name\":\"Vesting\",\"type\":\"event\"}]"

// lition is an auto generated Go binding around an Ethereum contract.
type lition struct {
	litionCaller     // Read-only binding to the contract
	litionTransactor // Write-only binding to the contract
	litionFilterer   // Log filterer for contract events
}

// litionCaller is an auto generated read-only Go binding around an Ethereum contract.
type litionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// litionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type litionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// litionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type litionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

func bindlition(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(litionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func Newlition(address common.Address, backend bind.ContractBackend) (*lition, error) {
	contract, err := bindlition(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &lition{litionCaller: litionCaller{contract: contract}, litionTransactor: litionTransactor{contract: contract}, litionFilterer: litionFilterer{contract: contract}}, nil
}

func (_lition *litionCaller) HasDeposited(opts *bind.CallOpts, id *big.Int, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _lition.contract.Call(opts, out, "has_deposited", id, user)
	return *ret0, err
}
