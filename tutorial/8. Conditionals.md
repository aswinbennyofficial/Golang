
### If Else
- `else` and `}` should be in the same line

```go
var x int =5;

if x>5{
fmt.Println("Greater than 5");

} else if x==5{
fmt.Println("Equal to 5");

} else{
	fmt.Println("Less than 5");

}
```


### Switch
```go
var x int =3;

switch x{
case 1:
	fmt.Println("1 says hi")
case 2:
	fmt.Println("2 says hi")
case 3:
	fmt.Println("3 says hi")
default:
	fmt.Println("Wrong choice")
}
```


### Tagless Switch
```go
var x int =4;

switch {
case x>5:
	fmt.Println("Greater than 5")
case x==5:
	fmt.Println("Equals 5")
default:
	fmt.Println("Less than 5")
}
```
