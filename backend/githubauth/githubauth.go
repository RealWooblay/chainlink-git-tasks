package githubauth

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/julienschmidt/httprouter"

	"chainlink-git-tasks/env"
	"chainlink-git-tasks/ghtasks"
	"chainlink-git-tasks/parser"
)

type MintContributorParams struct {
	Token string `bson:"token"`
}

type AddRepoParams struct {
	Token    string `bson:"token"`
	Username string `bson:"username"`
	OrgId    uint64 `bson:"orgId"`
	Owner    string `bson:"owner"`
	Repo     string `bson:"repo"`
}

type AuthenticatedUser struct {
	Id uint64 `bson:"id"`
}

func SetContributor(responseWriter http.ResponseWriter, httpRequest *http.Request, httpParams httprouter.Params) {
	// Parse request parameters
	var params MintContributorParams
	err := parser.ParseBodyParams(httpRequest.Body, &params)
	if err != nil {
		http.Error(responseWriter, "Invalid parameters!", http.StatusBadRequest)
		return
	}

	id, err := getUserId(params.Token)
	if err != nil {
		http.Error(responseWriter, "Github API error!", http.StatusBadRequest)
		return
	}

	// Set contributor NFT

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(map[string]interface{}{
		"id": id,
	})
}

func AddRepo(responseWriter http.ResponseWriter, httpRequest *http.Request, httpParams httprouter.Params) {
	// Parse request parameters
	var params AddRepoParams
	err := parser.ParseBodyParams(httpRequest.Body, &params)
	if err != nil {
		http.Error(responseWriter, "Invalid parameters!", http.StatusBadRequest)
		return
	}

	isContributor, err := isRepoContributor(params.Token, params.Username, params.Owner, params.Repo)
	if err != nil || !isContributor {
		http.Error(responseWriter, "Github API error!", http.StatusBadRequest)
		return
	}

	// Add repo to GHTasks contract
	
	providerUrl, err := env.GetString("BASE_SEPOLIA_PROVIDER")
	if err != nil {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}

	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}

	ghTasksAddress, err := env.GetString("GHTASKS_ADDRESS")
	if err != nil {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}
	
	address := common.HexToAddress(ghTasksAddress)
	instance, err := ghtasks.NewGhtasks(address, client)
	if err != nil {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}

	privKeyString, err := env.GetString("PRIVATE_KEY")
	if err != nil {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}

	privateKey, err := crypto.HexToECDSA(privKeyString)
	if err != nil {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}

	publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(84532))
	if err != nil {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}

    txOpts.Nonce = big.NewInt(int64(nonce))
    txOpts.Value = big.NewInt(0)
    txOpts.GasLimit = uint64(300000)
    txOpts.GasPrice = gasPrice

	tx, err := instance.SetOrgRepo(txOpts, big.NewInt(int64(params.OrgId)), params.Owner, params.Repo, true)
	if err != nil {
		http.Error(responseWriter, "Internal error!", http.StatusBadRequest)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(map[string]interface{}{
		"txHash": tx.Hash().Hex(),
	})
}

func getUserId(token string) (uint64, error) {
	url := "https://api.github.com/user"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != 200 {
		return 0, errors.New(res.Status)
	}

	var authenticatedUser AuthenticatedUser
	err = parser.ParseBodyParams(res.Body, &authenticatedUser)
	if err != nil {
		return 0, err
	}

	return authenticatedUser.Id, nil
}

func isRepoContributor(token string, username string, owner string, repo string) (bool, error) {
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/collaborators/" + username

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, err := client.Do(req)
	if err != nil {
		return false, err
	}

	if res.StatusCode != 204 {
		return false, nil
	}

	return true, nil
}
