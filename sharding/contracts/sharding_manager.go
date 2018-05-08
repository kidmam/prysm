// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// SMCABI is the input ABI used to generate the binding from.
const SMCABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_shardId\",\"type\":\"uint256\"},{\"name\":\"_period\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"},{\"name\":\"_chunkRoot\",\"type\":\"bytes32\"}],\"name\":\"submitVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deregisterNotary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerNotary\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"notaryRegistry\",\"outputs\":[{\"name\":\"deregisteredPeriod\",\"type\":\"uint256\"},{\"name\":\"poolIndex\",\"type\":\"uint256\"},{\"name\":\"deposited\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_shardId\",\"type\":\"uint256\"},{\"name\":\"_period\",\"type\":\"uint256\"},{\"name\":\"_chunkRoot\",\"type\":\"bytes32\"}],\"name\":\"addHeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastSubmittedCollation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastApprovedCollation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseNotary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"notaryPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shardId\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNotaryInCommittee\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"collationRecords\",\"outputs\":[{\"name\":\"chunkRoot\",\"type\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"isElected\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notaryPoolLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"shardId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"chunkRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"proposerAddress\",\"type\":\"address\"}],\"name\":\"HeaderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"NotaryRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"poolIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deregisteredPeriod\",\"type\":\"uint256\"}],\"name\":\"NotaryDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"NotaryReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"shardId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"chunkRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"notaryAddress\",\"type\":\"address\"}],\"name\":\"VoteSubmitted\",\"type\":\"event\"}]"

// SMCBin is the compiled bytecode used for deploying new contracts.
const SMCBin = `0x608060405234801561001057600080fd5b50610c43806100206000396000f3006080604052600436106100c45763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630c8da4cc81146100c95780634f33ffa0146100f357806358377bd11461011657806368e9513e1461012b5780636bdd32711461013357806375bd99121461017457806383ceeabe1461019257806397d369a2146101aa5780639910851d146101c2578063a81f4510146101d7578063b09f427e1461020b578063e9e0b68314610226578063f6f67d3614610269575b600080fd5b3480156100d557600080fd5b506100e160043561027e565b60408051918252519081900360200190f35b3480156100ff57600080fd5b50610114600435602435604435606435610290565b005b34801561012257600080fd5b50610114610436565b610114610561565b34801561013f57600080fd5b50610154600160a060020a0360043516610710565b604080519384526020840192909252151582820152519081900360600190f35b34801561018057600080fd5b50610114600435602435604435610733565b34801561019e57600080fd5b506100e1600435610885565b3480156101b657600080fd5b506100e1600435610897565b3480156101ce57600080fd5b506101146108a9565b3480156101e357600080fd5b506101ef6004356109d5565b60408051600160a060020a039092168252519081900360200190f35b34801561021757600080fd5b506101ef6004356024356109fd565b34801561023257600080fd5b50610241600435602435610a92565b60408051938452600160a060020a039092166020840152151582820152519081900360600190f35b34801561027557600080fd5b506100e1610adc565b60036020526000908152604090205481565b60008085101580156102a25750606485105b15156102ad57600080fd5b6005430484146102bc57600080fd5b60008581526005602052604090205484146102d657600080fd5b608783106102e357600080fd5b6000858152600460209081526040808320878452909152902054821461030857600080fd5b600160a060020a03331660009081526001602052604090206002015460ff16151561033257600080fd5b61033c8584610ae2565b1561034657600080fd5b33600160a060020a031661035a86856109fd565b600160a060020a03161461036d57600080fd5b6103778584610b05565b61038085610b29565b9050605a81106103e4576000858152600660209081526040808320879055600482528083208784529091529020600101805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790555b6040805183815260208101869052600160a060020a03331681830152905186917fc99370212b708f699fb6945a17eb34d0fc1ccd5b45d88f4d9682593a45d6e833919081900360600190a25050505050565b33600160a060020a038116600090815260016020819052604082209081015460029091015490919060ff16151561046c57600080fd5b82600160a060020a031660008381548110151561048557fe5b600091825260209091200154600160a060020a0316146104a457600080fd5b6104ac610b3e565b50600160a060020a038216600090815260016020526040902060054304908190556104d682610b61565b60008054839081106104e457fe5b600091825260209182902001805473ffffffffffffffffffffffffffffffffffffffff191690556002805460001901905560408051600160a060020a0386168152918201849052818101839052517f90e5afdc8fd31453dcf6e37154fa117ddf3b0324c96c65015563df9d5e4b5a759181900360600190a1505050565b33600160a060020a03811660009081526001602052604081206002015460ff161561058b57600080fd5b34683635c9adc5dea00000146105a057600080fd5b6105a8610b3e565b6105b0610bd2565b156106135750600254600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03841617905561065c565b61061b610bd9565b90508160008281548110151561062d57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055505b60028054600190810182556040805160608101825260008082526020808301878152838501868152600160a060020a038a168452918690529390912091518255915192810192909255519101805460ff1916911515919091179055600a5481106106c85760018101600a555b60408051600160a060020a03841681526020810183905281517fa4fe15c53db34d35a5117acc26c27a2653dc68e2dadfc21ed211e38b7864d7a7929181900390910190a15050565b600160208190526000918252604090912080549181015460029091015460ff1683565b600083101580156107445750606483105b151561074f57600080fd5b60054304821461075e57600080fd5b600083815260056020526040902054821161077857600080fd5b610780610b3e565b604080516060808201835283825233600160a060020a03908116602080850182815260008688018181528b8252600484528882208b8352845288822097518855915160019097018054925173ffffffffffffffffffffffffffffffffffffffff19909316979095169690961774ff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000091151591909102179092558784526005808352858520439190910490556003825284842093909355835185815290810186905280840192909252915185927f2d0a86178d2fd307b47be157a766e6bee19bc26161c32f9781ee0e818636f09c928290030190a2505050565b60056020526000908152604090205481565b60066020526000908152604090205481565b33600160a060020a038116600090815260016020819052604090912080820154600290910154909160ff9091161515146108e257600080fd5b600160a060020a038216600090815260016020526040902054151561090657600080fd5b600160a060020a038216600090815260016020526040902054613f0001600543041161093157600080fd5b600160a060020a03821660008181526001602081905260408083208381559182018390556002909101805460ff1916905551683635c9adc5dea000009082818181858883f1935050505015801561098c573d6000803e3d6000fd5b5060408051600160a060020a03841681526020810183905281517faee20171b64b7f3360a142659094ce929970d6963dcea8c34a9bf1ece8033680929181900390910190a15050565b60008054829081106109e357fe5b600091825260209091200154600160a060020a0316905081565b60008080808080600543049450600b54851115610a1e57600a549350610a24565b60095493505b6040805160001960058802018040808352602083018b90528284018c90529251918290036060019091209094509092508490811515610a5f57fe5b069050600081815481101515610a7157fe5b600091825260209091200154600160a060020a031698975050505050505050565b600460209081526000928352604080842090915290825290208054600190910154600160a060020a0381169074010000000000000000000000000000000000000000900460ff1683565b60025481565b60009182526003602052604090912054600160ff9290920360020a900481161490565b600091825260036020526040909120805460ff9290920360020a9091176001019055565b60009081526003602052604090205460ff1690565b600b546005430490811015610b5257610b5e565b600a54600955600b8190555b50565b6008546007541415610ba757600780546001810182556000919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68801819055610bc6565b806007600854815481101515610bb957fe5b6000918252602090912001555b50600880546001019055565b6008541590565b60006001600854111515610bec57600080fd5b600880546000190190819055600780549091908110610c0757fe5b90600052602060002001549050905600a165627a7a7230582025e4cd50e9d13e755d9e6ea3c2abfa8998c6141cc47c803fe8a649cade7c12140029`

// DeploySMC deploys a new Ethereum contract, binding an instance of SMC to it.
func DeploySMC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SMC, error) {
	parsed, err := abi.JSON(strings.NewReader(SMCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SMCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SMC{SMCCaller: SMCCaller{contract: contract}, SMCTransactor: SMCTransactor{contract: contract}, SMCFilterer: SMCFilterer{contract: contract}}, nil
}

// SMC is an auto generated Go binding around an Ethereum contract.
type SMC struct {
	SMCCaller     // Read-only binding to the contract
	SMCTransactor // Write-only binding to the contract
	SMCFilterer   // Log filterer for contract events
}

// SMCCaller is an auto generated read-only Go binding around an Ethereum contract.
type SMCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SMCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SMCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SMCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SMCSession struct {
	Contract     *SMC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SMCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SMCCallerSession struct {
	Contract *SMCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SMCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SMCTransactorSession struct {
	Contract     *SMCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SMCRaw is an auto generated low-level Go binding around an Ethereum contract.
type SMCRaw struct {
	Contract *SMC // Generic contract binding to access the raw methods on
}

// SMCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SMCCallerRaw struct {
	Contract *SMCCaller // Generic read-only contract binding to access the raw methods on
}

// SMCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SMCTransactorRaw struct {
	Contract *SMCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSMC creates a new instance of SMC, bound to a specific deployed contract.
func NewSMC(address common.Address, backend bind.ContractBackend) (*SMC, error) {
	contract, err := bindSMC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SMC{SMCCaller: SMCCaller{contract: contract}, SMCTransactor: SMCTransactor{contract: contract}, SMCFilterer: SMCFilterer{contract: contract}}, nil
}

// NewSMCCaller creates a new read-only instance of SMC, bound to a specific deployed contract.
func NewSMCCaller(address common.Address, caller bind.ContractCaller) (*SMCCaller, error) {
	contract, err := bindSMC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SMCCaller{contract: contract}, nil
}

// NewSMCTransactor creates a new write-only instance of SMC, bound to a specific deployed contract.
func NewSMCTransactor(address common.Address, transactor bind.ContractTransactor) (*SMCTransactor, error) {
	contract, err := bindSMC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SMCTransactor{contract: contract}, nil
}

// NewSMCFilterer creates a new log filterer instance of SMC, bound to a specific deployed contract.
func NewSMCFilterer(address common.Address, filterer bind.ContractFilterer) (*SMCFilterer, error) {
	contract, err := bindSMC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SMCFilterer{contract: contract}, nil
}

// bindSMC binds a generic wrapper to an already deployed contract.
func bindSMC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SMCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SMC *SMCRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SMC.Contract.SMCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SMC *SMCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMC.Contract.SMCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SMC *SMCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SMC.Contract.SMCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SMC *SMCCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SMC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SMC *SMCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SMC *SMCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SMC.Contract.contract.Transact(opts, method, params...)
}

// CollationRecords is a free data retrieval call binding the contract method 0xe9e0b683.
//
// Solidity: function collationRecords( uint256,  uint256) constant returns(chunkRoot bytes32, proposer address, isElected bool)
func (_SMC *SMCCaller) CollationRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	ChunkRoot [32]byte
	Proposer  common.Address
	IsElected bool
}, error) {
	ret := new(struct {
		ChunkRoot [32]byte
		Proposer  common.Address
		IsElected bool
	})
	out := ret
	err := _SMC.contract.Call(opts, out, "collationRecords", arg0, arg1)
	return *ret, err
}

// CollationRecords is a free data retrieval call binding the contract method 0xe9e0b683.
//
// Solidity: function collationRecords( uint256,  uint256) constant returns(chunkRoot bytes32, proposer address, isElected bool)
func (_SMC *SMCSession) CollationRecords(arg0 *big.Int, arg1 *big.Int) (struct {
	ChunkRoot [32]byte
	Proposer  common.Address
	IsElected bool
}, error) {
	return _SMC.Contract.CollationRecords(&_SMC.CallOpts, arg0, arg1)
}

// CollationRecords is a free data retrieval call binding the contract method 0xe9e0b683.
//
// Solidity: function collationRecords( uint256,  uint256) constant returns(chunkRoot bytes32, proposer address, isElected bool)
func (_SMC *SMCCallerSession) CollationRecords(arg0 *big.Int, arg1 *big.Int) (struct {
	ChunkRoot [32]byte
	Proposer  common.Address
	IsElected bool
}, error) {
	return _SMC.Contract.CollationRecords(&_SMC.CallOpts, arg0, arg1)
}

// CurrentVote is a free data retrieval call binding the contract method 0x0c8da4cc.
//
// Solidity: function currentVote( uint256) constant returns(bytes32)
func (_SMC *SMCCaller) CurrentVote(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SMC.contract.Call(opts, out, "currentVote", arg0)
	return *ret0, err
}

// CurrentVote is a free data retrieval call binding the contract method 0x0c8da4cc.
//
// Solidity: function currentVote( uint256) constant returns(bytes32)
func (_SMC *SMCSession) CurrentVote(arg0 *big.Int) ([32]byte, error) {
	return _SMC.Contract.CurrentVote(&_SMC.CallOpts, arg0)
}

// CurrentVote is a free data retrieval call binding the contract method 0x0c8da4cc.
//
// Solidity: function currentVote( uint256) constant returns(bytes32)
func (_SMC *SMCCallerSession) CurrentVote(arg0 *big.Int) ([32]byte, error) {
	return _SMC.Contract.CurrentVote(&_SMC.CallOpts, arg0)
}

// GetNotaryInCommittee is a free data retrieval call binding the contract method 0xb09f427e.
//
// Solidity: function getNotaryInCommittee(_shardId uint256, _index uint256) constant returns(address)
func (_SMC *SMCCaller) GetNotaryInCommittee(opts *bind.CallOpts, _shardId *big.Int, _index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SMC.contract.Call(opts, out, "getNotaryInCommittee", _shardId, _index)
	return *ret0, err
}

// GetNotaryInCommittee is a free data retrieval call binding the contract method 0xb09f427e.
//
// Solidity: function getNotaryInCommittee(_shardId uint256, _index uint256) constant returns(address)
func (_SMC *SMCSession) GetNotaryInCommittee(_shardId *big.Int, _index *big.Int) (common.Address, error) {
	return _SMC.Contract.GetNotaryInCommittee(&_SMC.CallOpts, _shardId, _index)
}

// GetNotaryInCommittee is a free data retrieval call binding the contract method 0xb09f427e.
//
// Solidity: function getNotaryInCommittee(_shardId uint256, _index uint256) constant returns(address)
func (_SMC *SMCCallerSession) GetNotaryInCommittee(_shardId *big.Int, _index *big.Int) (common.Address, error) {
	return _SMC.Contract.GetNotaryInCommittee(&_SMC.CallOpts, _shardId, _index)
}

// LastApprovedCollation is a free data retrieval call binding the contract method 0x97d369a2.
//
// Solidity: function lastApprovedCollation( uint256) constant returns(uint256)
func (_SMC *SMCCaller) LastApprovedCollation(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SMC.contract.Call(opts, out, "lastApprovedCollation", arg0)
	return *ret0, err
}

// LastApprovedCollation is a free data retrieval call binding the contract method 0x97d369a2.
//
// Solidity: function lastApprovedCollation( uint256) constant returns(uint256)
func (_SMC *SMCSession) LastApprovedCollation(arg0 *big.Int) (*big.Int, error) {
	return _SMC.Contract.LastApprovedCollation(&_SMC.CallOpts, arg0)
}

// LastApprovedCollation is a free data retrieval call binding the contract method 0x97d369a2.
//
// Solidity: function lastApprovedCollation( uint256) constant returns(uint256)
func (_SMC *SMCCallerSession) LastApprovedCollation(arg0 *big.Int) (*big.Int, error) {
	return _SMC.Contract.LastApprovedCollation(&_SMC.CallOpts, arg0)
}

// LastSubmittedCollation is a free data retrieval call binding the contract method 0x83ceeabe.
//
// Solidity: function lastSubmittedCollation( uint256) constant returns(uint256)
func (_SMC *SMCCaller) LastSubmittedCollation(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SMC.contract.Call(opts, out, "lastSubmittedCollation", arg0)
	return *ret0, err
}

// LastSubmittedCollation is a free data retrieval call binding the contract method 0x83ceeabe.
//
// Solidity: function lastSubmittedCollation( uint256) constant returns(uint256)
func (_SMC *SMCSession) LastSubmittedCollation(arg0 *big.Int) (*big.Int, error) {
	return _SMC.Contract.LastSubmittedCollation(&_SMC.CallOpts, arg0)
}

// LastSubmittedCollation is a free data retrieval call binding the contract method 0x83ceeabe.
//
// Solidity: function lastSubmittedCollation( uint256) constant returns(uint256)
func (_SMC *SMCCallerSession) LastSubmittedCollation(arg0 *big.Int) (*big.Int, error) {
	return _SMC.Contract.LastSubmittedCollation(&_SMC.CallOpts, arg0)
}

// NotaryPool is a free data retrieval call binding the contract method 0xa81f4510.
//
// Solidity: function notaryPool( uint256) constant returns(address)
func (_SMC *SMCCaller) NotaryPool(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SMC.contract.Call(opts, out, "notaryPool", arg0)
	return *ret0, err
}

// NotaryPool is a free data retrieval call binding the contract method 0xa81f4510.
//
// Solidity: function notaryPool( uint256) constant returns(address)
func (_SMC *SMCSession) NotaryPool(arg0 *big.Int) (common.Address, error) {
	return _SMC.Contract.NotaryPool(&_SMC.CallOpts, arg0)
}

// NotaryPool is a free data retrieval call binding the contract method 0xa81f4510.
//
// Solidity: function notaryPool( uint256) constant returns(address)
func (_SMC *SMCCallerSession) NotaryPool(arg0 *big.Int) (common.Address, error) {
	return _SMC.Contract.NotaryPool(&_SMC.CallOpts, arg0)
}

// NotaryPoolLength is a free data retrieval call binding the contract method 0xf6f67d36.
//
// Solidity: function notaryPoolLength() constant returns(uint256)
func (_SMC *SMCCaller) NotaryPoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SMC.contract.Call(opts, out, "notaryPoolLength")
	return *ret0, err
}

// NotaryPoolLength is a free data retrieval call binding the contract method 0xf6f67d36.
//
// Solidity: function notaryPoolLength() constant returns(uint256)
func (_SMC *SMCSession) NotaryPoolLength() (*big.Int, error) {
	return _SMC.Contract.NotaryPoolLength(&_SMC.CallOpts)
}

// NotaryPoolLength is a free data retrieval call binding the contract method 0xf6f67d36.
//
// Solidity: function notaryPoolLength() constant returns(uint256)
func (_SMC *SMCCallerSession) NotaryPoolLength() (*big.Int, error) {
	return _SMC.Contract.NotaryPoolLength(&_SMC.CallOpts)
}

// NotaryRegistry is a free data retrieval call binding the contract method 0x6bdd3271.
//
// Solidity: function notaryRegistry( address) constant returns(deregisteredPeriod uint256, poolIndex uint256, deposited bool)
func (_SMC *SMCCaller) NotaryRegistry(opts *bind.CallOpts, arg0 common.Address) (struct {
	DeregisteredPeriod *big.Int
	PoolIndex          *big.Int
	Deposited          bool
}, error) {
	ret := new(struct {
		DeregisteredPeriod *big.Int
		PoolIndex          *big.Int
		Deposited          bool
	})
	out := ret
	err := _SMC.contract.Call(opts, out, "notaryRegistry", arg0)
	return *ret, err
}

// NotaryRegistry is a free data retrieval call binding the contract method 0x6bdd3271.
//
// Solidity: function notaryRegistry( address) constant returns(deregisteredPeriod uint256, poolIndex uint256, deposited bool)
func (_SMC *SMCSession) NotaryRegistry(arg0 common.Address) (struct {
	DeregisteredPeriod *big.Int
	PoolIndex          *big.Int
	Deposited          bool
}, error) {
	return _SMC.Contract.NotaryRegistry(&_SMC.CallOpts, arg0)
}

// NotaryRegistry is a free data retrieval call binding the contract method 0x6bdd3271.
//
// Solidity: function notaryRegistry( address) constant returns(deregisteredPeriod uint256, poolIndex uint256, deposited bool)
func (_SMC *SMCCallerSession) NotaryRegistry(arg0 common.Address) (struct {
	DeregisteredPeriod *big.Int
	PoolIndex          *big.Int
	Deposited          bool
}, error) {
	return _SMC.Contract.NotaryRegistry(&_SMC.CallOpts, arg0)
}

// AddHeader is a paid mutator transaction binding the contract method 0x75bd9912.
//
// Solidity: function addHeader(_shardId uint256, _period uint256, _chunkRoot bytes32) returns()
func (_SMC *SMCTransactor) AddHeader(opts *bind.TransactOpts, _shardId *big.Int, _period *big.Int, _chunkRoot [32]byte) (*types.Transaction, error) {
	return _SMC.contract.Transact(opts, "addHeader", _shardId, _period, _chunkRoot)
}

// AddHeader is a paid mutator transaction binding the contract method 0x75bd9912.
//
// Solidity: function addHeader(_shardId uint256, _period uint256, _chunkRoot bytes32) returns()
func (_SMC *SMCSession) AddHeader(_shardId *big.Int, _period *big.Int, _chunkRoot [32]byte) (*types.Transaction, error) {
	return _SMC.Contract.AddHeader(&_SMC.TransactOpts, _shardId, _period, _chunkRoot)
}

// AddHeader is a paid mutator transaction binding the contract method 0x75bd9912.
//
// Solidity: function addHeader(_shardId uint256, _period uint256, _chunkRoot bytes32) returns()
func (_SMC *SMCTransactorSession) AddHeader(_shardId *big.Int, _period *big.Int, _chunkRoot [32]byte) (*types.Transaction, error) {
	return _SMC.Contract.AddHeader(&_SMC.TransactOpts, _shardId, _period, _chunkRoot)
}

// DeregisterNotary is a paid mutator transaction binding the contract method 0x58377bd1.
//
// Solidity: function deregisterNotary() returns()
func (_SMC *SMCTransactor) DeregisterNotary(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMC.contract.Transact(opts, "deregisterNotary")
}

// DeregisterNotary is a paid mutator transaction binding the contract method 0x58377bd1.
//
// Solidity: function deregisterNotary() returns()
func (_SMC *SMCSession) DeregisterNotary() (*types.Transaction, error) {
	return _SMC.Contract.DeregisterNotary(&_SMC.TransactOpts)
}

// DeregisterNotary is a paid mutator transaction binding the contract method 0x58377bd1.
//
// Solidity: function deregisterNotary() returns()
func (_SMC *SMCTransactorSession) DeregisterNotary() (*types.Transaction, error) {
	return _SMC.Contract.DeregisterNotary(&_SMC.TransactOpts)
}

// RegisterNotary is a paid mutator transaction binding the contract method 0x68e9513e.
//
// Solidity: function registerNotary() returns()
func (_SMC *SMCTransactor) RegisterNotary(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMC.contract.Transact(opts, "registerNotary")
}

// RegisterNotary is a paid mutator transaction binding the contract method 0x68e9513e.
//
// Solidity: function registerNotary() returns()
func (_SMC *SMCSession) RegisterNotary() (*types.Transaction, error) {
	return _SMC.Contract.RegisterNotary(&_SMC.TransactOpts)
}

// RegisterNotary is a paid mutator transaction binding the contract method 0x68e9513e.
//
// Solidity: function registerNotary() returns()
func (_SMC *SMCTransactorSession) RegisterNotary() (*types.Transaction, error) {
	return _SMC.Contract.RegisterNotary(&_SMC.TransactOpts)
}

// ReleaseNotary is a paid mutator transaction binding the contract method 0x9910851d.
//
// Solidity: function releaseNotary() returns()
func (_SMC *SMCTransactor) ReleaseNotary(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SMC.contract.Transact(opts, "releaseNotary")
}

// ReleaseNotary is a paid mutator transaction binding the contract method 0x9910851d.
//
// Solidity: function releaseNotary() returns()
func (_SMC *SMCSession) ReleaseNotary() (*types.Transaction, error) {
	return _SMC.Contract.ReleaseNotary(&_SMC.TransactOpts)
}

// ReleaseNotary is a paid mutator transaction binding the contract method 0x9910851d.
//
// Solidity: function releaseNotary() returns()
func (_SMC *SMCTransactorSession) ReleaseNotary() (*types.Transaction, error) {
	return _SMC.Contract.ReleaseNotary(&_SMC.TransactOpts)
}

// SubmitVote is a paid mutator transaction binding the contract method 0x4f33ffa0.
//
// Solidity: function submitVote(_shardId uint256, _period uint256, _index uint256, _chunkRoot bytes32) returns()
func (_SMC *SMCTransactor) SubmitVote(opts *bind.TransactOpts, _shardId *big.Int, _period *big.Int, _index *big.Int, _chunkRoot [32]byte) (*types.Transaction, error) {
	return _SMC.contract.Transact(opts, "submitVote", _shardId, _period, _index, _chunkRoot)
}

// SubmitVote is a paid mutator transaction binding the contract method 0x4f33ffa0.
//
// Solidity: function submitVote(_shardId uint256, _period uint256, _index uint256, _chunkRoot bytes32) returns()
func (_SMC *SMCSession) SubmitVote(_shardId *big.Int, _period *big.Int, _index *big.Int, _chunkRoot [32]byte) (*types.Transaction, error) {
	return _SMC.Contract.SubmitVote(&_SMC.TransactOpts, _shardId, _period, _index, _chunkRoot)
}

// SubmitVote is a paid mutator transaction binding the contract method 0x4f33ffa0.
//
// Solidity: function submitVote(_shardId uint256, _period uint256, _index uint256, _chunkRoot bytes32) returns()
func (_SMC *SMCTransactorSession) SubmitVote(_shardId *big.Int, _period *big.Int, _index *big.Int, _chunkRoot [32]byte) (*types.Transaction, error) {
	return _SMC.Contract.SubmitVote(&_SMC.TransactOpts, _shardId, _period, _index, _chunkRoot)
}

// SMCHeaderAddedIterator is returned from FilterHeaderAdded and is used to iterate over the raw logs and unpacked data for HeaderAdded events raised by the SMC contract.
type SMCHeaderAddedIterator struct {
	Event *SMCHeaderAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SMCHeaderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMCHeaderAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SMCHeaderAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SMCHeaderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMCHeaderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMCHeaderAdded represents a HeaderAdded event raised by the SMC contract.
type SMCHeaderAdded struct {
	ShardId         *big.Int
	ChunkRoot       [32]byte
	Period          *big.Int
	ProposerAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHeaderAdded is a free log retrieval operation binding the contract event 0x2d0a86178d2fd307b47be157a766e6bee19bc26161c32f9781ee0e818636f09c.
//
// Solidity: event HeaderAdded(shardId indexed uint256, chunkRoot bytes32, period uint256, proposerAddress address)
func (_SMC *SMCFilterer) FilterHeaderAdded(opts *bind.FilterOpts, shardId []*big.Int) (*SMCHeaderAddedIterator, error) {

	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _SMC.contract.FilterLogs(opts, "HeaderAdded", shardIdRule)
	if err != nil {
		return nil, err
	}
	return &SMCHeaderAddedIterator{contract: _SMC.contract, event: "HeaderAdded", logs: logs, sub: sub}, nil
}

// WatchHeaderAdded is a free log subscription operation binding the contract event 0x2d0a86178d2fd307b47be157a766e6bee19bc26161c32f9781ee0e818636f09c.
//
// Solidity: event HeaderAdded(shardId indexed uint256, chunkRoot bytes32, period uint256, proposerAddress address)
func (_SMC *SMCFilterer) WatchHeaderAdded(opts *bind.WatchOpts, sink chan<- *SMCHeaderAdded, shardId []*big.Int) (event.Subscription, error) {

	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _SMC.contract.WatchLogs(opts, "HeaderAdded", shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMCHeaderAdded)
				if err := _SMC.contract.UnpackLog(event, "HeaderAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SMCNotaryDeregisteredIterator is returned from FilterNotaryDeregistered and is used to iterate over the raw logs and unpacked data for NotaryDeregistered events raised by the SMC contract.
type SMCNotaryDeregisteredIterator struct {
	Event *SMCNotaryDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SMCNotaryDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMCNotaryDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SMCNotaryDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SMCNotaryDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMCNotaryDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMCNotaryDeregistered represents a NotaryDeregistered event raised by the SMC contract.
type SMCNotaryDeregistered struct {
	Notary             common.Address
	PoolIndex          *big.Int
	DeregisteredPeriod *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNotaryDeregistered is a free log retrieval operation binding the contract event 0x90e5afdc8fd31453dcf6e37154fa117ddf3b0324c96c65015563df9d5e4b5a75.
//
// Solidity: event NotaryDeregistered(notary address, poolIndex uint256, deregisteredPeriod uint256)
func (_SMC *SMCFilterer) FilterNotaryDeregistered(opts *bind.FilterOpts) (*SMCNotaryDeregisteredIterator, error) {

	logs, sub, err := _SMC.contract.FilterLogs(opts, "NotaryDeregistered")
	if err != nil {
		return nil, err
	}
	return &SMCNotaryDeregisteredIterator{contract: _SMC.contract, event: "NotaryDeregistered", logs: logs, sub: sub}, nil
}

// WatchNotaryDeregistered is a free log subscription operation binding the contract event 0x90e5afdc8fd31453dcf6e37154fa117ddf3b0324c96c65015563df9d5e4b5a75.
//
// Solidity: event NotaryDeregistered(notary address, poolIndex uint256, deregisteredPeriod uint256)
func (_SMC *SMCFilterer) WatchNotaryDeregistered(opts *bind.WatchOpts, sink chan<- *SMCNotaryDeregistered) (event.Subscription, error) {

	logs, sub, err := _SMC.contract.WatchLogs(opts, "NotaryDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMCNotaryDeregistered)
				if err := _SMC.contract.UnpackLog(event, "NotaryDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SMCNotaryRegisteredIterator is returned from FilterNotaryRegistered and is used to iterate over the raw logs and unpacked data for NotaryRegistered events raised by the SMC contract.
type SMCNotaryRegisteredIterator struct {
	Event *SMCNotaryRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SMCNotaryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMCNotaryRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SMCNotaryRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SMCNotaryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMCNotaryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMCNotaryRegistered represents a NotaryRegistered event raised by the SMC contract.
type SMCNotaryRegistered struct {
	Notary    common.Address
	PoolIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNotaryRegistered is a free log retrieval operation binding the contract event 0xa4fe15c53db34d35a5117acc26c27a2653dc68e2dadfc21ed211e38b7864d7a7.
//
// Solidity: event NotaryRegistered(notary address, poolIndex uint256)
func (_SMC *SMCFilterer) FilterNotaryRegistered(opts *bind.FilterOpts) (*SMCNotaryRegisteredIterator, error) {

	logs, sub, err := _SMC.contract.FilterLogs(opts, "NotaryRegistered")
	if err != nil {
		return nil, err
	}
	return &SMCNotaryRegisteredIterator{contract: _SMC.contract, event: "NotaryRegistered", logs: logs, sub: sub}, nil
}

// WatchNotaryRegistered is a free log subscription operation binding the contract event 0xa4fe15c53db34d35a5117acc26c27a2653dc68e2dadfc21ed211e38b7864d7a7.
//
// Solidity: event NotaryRegistered(notary address, poolIndex uint256)
func (_SMC *SMCFilterer) WatchNotaryRegistered(opts *bind.WatchOpts, sink chan<- *SMCNotaryRegistered) (event.Subscription, error) {

	logs, sub, err := _SMC.contract.WatchLogs(opts, "NotaryRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMCNotaryRegistered)
				if err := _SMC.contract.UnpackLog(event, "NotaryRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SMCNotaryReleasedIterator is returned from FilterNotaryReleased and is used to iterate over the raw logs and unpacked data for NotaryReleased events raised by the SMC contract.
type SMCNotaryReleasedIterator struct {
	Event *SMCNotaryReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SMCNotaryReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMCNotaryReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SMCNotaryReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SMCNotaryReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMCNotaryReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMCNotaryReleased represents a NotaryReleased event raised by the SMC contract.
type SMCNotaryReleased struct {
	Notary    common.Address
	PoolIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNotaryReleased is a free log retrieval operation binding the contract event 0xaee20171b64b7f3360a142659094ce929970d6963dcea8c34a9bf1ece8033680.
//
// Solidity: event NotaryReleased(notary address, poolIndex uint256)
func (_SMC *SMCFilterer) FilterNotaryReleased(opts *bind.FilterOpts) (*SMCNotaryReleasedIterator, error) {

	logs, sub, err := _SMC.contract.FilterLogs(opts, "NotaryReleased")
	if err != nil {
		return nil, err
	}
	return &SMCNotaryReleasedIterator{contract: _SMC.contract, event: "NotaryReleased", logs: logs, sub: sub}, nil
}

// WatchNotaryReleased is a free log subscription operation binding the contract event 0xaee20171b64b7f3360a142659094ce929970d6963dcea8c34a9bf1ece8033680.
//
// Solidity: event NotaryReleased(notary address, poolIndex uint256)
func (_SMC *SMCFilterer) WatchNotaryReleased(opts *bind.WatchOpts, sink chan<- *SMCNotaryReleased) (event.Subscription, error) {

	logs, sub, err := _SMC.contract.WatchLogs(opts, "NotaryReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMCNotaryReleased)
				if err := _SMC.contract.UnpackLog(event, "NotaryReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SMCVoteSubmittedIterator is returned from FilterVoteSubmitted and is used to iterate over the raw logs and unpacked data for VoteSubmitted events raised by the SMC contract.
type SMCVoteSubmittedIterator struct {
	Event *SMCVoteSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SMCVoteSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SMCVoteSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SMCVoteSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SMCVoteSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SMCVoteSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SMCVoteSubmitted represents a VoteSubmitted event raised by the SMC contract.
type SMCVoteSubmitted struct {
	ShardId       *big.Int
	ChunkRoot     [32]byte
	Period        *big.Int
	NotaryAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVoteSubmitted is a free log retrieval operation binding the contract event 0xc99370212b708f699fb6945a17eb34d0fc1ccd5b45d88f4d9682593a45d6e833.
//
// Solidity: event VoteSubmitted(shardId indexed uint256, chunkRoot bytes32, period uint256, notaryAddress address)
func (_SMC *SMCFilterer) FilterVoteSubmitted(opts *bind.FilterOpts, shardId []*big.Int) (*SMCVoteSubmittedIterator, error) {

	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _SMC.contract.FilterLogs(opts, "VoteSubmitted", shardIdRule)
	if err != nil {
		return nil, err
	}
	return &SMCVoteSubmittedIterator{contract: _SMC.contract, event: "VoteSubmitted", logs: logs, sub: sub}, nil
}

// WatchVoteSubmitted is a free log subscription operation binding the contract event 0xc99370212b708f699fb6945a17eb34d0fc1ccd5b45d88f4d9682593a45d6e833.
//
// Solidity: event VoteSubmitted(shardId indexed uint256, chunkRoot bytes32, period uint256, notaryAddress address)
func (_SMC *SMCFilterer) WatchVoteSubmitted(opts *bind.WatchOpts, sink chan<- *SMCVoteSubmitted, shardId []*big.Int) (event.Subscription, error) {

	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _SMC.contract.WatchLogs(opts, "VoteSubmitted", shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SMCVoteSubmitted)
				if err := _SMC.contract.UnpackLog(event, "VoteSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
