package ocaero_test

import (
	"context"
	"log"

	"github.com/Kochava/ocaero"
	"github.com/aerospike/aerospike-client-go"
	"go.opencensus.io/tag"
)

// This example shows how to use ocaero.Wrapper and register views with OpenCensus
func Example() {
	var (
		aeroClient    *aerospike.Client // TODO actually make the client connection
		ocAeroWrapper = ocaero.Wrap(aeroClient, "my-aero")
	)

	ocaero.RegisterAllViews()

	data, err := getData(context.Background(), ocAeroWrapper)

	if err != nil {
		log.Printf("Unable to query: %s", err.Error())
		return
	}

	log.Println("Got data:", data)

	data = data + " updated"

	err = setData(context.Background(), ocAeroWrapper)

	if err != nil {
		log.Printf("Unable to update data: %s", err.Error())
		return
	}

	log.Println("Updated data")
}

// This example shows how to add custom tags to existing ocaero views
func Example_customTags() {
	var (
		aeroClient    *aerospike.Client // TODO actually make the client connection
		ocAeroWrapper = ocaero.Wrap(aeroClient, "my-aero")
	)

	// appVersionTag is a tag represening the current version of the application
	appVersionTag, _ := tag.NewKey("app_version")

	// add the tag to the views
	ocaero.GoAerospikeLatencyView.TagKeys = append(
		ocaero.GoAerospikeLatencyView.TagKeys,
		appVersionTag,
	)

	ocaero.GoAerospikeCallsView.TagKeys = append(
		ocaero.GoAerospikeCallsView.TagKeys,
		appVersionTag,
	)

	// adding the tag to the application context will make the tag/value
	// available to any view its been added to
	ctx, _ := tag.New(context.Background(), tag.Insert(
		appVersionTag,
		"v1.0.0",
	))

	ocaero.RegisterAllViews()

	data, err := getData(ctx, ocAeroWrapper)

	if err != nil {
		log.Printf("Unable to query: %s", err.Error())
		return
	}

	log.Println("Got data:", data)

	data = data + " updated"

	err = setData(context.Background(), ocAeroWrapper)

	if err != nil {
		log.Printf("Unable to update data: %s", err.Error())
		return
	}

	log.Println("Updated data")
}

func setData(ctx context.Context, aeroClient ocaero.Client) error {
	return nil
}

func getData(ctx context.Context, aeroClient ocaero.Client) (string, error) {
	return "", nil
}
