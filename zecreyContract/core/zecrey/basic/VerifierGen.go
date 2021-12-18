// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basic

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VerifierABI is the input ABI used to generate the binding from.
const VerifierABI = "[{\"inputs\":[],\"name\":\"ScalarField\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"in_proof\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"proof_inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"num_proofs\",\"type\":\"uint256\"}],\"name\":\"verifyBatchProofs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"in_proof\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"proof_inputs\",\"type\":\"uint256[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierBin is the compiled bytecode used for deploying new contracts.
var VerifierBin = "0x608060405234801561001057600080fd5b50611b90806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063721ea4ac14610046578063c50e82631461017d578063daad1e63146102a2575b600080fd5b6101696004803603604081101561005c57600080fd5b810190602081018135600160201b81111561007657600080fd5b82018360208201111561008857600080fd5b803590602001918460208302840111600160201b831117156100a957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156100f857600080fd5b82018360208201111561010a57600080fd5b803590602001918460208302840111600160201b8311171561012b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506102bc945050505050565b604080519115158252519081900360200190f35b6101696004803603606081101561019357600080fd5b810190602081018135600160201b8111156101ad57600080fd5b8201836020820111156101bf57600080fd5b803590602001918460208302840111600160201b831117156101e057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561022f57600080fd5b82018360208201111561024157600080fd5b803590602001918460208302840111600160201b8311171561026257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505091359250610620915050565b6102aa610c6a565b60408051918252519081900360200190f35b6000806102c7610c8e565b905060006102d3610eb0565b90508351600160028351816102e457fe5b0403146102f057600080fd5b84516008146102fe57600080fd5b6103066118f6565b61030e611914565b600080600290508460008151811061032257fe5b60200260200101518360006004811061033757fe5b602002015284518590600190811061034b57fe5b60200260200101518360016004811061036057fe5b602002015260005b885181101561042e57855160018301928791811061038257fe5b60200260200101518560006003811061039757fe5b602002015285516001830192879181106103ad57fe5b6020026020010151856001600381106103c257fe5b602002015288518990829081106103d557fe5b6020026020010151856002600381106103ea57fe5b602002015260606040850160808760076107d05a03fa92508261040c57600080fd5b60608460c08660066107d05a03fa92508261042657600080fd5b600101610368565b5060006040518061030001604052808b60008151811061044a57fe5b602002602001015181526020018b60018151811061046457fe5b602002602001015181526020018b60028151811061047e57fe5b602002602001015181526020018b60038151811061049857fe5b602002602001015181526020018b6004815181106104b257fe5b602002602001015181526020018b6005815181106104cc57fe5b60200260200101518152602001886000600e81106104e657fe5b6020020151815260200161050a896001600e811061050057fe5b60200201516110a7565b815260408981015160208301526060808b0151918301919091526080808b01519183019190915260a0808b01519183019190915286519082015260c001610552866001610500565b815260c0890151602082015260e089015160408201526101008901516060820152608001886009602002015181526020018b60068151811061059057fe5b602002602001015181526020016105ba8c6007815181106105ad57fe5b60200260200101516110a7565b81526101408901516020820152610160890151604082015261018089015160608201526101a089015160809091015290506105f3611932565b6020816103008460086107d05a03fa93508361060e57600080fd5b516001149a9950505050505050505050565b6000816001141561063c5761063584846102bc565b9050610c63565b6000610646610c8e565b90506000610652610eb0565b9050836008028651146106965760405162461bcd60e51b8152600401808060200182810382526021815260200180611a0f6021913960400191505060405180910390fd5b838551816106a057fe5b06156106dd5760405162461bcd60e51b8152600401808060200182810382526021815260200180611a306021913960400191505060405180910390fd5b838551816106e757fe5b04600160028351816106f557fe5b0403146107335760405162461bcd60e51b815260040180806020018281038252602e815260200180611a51602e913960400191505060405180910390fd5b6060806107418888886110d1565b909250905060006107538585846116fc565b905060008760060260120167ffffffffffffffff8111801561077457600080fd5b5060405190808252806020026020018201604052801561079e578160200160208202803683370190505b50905060005b888110156108fd578481600202815181106107bb57fe5b60200260200101518282600602815181106107d257fe5b6020026020010181815250508481600202600101815181106107f057fe5b602002602001015182826006026001018151811061080a57fe5b6020026020010181815250508a816008026002018151811061082857fe5b602002602001015182826006026002018151811061084257fe5b6020026020010181815250508a816008026003018151811061086057fe5b602002602001015182826006026003018151811061087a57fe5b6020026020010181815250508a816008026004018151811061089857fe5b60200260200101518282600602600401815181106108b257fe5b6020026020010181815250508a81600802600501815181106108d057fe5b60200260200101518282600602600501815181106108ea57fe5b60209081029190910101526001016107a4565b5081518151829060068b0290811061091157fe5b6020908102919091010152610927826001610500565b81896006026001018151811061093957fe5b6020908102919091010152856002602002015181896006026002018151811061095e57fe5b6020908102919091010152856003602002015181896006026003018151811061098357fe5b602090810291909101015285600460200201518189600602600401815181106109a857fe5b602090810291909101015285600560200201518189600602600501815181106109cd57fe5b602090810291909101015281600260200201518189600602600601815181106109f257fe5b6020908102919091010152610a08826003610500565b818960060260070181518110610a1a57fe5b60209081029190910101528560066020020151818960060260080181518110610a3f57fe5b60209081029190910101528560076020020151818960060260090181518110610a6457fe5b602090810291909101015285600860200201518189600602600a0181518110610a8957fe5b602090810291909101015285600960200201518189600602600b0181518110610aae57fe5b602002602001018181525050838860020281518110610ac957fe5b60200260200101518189600602600c0181518110610ae357fe5b602002602001018181525050610b048489600202600101815181106105ad57fe5b8189600602600d0181518110610b1657fe5b602090810291909101015285600a60200201518189600602600e0181518110610b3b57fe5b602090810291909101015285600b60200201518189600602600f0181518110610b6057fe5b602090810291909101015285600c6020020151818960060260100181518110610b8557fe5b602090810291909101015285600d6020020151818960060260110181518110610baa57fe5b602002602001018181525050600081516020029050610bc7611932565b60c0820615610c075760405162461bcd60e51b815260040180806020018281038252602d815260200180611a7f602d913960400191505060405180910390fd5b602081836020860160086107d05a03fa985088610c555760405162461bcd60e51b81526004018080602001828103825260218152602001806119b16021913960400191505060405180910390fd5b516001149750505050505050505b9392505050565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190565b610c96611950565b7f206b76381e74bcad9bc7148b97554ad7516c164eab16e5967ca9a6bb9d8bf3cd81527f0d8497e9eba647f90aa20a17c331b646da41ae598026ffc49f6c8f8a17d1b53160208201527f19cbe45ba94e3550c761b231a7914944b1ffc2068d05b0bba8c73d2a7a0c914560408201527f11d7ef0ed898667f714fb5e2fbc63c7b570983ff12163035223043c47c1e9c5560608201527f0a77d1657fcbc52fd0a8311ea0919fdd775456119d21d672f63f21024fa03ded60808201527f08a685ed4c1554febfab8d9a2716bd526419032a65bf72e05e267d0996ee7d8e60a08201527f16a3bfe0860a127df4f504a176b957224c1458878d4f3e18aaacb91033a4425560c08201527f0b3c125198969e6dac79cb26aaed729a3bb451ce91503f66f54584e526ab51d760e08201527f1bbd24d63eaf4ceaf4b99eb23ef1cc635b657f541035f8d306bfdfb8f56ed2bc6101008201527f13feb9efda25a75ed9e7dc293f3a2c85cfd4d688565dc6c17e457d839f2981ce6101208201527f22ab3239ab21d30a06017b9c42c1d4ece66ae371d35617142f14e211d6bb224d6101408201527f09b05f1c3e1be3dc0374a5b5cfcef4b939077e12622bee4a134d9edcbfb7fd766101608201527f105246e2c6b4f2b11318bedc32b9464ad6681e27060decd070e468b7599b76406101808201527f2566e3d780269af6df31556f574dd3226872b834f032cec581b0912b374f8e2f6101a082015290565b604080516008808252610120820190925260609160208201610100803683370190505090507f20d320d02d3d02b3e8b4812645af41c1a76ee345b426c2bd9dd1f7c69b06cc9381600081518110610f0357fe5b6020026020010181815250507e935a190069620a4d67380d262a42daa0743e9c01da1d9f0775d697a78be25a81600181518110610f3c57fe5b6020026020010181815250507f127d72afd47dbbdbe1028dc1ba96dcf0bbb2eee3db501c7560d7446ba9d7ea4e81600281518110610f7657fe5b6020026020010181815250507f25e0a415c871cb1b8ceaa2359ca2e576993495865fcabea2c66f7b1781da61dd81600381518110610fb057fe5b6020026020010181815250507f10e7545ab241a7504b1b2ed3d50a913ca1cba9d97c2cffea9d7a01fcae98b31681600481518110610fea57fe5b6020026020010181815250507f1d32c35cfbf516bce383422cd9bc1a870c2ae5a58e05c7acec81ae2b7cef3aae8160058151811061102457fe5b6020026020010181815250507f099dd37da6c643bf297bc533ba7b7f4a9ca7783aeda988843673ae18cf8e98428160068151811061105e57fe5b6020026020010181815250507f037560f957178600c52f5507debfe8b66ea4ab5c630b2047253cfc2c0d5251198160078151811061109857fe5b60200260200101818152505090565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4790819006900390565b60608060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001905060008486518161110557fe5b04905060008567ffffffffffffffff8111801561112157600080fd5b5060405190808252806020026020018201604052801561114b578160200160208202803683370190505b5090508160010167ffffffffffffffff8111801561116857600080fd5b50604051908082528060200260200182016040528015611192578160200160208202803683370190505b50935060005b8681101561133557806111c45760018282815181106111b357fe5b6020026020010181815250506111eb565b834382900340816111d157fe5b068282815181106111de57fe5b6020026020010181815250505b8181815181106111f757fe5b602002602001015160001415611254576040805162461bcd60e51b815260206004820152601a60248201527f456e74726f70792073686f756c64206e6f74206265207a65726f000000000000604482015290519081900360640190fd5b838061125c57fe5b848061126457fe5b83838151811061127057fe5b60200260200101516001098660008151811061128857fe5b6020026020010151088560008151811061129e57fe5b60200260200101818152505060005b8381101561132c5784806112bd57fe5b85806112c557fe5b8a8387860201815181106112d557fe5b60200260200101518585815181106112e957fe5b60200260200101510987836001018151811061130157fe5b60200260200101510886826001018151811061131957fe5b60209081029190910101526001016112ad565b50600101611198565b5061133e6118f6565b60008760020260020167ffffffffffffffff8111801561135d57600080fd5b50604051908082528060200260200182016040528015611387578160200160208202803683370190505b5096508960008151811061139757fe5b6020026020010151876000815181106113ac57fe5b602002602001018181525050896001815181106113c557fe5b6020026020010151876001815181106113da57fe5b602090810291909101015260015b88811015611521578a81600802815181106113ff57fe5b60200260200101518360006003811061141457fe5b60200201528a518b906001600884020190811061142d57fe5b60200260200101518360016003811061144257fe5b6020020152835184908290811061145557fe5b60200260200101518360026003811061146a57fe5b602002015260408360608160076107d05a03fa9150826000602002015188826002028151811061149657fe5b602090810291909101015282600160200201518882600202600101815181106114bb57fe5b60200260200101818152505081611519576040805162461bcd60e51b815260206004820152601b60248201527f4661696c656420746f2063616c6c206120707265636f6d70696c650000000000604482015290519081900360640190fd5b6001016113e8565b5061152a611914565b8a60068151811061153757fe5b60200260200101518160006004811061154c57fe5b60200201528a518b90600790811061156057fe5b60200260200101518160016004811061157557fe5b602002015260015b898110156116a9578b816008026006018151811061159757fe5b6020026020010151846000600381106115ac57fe5b60200201528b518c90600760088402019081106115c557fe5b6020026020010151846001600381106115da57fe5b602002015284518590829081106115ed57fe5b60200260200101518460026003811061160257fe5b6020020152604082810160608660076107d05a03fa9250826116555760405162461bcd60e51b815260040180806020018281038252603d8152602001806119d2603d913960400191505060405180910390fd5b60408260808460066107d05a03fa9250826116a15760405162461bcd60e51b8152600401808060200182810382526037815260200180611aac6037913960400191505060405180910390fd5b60010161157d565b5080518851899060028c029081106116bd57fe5b60209081029190910101528060016020020151888a600202600101815181106116e257fe5b602002602001018181525050505050505050935093915050565b611704611914565b61170c611914565b6117146118f6565b6000805b85518110156118425786816002028151811061173057fe5b60200260200101518360006003811061174557fe5b6020020152865187906001600284020190811061175e57fe5b60200260200101518360016003811061177357fe5b6020020152855186908290811061178657fe5b60200260200101518360026003811061179b57fe5b6020020152604084810160608560076107d05a03fa9150816117ee5760405162461bcd60e51b8152600401808060200182810382526047815260200180611b146047913960600191505060405180910390fd5b60408460808660066107d05a03fa91508161183a5760405162461bcd60e51b81526004018080602001828103825260418152602001806119706041913960600191505060405180910390fd5b600101611718565b50825160408501526020830151606085015261185c6118f6565b87518152602080890151908201528551869060009061187757fe5b60200260200101518160026003811061188c57fe5b602002015260408160608160076107d05a03fa9150816118dd5760405162461bcd60e51b8152600401808060200182810382526031815260200180611ae36031913960400191505060405180910390fd5b8051855260209081015190850152509195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b604051806101c00160405280600e90602082028036833750919291505056fe4661696c656420746f2063616c6c206120707265636f6d70696c6520666f72204731206164646974696f6e20666f7220696e70757420616363756d756c61746f724661696c656420746f2063616c6c2070616972696e67732066756e6374696f6e734661696c656420746f2063616c6c206120707265636f6d70696c6520666f72204731206d756c7469706c69636174696f6e20666f722050726f6f662043496e76616c69642070726f6f6673206c656e67746820666f722061206261746368496e76616c696420696e70757473206c656e67746820666f7220612062617463684d69736d61746368696e67206e756d626572206f6620696e7075747320666f7220766572696679696e67206b6579496e70757473206c656e6774682073686f756c64206265206d756c7469706c65206f66203139322062797465734661696c656420746f2063616c6c206120707265636f6d70696c6520666f72204731206164646974696f6e20666f722050726f6f6620434661696c656420746f2063616c6c206120707265636f6d70696c6520666f72204731206d756c7469706c69636174696f6e4661696c656420746f2063616c6c206120707265636f6d70696c6520666f72204731206d756c7469706c69636174696f6e20666f7220696e70757420616363756d756c61746f72a26469706673582212207ada030cc0561a5e42bcc0d589170034938ca5653fba533df81bd948af60122664736f6c63430007060033"

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// ScalarField is a free data retrieval call binding the contract method 0xdaad1e63.
//
// Solidity: function ScalarField() pure returns(uint256)
func (_Verifier *VerifierCaller) ScalarField(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "ScalarField")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScalarField is a free data retrieval call binding the contract method 0xdaad1e63.
//
// Solidity: function ScalarField() pure returns(uint256)
func (_Verifier *VerifierSession) ScalarField() (*big.Int, error) {
	return _Verifier.Contract.ScalarField(&_Verifier.CallOpts)
}

// ScalarField is a free data retrieval call binding the contract method 0xdaad1e63.
//
// Solidity: function ScalarField() pure returns(uint256)
func (_Verifier *VerifierCallerSession) ScalarField() (*big.Int, error) {
	return _Verifier.Contract.ScalarField(&_Verifier.CallOpts)
}

// VerifyBatchProofs is a free data retrieval call binding the contract method 0xc50e8263.
//
// Solidity: function verifyBatchProofs(uint256[] in_proof, uint256[] proof_inputs, uint256 num_proofs) view returns(bool success)
func (_Verifier *VerifierCaller) VerifyBatchProofs(opts *bind.CallOpts, in_proof []*big.Int, proof_inputs []*big.Int, num_proofs *big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyBatchProofs", in_proof, proof_inputs, num_proofs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBatchProofs is a free data retrieval call binding the contract method 0xc50e8263.
//
// Solidity: function verifyBatchProofs(uint256[] in_proof, uint256[] proof_inputs, uint256 num_proofs) view returns(bool success)
func (_Verifier *VerifierSession) VerifyBatchProofs(in_proof []*big.Int, proof_inputs []*big.Int, num_proofs *big.Int) (bool, error) {
	return _Verifier.Contract.VerifyBatchProofs(&_Verifier.CallOpts, in_proof, proof_inputs, num_proofs)
}

// VerifyBatchProofs is a free data retrieval call binding the contract method 0xc50e8263.
//
// Solidity: function verifyBatchProofs(uint256[] in_proof, uint256[] proof_inputs, uint256 num_proofs) view returns(bool success)
func (_Verifier *VerifierCallerSession) VerifyBatchProofs(in_proof []*big.Int, proof_inputs []*big.Int, num_proofs *big.Int) (bool, error) {
	return _Verifier.Contract.VerifyBatchProofs(&_Verifier.CallOpts, in_proof, proof_inputs, num_proofs)
}

// VerifyProof is a free data retrieval call binding the contract method 0x721ea4ac.
//
// Solidity: function verifyProof(uint256[] in_proof, uint256[] proof_inputs) view returns(bool)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, in_proof []*big.Int, proof_inputs []*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", in_proof, proof_inputs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x721ea4ac.
//
// Solidity: function verifyProof(uint256[] in_proof, uint256[] proof_inputs) view returns(bool)
func (_Verifier *VerifierSession) VerifyProof(in_proof []*big.Int, proof_inputs []*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, in_proof, proof_inputs)
}

// VerifyProof is a free data retrieval call binding the contract method 0x721ea4ac.
//
// Solidity: function verifyProof(uint256[] in_proof, uint256[] proof_inputs) view returns(bool)
func (_Verifier *VerifierCallerSession) VerifyProof(in_proof []*big.Int, proof_inputs []*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, in_proof, proof_inputs)
}