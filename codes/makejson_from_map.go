package main

import(
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)



func main(){
	scanner:=bufio.NewScanner(os.Stdin)
	
	
	Name := make(map[string]string)

	fmt.Println("Enter the name: ")
	scanner.Scan()
	name:=scanner.Text()

	fmt.Println("Enter the address: ")
	scanner.Scan()
	address:=scanner.Text()

	Name[name]=address

	bytearray,err:=json.Marshal(Name)

	if(err!=nil){
		fmt.Println("error occcured on marshalling")
	}else{
		jsonString:=string(bytearray)
		fmt.Println(jsonString)
	}

}