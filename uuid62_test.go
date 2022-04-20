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
		str, err := V6()
		if err != nil {
			return "err should be nil (V6)"
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
		if uuidValue.Version() != uuid.V6 {
			return "uuid version is invalid"
		}
		ts, err := Timestamp(str)
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

func TestV7Nano(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("V7Nano", prop.ForAll(func() string {
		str, err := V7Nano()
		if err != nil {
			return "err should be nil (V6)"
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
		if uuidValue.Version() != uuid.V7 {
			return "uuid version is invalid"
		}
		ts, err := Timestamp(str)
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

func TestV7Micro(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("V7Micro", prop.ForAll(func() string {
		str, err := V7Micro()
		if err != nil {
			return "err should be nil (V7)"
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
		if uuidValue.Version() != uuid.V7 {
			return "uuid version is invalid"
		}
		ts, err := Timestamp(str)
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

func TestV7Milli(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("V7Milli", prop.ForAll(func() string {
		str, err := V7Milli()
		if err != nil {
			return "err should be nil (V7Milli)"
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
		if uuidValue.Version() != uuid.V7 {
			return "uuid version is invalid"
		}
		ts, err := Timestamp(str)
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
