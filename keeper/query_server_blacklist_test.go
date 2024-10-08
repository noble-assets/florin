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

package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/monerium/module-noble/v2/keeper"
	"github.com/monerium/module-noble/v2/types/blacklist"
	"github.com/monerium/module-noble/v2/utils"
	"github.com/monerium/module-noble/v2/utils/mocks"
	"github.com/stretchr/testify/require"
)

func TestBlacklistOwnerQuery(t *testing.T) {
	k, ctx := mocks.FlorinKeeper()
	server := keeper.NewBlacklistQueryServer(k)

	// ACT: Attempt to query owner with invalid request.
	_, err := server.Owner(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ARRANGE: Set owner in state.
	owner := utils.TestAccount()
	err = k.SetBlacklistOwner(ctx, owner.Address)
	require.NoError(t, err)

	// ACT: Attempt to query owner.
	res, err := server.Owner(ctx, &blacklist.QueryOwner{})
	// ASSERT: The query should've succeeded, with empty pending owner.
	require.NoError(t, err)
	require.Equal(t, owner.Address, res.Owner)
	require.Empty(t, res.PendingOwner)

	// ARRANGE: Set pending owner in state.
	pendingOwner := utils.TestAccount()
	err = k.SetBlacklistPendingOwner(ctx, pendingOwner.Address)
	require.NoError(t, err)

	// ACT: Attempt to query owner.
	res, err = server.Owner(ctx, &blacklist.QueryOwner{})
	// ASSERT: The query should've succeeded, with pending owner.
	require.NoError(t, err)
	require.Equal(t, owner.Address, res.Owner)
	require.Equal(t, pendingOwner.Address, res.PendingOwner)
}

func TestBlacklistAdminsQuery(t *testing.T) {
	k, ctx := mocks.FlorinKeeper()
	server := keeper.NewBlacklistQueryServer(k)

	// ACT: Attempt to query admins with invalid request.
	_, err := server.Admins(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query admins with no state.
	res, err := server.Admins(ctx, &blacklist.QueryAdmins{})
	// ASSERT: The query should've succeeded, returns empty.
	require.NoError(t, err)
	require.Empty(t, res.Admins)

	// ARRANGE: Set admin accounts in state.
	admin1, admin2 := utils.TestAccount(), utils.TestAccount()
	err = k.SetBlacklistAdmin(ctx, admin1.Address)
	require.NoError(t, err)
	err = k.SetBlacklistAdmin(ctx, admin2.Address)
	require.NoError(t, err)

	// ACT: Attempt to query admins.
	res, err = server.Admins(ctx, &blacklist.QueryAdmins{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Len(t, res.Admins, 2)
	require.Contains(t, res.Admins, admin1.Address)
	require.Contains(t, res.Admins, admin2.Address)
}

func TestBlacklistAdversariesQuery(t *testing.T) {
	k, ctx := mocks.FlorinKeeper()
	server := keeper.NewBlacklistQueryServer(k)

	// ACT: Attempt to query adversaries with invalid request.
	_, err := server.Adversaries(ctx, nil)
	// ASSERT: The query should've failed due to invalid request.
	require.ErrorContains(t, err, errors.ErrInvalidRequest.Error())

	// ACT: Attempt to query adversaries with no state.
	res, err := server.Adversaries(ctx, &blacklist.QueryAdversaries{})
	// ASSERT: The query should've succeeded, returns empty.
	require.NoError(t, err)
	require.Empty(t, res.Adversaries)

	// ARRANGE: Set adversaries in state.
	alice, bob := utils.TestAccount(), utils.TestAccount()
	err = k.SetAdversary(ctx, alice.Address)
	require.NoError(t, err)
	err = k.SetAdversary(ctx, bob.Address)
	require.NoError(t, err)

	// ACT: Attempt to query adversaries.
	res, err = server.Adversaries(ctx, &blacklist.QueryAdversaries{})
	// ASSERT: The query should've succeeded.
	require.NoError(t, err)
	require.Len(t, res.Adversaries, 2)
	require.Contains(t, res.Adversaries, alice.Address)
	require.Contains(t, res.Adversaries, bob.Address)
}
