package ocaero

import (
	"context"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

const (
	statusError = "ERROR"
	statusOK    = "OK"
)

// The following tags are applied to stats recorded by this package
var (
	// GoAeroInstanceName is the name of the aerospike instance.
	GoAeroInstanceName, _ = tag.NewKey("go_aerospike_instance_name")

	// GoAeroMethod is the client method called.
	GoAeroMethod = tag.MustNewKey("go_aerospike_method")

	// GoAeroStatus identifies the command status
	GoAeroStatus = tag.MustNewKey("go_aerospike_status")

	DefaultTags = []tag.Key{
		GoAeroMethod,
		GoAeroStatus,
	}
)

// The following measures are supported for use in custom views.
var (
	MeasureLatencyMs = stats.Int64("go.aerospike/latency", "The latency of calls in milliseconds", stats.UnitMilliseconds)
)

// Default distributions used by views in this package
var (
	DefaultMillisecondsDistribution = view.Distribution(
		0.0,
		0.001,
		0.005,
		0.01,
		0.05,
		0.1,
		0.5,
		1.0,
		1.5,
		2.0,
		2.5,
		5.0,
		10.0,
		25.0,
		50.0,
		100.0,
		200.0,
		400.0,
		600.0,
		800.0,
		1000.0,
		1500.0,
		2000.0,
		2500.0,
		5000.0,
		10000.0,
		20000.0,
		40000.0,
		100000.0,
		200000.0,
		500000.0,
	)
)

// Package ocaero provides some convenience views.
// You still need to register these views for data to actually be collected.
// You can use the RegisterAllViews function for this.
var (
	GoAerospikeLatencyView = &view.View{
		Name:        "go.aerospike/client/latency",
		Description: "The distribution of latency of various calls in milliseconds",
		Measure:     MeasureLatencyMs,
		Aggregation: DefaultMillisecondsDistribution,
		TagKeys:     DefaultTags,
	}

	GoAerospikeCallsView = &view.View{
		Name:        "go.aerospike/client/calls",
		Description: "The number of various calls of methods",
		Measure:     MeasureLatencyMs,
		Aggregation: view.Count(),
		TagKeys:     DefaultTags,
	}

	DefaultViews = []*view.View{GoAerospikeLatencyView, GoAerospikeCallsView}
)

// RegisterAllViews registers all the views to enable collection of stats
func RegisterAllViews() error {
	return view.Register(DefaultViews...)
}

func recordCall(ctx context.Context, method string, instanceName string) func(err error) {
	var startTime = time.Now()

	return func(err error) {
		var (
			timeSpentMs = time.Since(startTime).Milliseconds()
			tags        = []tag.Mutator{
				tag.Insert(GoAeroInstanceName, instanceName),
				tag.Insert(GoAeroMethod, method),
			}
		)

		if err != nil {
			tags = append(tags, tag.Insert(GoAeroStatus, statusError))
		} else {
			tags = append(tags, tag.Insert(GoAeroStatus, statusOK))
		}

		_ = stats.RecordWithTags(ctx, tags, MeasureLatencyMs.M(timeSpentMs))
	}
}
