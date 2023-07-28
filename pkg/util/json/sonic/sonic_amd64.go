package sonic

import (
	json "github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
)

type Sonic struct {
}

func (s *Sonic) Marshal(v interface{}) ([]byte, error) {
	return json.ConfigFastest.Marshal(v)
}

func (s *Sonic) Unmarshal(data []byte, v interface{}) error {
	return json.ConfigFastest.Unmarshal(data, v)
}

func (s *Sonic) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	// SIMD json marshal not support this feature
	return jsoniter.ConfigFastest.MarshalIndent(v, prefix, indent)
}
