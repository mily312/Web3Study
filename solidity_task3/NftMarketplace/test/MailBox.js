//引入必备的库
const { expect } = require("chai")
const { ethers } = require("hardhat")

describe("MailBox", function () {
    it("should get mailbox contract", async function () {
        //获取合约
        const mailBoxContract = await ethers.getContractFactory("MailBox")
    })

    it("should get numbers of letters in the box",async function () {
        // 部署合约
        const mailBoxContract = await ethers.getContractFactory("MailBox")
        const mailBox = await mailBoxContract.deploy();

        expect(await mailBox.totalLetters()).to.equal(0);
    })

    it("should add one when get new letter",async function () {
       
        const mailBoxContract = await ethers.getContractFactory("MailBox")
        const mailBox = await mailBoxContract.deploy();

        await mailBox.write("hello");
        expect(await mailBox.totalLetters()).to.equal(1);
    })

    it("should get all mail content",async function () {
       
        const mailBoxContract = await ethers.getContractFactory("MailBox")
        const mailBox = await mailBoxContract.deploy();

        await mailBox.write("hello");

        const allLetters = await mailBox.getAllLetters();
        expect(allLetters[0].message).to.equal("hello");
    })

    it("should get sender address",async function () {
       
        const mailBoxContract = await ethers.getContractFactory("MailBox")
        const mailBox = await mailBoxContract.deploy();

        await mailBox.write("hello");

        const allLetters = await mailBox.getAllLetters();
        expect(allLetters[0].sender).to.equal("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266");
    })
}
)