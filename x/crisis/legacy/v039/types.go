package v039

import sdk "github.com/Stride-Labs/cosmos-sdk/types"

const (
	ModuleName = "crisis"
)

type (
	GenesisState struct {
		ConstantFee sdk.Coin `json:"constant_fee" yaml:"constant_fee"`
	}
)
