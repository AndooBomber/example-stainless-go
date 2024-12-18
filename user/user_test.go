// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_test

import (
	"context"
	"errors"
	"os"
	"testing"

	examplestainless "github.com/stainless-sdks/example-stainless-go"
	"github.com/stainless-sdks/example-stainless-go/internal/testutil"
	"github.com/stainless-sdks/example-stainless-go/option"
	"github.com/stainless-sdks/example-stainless-go/user"
)

func TestUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.New(context.TODO(), user.UserNewParams{
		User: user.UserParam{
			ID:         examplestainless.F(int64(10)),
			Email:      examplestainless.F("john@email.com"),
			FirstName:  examplestainless.F("John"),
			LastName:   examplestainless.F("James"),
			Password:   examplestainless.F("12345"),
			Phone:      examplestainless.F("12345"),
			Username:   examplestainless.F("theUser"),
			UserStatus: examplestainless.F(int64(1)),
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

func TestUserGet(t *testing.T) {
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
	_, err := client.User.Get(context.TODO(), "username")
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	err := client.User.Update(
		context.TODO(),
		"username",
		user.UserUpdateParams{
			User: user.UserParam{
				ID:         examplestainless.F(int64(10)),
				Email:      examplestainless.F("john@email.com"),
				FirstName:  examplestainless.F("John"),
				LastName:   examplestainless.F("James"),
				Password:   examplestainless.F("12345"),
				Phone:      examplestainless.F("12345"),
				Username:   examplestainless.F("theUser"),
				UserStatus: examplestainless.F(int64(1)),
			},
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

func TestUserDelete(t *testing.T) {
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
	err := client.User.Delete(context.TODO(), "username")
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithList(t *testing.T) {
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
	_, err := client.User.NewWithList(context.TODO(), user.UserNewWithListParams{
		Items: []user.UserParam{{
			ID:         examplestainless.F(int64(10)),
			Email:      examplestainless.F("john@email.com"),
			FirstName:  examplestainless.F("John"),
			LastName:   examplestainless.F("James"),
			Password:   examplestainless.F("12345"),
			Phone:      examplestainless.F("12345"),
			Username:   examplestainless.F("theUser"),
			UserStatus: examplestainless.F(int64(1)),
		}},
	})
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Login(context.TODO(), user.UserLoginParams{
		Password: examplestainless.F("password"),
		Username: examplestainless.F("username"),
	})
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
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
	err := client.User.Logout(context.TODO())
	if err != nil {
		var apierr *examplestainless.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
