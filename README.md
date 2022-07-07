# **toggler**

Simple feature toggling package for Go applications.

## Installation

Add package to your application:

```
$ go get github.com/sarjon/toggler
```

## Example

```go
package main

import (
	"context"
	
	"github.com/sarjon/toggler"
	"github.com/sarjon/toggler/operator"
	"github.com/sarjon/toggler/storage"
)

func main() {
	inMemoryStorage, _ := storage.NewInMemory([]*toggler.Toggle{
		toggler.NewToggle("my_feature", toggler.NewOperatorCondition("ip", operator.NewEqual("127.0.0.1"))),
	})
	
	manager := toggler.NewManager(&toggler.ManagerArgs{
		Storage: inMemoryStorage,
	})
	
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ip", "127.0.0.1")
	
	active, _ := manager.Active("my_feature", ctx)
	if active {
		// my_feature is active when context "ip" matches configured IP.
    }
}
```

See more examples [here](./example/main.go).
