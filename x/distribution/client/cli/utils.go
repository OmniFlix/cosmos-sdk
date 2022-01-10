package cli

import (
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// ParseCommunityPoolSpendProposalWithDeposit reads and parses a CommunityPoolSpendProposalWithDeposit from a file.
func ParseCommunityPoolSpendProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (types.CommunityPoolSpendProposalWithDeposit, error) {
	proposal := types.CommunityPoolSpendProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// ParseCreatorPoolSpendProposalWithDeposit reads and parses a CreatorPoolSpendProposalWithDeposit from a file.
func ParseCreatorPoolSpendProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (types.CreatorPoolSpendProposalWithDeposit, error) {
	proposal := types.CreatorPoolSpendProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
