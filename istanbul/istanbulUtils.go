package istanbulUtils

import (
	log "github.com/sirupsen/logrus"
	lition "gitlab.com/lition/lition_contracts/contracts/go_wrapper"
)

func VoteWitness(event *lition.LitionStartMining) {
	log.Info("VoteWitness function invoked. Miner: ", event.Miner)
	// TODO: istanbul.vote
}

func UnvoteWitness(event *lition.LitionStopMining) {
	log.Info("UnvoteWitness function invoked. Miner: ", event.Miner)
	// TODO: istanbul.vote
}
