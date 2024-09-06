// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiacloudfunctions_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/brevdev/nvcf-go"
	"github.com/brevdev/nvcf-go/internal/testutil"
	"github.com/brevdev/nvcf-go/option"
)

func TestUserSecretManagementFunctionVersionUpdate(t *testing.T) {
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
	err := client.UserSecretManagement.Functions.Versions.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvidiacloudfunctions.UserSecretManagementFunctionVersionUpdateParams{
			Secrets: nvidiacloudfunctions.F([]nvidiacloudfunctions.UserSecretManagementFunctionVersionUpdateParamsSecret{{
				Name:  nvidiacloudfunctions.F("x"),
				Value: nvidiacloudfunctions.F("x"),
			}}),
		},
	)
	if err != nil {
		var apierr *nvidiacloudfunctions.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
