package main

import(
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)

type Person struct{
	Name map[string]string
}

func main(){
	scanner:=bufio.NewScanner(os.Stdin)
	
	var p1 Person
	p1.Name = make(map[string]string)

	fmt.Println("Enter the name: ")
	scanner.Scan()
	name:=scanner.Text()

	fmt.Println("Enter the address: ")
	scanner.Scan()
	address:=scanner.Text()

	p1.Name[name]=address

	bytearray,err:=json.Marshal(p1)

	if(err!=nil){
		fmt.Println("error occcured on marshalling")
	}else{
		jsonString:=string(bytearray)
		fmt.Println(jsonString)
	}

}