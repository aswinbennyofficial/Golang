## 17 File access


### Reading

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	//fd is file descriptor
	fd,err:=os.Open("/home/aswinbenny/VScode/Golang/Golang/codes/readme.txt")

	defer fd.Close() // Ensure the file is closed when the function exits

	bytearr:=make([]byte,1000) //making a byte array
	
	nb,err:=fd.Read(bytearr) // returns number of bytes read and error// Ensure the file is closed when the function exits

	fmt.Println("Number of bytes read : ",nb)
	fmt.Println("Content: ", string(bytearr[:nb]) )
	
	

  }
```

- read line by line
```go
// import bufio and os

fd, err := os.Open("readme.txt")

defer fd.Close()

scanner := bufio.NewScanner(fd)

for scanner.Scan() {
	str := scanner.Text()
}
```



### Writing
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fd,err:=os.Create("./codes/gowrite.txt")

	bytearr:=[]byte{1,2,3,4}
	nb,err:=fd.Write(bytearr) //writes bytes array

	nb,err:=fd.WriteString("Hello world") //writes string
	
  }
```

