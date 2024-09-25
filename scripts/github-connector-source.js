if(!secrets.githubToken) {
    throw Error("Secret error!");
}

const request = Functions.makeHttpRequest({
    url: `https://api.github.com/user`,
    method: "PUT",
    headers: {
        "Accept": "application/vnd.github+json",
        "Authorization": `Bearer ${secrets.githubToken}`,
        "X-GitHub-Api-Version": "2022-11-28"
    }
});

const response = await request;

if(response.error) {
    throw new Error("Response error!");
}

const githubId = response.data.id;

return Functions.encodeUint256(githubId);