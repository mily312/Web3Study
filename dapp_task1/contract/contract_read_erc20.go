package contract

import (
	"fmt"
	"log"
	"math"
	"math/big"

	token "go_ethereum_test/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ContractReadERC20(client *ethclient.Client) {
	//  fake Usdt (cUsdt) Address
	tokenAddress := common.HexToAddress("0x5E44bc3F480612d2aa9D3A847AEb96Ac348e1805")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xDFD3C6f4bc2edFb936c9f75ba5D424291C7fe450")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
