# uuid62

[![GoDoc](http://godoc.org/github.com/shibukawa/uuid62?status.svg)](http://godoc.org/github.com/shibukawa/uuid62)

Go implementation of short ID generator by using UUIDv4 and base62 inspired by [npm uuid62 package](https://www.npmjs.com/package/uuid62).

It always return strings whose length is 22. It is shorter than string represention of UUID (36).
And the results includes only alphabet and numeric characters (no symbols).

```go
package main

import (
	"fmt"
	"github.com/shibukawa/uuid62"
)

func main() {
	// it just returns string
	idString, err := uuid62.V1()
	if err != nil {
		fmt.Println(idString)
	}
	
	// it also provides API decode/encode between UUID
	uuid, err := Decode(idString)
	if err != nil {
		fmt.Println(idString)
	}
	idStringAgain := Encode(uuid) // idStringAgain == idString
	fmt.Println(idStringAgain)
}
```

## License

Apache 2