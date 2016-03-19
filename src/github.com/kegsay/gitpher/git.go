package main

import (
	"os/exec"
)

// RunGitCommand Execute an arbitrary git command. The output is returned.
func RunGitCommand(path string, args []string) (string, error) {
	result, err := exec.Command("git", args...).Output()
	if err != nil {
		return "", err
	}
	return string(result), nil
}
