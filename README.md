# cacshLib
cash library for cashing key, value data

how it works:
```go
package main

import (
	"fmt"
	"github.com/DenisFilisov/cacshLib"
)

func main() {
	cacshLib.Set("1", "test")
	fmt.Println(cacshLib.Get("1"))
	cacshLib.Delete("1")
	fmt.Println(cacshLib.Get("1"))
}

```