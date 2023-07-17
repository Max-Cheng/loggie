package sonic

import (
	"errors"
)

var (
	ErrNotSupportArch = errors.New("not support arch")
)

type Sonic struct {
}

func (s *Sonic) Marshal(v interface{}) ([]byte, error) {
	return nil, ErrNotSupportArch
}

func (s *Sonic) Unmarshal(data []byte, v interface{}) error {
	return ErrNotSupportArch
}
