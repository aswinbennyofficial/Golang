## Class

- No class feature in go

### Using recievers
```go
package main

import (
	"fmt"
)

// Define a new type MyInt based on int
type MyInt int

// Define a method Square for the MyInt type
func (p MyInt) Square() int {
	return int(p * p)
}

func main() {
	// Instantiate an object of type MyInt
	obj := MyInt(5)

	// Call the Square method on the MyInt object
	result := obj.Square()

	// Print the result
	fmt.Println(result)
}

```

```go
package main

import "fmt"

// Define a struct named 'Person'
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// Define a method 'FullName' for the 'Person' struct
func (p Person) FullName() string {
    return p.FirstName + " " + p.LastName
}

// Define a method 'IncreaseAge' for the 'Person' struct
func (p *Person) IncreaseAge(v int) {
    p.Age+=v
}

func main() {
    // Create an instance of the 'Person' struct
    person := Person{
        FirstName: "John",
        LastName:  "Doe",
        Age:       25,
    }

    // Access fields and call methods
    fmt.Println("First Name:", person.FirstName)
    fmt.Println("Last Name:", person.LastName)
    fmt.Println("Age:", person.Age)
    fmt.Println("Full Name:", person.FullName())

    // Call a method that modifies the struct
    person.IncreaseAge(100)
    fmt.Println("Updated Age:", person.Age) //125
}

```

