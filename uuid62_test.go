package uuid62

import (
	"fmt"
	"testing"

	"github.com/gofrs/uuid"
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
		if len(str) != 22 && len(str) != 21 {
			return fmt.Sprintf("length should be 21 or 22, but %d", len(str))
		}
		uuidValue, err := Decode(str)
		if err != nil {
			return "err should be nil (Decode)"
		}
		if len(uuidValue.Bytes()) != 16 {
			return "uuid is invalid"
		}
		if uuidValue.Version() != uuid.V1 {
			return "uuid version is invalid"
		}
		ts, err := Timestamp(str)
		if ts.IsZero() {
			return "V1 should support timestap"
		}
		if err != nil {
			return "V1 should support timestap"
		}
		return ""
	}))
	properties.TestingRun(t)
}

func TestV4(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("V4", prop.ForAll(func() string {
		str, err := V4()
		if err != nil {
			return "err should be nil (V4)"
		}
		if len(str) != 22 && len(str) != 21 {
			return fmt.Sprintf("length should be 21 or 22, but %d", len(str))
		}
		uuidValue, err := Decode(str)
		if err != nil {
			return "err should be nil (Decode)"
		}
		if len(uuidValue.Bytes()) != 16 {
			return "uuid is invalid"
		}
		if uuidValue.Version() != uuid.V4 {
			return "uuid version is invalid"
		}
		ts, err := Timestamp(str)
		if !ts.IsZero() {
			return "V4 can't support timestap"
		}
		if err == nil {
			return "V4 can't support timestap"
		}
		return ""
	}))
	properties.TestingRun(t)
}

func TestV6(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("V6", prop.ForAll(func() string {
		str1, err := V6()
		if err != nil {
			return "err should be nil (V6)"
		}
		str2, err := V6()
		if err != nil {
			return "err should be nil (V6)"
		}

		uuid1, err := Decode(str1)
		if err != nil {
			return "err should be nil (Decode)"
		}
		uuid2, err := Decode(str2)
		if err != nil {
			return "err should be nil (Decode)"
		}

		if (str1 < str2) != (uuid1.String() < uuid2.String()) {
			return "uuid62's string and original uuid should have same order"
		}

		if len(str1) != 22 && len(str1) != 21 {
			return fmt.Sprintf("length should be 21 or 22, but %d", len(str1))
		}
		if len(uuid1.Bytes()) != 16 {
			return "uuid is invalid"
		}
		if uuid1.Version() != uuid.V6 {
			return "uuid version is invalid"
		}

		ts, err := Timestamp(str1)
		if ts.IsZero() {
			return "V6 should support timestap"
		}
		if err != nil {
			return "V6 should support timestap"
		}

		return ""
	}))
	properties.TestingRun(t)
}

func TestV7(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("V7", prop.ForAll(func() string {
		str1, err := V7()
		if err != nil {
			return "err should be nil (V7)"
		}
		str2, err := V7()
		if err != nil {
			return "err should be nil (V7)"
		}
		if len(str1) != 22 && len(str1) != 21 {
			return fmt.Sprintf("length should be 21 or 22, but %d", len(str1))
		}

		uuid1, err := Decode(str1)
		if err != nil {
			return "err should be nil (Decode)"
		}
		uuid2, err := Decode(str2)
		if err != nil {
			return "err should be nil (Decode)"
		}
		if (str1 < str2) != (uuid1.String() < uuid2.String()) {
			return "uuid62's string and original uuid should have same order"
		}

		if len(uuid1.Bytes()) != 16 {
			return "uuid is invalid"
		}
		if uuid1.Version() != uuid.V7 {
			return "uuid version is invalid"
		}
		ts, err := Timestamp(str1)
		if !ts.IsZero() {
			return "V7 can't support timestap"
		}
		if err == nil {
			return "V7 can't support timestap"
		}
		return ""
	}))
	properties.TestingRun(t)
}
