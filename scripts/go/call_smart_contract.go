package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	url := os.Getenv("GOERLI_ALCHEMY_URL")
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Get the balance of an account
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Account balance:", balance) // 25893180161173005034

	// Get the latest known block
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest block:", block.Number().Uint64())

	// // Load the contract's ABI
	// // ContractABI is the ABI for the smart contract
	// var ContractABI = []byte(`[ABI encoded as a string]`)
	// contractAbi := bind.NewBoundContract(common.HexToAddress("0x1234567890abcdef"), ContractABI, client, client, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Call the contract's function
	// result, err := contractAbi.FunctionName(context.Background(), big.NewInt(123), "Hello, world")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	// Instantiate the contract and display its name
	address := common.HexToAddress("0x85Fe28EBe587C503055471813d146b365C5f146F")
	simpleStorage, err := NewMain(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	// Access simpleStorage properties
	name, err := simpleStorage.FavoriteNumber(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", name)

}
