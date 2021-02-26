package splitstore

import (
	"github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
)

type TrackingStore interface {
	Put(cid.Cid, abi.ChainEpoch) error
	PutBatch([]cid.Cid, abi.ChainEpoch) error
	Get(cid.Cid) (abi.ChainEpoch, error)
	Delete(cid.Cid) error
	ForEach(func(cid.Cid, abi.ChainEpoch) error) error
	Close() error
}