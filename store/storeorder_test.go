// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package store_test

import (
	"context"
	"errors"
	"os"
	"testing"

	examplestainless "github.com/stainless-sdks/example-stainless-go"
	"github.com/stainless-sdks/example-stainless-go/internal/testutil"
	"github.com/stainless-sdks/example-stainless-go/option"
)

func TestStoreOrderGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := examplestainless.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.Order.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreOrderDeleteOrder(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := examplestainless.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Store.Order.DeleteOrder(context.TODO(), int64(0))
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
