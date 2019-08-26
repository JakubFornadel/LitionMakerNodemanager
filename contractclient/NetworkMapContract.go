package contractclient

import (
	"math/big"

	log "github.com/sirupsen/logrus"

	"gitlab.com/lition/lition-maker-nodemanager/client"
	internalContract "gitlab.com/lition/lition-maker-nodemanager/contractclient/internalcontract"
	"gitlab.com/lition/lition/accounts/abi/bind"
	"gitlab.com/lition/lition/common"
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

type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

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
		log.Error("RegisterNode: ", err)
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
		log.Error("GetNodeDetails: ", err)
		return NodeDetails{}
	}

	return NodeDetails{details.N, details.R, details.P, details.E, details.Ip}
}

func (nmc *NetworkMapContractClient) GetNodeDetailsList() []NodeDetails {

	var list []NodeDetails

	if nmc.Ic == nil {
		return list
	}

	for i := 0; true; i++ {
		details, err := nmc.Ic.GetNodeDetails(nil, uint16(i))
		if err != nil {
			return list
		}
		if details.E != "" && len(details.E) > 0 {
			list = append(list, NodeDetails{details.N, details.R, details.P, details.E, details.Ip})
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
		log.Error("GetNodeCount", err)
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
		log.Error("UpdateNode: ", err)
		return ""
	}
	return tx.Hash().String()
}

func (nmc *NetworkMapContractClient) GetSignatureHashFromNotary(notary_block int64, miners []common.Address, blocks_mined []uint32, users []common.Address, user_gas []uint32, largest_tx uint32) []byte {
	if nmc.Ic == nil {
		return []byte{}
	}
	response, err := nmc.Ic.GetSignatureHashFromNotary(nil, big.NewInt(notary_block), miners, blocks_mined, users, user_gas, largest_tx)
	if err != nil {
		log.Error("GetSignatureHashFromNotary: ", err)
		return []byte{}
	}
	return response[:]
}

func (nmc *NetworkMapContractClient) GetSignatures(notary_block int64, index int) Signature {
	if nmc.Ic == nil {
		return Signature{}
	}
	result, err := nmc.Ic.GetSignatures(nil, big.NewInt(notary_block), big.NewInt(int64(index)))
	if err != nil {
		log.Error("GetSignatures: ", err)
		return Signature{}
	}
	return result
}

func (nmc *NetworkMapContractClient) GetSignaturesCount(notary_block int64) int {
	if nmc.Ic == nil {
		return 0
	}
	result, err := nmc.Ic.GetSignaturesCount(nil, big.NewInt(notary_block))
	if err != nil {
		log.Error("GetSignatures: ", err)
		return 0
	}
	return int(result.Int64())
}

func (nmc *NetworkMapContractClient) StoreSignature(notary_block int64, sig Signature) {
	if nmc.Ic == nil {
		return
	}
	_, err := nmc.Ic.StoreSignature(nmc.Auth, big.NewInt(notary_block), sig.V, sig.R, sig.S)
	if err != nil {
		log.Error("StoreSignature: ", err)
	}
}

type DeployContractHandler struct {
	binary string
}

func (d DeployContractHandler) Encode() string {

	return d.binary
}
