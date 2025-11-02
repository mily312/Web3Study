package contract

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
此初始化方法接收智能合约的地址，并返回可以开始与之交互的合约实例
*/
func ContractLoad(client *ethclient.Client) {
	//store合约部署地址
	address := common.HexToAddress("0x473F977946cb2C904Fb81b0A44879059D0fEE2C9")
	instance, err := NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
