// Package ocaero supplies a wrapper for the aerospile go client to report metrics with OpenCensus
package ocaero

// Client defines all the available Aerospike interactions
type Client interface {
	Querier
	PutBinser
}
