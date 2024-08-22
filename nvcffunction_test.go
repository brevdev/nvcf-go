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

func TestNvcfFunctionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Nvcf.Functions.New(context.TODO(), nvidiacloudfunctions.NvcfFunctionNewParams{
		InferenceURL:  nvidiacloudfunctions.F("https://example.com"),
		Name:          nvidiacloudfunctions.F("x"),
		APIBodyFormat: nvidiacloudfunctions.F(nvidiacloudfunctions.NvcfFunctionNewParamsAPIBodyFormatPredictV2),
		ContainerArgs: nvidiacloudfunctions.F("containerArgs"),
		ContainerEnvironment: nvidiacloudfunctions.F([]nvidiacloudfunctions.NvcfFunctionNewParamsContainerEnvironment{{
			Key:   nvidiacloudfunctions.F("key"),
			Value: nvidiacloudfunctions.F("value"),
		}, {
			Key:   nvidiacloudfunctions.F("key"),
			Value: nvidiacloudfunctions.F("value"),
		}, {
			Key:   nvidiacloudfunctions.F("key"),
			Value: nvidiacloudfunctions.F("value"),
		}}),
		ContainerImage: nvidiacloudfunctions.F("https://example.com"),
		Description:    nvidiacloudfunctions.F("description"),
		FunctionType:   nvidiacloudfunctions.F(nvidiacloudfunctions.NvcfFunctionNewParamsFunctionTypeDefault),
		Health: nvidiacloudfunctions.F(nvidiacloudfunctions.NvcfFunctionNewParamsHealth{
			ExpectedStatusCode: nvidiacloudfunctions.F(int64(0)),
			Port:               nvidiacloudfunctions.F(int64(0)),
			Protocol:           nvidiacloudfunctions.F(nvidiacloudfunctions.NvcfFunctionNewParamsHealthProtocolHTTP),
			Timeout:            nvidiacloudfunctions.F("PT10S"),
			Uri:                nvidiacloudfunctions.F("https://example.com"),
		}),
		HealthUri:            nvidiacloudfunctions.F("https://example.com"),
		HelmChart:            nvidiacloudfunctions.F("https://example.com"),
		HelmChartServiceName: nvidiacloudfunctions.F("helmChartServiceName"),
		InferencePort:        nvidiacloudfunctions.F(int64(0)),
		Models: nvidiacloudfunctions.F([]nvidiacloudfunctions.NvcfFunctionNewParamsModel{{
			Name:    nvidiacloudfunctions.F("name"),
			Uri:     nvidiacloudfunctions.F("https://example.com"),
			Version: nvidiacloudfunctions.F("version"),
		}}),
		Resources: nvidiacloudfunctions.F([]nvidiacloudfunctions.NvcfFunctionNewParamsResource{{
			Name:    nvidiacloudfunctions.F("name"),
			Uri:     nvidiacloudfunctions.F("https://example.com"),
			Version: nvidiacloudfunctions.F("version"),
		}}),
		Secrets: nvidiacloudfunctions.F([]nvidiacloudfunctions.NvcfFunctionNewParamsSecret{{
			Name:  nvidiacloudfunctions.F("x"),
			Value: nvidiacloudfunctions.F("x"),
		}}),
		Tags: nvidiacloudfunctions.F([]string{"string"}),
	})
	if err != nil {
		var apierr *nvidiacloudfunctions.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNvcfFunctionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Nvcf.Functions.List(context.TODO(), nvidiacloudfunctions.NvcfFunctionListParams{
		Visibility: nvidiacloudfunctions.F([]nvidiacloudfunctions.NvcfFunctionListParamsVisibility{nvidiacloudfunctions.NvcfFunctionListParamsVisibilityAuthorized}),
	})
	if err != nil {
		var apierr *nvidiacloudfunctions.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
