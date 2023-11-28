
### strings package
```go
import "strings"
```

```go
var str string="LAZY CAT jumps the wall";

fmt.Println( strings.ToLower(str) );
fmt.Println( strings.ToUpper(str) );

// str, old, new, number of instance to change
fmt.Println( strings.Replace(str,"CAT","DOG",2) ); 
fmt.Println( strings.ReplaceAll(str,"CAT","DOG") );

fmt.Println(strings.TrimSpace(str));

fmt.Println(strings.Contains(str,"DOG"));
```

### strconv package

```go
import "strconv"
```

- Atoi : convert string to integer
- Itoa : convert integer to string

```go
var str string="53";

i,err :=strconv.Atoi(str);
if err==nil {
	fmt.Println(i+5)
}

```

