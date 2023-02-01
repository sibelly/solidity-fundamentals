package main

func deploy() {
	// client, err := ethclient.Dial("https://mainnet.infura.io")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Deploy the multisignature wallet contract
	// bytecode := "0x606060..." // bytecode for the multisignature wallet contract
	// abi := "0x606060..."      // ABI for the multisignature wallet contract
	// deployerKey := "0x..."    // private key of the deployer account
	// deployerAuth := bind.NewKeyedTransactor(deployerKey)
	// deployerAuth.GasPrice = big.NewInt(20000000000)
	// deployerAuth.GasLimit = uint64(4712388)
	// deployerAuth.Value = big.NewInt(0)

	// // deploy the contract
	// address, _, _, err := bind.DeployContract(deployerAuth, abi, bytecode, client)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Get the instance of the deployed contract
	// instance, err := NewMultiSigWallet(address, client)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Add owners to the multisignature wallet
	// owner1 := common.HexToAddress("0x...")
	// owner2 := common.HexToAddress("0x...")
	// owner3 := common.HexToAddress("0x...")
	// _, err = instance.AddOwner(owner1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// _, err = instance.AddOwner(owner2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// _, err = instance.AddOwner(owner3)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Multisignature wallet created at address:", address.Hex())
}
