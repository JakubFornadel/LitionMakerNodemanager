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

type StartMiningEventListener struct {
	initialized    bool
	listening      bool
	scClient       *lition.Lition
	eventChannel   chan *lition.LitionStartMining
	eventSubs      event.Subscription
	filterChainId  []*big.Int
	stopChannel    chan struct{}
	stoppedChannel chan struct{}
	mutex          sync.Mutex
}

func NewStartMiningEventListener(scClient *lition.Lition, chainId *big.Int) (*StartMiningEventListener, error) {
	p := new(StartMiningEventListener)

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

func (listener *StartMiningEventListener) Init() error {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.initialized == true {
		return nil
	}

	var err error
	listener.eventChannel = make(chan *lition.LitionStartMining)
	listener.eventSubs, err = listener.scClient.WatchStartMining(&bind.WatchOpts{
		Context: context.Background(), Start: nil},
		listener.eventChannel,
		listener.filterChainId)

	if err != nil {
		log.Error(err)
		close(listener.eventChannel)
		return err
	}

	listener.stopChannel = make(chan struct{})
	listener.stoppedChannel = make(chan struct{})
	listener.initialized = true

	return nil
}

func (listener *StartMiningEventListener) DeInit() {
	listener.Stop()

	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.initialized == false {
		return
	}

	listener.eventSubs.Unsubscribe()
	close(listener.eventChannel)
	close(listener.stopChannel)
	close(listener.stoppedChannel)
	listener.initialized = false
}

func (listener *StartMiningEventListener) ReInit() {
	listener.DeInit()
	listener.Init()
}

func (listener *StartMiningEventListener) Start(f func(string)) error {
	if listener.initialized == false {
		return errors.New("Trying to Start \"StartMiningEventListener\" without previous initialization")
	}
	if listener.listening == true {
		log.Warning("Trying to Start \"StartMiningEventListener\", which is already listening.")
		return nil
	}

	log.Info("StartMiningEventListener start listening")
	listener.listening = true

	// close the stoppedchan when this func exits
	defer func() {
		close(listener.stoppedChannel)
		listener.listening = false
	}()

	for {
		select {
		case event := <-listener.eventChannel:
			log.Info("New \"StartMining\" event received.")
			f(event.Miner.String())
		case err := <-listener.eventSubs.Err():
			return err
		case <-listener.stopChannel:
			log.Info("Signal to stop StartMiningEventListener received.")
			return nil
		}
	}
}

func (listener *StartMiningEventListener) Stop() {
	listener.mutex.Lock()
	defer listener.mutex.Unlock()

	if listener.listening == false {
		return
	}

	close(listener.stopChannel)
	// wait for it to have stopped
	<-listener.stoppedChannel
	listener.stopChannel = make(chan struct{})
	listener.stoppedChannel = make(chan struct{})
	log.Info("StartMiningEventListener successfully stopped")
}
