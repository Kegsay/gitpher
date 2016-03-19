package main

import (
	"io/ioutil"
	"net/http"
)

// The base URL to Github User Content
const GithubUserContentUrl = "https://raw.githubusercontent.com/"

func GetGithubUserContent(contentPath string) (string, error) {
	githubURL := GithubUserContentUrl + contentPath
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
