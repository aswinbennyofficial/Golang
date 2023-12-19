## 21.4 Sync once

### Sync once
- `sync.Once`is a synchronization primitive provided by the `sync` package. It ensures that a particular function is only executed once, regardless of how many goroutines try to invoke it. It is commonly used for one-time initialization tasks.
- useful when you want to perform an expensive or critical initialization operation and want to ensure it's done only once, even in a concurrent environment with multiple goroutines.

```go
package main

import (
    "fmt"
    "sync"
)

var once sync.Once
var initializedValue int

func initialize() {
    initializedValue = 42
    fmt.Println("Initialization complete.")
}

func main() {
    // Execute the initialization function only once
    once.Do(initialize)

    // Subsequent calls to Do will not re-execute the initialization function
    once.Do(initialize)

    // Use the initialized value
    fmt.Println("Initialized Value:", initializedValue)
}

```
