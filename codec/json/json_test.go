package json

import (
	"testing"

	"github.com/starudream/go-lib/v2/fake"
	"github.com/starudream/go-lib/v2/x"
)

func TestJSON(t *testing.T) {
	type Birth struct {
		Year int `yaml:"year" fake:"{number:1900,2020}"`
		Mon  int `yaml:"mon" fake:"{number:1,12}"`
		Day  int `yaml:"day" fake:"{number:1,28}"`
	}
	type User struct {
		Name   string   `yaml:"name" fake:"{name}"`
		Age    int      `yaml:"age" fake:"{number:1,100}"`
		Male   bool     `yaml:"male,omitempty"`
		Weight float64  `yaml:"weight" fake:"{float64range:50,200}"`
		Birth  Birth    `yaml:"birth"`
		X      struct{} `yaml:"-"`
	}

	user1 := &User{}
	x.Must0(fake.F().Struct(user1))

	s1 := MustMarshalString(user1)
	t.Log(s1)

	user2 := MustUnmarshalStringTo[User](s1)
	t.Logf("%#v", user2)

	bs1 := MustMarshalIndent(user2)
	t.Log(string(bs1))

	bs2 := MustCompact(bs1)
	t.Log(string(bs2))

	t.Log(string(MustIndent(bs2)))

	t.Log(MustMarshalIndentString(user1, Colored()))

	m := map[string]any{
		"b": "2",
		"a": "1",
		"c": "3",
	}

	t.Log(MustMarshalString(m))
	t.Log(MustMarshalString(m, UnorderedMap()))
}
