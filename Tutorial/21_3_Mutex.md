## 21.3 Mutex

- `sync` package
- Lock() : To lock the mutex and acquire exclusive access to the critical section
- Unlock() : To unlock the mutex and release access to the shared resource

```go
package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment() {
	// Lock the mutex to protect the critical section
	mutex.Lock()
	defer mutex.Unlock() // Ensure the mutex is always unlocked, even if an error occurs

	// Access and modify the shared resource (counter)
	counter++
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	increment()
}

func main() {
	var wg sync.WaitGroup

	// Launch multiple goroutines to increment the counter concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print the final value of the counter
	fmt.Println("Final Counter Value:", counter)
}
```

