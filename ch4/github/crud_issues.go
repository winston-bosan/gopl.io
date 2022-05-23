package github

import (
	// "log"
	"net/http"
)

func GetIssue(owner, repo, issue_num string) (error) {
	get_url := Build_url_get(owner, repo, issue_num)
	_, err := http.Get(get_url)
	if err != nil { return err }
	return nil
}