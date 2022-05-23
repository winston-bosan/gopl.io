package github

import "fmt"

const api_url string = "https://api.github.com"

func Build_url_create(owner, repo string) string {
	x := fmt.Sprintf(api_url + "/repos/%s/%s/issues", owner, repo)
	return x
}

func Build_url_update(owner, repo, issue_num string) string {
	x := fmt.Sprintf(api_url + "/repos/%s/%s/issues/%s", owner, repo, issue_num)
	return x
}

func Build_url_lock(owner, repo, issue_num string) string {
	x := fmt.Sprintf(api_url + "/repos/%s/%s/issues/%s/lock", owner, repo, issue_num)
	return x
}

func Build_url_get(owner, repo, issue_num string) string {
	x := fmt.Sprintf(api_url + "/repos/%s/%s/issues/%s", owner, repo, issue_num)
	return x
}