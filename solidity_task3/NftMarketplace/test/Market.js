const { expect } = require('chai');
const { ethers } = require('hardhat');

describe("Market", async function () {
    // 五个变量  三个合约，两个账户
    let usdt, nft, market, account1, account2;

    beforeEach(async function () {
        [account1, account2] = await ethers.getSigners();

        let usdtFactory = await ethers.getContractFactory("cUsdt");
        usdt = await usdtFactory.deploy(account1.address);

        let nftFactory = await ethers.getContractFactory("NFTM");
        nft = await nftFactory.deploy(account1.address);

        let marketFactory = await ethers.getContractFactory("Market");
        market = await marketFactory.deploy(usdt.target, nft.target);

        //给账户1和账户2分别初始化
        await usdt.mint(account1.address, "1000000000000000000");
        await usdt.connect(account1).approve(market.target, "1000000000000000000");
        await nft.safeMint(account2.address);
        await nft.safeMint(account2.address);
    });

    it("market erc20 should be usdt", async function () {
        expect(await market.erc20()).to.equal(usdt.target);
    })

    it("market erc721 should be nft", async function () {
        expect(await market.erc721()).to.equal(nft.target);
    })

    it("account2 should have 2 nfts", async function () {
        expect(await nft.balanceOf(account2.address)).to.equal(2);
    })

    it("account1 should have usdt", async function () {
        expect(await usdt.balanceOf(account1.address)).to.equal("1000000000000000000");
    })

    it("account2 can list 2 nfts on market", async function () {
        const price = "0x0000000000000000000000000000000000000000000000000001c6bf52634000";

        expect(await nft.connect(account2)['safeTransferFrom(address,address,uint256,bytes)']
            (account2.address, market.target, 0, price)).to.emit(market, "CreateOrder");

        expect(await nft.connect(account2)['safeTransferFrom(address,address,uint256,bytes)']
            (account2.address, market.target, 1, price)).to.emit(market, "CreateOrder");

        expect(await nft.balanceOf(account2.address)).to.equal(0);
        expect(await nft.balanceOf(market.target)).to.equal(2);

        expect(await market.isListed(0)).to.equal(true);
        expect(await market.isListed(1)).to.equal(true);
    });

}
)