package uuid62

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/prop"
)

func TestV1(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("V1", prop.ForAll(func() string {
		str, err := V1()
		if err != nil {
			return "err should be nil (V1)"
		}
		if len(str) != 22 {
			return "length should be 22"
		}
		uuid, err := Decode(str)
		if err != nil {
			return "err should be nil (Decode)"
		}
		if len(uuid.Bytes()) != 16 {
			return "uuid is invalid"
		}
		return ""
	}))
	properties.TestingRun(t)
}

func TestV4(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("V4", prop.ForAll(func() string {
		str, err := V1()
		if err != nil {
			return "err should be nil (V4)"
		}
		if len(str) != 22 {
			return "length should be 22"
		}
		uuid, err := Decode(str)
		if err != nil {
			return "err should be nil (Decode)"
		}
		if len(uuid.Bytes()) != 16 {
			return "uuid is invalid"
		}
		return ""
	}))
	properties.TestingRun(t)
}