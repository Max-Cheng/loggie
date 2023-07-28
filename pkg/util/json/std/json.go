package std

import (
	"encoding/json"
)

type Std struct {
}

func (s *Std) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (s *Std) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (s *Std) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}
