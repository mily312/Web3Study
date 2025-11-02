package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go_ethereum_test/contract"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func main() {
	client, _ := getConnect()

	// 查询账户信息
	// getAccountInfo(client)

	// 查询区块
	// getBlockInfo(client)

	// 查询交易
	// transactions(client)

	//创建新钱包
	// walletGenerate(client)

	// Eth转账
	// transferETH(client)

	// 代币转账
	// transferToken(client)

	// 查询ERC20代币智能合约
	// contract.ContractReadERC20(client)

	// 监听新区快
	// blockSubscribe(client)

	// 部署新合约
	// contract.DeployContract(client)

	// 加载智能合约
	// contract.ContractLoad(client)

	// 写入&查询智能合约
	// contract.ContractWriteAndRead(client)

	// 读取事件日志
	contract.EventRead(client)
}

func getConnect() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/deZgaTL0_cXONkuTQ5RLO")
	// client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/deZgaTL0_cXONkuTQ5RLO")

	// client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/deZgaTL0_cXONkuTQ5RLO")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

	return client, err
}

// 查询账户信息
func getAccountInfo(client *ethclient.Client) {
	account := common.HexToAddress("0xDFD3C6f4bc2edFb936c9f75ba5D424291C7fe450")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
}

// 查询区块
func getBlockInfo(client *ethclient.Client) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 9537032

	blockNumber := big.NewInt(9537032)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144
}

// 查询交易
func transactions(client *ethclient.Client) {

	blockNumber := big.NewInt(9537032)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		// fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		// fmt.Println(tx.Value().String())    // 10000000000000000
		// fmt.Println(tx.Gas())               // 105000
		// fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		// fmt.Println(tx.Nonce())             // 110644
		// fmt.Println(tx.Data())              // []
		// fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		//读取发送方的地址
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if from, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println(from.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}

		// 读取事务收据
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
	}

}

func walletGenerate(client *ethclient.Client) {
	// 构建私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	// 构建公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	// 公共地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) //0x419781b4Ba0dBee10292665eBcB47cE40b56E817

	// 和上面生成的地址是一样的
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) //0x419781b4Ba0dBee10292665eBcB47cE40b56E817
}

// 发送eth
func transferETH(client *ethclient.Client) {
	// 根据私钥 生成公钥 接着生成地址
	privateKey, err := crypto.HexToECDSA("f49758a8e2440f468c3e539d52ef6ba8412569f58385f16b011b76e54e6ca8b6")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获得帐户的随机数(nonce)。 每笔交易都需要一个nonce。
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(10000000000000000) // in wei (0.01 eth)
	gasLimit := uint64(21000)              // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xA28cBabDD41B9EdCD3399dA77658B425FdECd991")

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	//使用发件人的私钥对事务进行签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

}

func transferToken(client *ethclient.Client) {
	// 根据私钥 生成公钥 接着生成地址
	privateKey, err := crypto.HexToECDSA("f49758a8e2440f468c3e539d52ef6ba8412569f58385f16b011b76e54e6ca8b6")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获得帐户的随机数(nonce)。 每笔交易都需要一个nonce。
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xA28cBabDD41B9EdCD3399dA77658B425FdECd991")
	// erc20 合约地址
	tokenAddress := common.HexToAddress("0x5E44bc3F480612d2aa9D3A847AEb96Ac348e1805")

	transferFnSignature := []byte("transfer(address,uint256)")

	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("1000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	To:   &toAddress,
	// 	Data: data,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(gasLimit)

	gasLimit := uint64(45000)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

// 监听新区快
func blockSubscribe(client *ethclient.Client) {
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
