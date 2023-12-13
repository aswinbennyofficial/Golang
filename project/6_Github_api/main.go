package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GitHubUser represents the structure of the GitHub user data we want to fetch and return as JSON.
type GitHubUser struct {
	Name   string `json:"name"`
	Userid string `json:"login"`
}

// getGitHubUser makes an HTTP GET request to the provided URL and decodes the JSON response into a GitHubUser struct.
func getGitHubUser(url string) (*GitHubUser, error) {
	// Make an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decode the JSON response into a GitHubUser struct using json.Unmarshal
	var user GitHubUser
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// handleGithub is an HTTP handler function that fetches GitHub user data and returns it as JSON.
func handleGithub(w http.ResponseWriter, r *http.Request) {
	// URL for fetching GitHub user data
	url := "https://api.github.com/users/aswinbennyofficial"

	// Fetch GitHub user data using the getGitHubUser function
	user, err := getGitHubUser(url)
	if err != nil {
		// If there is an error fetching or decoding data, log the error and return an internal server error response.
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert the GitHubUser struct to JSON using Marshal
	responseJSON, err := json.Marshal(user)
	if err != nil {
		// If there is an error marshaling the data to JSON, log the error and return an internal server error response.
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate that the response is in JSON format
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	//w.Write(responseJSON)
	fmt.Fprintf(w,string(responseJSON))
}

func main() {
	// Register the handleGithub function to handle requests to the "/v1/github" endpoint
	http.HandleFunc("/v1/github", handleGithub)

	// Start the HTTP server on port 8080
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
