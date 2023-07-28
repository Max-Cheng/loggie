/*
Copyright 2023 Loggie Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

func (g *Gojson) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (g *Gojson) MarshalToString(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
