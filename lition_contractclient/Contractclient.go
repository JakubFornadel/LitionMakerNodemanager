package litioncontractclient

import (
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var litionSC *lition
var ethClient *ethclient.Client

func init() {
	initEthClient()
}

func initEthClient() {
	var err error
	ethClient, err = ethclient.Dial("https://ropsten.infura.io/v3/5cd24295acd74cc1b9dd5ced7b4bb6a9")
	if err != nil {
		log.Fatal(err)
	}

	smartContract := common.HexToAddress("0x603a519f501fdAD1843B8A993Ec378C4DbBda097")

	litionSC, err = Newlition(smartContract, ethClient)
	if err != nil {
		log.Fatal(err)
	}
}

func IsAllowedUser(chanIDStr string, userAddressStr string) bool {
	// TODO: check if ethClient lost connection and call initEthClient() only if it did
	initEthClient()

	chanID, ok := new(big.Int).SetString(chanIDStr, 0)
	if ok == false {
		log.Error(fmt.Sprint("Unable to convert ", chanIDStr, " into big.Int"))
		return false
	}
	userAddres := common.HexToAddress(userAddressStr)

	canConnect, err := litionSC.HasDeposited(&bind.CallOpts{}, chanID, userAddres)
	if err != nil {
		log.Error(fmt.Sprint("HasDeposited func error: ", err))
		return false
	}

	if canConnect == true {
		log.Info(fmt.Sprint("User with public key ", userAddressStr, " is allowed to connect automatically."))
	} else {
		log.Info(fmt.Sprint("User with public key ", userAddressStr, " is not allowed to connect automatically."))
	}

	return canConnect
}
