require("@nomicfoundation/hardhat-toolbox");
require('dotenv').config()

// Replace this private key with your Goerli account private key
// To export your private key from Metamask, open Metamask and
// go to Account Details > Export Private Key
// Beware: NEVER put real Ether into testing accounts

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.17",
  networks: {
    goerli: {
      url:`https://eth-goerli.alchemyapi.io/v2/${process.env.GOERLI_API_KEY}`,
      accounts: [process.env.GOERLI_PRIVATE_KEY]
    }
  },
  etherscan: {
    apiKey: `${process.env.ETHERSCAN_API_KEY}`
  }
};
