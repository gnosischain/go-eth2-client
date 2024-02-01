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

package phase0

// Epoch is an epoch number.
type Epoch uint64

// CommitteeIndex is a committee index at a slot.
type CommitteeIndex uint64

// Version is a fork version.
type Version [4]byte

// DomainType is a domain type.
type DomainType [4]byte

// ForkDigest is a digest of fork data.
type ForkDigest [4]byte

// Domain is a signature domain.
type Domain [32]byte
