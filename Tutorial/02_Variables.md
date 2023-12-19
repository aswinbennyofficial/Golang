## 2. Variables

![image](https://github.com/aswinbennyofficial/Golang/assets/110408942/0b9ec672-dd5f-4a0f-8961-a2ed24da7853)

### Declare variables
```go
var str string = "Hello guys"  
fmt.Println("Hello world " + str)
```

- Short variable declaration
```go
numb := 1000  // only allowed inside a method
fmt.Println(numb)
```


### Rune
- Every single byte in a string is  a rune
- Go uses UTF 8
- Rune is a UTF 8 code point

### Public and Private
```go
const Num int8=100; //single first letter capital, it is a public variable
```

```go
const x=45
```


### Pointer
```go
var a int=10;


var ptr *int;
ptr=&a;

fmt.Println(*ptr);
```

---
### Memory Location
![image](https://github.com/aswinbennyofficial/Golang/assets/110408942/ec65f4aa-6c2c-4a3e-a7c5-f15f11555796)


- heap : global variable
- stack : local variable

- heap is persistant
