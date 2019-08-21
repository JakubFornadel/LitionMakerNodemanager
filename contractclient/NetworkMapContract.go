package contractclient

import (
	"log"

	"gitlab.com/lition/lition-maker-nodemanager/client"
	internalContract "gitlab.com/lition/lition-maker-nodemanager/contractclient/internalcontract"
	"gitlab.com/lition/lition/accounts/abi/bind"
)

type NodeDetails struct {
	Name      string `json:"nodeName,omitempty"`
	Role      string `json:"role,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	Enode     string `json:"enode,omitempty"`
	IP        string `json:"ip,omitempty"`
}

type NetworkMapContractClient struct {
	client.EthClient
	Auth *bind.TransactOpts
	Ic   *internalContract.Lition
}

type GetNodeDetailsParam int

func (nmc *NetworkMapContractClient) RegisterNode(name string, role string, publicKey string, enode string, ip string) string {

	if nmc.Ic == nil {
		return ""
	}

	nodeList := nmc.GetNodeDetailsList()
	for _, nodeDetails := range nodeList {
		if nodeDetails.Enode == enode {
			return "Exists"
		}
	}

	tx, err := nmc.Ic.RegisterNode(nmc.Auth, name, role, publicKey, enode, ip)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return tx.Hash().String()
}

func (nmc *NetworkMapContractClient) GetNodeDetails(i int) NodeDetails {

	if nmc.Ic == nil {
		return NodeDetails{}
	}

	details, err := nmc.Ic.GetNodeDetails(nil, uint16(i))
	if err != nil {
		log.Fatal(err)
		return NodeDetails{}
	}

	return NodeDetails{details.N, details.R, details.P, details.Ip, details.E}
}

func (nmc *NetworkMapContractClient) GetNodeDetailsList() []NodeDetails {

	var list []NodeDetails

	if nmc.Ic == nil {
		return list
	}

	for i := 0; true; i++ {
		details, err := nmc.Ic.GetNodeDetails(nil, uint16(i))
		if err != nil {
			log.Fatal(err)
			return list
		}
		if details.E != "" && len(details.E) > 0 {
			list = append(list, NodeDetails{details.N, details.R, details.P, details.Ip, details.E})
		} else {
			return list
		}
	}

	return list
}

func (nmc *NetworkMapContractClient) GetNodeCount() int {

	if nmc.Ic == nil {
		return 0
	}

	count, err := nmc.Ic.GetNodesCounter(nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return int(count.Int64())
}

func (nmc *NetworkMapContractClient) UpdateNode(name string, role string, publicKey string, enode string, ip string) string {

	if nmc.Ic == nil {
		return ""
	}
	tx, err := nmc.Ic.UpdateNode(nmc.Auth, name, role, publicKey, enode, ip)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return tx.Hash().String()
}

type DeployContractHandler struct {
	binary string
}

func (d DeployContractHandler) Encode() string {

	return d.binary
}
