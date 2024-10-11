// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiacloudfunctions_test

import (
	"context"
	"os"
	"testing"

	"github.com/brevdev/nvcf-go"
	"github.com/brevdev/nvcf-go/internal/testutil"
	"github.com/brevdev/nvcf-go/option"
)

func TestUsage(t *testing.T) {
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
	createFunctionResponse, err := client.Nvcf.Functions.New(context.TODO(), nvidiacloudfunctions.NvcfFunctionNewParams{
		InferenceURL: nvidiacloudfunctions.F("https://example.com"),
		Name:         nvidiacloudfunctions.F("x"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", createFunctionResponse.Function)
}
