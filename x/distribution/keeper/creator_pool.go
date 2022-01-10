package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// DistributeFromCreatorPool distributes funds from the distribution module account to
// a receiver address while updating the creator pool
func (k Keeper) DistributeFromCreatorPool(ctx sdk.Context, amount sdk.Coins, receiveAddr sdk.AccAddress) error {
	creatorPool := k.GetCreatorPool(ctx)

	// NOTE the creator pool isn't a module account, however its coins
	// are held in the distribution module account. Thus the creator pool
	// must be reduced separately from the SendCoinsFromModuleToAccount call
	newPool, negative := creatorPool.CreatorPool.SafeSub(sdk.NewDecCoinsFromCoins(amount...))
	if negative {
		return types.ErrBadDistribution
	}

	creatorPool.CreatorPool = newPool

	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiveAddr, amount)
	if err != nil {
		return err
	}

	k.SetCreatorPool(ctx, creatorPool)
	return nil
}
