
## 21.1 Weight Groups

- Add() : increments counter
- Done() : decrements counter
- Wait() : blocks until counter is 0
- uses "sync" package

```go
package main

import (
	"fmt"
	"sync"
)

func routineTwo(wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the counter when the goroutine completes
	fmt.Println("Routine two")
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup

	// Add 1 to the WaitGroup counter
	wg.Add(1)

	// Start the goroutine and pass the WaitGroup as an argument
	go routineTwo(&wg)

	fmt.Println("main routine")

	// Wait for the goroutine to finish by calling Wait on the WaitGroup
	wg.Wait()
}

```

