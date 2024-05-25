const fs = require("fs");
const path = require("path");
const {
    SecretsManager,
} = require("@chainlink/functions-toolkit");
const ethers = require("ethers");

const functionsConsumerAbi = require("./functionsConsumerAbi.json");

const consumerAddress = "0xa5e6a1D42779d9B95e0AeFDd87274dfFEC3852d8";
const subscriptionId = 46;

const routerAddress = "0xf9B8fc078197181C841c296C876945aaa425B278";
const donId = "fun-base-sepolia-1";
const explorerUrl = "https://sepolia.basescan.org/";
const gatewayUrls = [
    "https://01.functions-gateway.testnet.chain.link/",
    "https://02.functions-gateway.testnet.chain.link/",
];

const source = fs.readFileSync(path.resolve(__dirname, "source.js")).toString();

const args = ["RealWooblay", "chainlink-git-tasks", "2", "Onchain commit", "This was merged onchain!"];
const secrets = { githubToken: process.env.GITHUB_TOKEN };
const slotIdNumber = 0;
const expirationTimeMinutes = 15;
const gasLimit = 300000;

const request = async () => {
    const privateKey = process.env.PRIVATE_KEY;
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
        `Upload encrypted secret to gateways ${gatewayUrls}. slotId ${slotIdNumber}. Expiration in minutes: ${expirationTimeMinutes}`
    );

    const uploadResult = await secretsManager.uploadEncryptedSecretsToDON({
        encryptedSecretsHexstring: encryptedSecretsObj.encryptedSecrets,
        gatewayUrls: gatewayUrls,
        slotId: slotIdNumber,
        minutesUntilExpiration: expirationTimeMinutes,
    });

    if (!uploadResult.success) {
        throw new Error(`Encrypted secrets not uploaded to ${gatewayUrls}`);
    }

    console.log(
        `\n✅ Secrets uploaded properly to gateways ${gatewayUrls}! Gateways response: `,
        uploadResult
    );

    const donHostedSecretsVersion = parseInt(uploadResult.version); // fetch the reference of the encrypted secrets

    const functionsConsumer = new ethers.Contract(
        consumerAddress,
        functionsConsumerAbi,
        signer
    );

    const transaction = await functionsConsumer.sendRequest(
        source,
        "0x",
        slotIdNumber,
        donHostedSecretsVersion,
        args,
        [],
        subscriptionId,
        gasLimit,
        ethers.utils.formatBytes32String(donId)
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