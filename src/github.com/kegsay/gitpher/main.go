package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The JSON format for incoming /fetch requests
type FileFetcherJSONRequest struct {
	FilePath string
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

	http.HandleFunc("/fetch", fileFetcher)
	http.ListenAndServe(":8080", nil)
}
