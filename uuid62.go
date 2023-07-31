package uuid62

import (
	"errors"
	"time"

	"github.com/eknkc/basex"
	"github.com/gofrs/uuid/v5"
)

var base62 *basex.Encoding

func init() {
	base62, _ = basex.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
}

func V1() (string, error) {
	uuid, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	return Encode(uuid), nil
}

func V4() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return Encode(uuid), nil
}

func V6() (string, error) {
	uuid, err := uuid.NewV6()
	if err != nil {
		return "", err
	}
	return Encode(uuid), nil
}

func V6WithTimeStamp() (string, uint64, error) {
	v6uuid, err := uuid.NewV6()
	if err != nil {
		return "", 0, err
	}
	ts, _ := uuid.TimestampFromV6(v6uuid)
	return Encode(v6uuid), uint64(ts), nil
}

func V7() (string, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return Encode(uuid), nil
}

func Decode(uuid62 string) (uuid.UUID, error) {
	rawBytes, err := base62.Decode(uuid62)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.FromBytes(rawBytes)
}

func Encode(uuid uuid.UUID) string {
	return base62.Encode(uuid.Bytes())
}

func Timestamp(uuid62 string) (time.Time, error) {
	uuidObj, err := Decode(uuid62)
	if err != nil {
		return time.Time{}, nil
	}
	var ts uuid.Timestamp
	switch uuidObj.Version() {
	case uuid.V1:
		ts, err = uuid.TimestampFromV1(uuidObj)
	case uuid.V6:
		ts, err = uuid.TimestampFromV6(uuidObj)
	default:
		return time.Time{}, errors.New("timestamp is available only V1 and V6")
	}
	if err != nil {
		return time.Time{}, err
	}
	return ts.Time()
}
