const { SecretsManager } = require("@chainlink/functions-toolkit");
const ethers = require("ethers");
const ghTasksAbi = require("../foundry/out/GHTasks.sol/GHTasks.json");

const ghTasksAddress = "0x217e1FbcE6d1605d53111bd756cDcCb125A26A28";

const routerAddress = "0xf9B8fc078197181C841c296C876945aaa425B278";
const donId = "fun-base-sepolia-1";
const explorerUrl = "https://sepolia.basescan.org/";
const gatewayUrls = [
    "https://01.functions-gateway.testnet.chain.link/",
    "https://02.functions-gateway.testnet.chain.link/",
];

const secrets = { githubToken: process.env.OWNER_GITHUB_TOKEN };
const slotId = 0;
const expirationTimeMinutes = 5;

const request = async () => {
    const privateKey = process.env.OWNER_PRIVATE_KEY;
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

    const secretsManager = new SecretsManager({
        signer: signer,
        functionsRouterAddress: routerAddress,
        donId: donId,
    });
    await secretsManager.initialize();

    const encryptedSecretsObj = await secretsManager.encryptSecrets(secrets);

    console.log(
        `Upload encrypted secret to gateways ${gatewayUrls}. slotId ${slotId}. Expiration in minutes: ${expirationTimeMinutes}`
    );

    const uploadResult = await secretsManager.uploadEncryptedSecretsToDON({
        encryptedSecretsHexstring: encryptedSecretsObj.encryptedSecrets,
        gatewayUrls: gatewayUrls,
        slotId: slotId,
        minutesUntilExpiration: expirationTimeMinutes,
    });

    if (!uploadResult.success) {
        throw new Error(`Encrypted secrets not uploaded to ${gatewayUrls}`);
    }

    console.log(
        `\n✅ Secrets uploaded properly to gateways ${gatewayUrls}! Gateways response: `,
        uploadResult
    );

    const donHostedSecretsVersion = parseInt(uploadResult.version);

    const ghTasks = new ethers.Contract(
        ghTasksAddress,
        ghTasksAbi.abi,
        signer
    );

    const transaction = await ghTasks.setDonSecret(
        slotId,
        donHostedSecretsVersion
    );

    console.log(
        `\n✅ Set DON secrets sent! Transaction hash ${transaction.hash}. Waiting for a response...`
    );

    console.log(
        `See your request in the explorer ${explorerUrl}/tx/${transaction.hash}`
    );
}

request().catch((e) => {
    console.error(e);
    process.exit(1);
})