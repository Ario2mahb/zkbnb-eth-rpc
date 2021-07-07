package zecrey

import (
	"Zecrey-eth-rpc/_const"
	"Zecrey-eth-rpc/_rpc"
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"
)

const GovernanceAddr = "0xc879dFB2c3C1a6DCC2cB6EaEB4F395b821C9571a"

var (
	cli      *_rpc.ProviderClient
	authCli  *_rpc.AuthClient
	gasPrice *big.Int
)

func init() {
	cli, _ = _rpc.NewClient(_const.LocalNetwork)
	authCli, _ = _rpc.NewAuthClient(cli, _const.LocalSuperSk)
	gasPrice, _ = cli.SuggestGasPrice(context.Background())
}

func TestDeployGovernanceContract(t *testing.T) {
	elapse := time.Now()
	addr, txHash, err := DeployGovernanceContract(cli, authCli, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time.Since(elapse))
	fmt.Println("contract address:", addr)
	fmt.Println("deploy contract tx hash:", txHash)
}

func LoadGovernance() *Governance {
	instance, _ := LoadGovernanceInstance(cli, GovernanceAddr)
	return instance
}

func TestAddAsset(t *testing.T) {
	governance := LoadGovernance()
	txHash, err := AddAsset(cli, authCli, governance, ZecreyTokenAddr, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}

func TestUpdateRollupProvider(t *testing.T) {
	instance := LoadGovernance()
	txHash, err := UpdateRollupProvider(cli, authCli, instance, _const.LocalSuperAddress, true, gasPrice, _const.SuggestHighGasLimit)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txHash)
}

func TestGetSupportedAsset(t *testing.T) {
	instance := LoadGovernance()
	elapse := time.Now()
	assetAddr, err := GetSupportedAsset(instance, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(time.Since(elapse))
	fmt.Println(assetAddr)
}

func TestIsValidAsset(t *testing.T) {
	instance := LoadGovernance()
	isValid, err := IsValidAsset(instance, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(isValid)
}
