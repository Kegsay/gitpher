package main

import (
	"io/ioutil"
	"net/http"
)

const githubUserContentURL = "https://raw.githubusercontent.com/"

// GetGithubUserContent fetches a raw file from Github.
func GetGithubUserContent(contentPath string) ([]byte, error) {
	githubURL := githubUserContentURL + contentPath
	response, err := http.Get(githubURL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
