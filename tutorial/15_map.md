- key value pairs
- `[int]string` key is int and string is value

### Declaring
```go
var studentId map[int]string //initialise map variable
studentId =make(map[int]string) //make the map

studentId[1]="cat"
fmt.Println(studentId[1])

//or

var studentId=make(map[int]string)

studentId[1]="cat"
fmt.Println(studentId[1])


//or

studentId:= map[int]string {1:"cat"}

```


### Deleting
```go
delete(studentId,1) //name of the map, key
```

---
### Map functions

- Key Present?
```go
id,b:=studentId[1] 
// id gives value , b is boolen thats says whether key is present
```

- Length of Map
```go
len(studentId)
```

---
### Iterating a map
```go
var studentId= make(map[int]string)
	
studentId[1]="cat"
studentId[3]="bat"
studentId[4]="dog"

// i - key , e - value
for i,e:= range studentId{
	fmt.Println(i," value: ",e)	
}

```
