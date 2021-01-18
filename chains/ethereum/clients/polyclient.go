package clients

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go-chains/chains/evm/binding/poly/eccd_abi"
	"go-chains/chains/evm/binding/poly/eccm_abi"
	"go-chains/chains/evm/binding/poly/eccmp_abi"
	"go-chains/chains/evm/binding/poly/erc20_abi"
	"go-chains/chains/evm/binding/poly/oep4_abi"
	"go-chains/chains/log"
	"math/big"
)

func (c *EvmClient) DeployEthChainDataContract() (common.Address, *eccd_abi.EthCrossChainData, error) {
	auth := c.MakeAuth()
	contractAddress, tx, contract, err := eccd_abi.DeployEthCrossChainData(auth,
		c.EClient)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployEthChainDataContract, err: %v", err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (c *EvmClient) DeployECCMContract(eccdAddress string) (common.Address, *eccm_abi.EthCrossChainManager, error) {
	auth := c.MakeAuth()
	address := common.HexToAddress(eccdAddress)
	contractAddress, tx, contract, err := eccm_abi.DeployEthCrossChainManager(auth,
		c.EClient, address)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployECCMContract, err: %v", err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (c *EvmClient) DeployECCMPContract(eccmAddress string) (common.Address, *eccmp_abi.EthCrossChainManagerProxy, error) {
	auth := c.MakeAuth()
	address := common.HexToAddress(eccmAddress)
	contractAddress, tx, contract, err := eccmp_abi.DeployEthCrossChainManagerProxy(auth,
		c.EClient, address)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployECCMPContract, err: %v", err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (c *EvmClient) DeployERC20Template() (common.Address, *erc20_abi.ERC20Template, error) {
	auth := c.MakeAuth()
	contractAddress, tx, contract, err := erc20_abi.DeployERC20Template(auth,
		c.EClient)
	if err != nil {
		log.Fatal(err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (c *EvmClient) DeployOEP4(lockProxy string) (common.Address, *oep4_abi.OEP4Template, error) {
	auth := c.MakeAuth()
	lockProxyAddr := common.HexToAddress(lockProxy)
	contractAddress, tx, contract, err := oep4_abi.DeployOEP4Template(auth,
		c.EClient, lockProxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	c.WaitTransactionConfirm(tx.Hash())

	auth = c.MakeAuth()
	tx, err = contract.DeletageToProxy(auth, lockProxyAddr, big.NewInt(1e13))
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to DeletageToProxy: %v", err)
	}
	c.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

//func (c *EvmClient) GetAccInfo() (string, error) {
//	h := c.GetBlockNumber()
//	val, err := c.EClient.BalanceAt(context.Background(), c.Address(), big.NewInt(int64(h)))
//	if err != nil {
//		return "", err
//	}
//	ethInfo := fmt.Sprintf("eth: %d", val.Uint64())
//
//	//ontx, err := ontx.NewONTX(common.HexToAddress(ethInvoker.TConfiguration.EthOntx), ethInvoker.ETHUtil.ethclient)
//	//if err != nil {
//	//	return "", err
//	//}
//	//val, err = ontx.BalanceOf(nil, ethInvoker.EthTestSigner.Address)
//	//if err != nil {
//	//	return "", err
//	//}
//	//ontInfo := fmt.Sprintf("ontx: %d", val.Uint64())
//	//
//	//ongx, err := ongx_api.NewONGX(ethComm.HexToAddress(ethInvoker.TConfiguration.EthOngx), ethInvoker.ETHUtil.ethclient)
//	//if err != nil {
//	//	return "", err
//	//}
//	//val, err = ongx.BalanceOf(nil, ethInvoker.EthTestSigner.Address)
//	//if err != nil {
//	//	return "", err
//	//}
//	//ongInfo := fmt.Sprintf("ongx: %d", val.Uint64())
//	//
//	//oep4x, err := oep4_api.NewOEP4Template(ethComm.HexToAddress(ethInvoker.TConfiguration.EthOep4), ethInvoker.ETHUtil.ethclient)
//	//if err != nil {
//	//	return "", err
//	//}
//	//val, err = oep4x.BalanceOf(nil, ethInvoker.EthTestSigner.Address)
//	//if err != nil {
//	//	return "", err
//	//}
//	//oep4Info := fmt.Sprintf("oep4x: %d", val.Uint64())
//	//
//	//erc20, err := erc20_api.NewERC20(ethComm.HexToAddress(ethInvoker.TConfiguration.EthErc20), ethInvoker.ETHUtil.ethclient)
//	//if err != nil {
//	//	return "", err
//	//}
//	//val, err = erc20.BalanceOf(nil, ethInvoker.EthTestSigner.Address)
//	//if err != nil {
//	//	return "", err
//	//}
//	//erc20Info := fmt.Sprintf("erc20: %d", val.Uint64())
//	//
//	//btcx, err := btcx.NewBTCX(ethComm.HexToAddress(ethInvoker.TConfiguration.BtceContractAddress), ethInvoker.ETHUtil.ethclient)
//	//if err != nil {
//	//	return "", err
//	//}
//	//val, err = btcx.BalanceOf(nil, c.Address)
//	//if err != nil {
//	//	return "", err
//	//}
//	//btcxInfo := fmt.Sprintf("btcx: %d", val.Uint64())
//	//
//	//return fmt.Sprintf("ETHEREUM: acc: %s, asset: [ %s, %s, %s, %s, %s, %s ]",
//	//	c.Address.String(), ethInfo, ontInfo, ongInfo, oep4Info, erc20Info, btcxInfo), nil
//}
