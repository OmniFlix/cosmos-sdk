package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//nolint:interfacer
func NewGenesisState(
	params Params, fp FeePool, cp CreatorPool, dwis []DelegatorWithdrawInfo, pp sdk.ConsAddress, r []ValidatorOutstandingRewardsRecord,
	acc []ValidatorAccumulatedCommissionRecord, historical []ValidatorHistoricalRewardsRecord,
	cur []ValidatorCurrentRewardsRecord, dels []DelegatorStartingInfoRecord, slashes []ValidatorSlashEventRecord,
) *GenesisState {

	return &GenesisState{
		Params:                          params,
		FeePool:                         fp,
		CreatorPool:                     cp,
		DelegatorWithdrawInfos:          dwis,
		PreviousProposer:                pp.String(),
		OutstandingRewards:              r,
		ValidatorAccumulatedCommissions: acc,
		ValidatorHistoricalRewards:      historical,
		ValidatorCurrentRewards:         cur,
		DelegatorStartingInfos:          dels,
		ValidatorSlashEvents:            slashes,
	}
}

// get raw genesis raw message for testing
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		FeePool:                         InitialFeePool(),
		CreatorPool:                     InitialCreatorPool(),
		Params:                          DefaultParams(),
		DelegatorWithdrawInfos:          []DelegatorWithdrawInfo{},
		PreviousProposer:                "",
		OutstandingRewards:              []ValidatorOutstandingRewardsRecord{},
		ValidatorAccumulatedCommissions: []ValidatorAccumulatedCommissionRecord{},
		ValidatorHistoricalRewards:      []ValidatorHistoricalRewardsRecord{},
		ValidatorCurrentRewards:         []ValidatorCurrentRewardsRecord{},
		DelegatorStartingInfos:          []DelegatorStartingInfoRecord{},
		ValidatorSlashEvents:            []ValidatorSlashEventRecord{},
	}
}

// ValidateGenesis validates the genesis state of distribution genesis input
func ValidateGenesis(gs *GenesisState) error {
	if err := gs.Params.ValidateBasic(); err != nil {
		return err
	}
	if err := gs.FeePool.ValidateGenesis(); err != nil {
		return err
	}
	return gs.CreatorPool.ValidateGenesis()
}
