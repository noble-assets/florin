// Copyright 2024 Monerium ehf.
//
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

package utils

import (
	"github.com/cometbft/cometbft/crypto/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Account struct {
	Address string
	Invalid string
	Bytes   []byte
}

func TestAccount() Account {
	bytes := secp256k1.GenPrivKey().PubKey().Address().Bytes()
	address, _ := sdk.Bech32ifyAddressBytes("noble", bytes)
	invalid, _ := sdk.Bech32ifyAddressBytes("cosmos", bytes)

	return Account{
		Address: address,
		Invalid: invalid,
		Bytes:   bytes,
	}
}
