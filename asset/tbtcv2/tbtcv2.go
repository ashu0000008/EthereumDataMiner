package tbtcv2

const TBTCV2_ABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cachedChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cachedDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"tbtcv2 IERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"tbtcv2 IERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

const TBTCV2_ByteCode = "60c06040523480156200001157600080fd5b506040518060400160405280600781526020017f74425443207632000000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f74425443000000000000000000000000000000000000000000000000000000008152506200009e62000092620000f760201b60201c565b620000ff60201b60201c565b8160059080519060200190620000b69291906200026f565b508060069080519060200190620000cf9291906200026f565b504660808181525050620000e8620001c360201b60201c565b60a08181525050505062000520565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6005604051620001f79190620003dd565b60405180910390206040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525080519060200120463060405160200162000254959493929190620003f6565b60405160208183030381529060405280519060200120905090565b8280546200027d90620004bb565b90600052602060002090601f016020900481019282620002a15760008555620002ed565b82601f10620002bc57805160ff1916838001178555620002ed565b82800160010185558215620002ed579182015b82811115620002ec578251825591602001919060010190620002cf565b5b509050620002fc919062000300565b5090565b5b808211156200031b57600081600090555060010162000301565b5090565b6200032a8162000473565b82525050565b6200033b8162000487565b82525050565b600081546200035081620004bb565b6200035c818662000468565b945060018216600081146200037a57600181146200038c57620003c3565b60ff19831686528186019350620003c3565b620003978562000453565b60005b83811015620003bb578154818901526001820191506020810190506200039a565b838801955050505b50505092915050565b620003d781620004b1565b82525050565b6000620003eb828462000341565b915081905092915050565b600060a0820190506200040d600083018862000330565b6200041c602083018762000330565b6200042b604083018662000330565b6200043a6060830185620003cc565b6200044960808301846200031f565b9695505050505050565b60008190508160005260206000209050919050565b600081905092915050565b6000620004808262000491565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006002820490506001821680620004d457607f821691505b60208210811415620004eb57620004ea620004f1565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60805160a05161324d620005546000396000818161074f0152610c1e0152600081816107270152610a03015261324d6000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c8063771da5c5116100c3578063b4f94b2e1161007c578063b4f94b2e146103b7578063cae9ca51146103d5578063d505accf14610405578063dd62ed3e14610421578063f2fde38b14610451578063fc4e51f61461046d57610158565b8063771da5c5146102e157806379cc6790146102ff5780637ecebe001461031b5780638da5cb5b1461034b57806395d89b4114610369578063a9059cbb1461038757610158565b8063313ce56711610115578063313ce567146102335780633644e5151461025157806340c10f191461026f57806342966c681461028b57806370a08231146102a7578063715018a6146102d757610158565b806306fdde031461015d578063095ea7b31461017b5780631171bda9146101ab57806318160ddd146101c757806323b872dd146101e557806330adf81f14610215575b600080fd5b610165610489565b6040516101729190612796565b60405180910390f35b61019560048036038101906101909190611ef3565b610517565b6040516101a29190612667565b60405180910390f35b6101c560048036038101906101c09190611fbf565b61052e565b005b6101cf6105da565b6040516101dc9190612a18565b60405180910390f35b6101ff60048036038101906101fa9190611e06565b6105e0565b60405161020c9190612667565b60405180910390f35b61021d6106fa565b60405161022a9190612682565b60405180910390f35b61023b61071e565b6040516102489190612a33565b60405180910390f35b610259610723565b6040516102669190612682565b60405180910390f35b61028960048036038101906102849190611ef3565b610783565b005b6102a560048036038101906102a0919061208e565b610954565b005b6102c160048036038101906102bc9190611da1565b610961565b6040516102ce9190612a18565b60405180910390f35b6102df610979565b005b6102e9610a01565b6040516102f69190612a18565b60405180910390f35b61031960048036038101906103149190611ef3565b610a25565b005b61033560048036038101906103309190611da1565b610b36565b6040516103429190612a18565b60405180910390f35b610353610b4e565b6040516103609190612589565b60405180910390f35b610371610b77565b60405161037e9190612796565b60405180910390f35b6103a1600480360381019061039c9190611ef3565b610c05565b6040516103ae9190612667565b60405180910390f35b6103bf610c1c565b6040516103cc9190612682565b60405180910390f35b6103ef60048036038101906103ea9190611f2f565b610c40565b6040516103fc9190612667565b60405180910390f35b61041f600480360381019061041a9190611e55565b610cd7565b005b61043b60048036038101906104369190611dca565b610fbb565b6040516104489190612a18565b60405180910390f35b61046b60048036038101906104669190611da1565b610fe0565b005b6104876004803603810190610482919061200e565b6110d8565b005b6005805461049690612c56565b80601f01602080910402602001604051908101604052809291908181526020018280546104c290612c56565b801561050f5780601f106104e45761010080835404028352916020019161050f565b820191906000526020600020905b8154815290600101906020018083116104f257829003601f168201915b505050505081565b60006105243384846111ce565b6001905092915050565b610536611399565b73ffffffffffffffffffffffffffffffffffffffff16610554610b4e565b73ffffffffffffffffffffffffffffffffffffffff16146105aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a190612918565b60405180910390fd5b6105d582828573ffffffffffffffffffffffffffffffffffffffff166113a19092919063ffffffff16565b505050565b60045481565b600080600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146106e357828110156106cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c3906127f8565b60405180910390fd5b6106e2853385846106dd9190612b5d565b6111ce565b5b6106ee858585611427565b60019150509392505050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b601281565b60007f0000000000000000000000000000000000000000000000000000000000000000461415610775577f00000000000000000000000000000000000000000000000000000000000000009050610780565b61077d611718565b90505b90565b61078b611399565b73ffffffffffffffffffffffffffffffffffffffff166107a9610b4e565b73ffffffffffffffffffffffffffffffffffffffff16146107ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f690612918565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561086f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086690612998565b60405180910390fd5b61087b600083836117c0565b806004600082825461088d9190612b07565b9250508190555080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108e39190612b07565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516109489190612a18565b60405180910390a35050565b61095e33826117c5565b50565b60016020528060005260406000206000915090505481565b610981611399565b73ffffffffffffffffffffffffffffffffffffffff1661099f610b4e565b73ffffffffffffffffffffffffffffffffffffffff16146109f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ec90612918565b60405180910390fd5b6109ff600061192b565b565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610b275781811015610b10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b07906128b8565b60405180910390fd5b610b2683338484610b219190612b5d565b6111ce565b5b610b3183836117c5565b505050565b60036020528060005260406000206000915090505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60068054610b8490612c56565b80601f0160208091040260200160405190810160405280929190818152602001828054610bb090612c56565b8015610bfd5780601f10610bd257610100808354040283529160200191610bfd565b820191906000526020600020905b815481529060010190602001808311610be057829003601f168201915b505050505081565b6000610c12338484611427565b6001905092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000610c4c8484610517565b15610ccb578373ffffffffffffffffffffffffffffffffffffffff16638f4ffcb1338530866040518563ffffffff1660e01b8152600401610c90949392919061261b565b600060405180830381600087803b158015610caa57600080fd5b505af1158015610cbe573d6000803e3d6000fd5b5050505060019050610cd0565b600090505b9392505050565b42841015610d1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d1190612938565b60405180910390fd5b7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08160001c1115610d80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7790612978565b60405180910390fd5b601b8360ff161480610d955750601c8360ff16145b610dd4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dcb90612818565b60405180910390fd5b6000610dde610723565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9898989600360008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190610e5290612cb9565b919050558a604051602001610e6c9695949392919061269d565b60405160208183030381529060405280519060200120604051602001610e93929190612552565b604051602081830303815290604052805190602001209050600060018286868660405160008152602001604052604051610ed09493929190612751565b6020604051602081039080840390855afa158015610ef2573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610f6657508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610fa5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9c90612858565b60405180910390fd5b610fb08989896111ce565b505050505050505050565b6002602052816000526040600020602052806000526040600020600091509150505481565b610fe8611399565b73ffffffffffffffffffffffffffffffffffffffff16611006610b4e565b73ffffffffffffffffffffffffffffffffffffffff161461105c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161105390612918565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156110cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c3906127d8565b60405180910390fd5b6110d58161192b565b50565b6110e0611399565b73ffffffffffffffffffffffffffffffffffffffff166110fe610b4e565b73ffffffffffffffffffffffffffffffffffffffff1614611154576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161114b90612918565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff1663b88d4fde30868686866040518663ffffffff1660e01b81526004016111959594939291906125a4565b600060405180830381600087803b1580156111af57600080fd5b505af11580156111c3573d6000803e3d6000fd5b505050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561123e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123590612898565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156112ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112a5906127b8565b60405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405161138c9190612a18565b60405180910390a3505050565b600033905090565b6114228363a9059cbb60e01b84846040516024016113c09291906125f2565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506119ef565b505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161148e906129f8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611507576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114fe906129d8565b60405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611576576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161156d90612838565b60405180910390fd5b6115818383836117c0565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611608576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115ff906128d8565b60405180910390fd5b81816116149190612b5d565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546116a69190612b07565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161170a9190612a18565b60405180910390a350505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f600560405161174a919061253b565b60405180910390206040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152508051906020012046306040516020016117a59594939291906126fe565b60405160208183030381529060405280519060200120905090565b505050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561184c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611843906128f8565b60405180910390fd5b611858836000846117c0565b81816118649190612b5d565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008282546118b99190612b5d565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161191e9190612a18565b60405180910390a3505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000611a51826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611ab69092919063ffffffff16565b9050600081511115611ab15780806020019051810190611a719190611f96565b611ab0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611aa7906129b8565b60405180910390fd5b5b505050565b6060611ac58484600085611ace565b90509392505050565b606082471015611b13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b0a90612878565b60405180910390fd5b611b1c85611be2565b611b5b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b5290612958565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611b849190612524565b60006040518083038185875af1925050503d8060008114611bc1576040519150601f19603f3d011682016040523d82523d6000602084013e611bc6565b606091505b5091509150611bd6828286611bf5565b92505050949350505050565b600080823b905060008111915050919050565b60608315611c0557829050611c55565b600083511115611c185782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c4c9190612796565b60405180910390fd5b9392505050565b6000611c6f611c6a84612a73565b612a4e565b905082815260208101848484011115611c8757600080fd5b611c92848285612c14565b509392505050565b600081359050611ca981613176565b92915050565b600081519050611cbe8161318d565b92915050565b600081359050611cd3816131a4565b92915050565b60008083601f840112611ceb57600080fd5b8235905067ffffffffffffffff811115611d0457600080fd5b602083019150836001820283011115611d1c57600080fd5b9250929050565b600082601f830112611d3457600080fd5b8135611d44848260208601611c5c565b91505092915050565b600081359050611d5c816131bb565b92915050565b600081359050611d71816131d2565b92915050565b600081359050611d86816131e9565b92915050565b600081359050611d9b81613200565b92915050565b600060208284031215611db357600080fd5b6000611dc184828501611c9a565b91505092915050565b60008060408385031215611ddd57600080fd5b6000611deb85828601611c9a565b9250506020611dfc85828601611c9a565b9150509250929050565b600080600060608486031215611e1b57600080fd5b6000611e2986828701611c9a565b9350506020611e3a86828701611c9a565b9250506040611e4b86828701611d77565b9150509250925092565b600080600080600080600060e0888a031215611e7057600080fd5b6000611e7e8a828b01611c9a565b9750506020611e8f8a828b01611c9a565b9650506040611ea08a828b01611d77565b9550506060611eb18a828b01611d77565b9450506080611ec28a828b01611d8c565b93505060a0611ed38a828b01611cc4565b92505060c0611ee48a828b01611cc4565b91505092959891949750929550565b60008060408385031215611f0657600080fd5b6000611f1485828601611c9a565b9250506020611f2585828601611d77565b9150509250929050565b600080600060608486031215611f4457600080fd5b6000611f5286828701611c9a565b9350506020611f6386828701611d77565b925050604084013567ffffffffffffffff811115611f8057600080fd5b611f8c86828701611d23565b9150509250925092565b600060208284031215611fa857600080fd5b6000611fb684828501611caf565b91505092915050565b600080600060608486031215611fd457600080fd5b6000611fe286828701611d4d565b9350506020611ff386828701611c9a565b925050604061200486828701611d77565b9150509250925092565b60008060008060006080868803121561202657600080fd5b600061203488828901611d62565b955050602061204588828901611c9a565b945050604061205688828901611d77565b935050606086013567ffffffffffffffff81111561207357600080fd5b61207f88828901611cd9565b92509250509295509295909350565b6000602082840312156120a057600080fd5b60006120ae84828501611d77565b91505092915050565b6120c081612b91565b82525050565b6120cf81612ba3565b82525050565b6120de81612baf565b82525050565b6120f56120f082612baf565b612d02565b82525050565b60006121078385612acf565b9350612114838584612c14565b61211d83612d99565b840190509392505050565b600061213382612ab9565b61213d8185612acf565b935061214d818560208601612c23565b61215681612d99565b840191505092915050565b600061216c82612ab9565b6121768185612ae0565b9350612186818560208601612c23565b80840191505092915050565b6000815461219f81612c56565b6121a98186612ae0565b945060018216600081146121c457600181146121d557612208565b60ff19831686528186019350612208565b6121de85612aa4565b60005b83811015612200578154818901526001820191506020810190506121e1565b838801955050505b50505092915050565b600061221c82612ac4565b6122268185612aeb565b9350612236818560208601612c23565b61223f81612d99565b840191505092915050565b6000612257601b83612aeb565b915061226282612daa565b602082019050919050565b600061227a602683612aeb565b915061228582612dd3565b604082019050919050565b600061229d602183612aeb565b91506122a882612e22565b604082019050919050565b60006122c0600283612afc565b91506122cb82612e71565b600282019050919050565b60006122e3601b83612aeb565b91506122ee82612e9a565b602082019050919050565b6000612306601d83612aeb565b915061231182612ec3565b602082019050919050565b6000612329601183612aeb565b915061233482612eec565b602082019050919050565b600061234c602683612aeb565b915061235782612f15565b604082019050919050565b600061236f601d83612aeb565b915061237a82612f64565b602082019050919050565b6000612392601d83612aeb565b915061239d82612f8d565b602082019050919050565b60006123b5601f83612aeb565b91506123c082612fb6565b602082019050919050565b60006123d8601b83612aeb565b91506123e382612fdf565b602082019050919050565b60006123fb602083612aeb565b915061240682613008565b602082019050919050565b600061241e601283612aeb565b915061242982613031565b602082019050919050565b6000612441601d83612aeb565b915061244c8261305a565b602082019050919050565b6000612464601b83612aeb565b915061246f82613083565b602082019050919050565b6000612487601883612aeb565b9150612492826130ac565b602082019050919050565b60006124aa602a83612aeb565b91506124b5826130d5565b604082019050919050565b60006124cd601c83612aeb565b91506124d882613124565b602082019050919050565b60006124f0601e83612aeb565b91506124fb8261314d565b602082019050919050565b61250f81612bfd565b82525050565b61251e81612c07565b82525050565b60006125308284612161565b915081905092915050565b60006125478284612192565b915081905092915050565b600061255d826122b3565b915061256982856120e4565b60208201915061257982846120e4565b6020820191508190509392505050565b600060208201905061259e60008301846120b7565b92915050565b60006080820190506125b960008301886120b7565b6125c660208301876120b7565b6125d36040830186612506565b81810360608301526125e68184866120fb565b90509695505050505050565b600060408201905061260760008301856120b7565b6126146020830184612506565b9392505050565b600060808201905061263060008301876120b7565b61263d6020830186612506565b61264a60408301856120b7565b818103606083015261265c8184612128565b905095945050505050565b600060208201905061267c60008301846120c6565b92915050565b600060208201905061269760008301846120d5565b92915050565b600060c0820190506126b260008301896120d5565b6126bf60208301886120b7565b6126cc60408301876120b7565b6126d96060830186612506565b6126e66080830185612506565b6126f360a0830184612506565b979650505050505050565b600060a08201905061271360008301886120d5565b61272060208301876120d5565b61272d60408301866120d5565b61273a6060830185612506565b61274760808301846120b7565b9695505050505050565b600060808201905061276660008301876120d5565b6127736020830186612515565b61278060408301856120d5565b61278d60608301846120d5565b95945050505050565b600060208201905081810360008301526127b08184612211565b905092915050565b600060208201905081810360008301526127d18161224a565b9050919050565b600060208201905081810360008301526127f18161226d565b9050919050565b6000602082019050818103600083015261281181612290565b9050919050565b60006020820190508181036000830152612831816122d6565b9050919050565b60006020820190508181036000830152612851816122f9565b9050919050565b600060208201905081810360008301526128718161231c565b9050919050565b600060208201905081810360008301526128918161233f565b9050919050565b600060208201905081810360008301526128b181612362565b9050919050565b600060208201905081810360008301526128d181612385565b9050919050565b600060208201905081810360008301526128f1816123a8565b9050919050565b60006020820190508181036000830152612911816123cb565b9050919050565b60006020820190508181036000830152612931816123ee565b9050919050565b6000602082019050818103600083015261295181612411565b9050919050565b6000602082019050818103600083015261297181612434565b9050919050565b6000602082019050818103600083015261299181612457565b9050919050565b600060208201905081810360008301526129b18161247a565b9050919050565b600060208201905081810360008301526129d18161249d565b9050919050565b600060208201905081810360008301526129f1816124c0565b9050919050565b60006020820190508181036000830152612a11816124e3565b9050919050565b6000602082019050612a2d6000830184612506565b92915050565b6000602082019050612a486000830184612515565b92915050565b6000612a58612a69565b9050612a648282612c88565b919050565b6000604051905090565b600067ffffffffffffffff821115612a8e57612a8d612d6a565b5b612a9782612d99565b9050602081019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b6000612b1282612bfd565b9150612b1d83612bfd565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612b5257612b51612d0c565b5b828201905092915050565b6000612b6882612bfd565b9150612b7383612bfd565b925082821015612b8657612b85612d0c565b5b828203905092915050565b6000612b9c82612bdd565b9050919050565b60008115159050919050565b6000819050919050565b6000612bc482612b91565b9050919050565b6000612bd682612b91565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015612c41578082015181840152602081019050612c26565b83811115612c50576000848401525b50505050565b60006002820490506001821680612c6e57607f821691505b60208210811415612c8257612c81612d3b565b5b50919050565b612c9182612d99565b810181811067ffffffffffffffff82111715612cb057612caf612d6a565b5b80604052505050565b6000612cc482612bfd565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415612cf757612cf6612d0c565b5b600182019050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f417070726f766520746f20746865207a65726f20616464726573730000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f5472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6360008201527f6500000000000000000000000000000000000000000000000000000000000000602082015250565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b7f496e76616c6964207369676e6174757265202776272076616c75650000000000600082015250565b7f5472616e7366657220746f2074686520746f6b656e2061646472657373000000600082015250565b7f496e76616c6964207369676e6174757265000000000000000000000000000000600082015250565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f417070726f76652066726f6d20746865207a65726f2061646472657373000000600082015250565b7f4275726e20616d6f756e74206578636565647320616c6c6f77616e6365000000600082015250565b7f5472616e7366657220616d6f756e7420657863656564732062616c616e636500600082015250565b7f4275726e20616d6f756e7420657863656564732062616c616e63650000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f5065726d697373696f6e20657870697265640000000000000000000000000000600082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f496e76616c6964207369676e6174757265202773272076616c75650000000000600082015250565b7f4d696e7420746f20746865207a65726f20616464726573730000000000000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5472616e7366657220746f20746865207a65726f206164647265737300000000600082015250565b7f5472616e736665722066726f6d20746865207a65726f20616464726573730000600082015250565b61317f81612b91565b811461318a57600080fd5b50565b61319681612ba3565b81146131a157600080fd5b50565b6131ad81612baf565b81146131b857600080fd5b50565b6131c481612bb9565b81146131cf57600080fd5b50565b6131db81612bcb565b81146131e657600080fd5b50565b6131f281612bfd565b81146131fd57600080fd5b50565b61320981612c07565b811461321457600080fd5b5056fea2646970667358221220b66e3ef21ffde9263f6eb3247e39c2adf91afb2daa59dd552c43bf8eace893db64736f6c63430008040033"