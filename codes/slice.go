/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

*/

package main

import(
	"fmt"
	"slices"
	"strconv"
)

func main(){
	// initialising a slice with default size of 3
	var sliceVar=make([]int,3)
	// index for slice
	var i int=0

	//pre initialising slice with 0
	sliceVar[0]=0
	sliceVar[1]=0
	sliceVar[2]=0

	for true{

		var n string
		fmt.Scan(&n)
		
		if(n=="X" || n=="x") {
			break
		}

		//converting string to integer
		e,err:=strconv.Atoi(n)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		//fill the empty slice portion
		if(i<=2){
			sliceVar[i]=e
			i++
		} else{

			sliceVar=append(sliceVar,e)
		}
		slices.Sort(sliceVar)
		fmt.Println(sliceVar)

	}


}