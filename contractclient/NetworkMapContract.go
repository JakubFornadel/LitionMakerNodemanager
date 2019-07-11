package contractclient

import (
	"gitlab.com/lition/quorum-maker-nodemanager/client"
	"gitlab.com/lition/quorum-maker-nodemanager/contracthandler"
)

const registerNodeFunSig = "0x82cb1a2a"
const updateNodeFunSig = "0xe1d33203"
const getNodeDetailsFunSig = "0x7f11a8ed"

type NodeDetails struct {
	Name      string `json:"nodeName,omitempty"`
	Role      string `json:"role,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	Enode     string `json:"enode,omitempty"`
	IP        string `json:"ip,omitempty"`
}

type NetworkMapContractClient struct {
	client.EthClient
	ContractParam contracthandler.ContractParam
}

type GetNodeDetailsParam int

func (nmc *NetworkMapContractClient) SetContractParam(cp contracthandler.ContractParam) {
	nmc.ContractParam = cp
}

func (nmc *NetworkMapContractClient) RegisterNode(name string, role string, publicKey string, enode string, ip string) string {

	if nmc.ContractParam.To == "" || nmc.ContractParam.From == "" {
		return ""
	}

	nd := NodeDetails{name, role, publicKey, enode, ip}
	nodeList := nmc.GetNodeDetailsList()
	for _, nodeDetails := range nodeList {
		if nodeDetails.Enode == enode {
			return "Exists"
		}
	}
	return nmc.SendTransaction(nmc.ContractParam, RegisterUpdateNodeFuncHandler{nd, registerNodeFunSig})

}

func (nmc *NetworkMapContractClient) GetNodeDetails(i int) NodeDetails {

	if nmc.ContractParam.To == "" || nmc.ContractParam.From == "" {
		return NodeDetails{}
	}

	encoderDecoder := GetNodeDetailsFuncHandler{index: i, funcSig: getNodeDetailsFunSig}
	nmc.EthCall(nmc.ContractParam, encoderDecoder, &encoderDecoder)

	return encoderDecoder.result
}

func (nmc *NetworkMapContractClient) GetNodeDetailsList() []NodeDetails {

	if nmc.ContractParam.To == "" || nmc.ContractParam.From == "" {
		return []NodeDetails{}
	}

	var list []NodeDetails

	for i := 0; true; i++ {
		encoderDecoder := GetNodeDetailsFuncHandler{index: i, funcSig: getNodeDetailsFunSig}
		nmc.EthCall(nmc.ContractParam, encoderDecoder, &encoderDecoder)

		if encoderDecoder.result.Enode != "" && len(encoderDecoder.result.Enode) > 0 {
			list = append(list, encoderDecoder.result)
		} else {
			return list
		}
	}

	return list
}

func (nmc *NetworkMapContractClient) UpdateNode(name string, role string, publicKey string, enode string, ip string) string {

	if nmc.ContractParam.To == "" || nmc.ContractParam.From == "" {
		return ""
	}

	nd := NodeDetails{name, role, publicKey, enode, ip}
	return nmc.SendTransaction(nmc.ContractParam, RegisterUpdateNodeFuncHandler{nd, updateNodeFunSig})
}

type RegisterUpdateNodeFuncHandler struct {
	nd      NodeDetails
	funcSig string
}

func (h RegisterUpdateNodeFuncHandler) Encode() string {

	sig := "string,string,string,string,string"

	param := []interface{}{h.nd.Name, h.nd.Role, h.nd.PublicKey, h.nd.Enode, h.nd.IP}

	data := h.funcSig + contracthandler.FunctionProcessor{sig}.Encode(param)

	return data
}

type GetNodeDetailsFuncHandler struct {
	index   int
	funcSig string
	result  NodeDetails
}

func (g *GetNodeDetailsFuncHandler) Decode(r string) {
	var nd NodeDetails

	if r == "" || len(r) < 1 {
		g.result = nd
		return
	}

	sig := "string,string,string,string,string,uint16"

	resultArray := contracthandler.FunctionProcessor{sig}.Decode(r)

	g.result = NodeDetails{resultArray[0].(string), resultArray[1].(string), resultArray[2].(string), resultArray[4].(string), resultArray[3].(string)}
}

func (g GetNodeDetailsFuncHandler) Encode() string {

	sig := "uint16"

	param := []interface{}{g.index}

	return g.funcSig + contracthandler.FunctionProcessor{sig}.Encode(param)
}

type DeployContractHandler struct {
	binary string
}

func (d DeployContractHandler) Encode() string {

	return d.binary
}
