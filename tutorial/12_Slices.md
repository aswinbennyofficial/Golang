## 12. Slices

- Window of an underlying array
- Size is variable upto the length of underlying array
- **Pointer** : indicates start of the slice
- **Length** : number of elements in the slice
- **Capacity** : maximum number of elements in the slice
- Editing slice also affects underlying array


```go
var arr [5]int=[...]int{1,2,3,4,5}

s1:=arr[1:3] // does not include index 3
```

### Initialising a slice
- Creates an underlying array and creates a slice to reference it in a single command
```go
s1:=[]int{1,2,3,4}
```


---
### Len and Capacity
- **len()** : shows the length of the slice
- **capcity()** : show the maximum size it can extend
	- eg, If arr is of size 5
	- arr[1:3] has **capacity of 4** (ie. 1,2,3,4) and **length 2**
	- arr[0:3] has **capacity of 5** (ie. 0,1,2,3,4) and **length 2**

```go
var arr [5]int=[...]int{1,2,3,4,5}

s1:=arr[0:3]
fmt.Println( len(s1),"  ", cap(s1) )

```

