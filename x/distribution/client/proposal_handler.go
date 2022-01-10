package client

import (
	"github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	"github.com/cosmos/cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	// CommunityPoolSpendProposalHandler is the community spend proposal handler.
	CommunityPoolSpendProposalHandler = govclient.NewProposalHandler(cli.GetCmdCommunityPoolSpendSubmitProposal, rest.ProposalRESTHandler)

	// CreatorPoolSpendProposalHandler is the creator spend proposal handler.
	CreatorPoolSpendProposalHandler = govclient.NewProposalHandler(cli.GetCmdCreatorPoolSpendSubmitProposal, rest.ProposalRESTHandler)
)
