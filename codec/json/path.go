package json

import (
	"github.com/goccy/go-json"

	"github.com/starudream/go-lib/v2/x"
)

func Get[T any](data []byte, path string) (v T, err error) {
	return v, CreatePath(path).Unmarshal(data, &v)
}

func MustGet[T any](data []byte, path string) (v T) {
	return x.Must1(v, CreatePath(path).Unmarshal(data, &v))
}

func CreatePath(path string) *json.Path {
	return x.Must1(json.CreatePath(path))
}
