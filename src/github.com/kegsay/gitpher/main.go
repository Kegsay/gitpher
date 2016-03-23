package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

var port = flag.Int("port", 8080, "The port to listen on")

func main() {
	fmt.Println("         ,_---~~~~~----._         ")
	fmt.Println("  _,,_,*^____      _____``*g*\"*, ")
	fmt.Println(" / __/ /'     ^.  /      \\ ^@q   f ")
	fmt.Println("[  @f | @))    |  | @))   l  0 _/  ")
	fmt.Println(" \\`/   \\~____ / __ \\_____/    \\   ")
	fmt.Println("  |           _l__l_           I   ")
	fmt.Println("  }          [______]           I  ")
	fmt.Println("  ]            | | |            |  ")
	fmt.Println("  ]             ~ ~             |  ")
	fmt.Println("  |                            |   ")
	fmt.Println("   |                           |  ")
	fmt.Println("              Gitpher")
	fmt.Println()

	flag.Parse()

	http.HandleFunc("/fetch", fileFetcher)
	http.HandleFunc("/git", gitHandler)
	fmt.Println("Listening on port", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}

// FileFetcherRequest is the JSON format for incoming /fetch requests.
type FileFetcherRequest struct {
	FilePath string `json:"filepath"` // The path to get from github user content.
}

// GitCommandRequest is the JSON format for incoming /git requests.
type GitCommandRequest struct {
	Args []string `json:"args"` // Argv[1:] for the git command to run.
}

func fileFetcher(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Must be POST")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var fileRequest FileFetcherRequest
	err := decoder.Decode(&fileRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Request body not JSON")
		return
	}
	if fileRequest.FilePath == "" {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Missing 'filepath'")
		return
	}

	fmt.Println("Fetching ", fileRequest.FilePath)
	contents, err := GetGithubUserContent(fileRequest.FilePath)

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "Upstream github request failed")
	}
	w.Write(contents)
}

func gitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Must be POST")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var gitRequest GitCommandRequest
	err := decoder.Decode(&gitRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Request body not JSON")
		return
	}
	if len(gitRequest.Args) == 0 {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "No 'args' specified")
		return
	}

	fmt.Println("Executing ", gitRequest.Args)
	result, err := RunGitCommand(gitRequest.Args)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Problem executing command %s", err.Error())
		// Should probably actually define a custom error type here, or shunt the stderr into a meaningful message on failure.
		// But this shows you roughly how you would do that.
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Println("Stderr: ", string(exitErr.Stderr))
		}
		return
	}
	io.WriteString(w, result)
}
