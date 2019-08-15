pragma solidity^0.5.0;
contract NetworkManagerContract {

    uint nodeCounter;

    struct NodeDetails {
        string nodeName;
        string role;
        string publickey;
        string enode;
        string ip;
    }

    mapping (string => NodeDetails)nodes;
    string[] enodeList;

    event print(string nodeName, string role,string publickey, string enode, string ip);

    function registerNode(string memory n, string memory r, string memory p, string memory e, string memory ip) public {

        nodes[e].publickey = p;
        nodes[e].nodeName = n;
        nodes[e].role = r;
        nodes[e].ip = ip;
        enodeList.push(e);
        emit print(n, r, p, e, ip);

    }

    function getNodeDetails(uint16 _index) public view returns (string memory n, string memory r, string memory p, string memory ip, string memory e, uint i) {
        NodeDetails memory nodeInfo = nodes[enodeList[_index]];
        return (
            nodeInfo.nodeName,
            nodeInfo.role,
            nodeInfo.publickey,
            nodeInfo.ip,
            enodeList[_index],
            _index
        );
    }

    function getNodesCounter() public view  returns (uint) {
        return enodeList.length;
    }

    function updateNode(string memory n, string memory r, string memory p, string memory e, string memory ip) public {

        nodes[e].publickey = p;
        nodes[e].nodeName = n;
        nodes[e].role = r;
        nodes[e].ip = ip;
        emit print(n, r, p, e, ip);
    }

    function getNodeList(uint16 i)  public  view   returns (string memory n, string memory r,string memory p,string memory ip,string memory e) {

        NodeDetails memory nodeInfo = nodes[enodeList[i]];
        return (
            nodeInfo.nodeName,
            nodeInfo.role,
            nodeInfo.publickey,
	        nodeInfo.ip,
            enodeList[i]
        );
    }

    function get_signature_hash_from_notary(uint32 notary_block, address[] memory miners,
                                  uint32[] memory blocks_mined, address[] memory users,
                                  uint32[] memory user_gas, uint32 largest_tx)
                                      public pure returns (bytes32) {
       bytes memory prefix = "\x19Ethereum Signed Message:\n32";
       string encoded_notary = abi.encodePacked(notary_block, miners, blocks_mined, users, user_gas, largest_tx);
       bytes32 notary_hash keccak256(encoded_notary);
       return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", _hash));
    }

    struct signatures {
       uint8[] vs;
       bytes32[] rs;
       bytes32[] ss;
    }

    mapping (uint256 => signatures) public sigs;

    function store_signature(uint256 block_no, uint8 v, bytes32 r, bytes32 s) public {
       sigs[block_no].vs.push(v);
       sigs[block_no].rs.push(r);
       sigs[block_no].ss.push(s);
    }

    function get_signatures_count(uint256 block_no) public view returns (uint256) {
       return sigs[block_no].vs.length;
    }

    function get_signatures(uint256 block_no, uint256 index) public view returns (uint8[] memory v, bytes32[] memory r, bytes32[] memory s) {
       v = sigs[block_no].vs[index];
       r = sigs[block_no].rs[index];
       s = sigs[block_no].ss[index];
    }
}
