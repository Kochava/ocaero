package ocaero

import (
	"context"

	as "github.com/aerospike/aerospike-client-go"
)

// PutBinser allows a caller to set bins in Aerospike
type PutBinser interface {
	PutBins(ctx context.Context, policy *as.WritePolicy, key *as.Key, bins ...*as.Bin) error
}
