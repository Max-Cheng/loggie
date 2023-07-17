package sonic

import (
	json "github.com/bytedance/sonic"
)

type Sonic struct {
}

func (s *Sonic) Marshal(v interface{}) ([]byte, error) {
	return json.ConfigFastest.Marshal(v)
}

func (s *Sonic) Unmarshal(data []byte, v interface{}) error {
	return json.ConfigFastest.Unmarshal(data, v)
}
