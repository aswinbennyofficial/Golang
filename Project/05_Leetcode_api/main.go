package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Submission struct {
	Difficulty  string `json:"difficulty"`
	Count       int    `json:"count"`
	Submissions int    `json:"submissions"`
}

type SubmissionData struct {
	TotalSolved      int          `json:"totalSolved"`
	//TotalSubmissions []Submission `json:"totalSubmissions"`
	EasySolved       int          `json:"easySolved"`
	MediumSolved     int          `json:"mediumSolved"`
	HardSolved       int          `json:"hardSolved"`
}

func handleLeetcode(w http.ResponseWriter, r *http.Request) {

	url := "https://leetcode-api-faisalshohag.vercel.app"+r.URL.Path // Replace with your actual API endpoint

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var user SubmissionData

	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	// Now 'user' contains the decoded JSON data, you can use it as needed.

	// // Example: Print totalSolved, easySolved, mediumSolved, hardSolved
	// fmt.Fprintf(w,"Total Solved: %d\n", user.TotalSolved)
	// fmt.Fprintf(w,"Easy Solved: %d\n", user.EasySolved)
	// fmt.Fprintf(w,"Medium Solved: %d\n", user.MediumSolved)
	// fmt.Fprintf(w,"Hard Solved: %d\n", user.HardSolved)

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
	http.HandleFunc("/", handleLeetcode)
	// example localhost:8080/aswinbenny

	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
