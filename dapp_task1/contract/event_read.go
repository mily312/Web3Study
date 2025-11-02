package contract

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EventRead(client *ethclient.Client) {
	contractAddress := common.HexToAddress("0x473F977946cb2C904Fb81b0A44879059D0fEE2C9")
	// 日志过滤条件
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(9543054),
		ToBlock:   big.NewInt(9543063),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	// 从链上查询数据
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 返回的所有日志将是ABI编码，为了解码日志，我们需要导入我们智能合约的ABI
	contractAbi, err := abi.JSON(strings.NewReader(string(StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("signature topics=", hash.Hex())
}
