// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package examplestainless_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/AndooBomber/example-stainless-go"
	"github.com/AndooBomber/example-stainless-go/internal/testutil"
	"github.com/AndooBomber/example-stainless-go/option"
	"github.com/AndooBomber/example-stainless-go/shared"
)

func TestStoreNewOrderWithOptionalParams(t *testing.T) {
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
	_, err := client.Store.NewOrder(context.TODO(), examplestainless.StoreNewOrderParams{
		Order: shared.OrderParam{
			ID:       examplestainless.F(int64(10)),
			Complete: examplestainless.F(true),
			PetID:    examplestainless.F(int64(198772)),
			Quantity: examplestainless.F(int64(7)),
			ShipDate: examplestainless.F(time.Now()),
			Status:   examplestainless.F(shared.OrderStatusPlaced),
		},
	})
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreInventory(t *testing.T) {
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
	_, err := client.Store.Inventory(context.TODO())
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
