## 18 Functions

### Can have multiple return values

```go
package main

import(
	"fmt"
)

func cat(x int) (int, int, int){
	return x+1,5,8
}

func main(){
	a,b,c:=cat(12)
	fmt.Println(a,b,c)
}
```


### Pass by reference
```go
func cat(x *int){
	*x = *x + 1
	
}

func main(){
	var z int=8
	cat(&z)

	fmt.Println(z)
}
```

- arrays
```go
func cat(x *[3]int){
	(*x)[0]=(*x)[0]+10
	
}

func main(){
	var z [3]int=[3]int{1,2,3}
	cat(&z)

	fmt.Println(z[0])
}
```

- slices are always passed by reference
```go
func cat(x []int){
	x[0]=x[0]+10
	
}

func main(){
	z:=[]int{1,2,3}
	cat(z)

	fmt.Println(z)
}
```


### Function as a variable
```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func main() {
    // Declare a variable 'operation' of type func(int, int) int
    var operation func(int, int) int

    // Assign the 'add' function to the 'operation' variable
    operation = add

    // Use the 'operation' variable as a function
    result := operation(3, 4)
    fmt.Println("Result of add:", result)

    // Assign the 'subtract' function to the 'operation' variable
    operation = subtract

    // Use the 'operation' variable as a function again
    result = operation(7, 2)
    fmt.Println("Result of subtract:", result)
}

```


### Anonymous functions
```go
package main

import "fmt"

func main() {
    // Define an anonymous function and assign it to the variable 'add'
    add := func(a, b int) int {
        return a + b
    }

    // Use the anonymous function
    result := add(3, 4)
    fmt.Println("Result of add:", result)

    // You can also directly call the anonymous function without assigning it to a variable
    subtract := func(a, b int) int {
        return a - b
    }

    // Use the anonymous function directly
    result = subtract(7, 2)
    fmt.Println("Result of subtract:", result)
}

```


### Function returning function
```go
package main

import "fmt"

func generateMultiplier(factor int) func(int) int {
    // This function returns another function
    return func(x int) int {
        return factor * x
    }
}

func main() {
    // Call generateMultiplier to create a new function that multiplies by 3
    multiplyBy3 := generateMultiplier(3)

    // Use the generated function
    result := multiplyBy3(4)
    fmt.Println("Result:", result)  // Output: 12

    // Call generateMultiplier to create a new function that multiplies by 5
    multiplyBy5 := generateMultiplier(5)

    // Use the generated function
    result = multiplyBy5(4)
    fmt.Println("Result:", result)  // Output: 20
}

```


### Defer
- Evaluated at the same time but excecuted when all other codes are completed
```go
func main(){
	i:=1
	defer fmt.Println(i+1) 
	i++
	fmt.Println("Hello")
	//hello 2
}
```
