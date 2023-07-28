package json

import (
	"github.com/loggie-io/loggie/pkg/util/json/gojson"
	"github.com/loggie-io/loggie/pkg/util/json/jsoniter"
	"github.com/loggie-io/loggie/pkg/util/json/sonic"
	"github.com/loggie-io/loggie/pkg/util/json/std"
)

const (
	Decoderjsoniter = "jsoniter"
	Decodersonic    = "sonic"
	Decoderstd      = "std"
	Decodergojson   = "go-json"
)

var decoderSet = map[string]JSON{
	Decoderjsoniter: &jsoniter.Jsoniter{},
	Decodersonic:    &sonic.Sonic{},
	Decoderstd:      &std.Std{},
	Decodergojson:   &gojson.Gojson{},
}

func init() {
	for name, decoder := range decoderSet {
		Register(name, decoder)
	}
}

type JSON interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
}

var JSONFactory = make(map[string]JSON)

func Register(name string, factory JSON) {
	JSONFactory[name] = factory
}

func SetDefaultEngine(name string) {
	JSONFactory[name] = JSONFactory["default"]
}

func Marshal(v interface{}) ([]byte, error) {
	return JSONFactory["default"].Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return JSONFactory["default"].Unmarshal(data, v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return JSONFactory["default"].MarshalIndent(v, prefix, indent)
}
