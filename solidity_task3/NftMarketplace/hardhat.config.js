require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks: {
    hardhat: {
    },
    sepolia: {
      url: "https://eth-sepolia.g.alchemy.com/v2/deZgaTL0_cXONkuTQ5RLO",
      accounts: ["7c8290252a2f83ab27f667adf7e91d6ddb0f74dbadf0a608f152a267e32c6163", 
            "f49758a8e2440f468c3e539d52ef6ba8412569f58385f16b011b76e54e6ca8b6"]
    },
    localhost: {
      url: "http://127.0.0.1:8545",
    }
  },
};
