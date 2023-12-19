package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strings"
)

type AuthorInfo struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type PublisherInfo struct{
	Title string `json:"title"`
	Author AuthorInfo `json:"author"`
	Published bool `json:"published"`
	Tags []string `json:"tags"`
	Content string `json:"content"`
}

func handlePublisher(w http.ResponseWriter, r *http.Request){
	// plain json data
	jsonPlainData:=`{"title":"Sample Article","author":{"name":"John Doe","email":"john.doe@example.com"},"published":true,"tags":["programming","web development","golang"],"content":"This is a sample article content. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed et sapien ac nisl venenatis gravida."}`

	// making a object of struct PublisherInfo
	var publisherStruct PublisherInfo

	// decoding plain json and saves it to the struct object
	json.NewDecoder(strings.NewReader(jsonPlainData)).Decode(&publisherStruct)

	fmt.Fprintf(w,publisherStruct.Author.Name)


}

func main(){
	http.HandleFunc("/",handlePublisher)

	fmt.Printf("Listening to 8080")
	http.ListenAndServe(":8080",nil)

}