const {
    SecretsManager
} = require("@chainlink/functions-toolkit");
require("@chainlink/env-enc").config();
const ethers = require("ethers");
require('dotenv').config();

//////// MAKE REQUEST ////////

const consumerAddress = "0xab093521a504dcCa5Bf6efAb646f0c2dfc22f84d";
const subscriptionId = 50;
const secrets = { githubToken: process.env.GITHUB_TOKEN };
const gatewayUrls = ["https://01.functions-gateway.testnet.chain.link/", "https://02.functions-gateway.testnet.chain.link/"]
const expirationTimeMinutes = 5;
const privateKey = process.env.PRIVATE_KEY; // fetch PRIVATE_KEY

async function main() {
    if (!privateKey)
        throw new Error(
            "private key not provided - check your environment variables"
        );

    const rpcUrl = "https://base-sepolia-rpc.publicnode.com"; // fetch Sepolia RPC URL

    if (!rpcUrl)
        throw new Error("rpcUrl not provideds");

    const provider = new ethers.providers.JsonRpcProvider(rpcUrl);

    const wallet = new ethers.Wallet(privateKey);
    const signer = wallet.connect(provider); // create ethers signer for signing transactions
    const slotIdNumber = 0;

    console.log("\nMake request...");

    // First encrypt secrets and upload the encrypted secrets to the DON
    const secretsManager = new SecretsManager({
        signer: signer,
        functionsRouterAddress: "0xf9B8fc078197181C841c296C876945aaa425B278",
        donId: "fun-base-sepolia-1",
    });
    await secretsManager.initialize();

    // Encrypt secrets and upload to DON
    const encryptedSecretsObj = await secretsManager.encryptSecrets(secrets);

    console.log(
        `Upload encrypted secret to gateways ${gatewayUrls}. slotId ${slotIdNumber}. Expiration in minutes: ${expirationTimeMinutes}`
    );
    // Upload secrets
    const uploadResult = await secretsManager.uploadEncryptedSecretsToDON({
        encryptedSecretsHexstring: encryptedSecretsObj.encryptedSecrets,
        gatewayUrls: gatewayUrls,
        slotId: slotIdNumber,
        minutesUntilExpiration: expirationTimeMinutes,
    });

    if (!uploadResult.success)
        throw new Error(`Encrypted secrets not uploaded to ${gatewayUrls}`);

    console.log(
        `\n✅ Secrets uploaded properly to gateways ${gatewayUrls}! Gateways response: `,
        uploadResult
    );

    const donHostedSecretsVersion = parseInt(uploadResult.version); // fetch the reference of the encrypted secrets
}

main()