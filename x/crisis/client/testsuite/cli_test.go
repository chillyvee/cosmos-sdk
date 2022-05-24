// +build norace

package testutil

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Stride-Labs/cosmos-sdk/testutil/network"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	suite.Run(t, NewIntegrationTestSuite(cfg))
}
