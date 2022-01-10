package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

func TestValidateGenesisCreatorPool(t *testing.T) {

	fp := types.InitialCreatorPool()
	require.Nil(t, fp.ValidateGenesis())

	fp2 := types.CreatorPool{CreatorPool: sdk.DecCoins{{Denom: "stake", Amount: sdk.NewDec(-1)}}}
	require.NotNil(t, fp2.ValidateGenesis())
}