package main

func main2() {
	// client, err := ethclient.Dial("https://mainnet.infura.io")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Address of the deployed multisig contract
	// contractAddress := common.HexToAddress("0xE11BA2b4D45Eaed5996Cd0823791E0C93114882d")

	// // Create a new instance of the contract
	// contract, err := NewMultiSig(contractAddress, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	/*
		It is important to note that the NewMultiSig method is not a built-in method
		in go-ethereum library, it is generated by the library using the ABI of the
		contract and it should match the contract's name and its functions signature.
	*/

	// Call the 'addOwner' function of the contract
	// addOwnerTx, err := contract.AddOwner(
	// 	auth,
	// 	common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Wait for the transaction to be mined
	// addOwnerReceipt, err := bind.WaitMined(context.Background(), client, addOwnerTx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Confirm that the transaction was successful
	// if addOwnerReceipt.Status != 1 {
	// 	log.Println("Failed to add owner")
	// } else {
	// 	log.Println("Successfully added owner")
	// }

	// // Call the 'submitTransaction' function of the contract
	// submitTransactionTx, err := contract.SubmitTransaction(
	// 	auth,
	// 	common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
	// 	big.NewInt(1000000000000000000),
	// 	[]byte("transfer"),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Wait for the transaction to be mined
	// submitTransactionReceipt, err := bind.WaitMined(context.Background(), client, submitTransactionTx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Confirm that the transaction was successful
	// if submitTransactionReceipt.Status != 1 {
	// 	log.Println("Failed to submit transaction")
	// } else {
	// 	log.Println("Successfully submitted transaction")
	// }
}
