package main

import(
	"fmt"
	"net/http"
)

func githubHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	fmt.Fprintf(w,"request success")
}

func main(){
	http.HandleFunc("/v1/github",githubHandler)

	http.ListenAndServe(":8080",nil)
}