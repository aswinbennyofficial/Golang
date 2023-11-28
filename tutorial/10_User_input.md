## 10. User input

### Scan Function
```go
_, err := fmt.Scan(&n) //number of items scanned and error

```


```go
var n int ;

fmt.Println("Enter the number: ")
fmt.Scan(&n);

fmt.Println("Entered number is ",n)
```


### Inserting String with spaces
```go
package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	var str string;
	var strinp string;
	fmt.Println("Enter the text: ")
	scanner :=bufio.NewScanner(os.Stdin)
	
	scanner.Scan()
	str=scanner.Text()

	scanner.Scan()
	strinp=scanner.Text()
	
	fmt.Println("The entered text is : ",str," ",strinp)

}
```


---

```go
reader := bufio.NewReader(os.Stdin)  
input, _ := reader.ReadString('\n')  
fmt.Print(input)
```