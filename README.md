# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a script that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.js

npx hardhat compile
npx hardhat run scripts/deploy.js --network goerli
npx hardhat verify --network goerli DEPLOYED_CONTRACT_ADDRESS
```

```
go run scripts/call_smart_contract.go
```

## Generate the abi and bin smart contract code
- Install abigen tool
```
# Download go-ethereum, build and install the devtools (which includes abigen)
git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
make devtools

# Run abigen and print the version
abigen -version
abigen -help
```

- Install solc tool
```
# Requires node installed
npm install -g solc
```

- Build the .sol file
```
npx solc --abi contracts/SimpleStorage.sol -o build
npx solc --bin contracts/SimpleStorage.sol -o build

abigen --abi build/SimpleStorage.abi --bin build/SimpleStorage.bin --pkg main --out SimpleStorage.go
```
