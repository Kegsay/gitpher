package main

import (
	"io/ioutil"
	"net/http"
)

const GITHUB_URL = "https://raw.githubusercontent.com/"

func GetGithubUserContent(contentPath string) (string, error) {
	githubUrl := GITHUB_URL + contentPath
	response, err := http.Get(githubUrl)
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
