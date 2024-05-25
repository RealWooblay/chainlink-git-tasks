const ethers = require("ethers");
const ghTasksAbi = require("../foundry/out/GHTasks.sol/GHTasks.json");

const ghTasksAddress = "0x217e1FbcE6d1605d53111bd756cDcCb125A26A28";

const explorerUrl = "https://sepolia.basescan.org/";

const orgId = 0;
const sourceId = 0; // Merge
const args = [
    "RealWooblay", 
    "chainlink-git-tasks", 
    "3",
    "Onchain commit", 
    "This was merged onchain!"
];
const gasLimit = 300000;

const request = async () => {
    const privateKey = process.env.CONTRIBUTOR_PRIVATE_KEY;
    if (!privateKey) {
        throw new Error(
            "private key not provided - check your environment variables"
        );
    }

    const rpcUrl = process.env.BASE_SEPOLIA_RPC_URL;

    if (!rpcUrl) {
        throw new Error(`rpcUrl not provided  - check your environment variables`);
    }

    const provider = new ethers.providers.JsonRpcProvider(rpcUrl);

    const wallet = new ethers.Wallet(privateKey);
    const signer = wallet.connect(provider);
   
    const ghTasks = new ethers.Contract(
        ghTasksAddress,
        ghTasksAbi.abi,
        signer
    );

    const transaction = await ghTasks.request(
        orgId,
        sourceId,
        args,
        gasLimit
    );

    console.log(
        `\n✅ Functions request sent! Transaction hash ${transaction.hash}. Waiting for a response...`
    );

    console.log(
        `See your request in the explorer ${explorerUrl}/tx/${transaction.hash}`
    );
}

request().catch((e) => {
    console.error(e);
    process.exit(1);
})