// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiacloudfunctions_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/nvidia-cloud-functions-go"
	"github.com/stainless-sdks/nvidia-cloud-functions-go/internal/testutil"
	"github.com/stainless-sdks/nvidia-cloud-functions-go/option"
)

func TestAssetNew(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nvidiacloudfunctions.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Assets.New(context.TODO(), nvidiacloudfunctions.AssetNewParams{
		ContentType: nvidiacloudfunctions.F("contentType"),
		Description: nvidiacloudfunctions.F("description"),
	})
	if err != nil {
		var apierr *nvidiacloudfunctions.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssetList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nvidiacloudfunctions.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Assets.List(context.TODO())
	if err != nil {
		var apierr *nvidiacloudfunctions.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
