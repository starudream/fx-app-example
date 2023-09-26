package yaml

import (
	"testing"

	"github.com/starudream/go-lib/v2/fake"
	"github.com/starudream/go-lib/v2/x"
)

func TestYAML(t *testing.T) {
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

	t.Log(string(x.Must1(YAMLToJSON(bs1))))

	t.Log(MustMarshalIndentString(user2, UseSingleQuote(true), WithComment(CommentMap{
		"$.name":       []*Comment{HeadComment(" name")},
		"$.age":        []*Comment{HeadComment(" age ")},
		"$.birth.year": []*Comment{HeadComment(" birth year")},
	})))
}
