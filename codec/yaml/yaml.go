package yaml

import (
	"github.com/goccy/go-yaml"

	"github.com/starudream/go-lib/v2/x"
)

type (
	EncodeOption = yaml.EncodeOption
	DecodeOption = yaml.DecodeOption

	Comment    = yaml.Comment
	CommentMap = yaml.CommentMap
)

var (
	Flow                       = yaml.Flow
	Indent                     = yaml.Indent
	IndentSequence             = yaml.IndentSequence
	UseSingleQuote             = yaml.UseSingleQuote
	UseLiteralStyleIfMultiline = yaml.UseLiteralStyleIfMultiline
	WithComment                = yaml.WithComment

	CommentToMap  = yaml.CommentToMap
	Strict        = yaml.Strict
	UseOrderedMap = yaml.UseOrderedMap

	HeadComment = yaml.HeadComment
	FootComment = yaml.FootComment
	LineComment = yaml.LineComment

	NewEncoder = yaml.NewEncoder
	NewDecoder = yaml.NewDecoder

	JSONToYAML = yaml.JSONToYAML
	YAMLToJSON = yaml.YAMLToJSON
)

func Marshal(v any, opts ...EncodeOption) ([]byte, error) {
	return yaml.MarshalWithOptions(v, opts...)
}

func MustMarshal(v any, opts ...EncodeOption) []byte {
	return x.Must1(Marshal(v, opts...))
}

func MustMarshalString(v any, opts ...EncodeOption) string {
	return string(MustMarshal(v, opts...))
}

func MarshalIndent(v any, spaces int, opts ...EncodeOption) ([]byte, error) {
	return yaml.MarshalWithOptions(v, append(opts, Indent(spaces), IndentSequence(true), UseLiteralStyleIfMultiline(true))...)
}

func MustMarshalIndent(v any, opts ...EncodeOption) []byte {
	return x.Must1(MarshalIndent(v, 2, opts...))
}

func MustMarshalIndentString(v any, opts ...EncodeOption) string {
	return string(MustMarshalIndent(v, opts...))
}

func Unmarshal(bs []byte, v any, opts ...DecodeOption) error {
	return yaml.UnmarshalWithOptions(bs, v, opts...)
}

func UnmarshalString(s string, v any, opts ...DecodeOption) error {
	return yaml.UnmarshalWithOptions([]byte(s), v, opts...)
}

func UnmarshalTo[T any](bs []byte, opts ...DecodeOption) (m T, err error) {
	return m, Unmarshal(bs, &m, opts...)
}

func UnmarshalStringTo[T any](s string, opts ...DecodeOption) (m T, err error) {
	return m, UnmarshalString(s, &m, opts...)
}

func MustUnmarshalTo[T any](bs []byte, opts ...DecodeOption) (m T) {
	return x.Must1(m, Unmarshal(bs, &m, opts...))
}

func MustUnmarshalStringTo[T any](s string, opts ...DecodeOption) (m T) {
	return x.Must1(m, UnmarshalString(s, &m, opts...))
}
