package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gitlab.com/lition/quorum-maker-nodemanager/util"
)

type contractJSON struct {
	Abi       []interface{} `json:"abi"`
	Interface []interface{} `json:"interface"`
	Bytecode  string        `json:"bytecode"`
}

type genesisJSON struct {
	Config configField `json:"config"`
}

type configField struct {
	ChainId int `json:"chainId"`
}

type accountPassword struct {
	Password string `json:"password"`
}

type connectedIP struct {
	IP          string `json:"ip"`
	Whitelisted bool   `json:"whitelisted"`
	Count       int    `json:"count"`
}

type IPList struct {
	WhiteList     []string      `json:"whiteList"`
	ConnectedList []connectedIP `json:"connectedList"`
}

var pendCount = 0
var nameMap = map[string]string{}
var channelMap = make(map[string](chan string))

func (nsi *NodeServiceImpl) ProposeValidator(w http.ResponseWriter, r *http.Request) {
	var request ProposeValidatorRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	// TODO: check input parameters
	//w.WriteHeader(http.BadRequest)
	//w.Write([]byte("Invalid call arguments"))

	response := nsi.proposeValidator(nsi.Url, request.ValidatorAddress, request.Vote)
	json.NewEncoder(w).Encode(response)
}

// This wrapper is used in event listener for automatic voting
func (nsi *NodeServiceImpl) VoteValidator(validatorAddress string) {
	log.Info("Aut. VoteValidator function invoked. Validator: ", validatorAddress)
	nsi.proposeValidator(nsi.Url, validatorAddress, true)
}

// This wrapper is used in event listener for automatic unvoting
func (nsi *NodeServiceImpl) UnvoteValidator(validatorAddress string) {
	log.Info("Aut. UnvoteValidator function invoked. Validator: ", validatorAddress)
	nsi.proposeValidator(nsi.Url, validatorAddress, false)
}

func (nsi *NodeServiceImpl) GetNmcAddress(w http.ResponseWriter, r *http.Request) {
	var request JoinNetworkRequest
	_ = json.NewDecoder(r.Body).Decode(&request)
	accAddress := request.AccPubKey

	hasVested, err := nsi.LitionContractClient.AccHasVested(accAddress)
	if err != nil {
		log.Error("GetNmcAddress AccHasVested err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Lition SC read error"))
	}

	hasDeposited, err := nsi.LitionContractClient.AccHasDeposited(accAddress)
	if err != nil {
		log.Error("GetNmcAddress AccHasDeposited err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Lition SC read error"))
	}

	if hasVested || hasDeposited {
		response := nsi.getNmcAddress()
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Permissoon denied"))
}

func (nsi *NodeServiceImpl) GetGenesisHandler(w http.ResponseWriter, r *http.Request) {
	var request JoinNetworkRequest
	_ = json.NewDecoder(r.Body).Decode(&request)
	enode := request.EnodeID
	foreignIP := request.IPAddress
	nodename := request.Nodename
	accPubKey := request.AccPubKey
	chainID := request.ChainID

	log.Info(fmt.Sprint("Join request received from node: ", nodename, " with IP: ", foreignIP, ", enode: ", enode, ", accPubKey: ", accPubKey, " and chainID: ", chainID))

	hasVested, err := nsi.LitionContractClient.AccHasVested(accPubKey)
	if err != nil {
		log.Error("GetGenesis AccHasVested call error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Lition SC read error"))
		return
	}

	if hasVested == true {
		response := nsi.getGenesis(nsi.Url)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Warning("Access denied. Acc has not vested.")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Access denied"))
}

func (nsi *NodeServiceImpl) JoinRequestResponseHandler(w http.ResponseWriter, r *http.Request) {
	var request JoinNetworkResponse
	_ = json.NewDecoder(r.Body).Decode(&request)
	enode := request.EnodeID
	status := request.Status
	response := nsi.joinRequestResponse(enode, status)
	channelMap[enode] <- status
	delete(channelMap, enode)
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) GetCurrentNodeHandler(w http.ResponseWriter, r *http.Request) {
	response := nsi.getCurrentNode(nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) GetOtherPeerHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["peer_id"] == "all" {
		response := nsi.getOtherPeers(nsi.Url)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	} else {
		response := nsi.getOtherPeer(params["peer_id"], nsi.Url)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}

func (nsi *NodeServiceImpl) GetLatestBlockInfoHandler(w http.ResponseWriter, r *http.Request) {
	count := r.FormValue("number")
	reference := r.FormValue("reference")
	response := nsi.getLatestBlockInfo(count, reference, nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) GetBlockInfoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	block, err := strconv.ParseInt(params["block_no"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	response := nsi.getBlockInfo(block, nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) GetTransactionInfoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["txn_hash"] == "pending" {
		response := nsi.getPendingTransactions(nsi.Url)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	} else {
		response := nsi.getTransactionInfo(params["txn_hash"], nsi.Url)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}

func (nsi *NodeServiceImpl) DeleteTransactionPayloadHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response := nsi.deleteTransactionPayload(params["txn_hash"], nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}

func (nsi *NodeServiceImpl) GetLatestTransactionInfoHandler(w http.ResponseWriter, r *http.Request) {
	count := r.FormValue("number")
	response := nsi.getLatestTransactionInfo(count, nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) GetTransactionReceiptHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response := nsi.getTransactionReceipt(params["txn_hash"], nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) DeployContractHandler(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
	var private bool
	var publicKeys []string

	count := r.FormValue("count")
	countInt, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println(err)
	}
	fileNames := make([]string, countInt)
	boolVal := r.FormValue("private")
	if boolVal == "true" {
		private = true
	} else {
		private = false
	}

	keys := r.FormValue("privateFor")
	publicKeys = strings.Split(keys, ",")

	for i := 0; i < countInt; i++ {
		keyVal := "file" + strconv.Itoa(i+1)

		file, header, err := r.FormFile(keyVal)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		name := strings.Split(header.Filename, ".")

		fileNames[i] = name[0] + ".sol"

		io.Copy(&Buf, file)

		contents := Buf.String()

		fileContent := []byte(contents)
		err = ioutil.WriteFile("./"+name[0]+".sol", fileContent, 0775)
		if err != nil {
			fmt.Println(err)
		}

		Buf.Reset()
	}

	response := nsi.deployContract(publicKeys, fileNames, private, nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) CreateNetworkScriptCallHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateNetworkScriptArgs
	_ = json.NewDecoder(r.Body).Decode(&request)
	fmt.Println(request)
	response := nsi.createNetworkScriptCall(request.Nodename, request.CurrentIP, request.RPCPort, request.WhisperPort, request.ConstellationPort, request.NodeManagerPort)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) JoinNetworkScriptCallHandler(w http.ResponseWriter, r *http.Request) {
	var request JoinNetworkScriptArgs
	_ = json.NewDecoder(r.Body).Decode(&request)
	fmt.Println(request)
	response := nsi.joinRequestResponseCall(request.Nodename, request.CurrentIP, request.RPCPort, request.WhisperPort, request.ConstellationPort, request.NodeManagerPort, request.MasterNodeManagerPort, request.MasterIP)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) ResetHandler(w http.ResponseWriter, r *http.Request) {
	response := nsi.resetCurrentNode()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) RestartHandler(w http.ResponseWriter, r *http.Request) {
	response := nsi.restartCurrentNode()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) LatestBlockHandler(w http.ResponseWriter, r *http.Request) {
	response := nsi.latestBlockDetails(nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) LatencyHandler(w http.ResponseWriter, r *http.Request) {
	response := nsi.latency(nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

//func (nsi *NodeServiceImpl) LogsHandler(w http.ResponseWriter, r *http.Request) {
//	response := nsi.logs()
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	json.NewEncoder(w).Encode(response)
//}

func (nsi *NodeServiceImpl) TransactionSearchHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response := nsi.transactionSearchDetails(params["txn_hash"], nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) MailServerConfigHandler(w http.ResponseWriter, r *http.Request) {
	var request MailServerConfig
	_ = json.NewDecoder(r.Body).Decode(&request)
	response := nsi.emailServerConfig(request.Host, request.Port, request.Username, request.Password, request.RecipientList, nsi.Url)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) OptionsHandler(w http.ResponseWriter, r *http.Request) {
	response := "Options Handled"
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) GetChartDataHandler(w http.ResponseWriter, r *http.Request) {
	response := nsi.GetChartData(nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) GetContractListHandler(w http.ResponseWriter, r *http.Request) {
	response := nsi.ContractList()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) GetContractCountHandler(w http.ResponseWriter, r *http.Request) {
	response := nsi.ContractCount()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) ContractDetailsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
	contractAddress := r.FormValue("address")
	contractName := r.FormValue("name")
	description := r.FormValue("description")
	file, header, err := r.FormFile("abi")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	io.Copy(&Buf, file)
	content := Buf.String()

	var jsonContent contractJSON

	json.Unmarshal([]byte(content), &jsonContent)
	abiContent, _ := json.Marshal(jsonContent.Abi)
	abiString := make([]string, len(abiContent))
	for i := 0; i < len(abiContent); i++ {
		abiString[i] = string(abiContent[i])
	}
	abiData := fmt.Sprint(strings.Join(abiString, ""))

	interfaceContent, _ := json.Marshal(jsonContent.Interface)
	interfaceString := make([]string, len(interfaceContent))
	for i := 0; i < len(interfaceContent); i++ {
		interfaceString[i] = string(interfaceContent[i])
	}
	interfaceData := fmt.Sprint(strings.Join(interfaceString, ""))
	bytecodeData := jsonContent.Bytecode
	var data string
	if len(abiData) != 4 {
		data = abiData
	} else if len(interfaceData) != 4 {
		data = interfaceData
	} else {
		data = content
		data = strings.Replace(data, "\n", "", -1)
	}

	jsonString := util.ComposeJSON(data, bytecodeData, contractAddress)

	path := "./contracts"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0775)
	}
	path = "./contracts/" + contractAddress + "_" + name[0]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0775)
	}

	filePath := path + "/" + name[0] + ".json"
	jsByte := []byte(jsonString)
	err = ioutil.WriteFile(filePath, jsByte, 0775)
	if err != nil {
		fmt.Println(err)
	}

	Buf.Reset()
	response := nsi.updateContractDetails(contractAddress, contractName, data, description)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) AttachedNodeDetailsHandler(w http.ResponseWriter, r *http.Request) {
	var successResponse SuccessResponse
	var Buf bytes.Buffer
	gethLogsDirectory := r.FormValue("gethPath")
	constellationLogsDirectory := r.FormValue("constellationPath")

	file, _, err := r.FormFile("genesis")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	io.Copy(&Buf, file)
	content := Buf.String()

	filePath := "/home/node/genesis.json"
	jsByte := []byte(content)
	err = ioutil.WriteFile(filePath, jsByte, 0775)
	if err != nil {
		fmt.Println(err)
	}

	var jsonContent genesisJSON
	json.Unmarshal([]byte(content), &jsonContent)
	chainIdAppend := fmt.Sprint("NETWORK_ID=", jsonContent.Config.ChainId, "\n")
	util.AppendStringToFile("/home/setup.conf", chainIdAppend)
	util.InsertStringToFile("/home/start.sh", "	   -v "+gethLogsDirectory+":/home/node/qdata/gethLogs \\\n", 13)
	util.InsertStringToFile("/home/start.sh", "	   -v "+constellationLogsDirectory+":/home/node/qdata/constellationLogs \\\n", 13)

	Buf.Reset()
	fmt.Println("Updates have been saved. Please press Ctrl+C to exit from this container and run start.sh to apply changes")
	state := currentState()
	if state == "NI" {
		util.DeleteProperty("STATE=NI", "/home/setup.conf")
		stateInitialized := fmt.Sprint("STATE=I\n")
		util.AppendStringToFile("/home/setup.conf", stateInitialized)
	}
	successResponse.Status = "Updates have been saved. Please press Ctrl+C from CLI to exit from this container and run start.sh to apply changes"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(successResponse)
}

func (nsi *NodeServiceImpl) InitializationHandler(w http.ResponseWriter, r *http.Request) {
	response := nsi.returnCurrentInitializationState()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var request accountPassword
	_ = json.NewDecoder(r.Body).Decode(&request)
	password := request.Password
	response := nsi.createAccount(password, nsi.Url)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
	json.NewEncoder(w).Encode(response)
}

func (nsi *NodeServiceImpl) GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	response := nsi.getAccounts(nsi.Url)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
