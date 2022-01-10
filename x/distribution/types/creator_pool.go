package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// zero fee pool
func InitialCreatorPool() CreatorPool {
	return CreatorPool{
		CreatorPool: sdk.DecCoins{},
	}
}

// ValidateGenesis validates the fee pool for a genesis state
func (f CreatorPool) ValidateGenesis() error {
	if f.CreatorPool.IsAnyNegative() {
		return fmt.Errorf("negative CreatorPool in distribution creator pool, is %v",
			f.CreatorPool)
	}

	return nil
}