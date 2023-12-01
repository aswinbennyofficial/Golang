## Dotenv


### Make a .env file
```.env
API_KEY="dsjkfkjdslkjflkjdshf"

```


### Import the packages
```go
import (
    "fmt"
    "github.com/joho/godotenv"
    "log"
    "os"
)
```


### Load env file in the program
```go
err := godotenv.Load(".env")
if err != nil {
        log.Fatalf("Error loading environment variables file")
    }
```

### Get the variables
```go
key := os.Getenv("API_KEY")

fmt.Println(key)
```
