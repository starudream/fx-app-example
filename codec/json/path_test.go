package json

import (
	"testing"
)

func TestPath(t *testing.T) {
	bs := []byte(`{"name":"Brant Doyle","age":59,"male":true,"weight":125.37156563310717}`)

	t.Log(MustGet[string](bs, "$.name"))
	t.Log(MustGet[float64](bs, "$.weight"))
}
