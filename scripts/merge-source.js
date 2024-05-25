const owner = args[0];
const repo = args[1];
const pullNumber = args[2];

const commitTitle = args[3];
const commitMessage = args[4];

if(!secrets.githubToken) {
    throw Error(
        "Secret error!"
    );
}

const request = Functions.makeHttpRequest({
    url: `https://api.github.com/repos/${owner}/${repo}/pulls/${pullNumber}/merge`,
    method: "PUT",
    headers: {
        "Accept": "application/vnd.github+json",
        "Authorization": `Bearer ${secrets.githubToken}`,
        "X-GitHub-Api-Version": "2022-11-28"
    },
    data: {
        commit_title: commitTitle,
        commit_message: commitMessage
    },
});

const response = await request;

if(response.error) {
    console.log("GitHub request error: ", response.error);
    return Functions.encodeUint256(0);
}

const sha = response.data.sha;

return Functions.encodeString(sha);