package main

import (
	
	"fmt"
	"log"
	"net/http"
	
)



func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/refresh", refreshTokenHandler)

	fmt.Println("Starting server in port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Printf("Error starting the server at port 8080..")
	}
}

