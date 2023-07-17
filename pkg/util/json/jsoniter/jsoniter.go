package jsoniter

import (
	json "github.com/json-iterator/go"
)

type Jsoniter struct {
}

func (j *Jsoniter) Marshal(v interface{}) ([]byte, error) {
	return json.ConfigFastest.Marshal(v)
}

func (j *Jsoniter) Unmarshal(data []byte, v interface{}) error {
	return json.ConfigFastest.Unmarshal(data, v)
}
