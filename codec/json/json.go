package json

import (
	"bytes"

	"github.com/goccy/go-json"

	"github.com/starudream/go-lib/v2/x"
)

type (
	EncodeOption = json.EncodeOptionFunc
	DecodeOption = json.DecodeOptionFunc
)

var (
	NewEncoder = json.NewEncoder
	NewDecoder = json.NewDecoder
)

func Marshal(v any, opts ...EncodeOption) ([]byte, error) {
	return json.MarshalWithOption(v, opts...)
}

func MustMarshal(v any, opts ...EncodeOption) []byte {
	return x.Must1(Marshal(v, opts...))
}

func MustMarshalString(v any, opts ...EncodeOption) string {
	return string(MustMarshal(v, opts...))
}

func MarshalIndent(v any, prefix, indent string, opts ...EncodeOption) ([]byte, error) {
	return json.MarshalIndentWithOption(v, prefix, indent, opts...)
}

func MustMarshalIndent(v any, opts ...EncodeOption) []byte {
	return x.Must1(MarshalIndent(v, "", "  ", opts...))
}

func MustMarshalIndentString(v any, opts ...EncodeOption) string {
	return string(MustMarshalIndent(v, opts...))
}

func Unmarshal(bs []byte, v any, opts ...DecodeOption) error {
	return json.UnmarshalWithOption(bs, v, opts...)
}

func UnmarshalString(s string, v any, opts ...DecodeOption) error {
	return json.UnmarshalWithOption([]byte(s), v, opts...)
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

func Compact(src []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.Compact(buf, src)
	return buf.Bytes(), err
}

func MustCompact(src []byte) []byte {
	return x.Must1(Compact(src))
}

func Indent(src []byte, prefix, indent string) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := json.Indent(buf, src, prefix, indent)
	return buf.Bytes(), err
}

func MustIndent(src []byte) []byte {
	return x.Must1(Indent(src, "", "  "))
}

func Colored() EncodeOption {
	return json.Colorize(json.DefaultColorScheme)
}

func UnorderedMap() EncodeOption {
	return json.UnorderedMap()
}
