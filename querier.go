package ocaero

import (
	"context"

	as "github.com/aerospike/aerospike-client-go"
)

// Querier allows a caller to perform an Aerospike query
type Querier interface {
	Query(ctx context.Context, policy *as.QueryPolicy, stmt *as.Statement) (*as.Recordset, error)
}
