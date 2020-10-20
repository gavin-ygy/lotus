// +build !debug
// +build !2k
// +build !testground

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

const UpgradeActorsV2Height = 30

const UpgradeTapeHeight = 45

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 22400

// TODO: Confirm epoch
const UpgradeKumquatHeight = 22450

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true
}

const BlockDelaySecs = uint64(builtin0.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)