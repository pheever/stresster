package core

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

var unmarshallers = map[string]func([]byte, interface{}) error{
	".json": json.Unmarshal,
	".yaml": yaml.Unmarshal,
	".yml":  yaml.Unmarshal,
}
