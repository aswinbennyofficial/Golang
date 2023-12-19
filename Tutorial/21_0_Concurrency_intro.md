## 21.0 Concurrency

- Concurrency not equals parrallelism
- Go routines : lightweight threads managed by the Go runtime, allowing concurrent execution of functions.
- Go runtime scheduler : Schedule goroutines inside an OS thread


- `go` keyword is used to create a new go routine
- sleep the main func for 100ms
	`time.Sleep(100 * time.Millisecond)`

```go
package main

import (
	"fmt"
	"time"
)

func myFunction(id int) {
	fmt.Printf("Goroutine %d started\n", id)
	time.Sleep(time.Second * time.Duration(id)) // Simulate work with sleep
	fmt.Printf("Goroutine %d completed\n", id)
}

func main() {
	// Launch multiple goroutines with a delay using time.Sleep
	for i := 1; i <= 5; i++ {
		go myFunction(i)
		time.Sleep(time.Second) // Introduce a delay between goroutine launches
	}

	// Allow some time for goroutines to complete before exiting
	time.Sleep(10 * time.Second )
}

```