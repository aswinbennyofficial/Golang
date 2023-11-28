```go
import "fmt"

//or 


import(
	"fmt"
	"strings"
)
```


- Provide alias to packages
```go
package main

import (
	k "fmt"
)

func main() {
	k.Println("Hello world");
}
```