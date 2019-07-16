package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gitlab.com/lition/quorum-maker-nodemanager/client"
	"gitlab.com/lition/quorum-maker-nodemanager/contractclient"
	litionContractClient "gitlab.com/lition/quorum-maker-nodemanager/lition_contractclient"
	"gitlab.com/lition/quorum-maker-nodemanager/service"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	nodeUrl := flag.String("nodeUrl", "http://localhost:22000", "descrciption...")
	listenPort := flag.Int("listenPort", 8000, "descrciption...")
	infuraURL := flag.String("infuraURL", "wss://ropsten.infura.io/ws", "descrciption...")
	contractAddress := flag.String("contractAddress", "0xF4f9c1c8D66C8c9c09456BaD6a9890C3caa768c3", "descrciption...")
	privateKey := flag.String("privateKey", "", "descrciption...")
	chainID := flag.Int("chainID", 0, "descrciption...")
	miningFlag := flag.Bool("miningFlag", false, "descrciption...")
	flag.Parse()

	// TODO: rework pk processing
	if *miningFlag == true && *privateKey == "" {
		log.Fatal("NodeManager misconfiguration. When miningFlag == true, also private key must be provided.")
	}

	listenPortStr := ":" + strconv.Itoa(*listenPort)

	// Init Lition Smartcontract client
	litionContractClient, err := litionContractClient.NewContractClient(*infuraURL, *contractAddress, *privateKey, big.NewInt(int64(*chainID)))
	if err != nil {
		log.Fatal("Unable to init Lition smart contract client")
	}
	// Init Lition Smartcontract event listeners
	if *miningFlag == true {
		err := litionContractClient.InitListeners()
		if err != nil {
			log.Fatal("Unable to init Lition smart contract event listeners")
		}
	}

	router := mux.NewRouter()
	nodeService := service.NodeServiceImpl{*nodeUrl, litionContractClient}

	ticker := time.NewTicker(86400 * time.Second)
	go func() {
		for range ticker.C {
			log.Debug("Rotating log for Geth and Constellation.")
			nodeService.LogRotaterGeth()
			nodeService.LogRotaterConst()
		}
	}()

	go func() {
		nodeService.CheckGethStatus(*nodeUrl)
		//log.Info("Deploying Network Manager Contract")
		nodeService.NetworkManagerContractDeployer(*nodeUrl)
		nodeService.RegisterNodeDetails(*nodeUrl)
		nodeService.ContractCrawler(*nodeUrl)
		nodeService.ABICrawler(*nodeUrl)

		// Let lition SC know that this node wants to start mining
		if *miningFlag == true {
			err := litionContractClient.StartMining()
			if err != nil {
				log.Fatal("Unable to start mining. Errr: ", err)
			}

			// Start standalone event listeners
			go litionContractClient.Start_StartMiningEventListener(nodeService.VoteValidator)
			go litionContractClient.Start_StopMiningEventListener(nodeService.UnvoteValidator)
		}
	}()

	networkMapService := contractclient.NetworkMapContractClient{EthClient: client.EthClient{*nodeUrl}}
	router.HandleFunc("/txn/{txn_hash}", nodeService.GetTransactionInfoHandler).Methods("GET")
	router.HandleFunc("/rmpld/{txn_hash}", nodeService.DeleteTransactionPayloadHandler).Methods("GET")
	router.HandleFunc("/txn", nodeService.GetLatestTransactionInfoHandler).Methods("GET")
	router.HandleFunc("/block/{block_no}", nodeService.GetBlockInfoHandler).Methods("GET")
	router.HandleFunc("/block", nodeService.GetLatestBlockInfoHandler).Methods("GET")
	router.HandleFunc("/genesis", nodeService.GetGenesisHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/peer/{peer_id}", nodeService.GetOtherPeerHandler).Methods("GET")
	router.HandleFunc("/peer", nodeService.GetCurrentNodeHandler).Methods("GET")
	router.HandleFunc("/nmcAddress", nodeService.GetNmcAddress).Methods("POST")
	router.HandleFunc("/txnrcpt/{txn_hash}", nodeService.GetTransactionReceiptHandler).Methods("GET")
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

	router.PathPrefix("/contracts").Handler(http.StripPrefix("/contracts", http.FileServer(http.Dir("/root/quorum-maker/contracts"))))
	router.PathPrefix("/geth").Handler(http.StripPrefix("/geth", http.FileServer(http.Dir("/home/node/qdata/gethLogs"))))
	router.PathPrefix("/constellation").Handler(http.StripPrefix("/constellation", http.FileServer(http.Dir("/home/node/qdata/constellationLogs"))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", NewFileServer("NodeManagerUI")))

	log.Info(fmt.Sprintf("Node Manager listening on %s...", listenPortStr))

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0" + listenPortStr,

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

	if *miningFlag == true {
		// Stop mining
		litionContractClient.StopMining()
		// Unvote itself
		nodeService.UnvoteValidator(litionContractClient.GetPubKeyStr())
	}

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
