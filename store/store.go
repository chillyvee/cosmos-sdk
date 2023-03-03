package store

import (
	"os"

	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/cache"
	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	"github.com/cosmos/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	// Developers may enable rootmulti logging by setting env LOG_ROOTMULTI=1
	if len(os.Getenv("LOG_ROOTMULTI")) > 0 {
		// Enable rootmulti logging
		tmlogger := log.NewTMLogger(os.Stderr).With("module", "rootmulti")
		return rootmulti.NewStore(db, tmlogger)
	} else {
		// Original design disables all rootmulti logging
		return rootmulti.NewStore(db, log.NewNopLogger())
	}
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
