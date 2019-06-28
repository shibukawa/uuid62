package uuid62

import (
	"github.com/eknkc/basex"
	"github.com/gofrs/uuid"
)

var base62 *basex.Encoding

func init() {
	base62, _ = basex.NewEncoding("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
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