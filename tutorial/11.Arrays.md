
### Initialising and declaring

```go
var arr [5]int
arr[0]=8

fmt.Println(arr[0])
```

```go
var arr [5]int
arr=[5]int{1,2,3,4,5}

//or directly
var arr [5]int=[5]int{1,2,3,4,5}

//or
var arr [5]int=[...]int{1,2,3,4,5}
```

- integer array is initialised to 0 by default

### Traversing
```go
var arr [5]int=[...]int{1,2,3,4,5}

// i is the index and e is the element in that index
for i,e := range arr{
	fmt.Println(e, " is present at index ", i)
}

//or

for i:=0;i<len(arr);i++{
	fmt.Println(arr[i], " is present at index ", i)
}
```
