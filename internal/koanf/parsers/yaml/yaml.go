package yaml

import (
	"github.com/starudream/go-lib/v2/codec/yaml"
)

type YAML struct{}

func Parser() *YAML {
	return &YAML{}
}

func (p *YAML) Unmarshal(b []byte) (map[string]interface{}, error) {
	var out map[string]interface{}
	if err := yaml.Unmarshal(b, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (p *YAML) Marshal(o map[string]interface{}) ([]byte, error) {
	return yaml.Marshal(o)
}
