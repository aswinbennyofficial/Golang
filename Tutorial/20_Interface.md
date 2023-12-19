## 20 Interfaces

- Specifies method signatures
- Abstracts implementation

```go
package main

import "fmt"

// Shape interface defines a method Area() that returns an int
type Shape interface {
	Area() int
}

// Square type implements the Shape interface
type Square struct {
	SideLength int
}

// Area method for Square
func (s Square) Area() int {
	return s.SideLength * s.SideLength
}

// Circle type also implements the Shape interface
type Circle struct {
	Radius int
}

// Area method for Circle
func (c Circle) Area() int {
	return 3 * c.Radius * c.Radius // a simple approximation of pi for demonstration
}

func main() {
	// Create instances of Square and Circle
	square := Square{SideLength: 4}
	circle := Circle{Radius: 3}

	// Use the Area method with different shapes
	fmt.Println("Square Area:", square.Area())
	fmt.Println("Circle Area:", circle.Area())

	// Create a slice of shapes
	shapes := []Shape{square, circle}

	// Iterate through the shapes and print their areas
	for _, shape := range shapes {
		fmt.Println("Shape Area:", shape.Area())
	}
}

```
