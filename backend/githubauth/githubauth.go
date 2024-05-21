package githubauth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"chainlink-git-tasks/parser"
)

type MintContributorParams struct {
	Token    string `bson:"token"`
}

type AddRepoParams struct {
	Token    string `bson:"token"`
	Username string `bson:"username"`
	Owner    string `bson:"owner"`
	Repo     string `bson:"repo"`
	OrgId    uint64 `bson:"orgId"`
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

	// Add repo to smart contract consumer

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(map[string]interface{}{})
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
