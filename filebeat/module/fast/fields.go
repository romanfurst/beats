// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package fast

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "fast", asset.ModuleFieldsPri, AssetFast); err != nil {
		panic(err)
	}
}

// AssetFast returns asset data.
// This is the base64 encoded gzipped contents of module/fast.
func AssetFast() string {
	return "eJycjj2OgzAQhXuf4okeDuBiy+32ENYyOE7GP/KYSNw+AhIEyFGkTOd5nu97LW40aQxGigKKK0wazfxsFNCT/GeXiotB40cBWH7iL/YjkwIGR9yLXpIWwXjaWPOUKZGGzXFMz02FeKTsSRzttqvB3gLX+Z2rcrRgF0i6XXYWHqR0Jz4kL7VhZ+SUJFMuGleJoasdemezWbuVPFJV6EnEWPpG6cV+ED4CAAD//w/xeu0="
}
