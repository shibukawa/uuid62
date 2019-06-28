# uuid62

[![GoDoc](http://godoc.org/github.com/shibukawa/uuid62?status.svg)](http://godoc.org/github.com/shibukawa/uuid62)

Go implementation of short ID generator by using UUIDv4 and base62 inspired by [npm uuid62 package](https://www.npmjs.com/package/uuid62).

It always return strings whose length is 22.

```go
package main

import (
	"fmt"
	"github.com/shibukawa/uuid62"
)

func main() {
	id, err := uuid62.V1()
	if err != nil {
		fmt.Println(id)
	}
}
```

## License

Apache 2