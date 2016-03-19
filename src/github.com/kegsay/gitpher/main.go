package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

// FileFetcherJSONRequest The JSON format for incoming /fetch requests
type FileFetcherJSONRequest struct {
	FilePath string
}

// GitCommandJSONRequest The JSON format for incoming /git requests
type GitCommandJSONRequest struct {
	Args []string
}

func fileFetcher(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Must be POST")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var fileRequest FileFetcherJSONRequest
	err := decoder.Decode(&fileRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Request body not JSON")
		return
	}
	if fileRequest.FilePath == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Missing 'filepath'")
		return
	}

	fmt.Println("Fetching ", fileRequest.FilePath)
	contents, err := GetGithubUserContent(fileRequest.FilePath)

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "Upstream github request failed")
	}
	fmt.Fprintf(w, contents)
}

func gitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Must be POST")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var gitRequest GitCommandJSONRequest
	err := decoder.Decode(&gitRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Request body not JSON")
		return
	}
	if len(gitRequest.Args) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No 'args' specified")
		return
	}

	fmt.Println("Executing ", gitRequest.Args)
	result, err := RunGitCommand("", gitRequest.Args)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "Problem executing command ", err)
		return
	}
	fmt.Fprintf(w, result)
}

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

	port := flag.Int("port", 8080, "The port to listen on")
	flag.Parse()

	fmt.Println("Listening on port", *port)
	http.HandleFunc("/fetch", fileFetcher)
	http.HandleFunc("/git", gitHandler)
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
