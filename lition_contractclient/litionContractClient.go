package litionContractClient

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	log "github.com/sirupsen/logrus"
	lition "gitlab.com/lition/lition_contracts/contracts/go_wrapper"
	eventListener "gitlab.com/lition/quorum-maker-nodemanager/lition_contractclient/eventListener"
)

// ContractClient contains variables needed for communication with lition smart contract
type ContractClient struct {
	ethClient                *ethclient.Client
	privateKey               *ecdsa.PrivateKey
	auth                     *bind.TransactOpts
	scAddress                common.Address
	scClient                 *lition.Lition
	chainID                  *big.Int // chainID on top of which all sc calls are made
	startMiningEventListener *eventListener.StartMiningEventListener
	stopMiningEventListener  *eventListener.StopMiningEventListener
}

func NewContractClient(ethClientURL string, scAddress string, privateKey string, chainID *big.Int) (*ContractClient, error) {
	contractClient := new(ContractClient)
	ethClient, err := ethclient.Dial(ethClientURL)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	contractClient.ethClient = ethClient

	contractClient.privateKey = nil
	contractClient.auth = nil
	if privateKey != "" {
		pPrivateKey, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		contractClient.privateKey = pPrivateKey
		contractClient.auth = bind.NewKeyedTransactor(contractClient.privateKey)
	}

	contractClient.scAddress = common.HexToAddress(scAddress)
	contractClient.chainID = chainID

	pScClient, err := lition.NewLition(contractClient.scAddress, contractClient.ethClient)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	contractClient.scClient = pScClient

	contractClient.startMiningEventListener = nil
	contractClient.stopMiningEventListener = nil

	return contractClient, nil
}
func (contractClient *ContractClient) InitListeners() error {
	var err error
	contractClient.startMiningEventListener, err = eventListener.NewStartMiningEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}
	contractClient.stopMiningEventListener, err = eventListener.NewStopMiningEventListener(contractClient.scClient, contractClient.chainID)
	if err != nil {
		return err
	}

	return nil
}

func (contractClient *ContractClient) DeInit() {
	if contractClient.startMiningEventListener != nil {
		contractClient.startMiningEventListener.DeInit()
	}
	if contractClient.stopMiningEventListener != nil {
		contractClient.stopMiningEventListener.DeInit()
	}

	contractClient.chainID = nil
	contractClient.startMiningEventListener = nil
	contractClient.stopMiningEventListener = nil
	contractClient.auth = nil
	contractClient.ethClient.Close()
}

func (contractClient *ContractClient) Start_StartMiningEventListener(f func(string)) {
	listener := contractClient.startMiningEventListener
	if listener == nil {
		log.Fatal("Trying to start \"StartMining\" listener without previous initialization")
		return
	}

	// Infinite loop - try to initialze listeners until it succeeds
	for {
		retErr := listener.Start(f)
		// Listener was manually stopped, do not try to start it again
		if retErr == nil {
			return
		}
		log.Error("Start StartMiningEventListener err: \"", retErr, "\". Try to reinit.")

		// Wait some time before trying to reinit and start listener again
		time.Sleep(1 * time.Second)
		listener.ReInit()
	}
}

func (contractClient *ContractClient) Start_StopMiningEventListener(f func(string)) {
	listener := contractClient.stopMiningEventListener
	if listener == nil {
		log.Fatal("Trying to start \"StopMining\" listener without previous initialization")
		return
	}

	// Infinite loop - try to initialze listeners until it succeeds
	for {
		retErr := listener.Start(f)
		// Listener was manually stopped, do not try to start it again
		if retErr == nil {
			return
		}
		log.Error("Start StopMiningEventListener err: \"", retErr, "\". Try to reinit.")

		// Wait some time before trying to reinit and start listener again
		time.Sleep(1 * time.Second)
		listener.ReInit()
	}
}

func (contractClient *ContractClient) StartMining() error {
	if contractClient.auth == nil {
		return errors.New("PrivateKey must be provided for smart contract writes")
	}

	tx, err := contractClient.scClient.StartMining(contractClient.auth, contractClient.chainID)
	if err != nil {
		return err
	}
	log.Info("Transaction \"startMining\" TX: ", tx.Hash())
	return nil
}

func (contractClient *ContractClient) StopMining() error {
	if contractClient.auth == nil {
		return errors.New("PrivateKey must be provided for smart contract write")
	}

	tx, err := contractClient.scClient.StopMining(contractClient.auth, contractClient.chainID)
	if err != nil {
		return err
	}
	log.Info("Transaction \"stopMining\" TX: ", tx.Hash())
	return nil
}

func (contractClient *ContractClient) AccHasVested(userAddressStr string) (bool, error) {
	userAddress := common.HexToAddress(userAddressStr)

	hasVested, err := contractClient.scClient.HasVested(&bind.CallOpts{}, contractClient.chainID, userAddress)
	if err != nil {
		return false, err
	}

	return hasVested, nil
}

func (contractClient *ContractClient) AccHasDeposited(userAddressStr string) (bool, error) {
	userAddress := common.HexToAddress(userAddressStr)

	hasDeposited, err := contractClient.scClient.HasDeposited(&bind.CallOpts{}, contractClient.chainID, userAddress)
	if err != nil {
		return false, err
	}

	return hasDeposited, nil
}
