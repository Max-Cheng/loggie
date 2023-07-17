package gojson

import (
	json "github.com/goccy/go-json"
)

type Gojson struct{}

func (g *Gojson) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (g *Gojson) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
