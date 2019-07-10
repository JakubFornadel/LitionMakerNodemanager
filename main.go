package main

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gitlab.com/lition/quorum-maker-nodemanager/client"
	"gitlab.com/lition/quorum-maker-nodemanager/contractclient"
	litionContractClient "gitlab.com/lition/quorum-maker-nodemanager/lition_contractclient"
	"gitlab.com/lition/quorum-maker-nodemanager/service"

	"github.com/magiconair/properties"
	"gitlab.com/lition/quorum-maker-nodemanager/util"
)

var nodeUrl = "http://localhost:22000"
var listenPort = ":8000"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	if len(os.Args) > 1 {
		nodeUrl = os.Args[1]
	}

	if len(os.Args) > 2 {
		listenPort = ":" + os.Args[2]
	}

	// Read Lition Smartcontract client related config parameters from file
	infuraURL, contractAddress, chainID, privateKey, miningFlag := getContractConfig()
	// Init Lition Smartcontract client
	litionContractClient, err := litionContractClient.NewContractClient(infuraURL, contractAddress, privateKey, chainID)
	if err != nil {
		log.Fatal("Unable to init Lition smart contract client")
	}
	// Init Lition Smartcontract event listeners
	if miningFlag == true {
		err := litionContractClient.InitListeners()
		if err != nil {
			log.Fatal("Unable to init Lition smart contract event listeners")
		}
	}

	router := mux.NewRouter()
	nodeService := service.NodeServiceImpl{nodeUrl, litionContractClient}

	ticker := time.NewTicker(86400 * time.Second)
	go func() {
		for range ticker.C {
			log.Debug("Rotating log for Geth and Constellation.")
			nodeService.LogRotaterGeth()
			nodeService.LogRotaterConst()
		}
	}()

	// TODO: remove when done testing
	// Let lition SC know that this node wants to start mining
	if miningFlag == true {
		litionContractClient.StartMining()

		// Start standalone event listeners
		go litionContractClient.Start_StartMiningEventListener(nodeService.VoteValidator)
		go litionContractClient.Start_StopMiningEventListener(nodeService.UnvoteValidator)
	}

	go func() {
		nodeService.CheckGethStatus(nodeUrl)
		//log.Info("Deploying Network Manager Contract")
		nodeService.NetworkManagerContractDeployer(nodeUrl)
		nodeService.RegisterNodeDetails(nodeUrl)
		nodeService.ContractCrawler(nodeUrl)
		nodeService.ABICrawler(nodeUrl)
		nodeService.IPWhitelister()

		// Let lition SC know that this node wants to start mining
		if miningFlag == true {
			litionContractClient.StartMining()

			// Start standalone event listeners
			go litionContractClient.Start_StartMiningEventListener(nodeService.VoteValidator)
			go litionContractClient.Start_StopMiningEventListener(nodeService.UnvoteValidator)
		}
	}()

	networkMapService := contractclient.NetworkMapContractClient{EthClient: client.EthClient{nodeUrl}}
	router.HandleFunc("/txn/{txn_hash}", nodeService.GetTransactionInfoHandler).Methods("GET")
	router.HandleFunc("/rmpld/{txn_hash}", nodeService.DeleteTransactionPayloadHandler).Methods("GET")
	router.HandleFunc("/txn", nodeService.GetLatestTransactionInfoHandler).Methods("GET")
	router.HandleFunc("/block/{block_no}", nodeService.GetBlockInfoHandler).Methods("GET")
	router.HandleFunc("/block", nodeService.GetLatestBlockInfoHandler).Methods("GET")
	router.HandleFunc("/genesis", nodeService.GetGenesisHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/peer/{peer_id}", nodeService.GetOtherPeerHandler).Methods("GET")
	router.HandleFunc("/peer", nodeService.JoinNetworkHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/peer", nodeService.GetCurrentNodeHandler).Methods("GET")
	router.HandleFunc("/txnrcpt/{txn_hash}", nodeService.GetTransactionReceiptHandler).Methods("GET")
	router.HandleFunc("/pendingJoinRequests", nodeService.PendingJoinRequestsHandler).Methods("GET")
	router.HandleFunc("/joinRequestResponse", nodeService.JoinRequestResponseHandler).Methods("POST")
	router.HandleFunc("/joinRequestResponse", nodeService.OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/createNetwork", nodeService.CreateNetworkScriptCallHandler).Methods("POST")
	router.HandleFunc("/createNetwork", nodeService.OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/joinNetwork", nodeService.JoinNetworkScriptCallHandler).Methods("POST")
	router.HandleFunc("/joinNetwork", nodeService.OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/deployContract", nodeService.DeployContractHandler).Methods("POST")
	router.HandleFunc("/reset", nodeService.ResetHandler).Methods("GET")
	router.HandleFunc("/restart", nodeService.RestartHandler).Methods("GET")
	router.HandleFunc("/latestBlock", nodeService.LatestBlockHandler).Methods("GET")
	router.HandleFunc("/latency", nodeService.LatencyHandler).Methods("GET")
	router.HandleFunc("/proposeValidator", nodeService.ProposeValidator).Methods("POST")
	//router.HandleFunc("/logs", nodeService.LogsHandler).Methods("GET")
	router.HandleFunc("/txnsearch/{txn_hash}", nodeService.TransactionSearchHandler).Methods("GET")
	router.HandleFunc("/mailserver", nodeService.MailServerConfigHandler).Methods("POST")
	router.HandleFunc("/mailserver", nodeService.OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/registerNode", networkMapService.RegisterNodeRequestHandler).Methods("POST")
	router.HandleFunc("/updateNode", networkMapService.UpdateNodeHandler).Methods("POST")
	router.HandleFunc("/updateNode", networkMapService.OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/getNodeDetails/{index}", networkMapService.GetNodeDetailsResponseHandler).Methods("GET")
	router.HandleFunc("/getNodeList", networkMapService.GetNodeListSelfResponseHandler).Methods("GET")
	router.HandleFunc("/activeNodes", networkMapService.ActiveNodesHandler).Methods("GET")
	router.HandleFunc("/chartData", nodeService.GetChartDataHandler).Methods("GET")
	router.HandleFunc("/contractList", nodeService.GetContractListHandler).Methods("GET")
	router.HandleFunc("/contractCount", nodeService.GetContractCountHandler).Methods("GET")
	router.HandleFunc("/updateContractDetails", nodeService.ContractDetailsUpdateHandler).Methods("POST")
	router.HandleFunc("/attachedNodeDetails", nodeService.AttachedNodeDetailsHandler).Methods("POST")
	router.HandleFunc("/initialized", nodeService.InitializationHandler).Methods("GET")
	router.HandleFunc("/createAccount", nodeService.CreateAccountHandler).Methods("POST")
	router.HandleFunc("/createAccount", nodeService.OptionsHandler).Methods("OPTIONS")
	router.HandleFunc("/getAccounts", nodeService.GetAccountsHandler).Methods("GET")
	router.HandleFunc("/getWhitelist", nodeService.GetWhitelistedIPsHandler).Methods("GET")
	router.HandleFunc("/updateWhitelist", nodeService.UpdateWhitelistHandler).Methods("POST")
	router.HandleFunc("/updateWhitelist", nodeService.OptionsHandler).Methods("OPTIONS")

	router.PathPrefix("/contracts").Handler(http.StripPrefix("/contracts", http.FileServer(http.Dir("/root/quorum-maker/contracts"))))
	router.PathPrefix("/geth").Handler(http.StripPrefix("/geth", http.FileServer(http.Dir("/home/node/qdata/gethLogs"))))
	router.PathPrefix("/constellation").Handler(http.StripPrefix("/constellation", http.FileServer(http.Dir("/home/node/qdata/constellationLogs"))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", NewFileServer("NodeManagerUI")))

	log.Info(fmt.Sprintf("Node Manager listening on %s...", listenPort))

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0" + listenPort,

		//WriteTimeout: 15 * time.Second,
		//ReadTimeout:  15 * time.Second,
		//IdleTimeout:  time.Second * 60,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	// Deinit lition smart contract cliet
	litionContractClient.DeInit()

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("Node Manager Shutting down")
	os.Exit(0)
}

type MyFileServer struct {
	name    string
	handler http.Handler
}

func NewFileServer(file string) *MyFileServer {

	return &MyFileServer{file, http.FileServer(http.Dir(file))}

}
func (mf *MyFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	_, err := os.Open(mf.name + "/" + r.URL.Path)
	if err != nil {
		r.URL.Path = "/"
	}

	mf.handler.ServeHTTP(w, r)
}

func getContractConfig() (infuraURL string, contractAddress string, chainID *big.Int, privateKey string, miningFlag bool) {
	// Default values - should not change
	pInfuraURL := "wss://ropsten.infura.io/ws"
	pContractAddress := "0xF4f9c1c8D66C8c9c09456BaD6a9890C3caa768c3"

	// TODO: remove when testing done
	return pInfuraURL, pContractAddress, big.NewInt(0), "5C5D06D3A4F0EB0B90F703CF345C8B4FE209FB0958E884312962F3A24D8218FE", true

	p := properties.MustLoadFile("/home/setup.conf", properties.UTF8)

	if util.PropertyExists("ROLE", "/home/setup.conf") == "" {
		log.Fatal("\"ROLE\" must be present in config")
	}
	pMiningFlag, err := regexp.MatchString("[Vv][Aa][Ll][Ii][Dd][Aa][Tt][Oo][Rr]", util.MustGetString("ROLE", p))
	if err != nil {
		log.Fatal("Unable to parse \"ROLE\" config parameter.")
	}

	if util.PropertyExists("CHAIN_ID", "/home/setup.conf") == "" {
		log.Fatal("\"CHAIN_ID\" must be present in config")
	}
	pChainID := new(big.Int)
	pChainID, ok := chainID.SetString(util.MustGetString("CHAIN_ID", p), 10)
	if ok == false {
		log.Fatal("Unable to parse \"CHAIN_ID\" config parameter.")
	}

	// TODO: read private key from file
	pPrivateKey := "5C5D06D3A4F0EB0B90F703CF345C8B4FE209FB0958E884312962F3A24D8218FE"

	return pInfuraURL, pContractAddress, pChainID, pPrivateKey, pMiningFlag
}
