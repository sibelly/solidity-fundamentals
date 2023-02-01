const hre = require("hardhat")

async function main() {
  await hre.run("compile")
  const SimpleStorage = await hre.ethers.getContractFactory("SimpleStorage")
  const simpleStorage = await SimpleStorage.deploy()
  await simpleStorage.deployed();

  const transactionResponse = await simpleStorage.store(1)
  const transactionReceipt = await transactionResponse.wait()

  console.log(transactionReceipt.events[0].args.oldNumber.toString())
  console.log(transactionReceipt.events[0].args.newNumber.toString())
  console.log(transactionReceipt.events[0].args.addedNumber.toString())
  console.log(transactionReceipt.events[0].args.sender)
  console.log( "contract address => ", simpleStorage.address)

}

main().catch((error) => {
  console.error(error)
  process.exitCode = 1
})