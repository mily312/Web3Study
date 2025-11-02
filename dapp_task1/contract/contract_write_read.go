package contract

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ContractWriteAndRead(client *ethclient.Client) {
	//获取合约实例
	storeContract, err := NewStore(common.HexToAddress("0x473F977946cb2C904Fb81b0A44879059D0fEE2C9"), client)
	if err != nil {
		log.Fatal(err)
	}

	// 生成私钥
	privateKey, err := crypto.HexToECDSA("f49758a8e2440f468c3e539d52ef6ba8412569f58385f16b011b76e54e6ca8b6")
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
}
