package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func githubHandler(w http.ResponseWriter, r *http.Request) {
	var url string = "https://api.github.com/users/aswinbennyofficial"

	var data map[string]interface{} // Assuming you want to decode JSON into a map

	err := getJson(url, &data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	// Access the data retrieved from the GitHub API
	fmt.Fprintf(w, "GitHub User ID: %.f", data["id"])
}

func main() {
	http.HandleFunc("/v1/github", githubHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
