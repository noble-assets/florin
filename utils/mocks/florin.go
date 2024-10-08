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

package mocks

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/monerium/module-noble/v2"
	"github.com/monerium/module-noble/v2/keeper"
	"github.com/monerium/module-noble/v2/types"
)

func FlorinKeeper() (*keeper.Keeper, sdk.Context) {
	return FlorinWithKeepers(BankKeeper{})
}

func FlorinWithKeepers(bank BankKeeper) (*keeper.Keeper, sdk.Context) {
	key := storetypes.NewKVStoreKey(types.ModuleName)
	tkey := storetypes.NewTransientStoreKey("transient_florin")

	reg := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(reg)
	cdc := codec.NewProtoCodec(reg)

	k := keeper.NewKeeper(
		"authority",
		runtime.NewKVStoreService(key),
		runtime.ProvideEventService(),
		cdc,
		address.NewBech32Codec("noble"),
		bank,
	)

	bank = bank.WithSendCoinsRestriction(k.SendRestrictionFn)
	k.SetBankKeeper(bank)

	ctx := testutil.DefaultContext(key, tkey)
	florin.InitGenesis(ctx, k, *types.DefaultGenesisState())
	_ = k.MaxMintAllowance.Remove(ctx, "ueure")

	return k, ctx
}

//

func init() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("noble", "noblepub")
	config.Seal()
}
