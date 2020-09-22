// Copyright © 2020 Attestant Limited.
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

package prysmgrpc_test

import (
	"context"
	"os"
	"testing"

	client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/prysmgrpc"
	"github.com/stretchr/testify/require"
)

func TestProposerDuties(t *testing.T) {
	tests := []struct {
		name       string
		epoch      uint64
		validators []client.ValidatorIDProvider
		expected   int
	}{
		{
			name:     "Old",
			epoch:    1,
			expected: 32,
		},
		{
			name:     "Current",
			epoch:    10989,
			expected: 32,
		},
		{
			name:  "GoodWithValidators",
			epoch: 10758,
			validators: []client.ValidatorIDProvider{
				&testValidatorIDProvider{
					index:  46752,
					pubKey: "0x93972e226623bd0c19d98d7774bb5da8230d37d3e629352d969fa7b76d03a5960c6f7e494b186940d3e4c8464d24d74a",
				},
			},
			expected: 1,
		},
	}

	service, err := prysmgrpc.New(context.Background(), prysmgrpc.WithAddress(os.Getenv("PRYSMGRPC_ADDRESS")))
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			duties, err := service.ProposerDuties(context.Background(), test.epoch, test.validators)
			require.NoError(t, err)
			require.NotNil(t, duties)
			require.Equal(t, test.expected, len(duties))
		})
	}
}