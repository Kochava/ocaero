# ocaero

OpenCensus Aerospike client wrapper for Go

## Contexts

The contexts supplied to `ocaero` client calls are used to supply tags
to Open Census measurements. Canceling a context will not cause the underlying
aerospike calls to be cancelled.

## Migration

To interact with this wrapper instead of the regular aerospike client, call `ocaero.Wrap()` and supply a context to client method calls.

```go
var (
    ctx = context.Background()
    ocaeroClient = ocaero.Wrap(aeroClient, "my-aero")
    policy = as.NewQueryPolicy()
    statement = as.NewStatement("my-namespace", "my-set")
)

ocaeroClient.Query(ctx, policy, statement)
```

## TODO

- Add remaining Aerospike client functions
- Add Open Census tracing