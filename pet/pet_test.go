// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pet_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	examplestainless "github.com/stainless-sdks/example-stainless-go"
	"github.com/stainless-sdks/example-stainless-go/internal/testutil"
	"github.com/stainless-sdks/example-stainless-go/option"
	"github.com/stainless-sdks/example-stainless-go/pet"
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
	_, err := client.Pets.New(context.TODO(), pet.PetNewParams{
		Pet: pet.PetParam{
			Name:      examplestainless.F("doggie"),
			PhotoURLs: examplestainless.F([]string{"string"}),
			ID:        examplestainless.F(int64(10)),
			Category: examplestainless.F(pet.CategoryParam{
				ID:   examplestainless.F(int64(1)),
				Name: examplestainless.F("Dogs"),
			}),
			Status: examplestainless.F(pet.PetStatusAvailable),
			Tags: examplestainless.F([]pet.PetTagParam{{
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
	_, err := client.Pets.Update(context.TODO(), pet.PetUpdateParams{
		Pet: pet.PetParam{
			Name:      examplestainless.F("doggie"),
			PhotoURLs: examplestainless.F([]string{"string"}),
			ID:        examplestainless.F(int64(10)),
			Category: examplestainless.F(pet.CategoryParam{
				ID:   examplestainless.F(int64(1)),
				Name: examplestainless.F("Dogs"),
			}),
			Status: examplestainless.F(pet.PetStatusAvailable),
			Tags: examplestainless.F([]pet.PetTagParam{{
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
	_, err := client.Pets.FindByStatus(context.TODO(), pet.PetFindByStatusParams{
		Status: examplestainless.F(pet.PetFindByStatusParamsStatusAvailable),
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
	_, err := client.Pets.FindByTags(context.TODO(), pet.PetFindByTagsParams{
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
		pet.PetUpdateByIDParams{
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
		pet.PetUploadImageParams{
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
