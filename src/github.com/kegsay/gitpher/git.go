package main

import (
	"os/exec"
)

// RunGitCommand executes an arbitrary git command.
// The standard output is returned. Standard error is discarded.
// Args should be argv[1:] to pass to git.
func RunGitCommand(args []string) (string, error) {
	result, err := exec.Command("git", args...).Output()
	if err != nil {
		return "", err
	}
	return string(result), nil
}
