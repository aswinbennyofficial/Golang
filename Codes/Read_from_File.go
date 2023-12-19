/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line. 

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Firstname string
	Lastname  string
}

func main() {
	var arr []Person
	var filename string
	fmt.Println("Enter the name of file to be read? (eg. readme.txt): ")
	fmt.Scan(&filename)

	fd, err := os.Open(filename)
	if err != nil {
		fmt.Println("Couldnt open file ", err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		str := scanner.Text()
		words := strings.Split(str, " ")
		arr = append(arr, Person{Firstname: words[0], Lastname: words[1]})
	}

	for i, person := range arr {
		fmt.Printf("%d %s %s\n", i+1, person.Firstname, person.Lastname)
	}

}
