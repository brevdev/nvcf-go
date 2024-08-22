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

func TestEnvelopeFunctionInvocationFunctionInvokeWithOptionalParams(t *testing.T) {
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
	_, err := client.EnvelopeFunctionInvocation.Functions.Invoke(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvidiacloudfunctions.EnvelopeFunctionInvocationFunctionInvokeParams{
			RequestBody: nvidiacloudfunctions.F[any](map[string]interface{}{}),
			RequestHeader: nvidiacloudfunctions.F(nvidiacloudfunctions.EnvelopeFunctionInvocationFunctionInvokeParamsRequestHeader{
				InputAssetReferences: nvidiacloudfunctions.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				MeteringData: nvidiacloudfunctions.F([]nvidiacloudfunctions.EnvelopeFunctionInvocationFunctionInvokeParamsRequestHeaderMeteringData{{
					Key:   nvidiacloudfunctions.F("key"),
					Value: nvidiacloudfunctions.F("value"),
				}, {
					Key:   nvidiacloudfunctions.F("key"),
					Value: nvidiacloudfunctions.F("value"),
				}, {
					Key:   nvidiacloudfunctions.F("key"),
					Value: nvidiacloudfunctions.F("value"),
				}}),
				PollDurationSeconds: nvidiacloudfunctions.F(int64(0)),
			}),
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
