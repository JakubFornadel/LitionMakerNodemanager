package eventListener

import (
	"context"
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"

	log "github.com/sirupsen/logrus"
	lition "gitlab.com/lition/lition_contracts/contracts/go_wrapper"
)

type StopMiningEventListener struct {
	initialized    bool
	listening      bool
	scClient       *lition.Lition
	eventChannel   chan *lition.LitionStopMining
	eventSubs      event.Subscription
	filterChainId  []*big.Int
	stopChannel    chan struct{}
	stoppedChannel chan struct{}
	mutex          sync.Mutex
}

func NewStopMiningEventListener(scClient *lition.Lition, chainId *big.Int) (*StopMiningEventListener, error) {
	p := new(StopMiningEventListener)

	p.initialized = false
	p.listening = false
	p.mutex = sync.Mutex{}
	p.scClient = scClient
	p.filterChainId = []*big.Int{chainId}
	err := p.Init()
	if err == nil {
		return p, nil
	}

	return nil, err
}

func (listener *StopMiningEventListener) Init() error {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.initialized == true {
		return nil
	}

	var err error
	listener.eventChannel = make(chan *lition.LitionStopMining)
	listener.eventSubs, err = listener.scClient.WatchStopMining(&bind.WatchOpts{
		Context: context.Background(), Start: nil},
		listener.eventChannel,
		listener.filterChainId)

	if err != nil {
		log.Error(err)
		close(listener.eventChannel)
		return err
	}

	listener.stopChannel = make(chan struct{})
	listener.stoppedChannel = nil // Stopped channel is created and deleted only in start function
	listener.initialized = true

	return nil
}

func (listener *StopMiningEventListener) DeInit() {
	listener.Stop()

	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.initialized == false {
		return
	}

	listener.eventSubs.Unsubscribe()
	close(listener.eventChannel)
	close(listener.stopChannel)
	// close(listener.stoppedChannel) // stoppned channel is already closed when stopped listening
	listener.initialized = false
}

func (listener *StopMiningEventListener) ReInit() error {
	listener.DeInit()
	return listener.Init()
}

func (listener *StopMiningEventListener) Start(f func(string)) error {
	if listener.initialized == false {
		return errors.New("Trying to Start \"StopMiningEventListener\" without previous initialization")
	}
	if listener.listening == true {
		log.Warning("Trying to Start \"StopMiningEventListener\", which is already listening.")
		return nil
	}

	listener.stoppedChannel = make(chan struct{})
	listener.listening = true
	log.Info("StopMiningEventListener start listening")

	// close the stoppedchan when this func exits
	defer func() {
		close(listener.stoppedChannel)
		listener.listening = false
	}()

	for {
		select {
		case event := <-listener.eventChannel:
			log.Info("New \"StopMining\" event received.")
			f(event.Miner.String())
		case err := <-listener.eventSubs.Err():
			return err
		case <-listener.stopChannel:
			log.Info("Signal to stop StopMiningEventListener received.")
			return nil
		}
	}
}

func (listener *StopMiningEventListener) Stop() {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.listening == false {
		return
	}

	close(listener.stopChannel)
	// wait for it to have stopped
	<-listener.stoppedChannel
	listener.stopChannel = make(chan struct{})
	log.Info("StopMiningEventListener successfully stopped")
}
