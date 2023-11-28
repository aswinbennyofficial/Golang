- Uses **make()**
- Two argument
	- Takes arguments : type and length/capacity
	- `make([]int, 10)
- Three argument
	- Takes arguments : type ,length, capacity 
	- `make([]int, 10, 15)`

- append() : expand size of slice by adding elements to end
	- `s1=append(s1,100)` 100 is the value to be inserted

```go
var s1=make( []int,10 )
```

```go
var s1=make([]int,3)

s1[0]=4
s1=append(s1, 89)

fmt.Println(s1)
```
