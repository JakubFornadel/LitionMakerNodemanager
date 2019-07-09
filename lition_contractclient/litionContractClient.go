package litionContractClient

import (
	"crypto/ecdsa"
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
	startMiningEventListener *eventListener.StartMiningEventListener
	stopMiningEventListener  *eventListener.StopMiningEventListener
}

func NewContractClient(ethClientURL string, scAddress string, privateKey string) (*ContractClient, error) {
	contractClient := new(ContractClient)
	ethClient, err := ethclient.Dial(ethClientURL)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	contractClient.ethClient = ethClient

	pPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	contractClient.privateKey = pPrivateKey

	contractClient.scAddress = common.HexToAddress(scAddress)

	pScClient, err := lition.NewLition(contractClient.scAddress, contractClient.ethClient)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	contractClient.scClient = pScClient

	contractClient.auth = bind.NewKeyedTransactor(contractClient.privateKey)
	contractClient.startMiningEventListener = nil
	contractClient.stopMiningEventListener = nil

	return contractClient, nil
}
func (contractClient *ContractClient) InitListeners(chainId *big.Int) error {
	var err error
	contractClient.startMiningEventListener, err = eventListener.NewStartMiningEventListener(contractClient.scClient, chainId)
	if err != nil {
		return err
	}
	contractClient.stopMiningEventListener, err = eventListener.NewStopMiningEventListener(contractClient.scClient, chainId)
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

	contractClient.ethClient.Close()
}

func (contractClient *ContractClient) Start_StartMiningEventListener(f func(*lition.LitionStartMining)) {
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
		log.Error("Start StartMiningEventListener err: ", retErr, "Try to reinit.")

		// Wait some time before trying to reinit and start listener again
		time.Sleep(1 * time.Second)
		listener.ReInit()
	}
}

func (contractClient *ContractClient) Start_StopMiningEventListener(f func(*lition.LitionStopMining)) {
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
		log.Error("Start StopMiningEventListener err: ", retErr, "Try to reinit.")

		// Wait some time before trying to reinit and start listener again
		time.Sleep(1 * time.Second)
		listener.ReInit()
	}
}

func (contractClient *ContractClient) StartMining(chainId *big.Int) {
	tx, err := contractClient.scClient.StartMining(contractClient.auth, chainId)
	if err != nil {
		log.Error(err)
	}
	log.Info("Transaction \"startMining\" TX: ", tx.Hash())
}

func (contractClient *ContractClient) StopMining(chainId *big.Int) {
	tx, err := contractClient.scClient.StopMining(contractClient.auth, chainId)
	if err != nil {
		log.Error(err)
	}
	log.Info("Transaction \"stopMining\" TX: ", tx.Hash())
}
