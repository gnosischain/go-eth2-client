// Copyright © 2023 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deneb

import (
	"bytes"
	"encoding/json"

	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

// signedBlindedBlobSidecarYAML is the spec representation of the struct.
type signedBlindedBlobSidecarYAML struct {
	Message   *BlindedBlobSidecar `json:"message"`
	Signature string              `json:"signature"`
}

// MarshalJSON implements json.Marshaler.
func (s *SignedBlindedBlobSidecar) MarshalYAML() ([]byte, error) {
	yamlBytes, err := yaml.MarshalWithOptions(&signedBlindedBlobSidecarYAML{
		Message:   s.Message,
		Signature: s.Signature.String(),
	}, yaml.Flow(true))
	if err != nil {
		return nil, err
	}
	return bytes.ReplaceAll(yamlBytes, []byte(`"`), []byte(`'`)), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *SignedBlindedBlobSidecar) UnmarshalYAML(input []byte) error {
	var data signedBlindedBlobSidecarYAML
	if err := yaml.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "failed to unmarshal YAML")
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "failed to marshal JSON")
	}

	return s.UnmarshalJSON(bytes)
}
