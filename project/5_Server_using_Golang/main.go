package main

import(
	"fmt"
	// "log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path !="/hello"{
		http.Error(w,"404 Not found",http.StatusNotFound)
		return
	}

	if r.Method !="GET"{
		http.Error(w,"method not supported",http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w,"hello")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err:=r.ParseForm(); err!=nil{
		fmt.Fprintf(w,"Form parse error")
		return
	}

	fmt.Fprintf(w,"Post request successful")
	name:=r.FormValue("name")
	address:=r.FormValue("address")

	fmt.Fprintf(w,"Name: %s\nAddress: %s\n ",name,address)
}


func main(){
	fileserver:=http.FileServer(http.Dir("./project/5_Server_using_Golang/static"))
	
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("Server started..")

	err:=http.ListenAndServe(":8080",nil)
	
	if err!=nil{
		fmt.Println("Error ",err)
	}else{
		fmt.Println("Server started")
	}


}
