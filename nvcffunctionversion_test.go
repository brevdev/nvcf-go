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

func TestNVCFFunctionVersionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.NVCF.Functions.Versions.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvcf.NVCFFunctionVersionNewParams{
			InferenceURL:  nvcf.F("https://example.com"),
			Name:          nvcf.F("x"),
			APIBodyFormat: nvcf.F(nvcf.NVCFFunctionVersionNewParamsAPIBodyFormatPredictV2),
			ContainerArgs: nvcf.F("containerArgs"),
			ContainerEnvironment: nvcf.F([]nvcf.NVCFFunctionVersionNewParamsContainerEnvironment{{
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
			FunctionType:   nvcf.F(nvcf.NVCFFunctionVersionNewParamsFunctionTypeDefault),
			Health: nvcf.F(nvcf.NVCFFunctionVersionNewParamsHealth{
				ExpectedStatusCode: nvcf.F(int64(0)),
				Port:               nvcf.F(int64(0)),
				Protocol:           nvcf.F(nvcf.NVCFFunctionVersionNewParamsHealthProtocolHTTP),
				Timeout:            nvcf.F("PT10S"),
				Uri:                nvcf.F("https://example.com"),
			}),
			HealthUri:            nvcf.F("https://example.com"),
			HelmChart:            nvcf.F("https://example.com"),
			HelmChartServiceName: nvcf.F("helmChartServiceName"),
			InferencePort:        nvcf.F(int64(0)),
			Models: nvcf.F([]nvcf.NVCFFunctionVersionNewParamsModel{{
				Name:    nvcf.F("name"),
				Uri:     nvcf.F("https://example.com"),
				Version: nvcf.F("version"),
			}}),
			Resources: nvcf.F([]nvcf.NVCFFunctionVersionNewParamsResource{{
				Name:    nvcf.F("name"),
				Uri:     nvcf.F("https://example.com"),
				Version: nvcf.F("version"),
			}}),
			Secrets: nvcf.F([]nvcf.NVCFFunctionVersionNewParamsSecret{{
				Name:  nvcf.F("x"),
				Value: nvcf.F("x"),
			}}),
			Tags: nvcf.F([]string{"string"}),
		},
	)
	if err != nil {
		var apierr *nvcf.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNVCFFunctionVersionList(t *testing.T) {
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
	_, err := client.NVCF.Functions.Versions.List(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *nvcf.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
