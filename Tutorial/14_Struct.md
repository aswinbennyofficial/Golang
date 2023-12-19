## 14. Struct

### Initialise

```go
package main

import(
	"fmt"
)
type Person struct{
	name string
	phone int
	address string
}

func main(){

	var p1 Person
	p1.name="Aswin"
	p1.phone=989695996
	p1.address="India, Asia" 

	fmt.Println(p1.name)

}
```


```go
//or
var p1 =new(Person) //initialises to 0
```


```go
//or
package main

import(
	"fmt"
)
type Person struct{
	name string
	phone int
	address string
}

func main(){

	var p1 =Person{
		name:"Aswin", 
		phone:89765800,
	}
	

	fmt.Println(p1.name)

}
```

