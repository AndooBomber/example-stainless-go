// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package examplestainless_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/AndooBomber/example-stainless-go"
	"github.com/AndooBomber/example-stainless-go/internal/testutil"
	"github.com/AndooBomber/example-stainless-go/option"
)

func TestPetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.New(context.TODO(), examplestainless.PetNewParams{
		Pet: examplestainless.PetParam{
			Name:      examplestainless.F("doggie"),
			PhotoURLs: examplestainless.F([]string{"string"}),
			ID:        examplestainless.F(int64(10)),
			Category: examplestainless.F(examplestainless.CategoryParam{
				ID:   examplestainless.F(int64(1)),
				Name: examplestainless.F("Dogs"),
			}),
			First:  examplestainless.F("dogger"),
			Status: examplestainless.F(examplestainless.PetStatusAvailable),
			Tags: examplestainless.F([]examplestainless.PetTagParam{{
				ID:   examplestainless.F(int64(0)),
				Name: examplestainless.F("name"),
			}}),
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

func TestPetGet(t *testing.T) {
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
	_, err := client.Pets.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.Update(context.TODO(), examplestainless.PetUpdateParams{
		Pet: examplestainless.PetParam{
			Name:      examplestainless.F("doggie"),
			PhotoURLs: examplestainless.F([]string{"string"}),
			ID:        examplestainless.F(int64(10)),
			Category: examplestainless.F(examplestainless.CategoryParam{
				ID:   examplestainless.F(int64(1)),
				Name: examplestainless.F("Dogs"),
			}),
			First:  examplestainless.F("dogger"),
			Status: examplestainless.F(examplestainless.PetStatusAvailable),
			Tags: examplestainless.F([]examplestainless.PetTagParam{{
				ID:   examplestainless.F(int64(0)),
				Name: examplestainless.F("name"),
			}}),
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

func TestPetDelete(t *testing.T) {
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
	err := client.Pets.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByStatus(context.TODO(), examplestainless.PetFindByStatusParams{
		Status: examplestainless.F(examplestainless.PetFindByStatusParamsStatusAvailable),
	})
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByTags(context.TODO(), examplestainless.PetFindByTagsParams{
		Tags: examplestainless.F([]string{"string"}),
	})
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateByIDWithOptionalParams(t *testing.T) {
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
	err := client.Pets.UpdateByID(
		context.TODO(),
		int64(0),
		examplestainless.PetUpdateByIDParams{
			Name:   examplestainless.F("name"),
			Status: examplestainless.F("status"),
		},
	)
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUploadImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.UploadImage(
		context.TODO(),
		int64(0),
		examplestainless.PetUploadImageParams{
			Image:              io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			AdditionalMetadata: examplestainless.F("additionalMetadata"),
		},
	)
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
