package main

import (
	"io/ioutil"
	"net/http"
)

// GithubUserContentURL The base URL to Github User Content
const GithubUserContentURL = "https://raw.githubusercontent.com/"

// GetGithubUserContent Fetch a raw file from Github.
func GetGithubUserContent(contentPath string) (string, error) {
	githubURL := GithubUserContentURL + contentPath
	response, err := http.Get(githubURL)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
