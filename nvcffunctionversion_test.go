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

func TestNvcfFunctionVersionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Nvcf.Functions.Versions.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvidiacloudfunctions.NvcfFunctionVersionNewParams{
			InferenceURL:  nvidiacloudfunctions.F("https://example.com"),
			Name:          nvidiacloudfunctions.F("x"),
			APIBodyFormat: nvidiacloudfunctions.F(nvidiacloudfunctions.NvcfFunctionVersionNewParamsAPIBodyFormatPredictV2),
			ContainerArgs: nvidiacloudfunctions.F("containerArgs"),
			ContainerEnvironment: nvidiacloudfunctions.F([]nvidiacloudfunctions.NvcfFunctionVersionNewParamsContainerEnvironment{{
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
			FunctionType:   nvidiacloudfunctions.F(nvidiacloudfunctions.NvcfFunctionVersionNewParamsFunctionTypeDefault),
			Health: nvidiacloudfunctions.F(nvidiacloudfunctions.NvcfFunctionVersionNewParamsHealth{
				ExpectedStatusCode: nvidiacloudfunctions.F(int64(0)),
				Port:               nvidiacloudfunctions.F(int64(0)),
				Protocol:           nvidiacloudfunctions.F(nvidiacloudfunctions.NvcfFunctionVersionNewParamsHealthProtocolHTTP),
				Timeout:            nvidiacloudfunctions.F("PT10S"),
				Uri:                nvidiacloudfunctions.F("https://example.com"),
			}),
			HealthUri:            nvidiacloudfunctions.F("https://example.com"),
			HelmChart:            nvidiacloudfunctions.F("https://example.com"),
			HelmChartServiceName: nvidiacloudfunctions.F("helmChartServiceName"),
			InferencePort:        nvidiacloudfunctions.F(int64(0)),
			Models: nvidiacloudfunctions.F([]nvidiacloudfunctions.NvcfFunctionVersionNewParamsModel{{
				Name:    nvidiacloudfunctions.F("name"),
				Uri:     nvidiacloudfunctions.F("https://example.com"),
				Version: nvidiacloudfunctions.F("version"),
			}}),
			Resources: nvidiacloudfunctions.F([]nvidiacloudfunctions.NvcfFunctionVersionNewParamsResource{{
				Name:    nvidiacloudfunctions.F("name"),
				Uri:     nvidiacloudfunctions.F("https://example.com"),
				Version: nvidiacloudfunctions.F("version"),
			}}),
			Secrets: nvidiacloudfunctions.F([]nvidiacloudfunctions.NvcfFunctionVersionNewParamsSecret{{
				Name:  nvidiacloudfunctions.F("x"),
				Value: nvidiacloudfunctions.F("x"),
			}}),
			Tags: nvidiacloudfunctions.F([]string{"string"}),
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

func TestNvcfFunctionVersionList(t *testing.T) {
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
	_, err := client.Nvcf.Functions.Versions.List(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *nvidiacloudfunctions.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
