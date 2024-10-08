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

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/monerium/module-noble/v2/types/blacklist"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	blacklist.RegisterLegacyAminoCodec(cdc)

	cdc.RegisterConcrete(&MsgAcceptOwnership{}, "florin/AcceptOwnership", nil)
	cdc.RegisterConcrete(&MsgAddAdminAccount{}, "florin/AddAdminAccount", nil)
	cdc.RegisterConcrete(&MsgAddSystemAccount{}, "florin/AddSystemAccount", nil)
	cdc.RegisterConcrete(&MsgAllowDenom{}, "florin/AllowDenom", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "florin/Burn", nil)
	cdc.RegisterConcrete(&MsgMint{}, "florin/Mint", nil)
	cdc.RegisterConcrete(&MsgRecover{}, "florin/Recover", nil)
	cdc.RegisterConcrete(&MsgRemoveAdminAccount{}, "florin/RemoveAdminAccount", nil)
	cdc.RegisterConcrete(&MsgRemoveSystemAccount{}, "florin/RemoveSystemAccount", nil)
	cdc.RegisterConcrete(&MsgSetMaxMintAllowance{}, "florin/SetMaxMintAllowance", nil)
	cdc.RegisterConcrete(&MsgSetMintAllowance{}, "florin/SetMintAllowance", nil)
	cdc.RegisterConcrete(&MsgTransferOwnership{}, "florin/TransferOwnership", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	blacklist.RegisterInterfaces(registry)

	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgAcceptOwnership{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgAddAdminAccount{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgAddSystemAccount{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgAllowDenom{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgBurn{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgMint{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgRecover{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgRemoveAdminAccount{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgRemoveSystemAccount{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSetMaxMintAllowance{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSetMintAllowance{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgTransferOwnership{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var amino = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
