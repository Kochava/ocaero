package ocaero

import (
	"context"

	as "github.com/aerospike/aerospike-client-go"
)

// ensure Wrapper implements Client
var _ Client = &Wrapper{}

// Wrap uses an Aerospike Client to make a Wrapper
func Wrap(aeroClient *as.Client, instanceName string) *Wrapper {
	return &Wrapper{Client: aeroClient, instanceName: instanceName}
}

// Wrapper reports method call latency with OpenCensus
//
// Contexts
//
// Contexts are supplied Wrapper methods but aren't supported by the underlying
// Aerospike client. Instead they are supplied to the OpenCensus metrics.
type Wrapper struct {
	*as.Client
	instanceName string
}

// Query routes requests to the aerospike client and reports latency to OpenCensus metrics
func (wrapper *Wrapper) Query(ctx context.Context, policy *as.QueryPolicy, stmt *as.Statement) (recordSet *as.Recordset, err error) {
	var recordCallFunc = recordCall(ctx, "go.aerospike.query", wrapper.instanceName)
	defer func() {
		recordCallFunc(err)
	}()

	recordSet, err = wrapper.Client.Query(policy, stmt)
	return recordSet, err
}

// PutBins routes requests to the aerospike client and reports latency to OpenCensus metrics
func (wrapper *Wrapper) PutBins(ctx context.Context, policy *as.WritePolicy, key *as.Key, bins ...*as.Bin) (err error) {
	var recordCallFunc = recordCall(ctx, "go.aerospike.put_bins", wrapper.instanceName)
	defer func() {
		recordCallFunc(err)
	}()

	err = wrapper.Client.PutBins(policy, key, bins...)
	return err
}
