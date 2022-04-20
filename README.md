# uuid62

[![GoDoc](http://godoc.org/github.com/shibukawa/uuid62?status.svg)](http://godoc.org/github.com/shibukawa/uuid62)

Go implementation of short ID generator by using UUIDv1, V4, V6, V7 and base62 inspired by [npm uuid62 package](https://www.npmjs.com/package/uuid62).

It always return strings whose length is 21 or 22. It is shorter than string represention of UUID (36).
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

	ts, err :=
}
```

## API

### Generator Functions

- `func uuid62.V1() (string, error)`
- `func uuid62.V4() (string, error)`
- `func uuid62.V6() (string, error)`
- `func uuid62.V7Nano() (string, error)`
- `func uuid62.V7Micro() (string, error)`
- `func uuid62.V7Milli() (string, error)`

They generates UUID string in base62 format.

### Helper Functions

- `func uuid62.Decode(uuidString string) (uuid.UUID, error)`

Convert string to `github.com/gofrs/uuid`'s UUID object.

- `func uuid62.Encode(uuid.UUID uuidObj) (string, error)`

Convert `github.com/gofrs/uuid`'s UUID object to string.

- `func uuid62.Timestamp(uuidString string) (time.Time, error)`

Get timestamp from uuid string. It accepts only V1 and V6.

## License

Apache 2
