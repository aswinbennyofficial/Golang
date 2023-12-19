## 16. JSON

- import "encoding/json"



- **Marshal()** returns JSON representation as `[] bytes`

```go
package main

import (
	"encoding/json"
	"fmt"
)
type Person struct{
	name string
	phone int
	address string
}

func main() {
	var p1 =Person{name:"Aswin",phone:9087980,address:"CA, US"}
	bytearr,err:=json.Marshal(p1)
  }
```


- **Unmarshal()** returns struct from `[]bytes`

```go
var p2 Person
err:=json.Unmarshal(bytearr,&p2)

fmt.Println(p2.Name)
```


----




```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string // capitalize the first letter
	Phone   int    // Assuming phone is an int
	Address string
}

func main() {
	var p1 = Person{Name: "Aswin", Phone: 9087980, Address: "CA, US"}

	// Marshal the struct to JSON
	bytearr, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Unmarshal the JSON back into a new Person variable
	var p2 Person
	err = json.Unmarshal(bytearr, &p2)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the Name field of the p2 instance
	fmt.Println(p2.Name)
}

```
