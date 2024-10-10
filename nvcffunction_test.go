// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/brevdev/nvcf-go"
	"github.com/brevdev/nvcf-go/internal/testutil"
	"github.com/brevdev/nvcf-go/option"
)

func TestNVCFFunctionNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nvcf.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.NVCF.Functions.New(context.TODO(), nvcf.NVCFFunctionNewParams{
		InferenceURL:  nvcf.F("https://example.com"),
		Name:          nvcf.F("x"),
		APIBodyFormat: nvcf.F(nvcf.NVCFFunctionNewParamsAPIBodyFormatPredictV2),
		ContainerArgs: nvcf.F("containerArgs"),
		ContainerEnvironment: nvcf.F([]nvcf.NVCFFunctionNewParamsContainerEnvironment{{
			Key:   nvcf.F("key"),
			Value: nvcf.F("value"),
		}, {
			Key:   nvcf.F("key"),
			Value: nvcf.F("value"),
		}, {
			Key:   nvcf.F("key"),
			Value: nvcf.F("value"),
		}}),
		ContainerImage: nvcf.F("https://example.com"),
		Description:    nvcf.F("description"),
		FunctionType:   nvcf.F(nvcf.NVCFFunctionNewParamsFunctionTypeDefault),
		Health: nvcf.F(nvcf.NVCFFunctionNewParamsHealth{
			ExpectedStatusCode: nvcf.F(int64(0)),
			Port:               nvcf.F(int64(0)),
			Protocol:           nvcf.F(nvcf.NVCFFunctionNewParamsHealthProtocolHTTP),
			Timeout:            nvcf.F("PT10S"),
			Uri:                nvcf.F("https://example.com"),
		}),
		HealthUri:            nvcf.F("https://example.com"),
		HelmChart:            nvcf.F("https://example.com"),
		HelmChartServiceName: nvcf.F("helmChartServiceName"),
		InferencePort:        nvcf.F(int64(0)),
		Models: nvcf.F([]nvcf.NVCFFunctionNewParamsModel{{
			Name:    nvcf.F("name"),
			Uri:     nvcf.F("https://example.com"),
			Version: nvcf.F("version"),
		}}),
		Resources: nvcf.F([]nvcf.NVCFFunctionNewParamsResource{{
			Name:    nvcf.F("name"),
			Uri:     nvcf.F("https://example.com"),
			Version: nvcf.F("version"),
		}}),
		Secrets: nvcf.F([]nvcf.NVCFFunctionNewParamsSecret{{
			Name:  nvcf.F("x"),
			Value: nvcf.F("x"),
		}}),
		Tags: nvcf.F([]string{"string"}),
	})
	if err != nil {
		var apierr *nvcf.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNVCFFunctionListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nvcf.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.NVCF.Functions.List(context.TODO(), nvcf.NVCFFunctionListParams{
		Visibility: nvcf.F([]nvcf.NVCFFunctionListParamsVisibility{nvcf.NVCFFunctionListParamsVisibilityAuthorized}),
	})
	if err != nil {
		var apierr *nvcf.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
