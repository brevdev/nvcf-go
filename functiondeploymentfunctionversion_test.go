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

func TestFunctionDeploymentFunctionVersionNew(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvidiacloudfunctions.FunctionDeploymentFunctionVersionNewParams{
			DeploymentSpecifications: nvidiacloudfunctions.F([]nvidiacloudfunctions.FunctionDeploymentFunctionVersionNewParamsDeploymentSpecification{{
				GPU:                   nvidiacloudfunctions.F("gpu"),
				InstanceType:          nvidiacloudfunctions.F("instanceType"),
				MaxInstances:          nvidiacloudfunctions.F(int64(0)),
				MinInstances:          nvidiacloudfunctions.F(int64(0)),
				Attributes:            nvidiacloudfunctions.F([]string{"string"}),
				AvailabilityZones:     nvidiacloudfunctions.F([]string{"string", "string", "string"}),
				Backend:               nvidiacloudfunctions.F("backend"),
				Clusters:              nvidiacloudfunctions.F([]string{"string"}),
				Configuration:         nvidiacloudfunctions.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvidiacloudfunctions.F(int64(1)),
				PreferredOrder:        nvidiacloudfunctions.F(int64(1)),
				Regions:               nvidiacloudfunctions.F([]string{"string"}),
			}, {
				GPU:                   nvidiacloudfunctions.F("gpu"),
				InstanceType:          nvidiacloudfunctions.F("instanceType"),
				MaxInstances:          nvidiacloudfunctions.F(int64(0)),
				MinInstances:          nvidiacloudfunctions.F(int64(0)),
				Attributes:            nvidiacloudfunctions.F([]string{"string"}),
				AvailabilityZones:     nvidiacloudfunctions.F([]string{"string", "string", "string"}),
				Backend:               nvidiacloudfunctions.F("backend"),
				Clusters:              nvidiacloudfunctions.F([]string{"string"}),
				Configuration:         nvidiacloudfunctions.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvidiacloudfunctions.F(int64(1)),
				PreferredOrder:        nvidiacloudfunctions.F(int64(1)),
				Regions:               nvidiacloudfunctions.F([]string{"string"}),
			}, {
				GPU:                   nvidiacloudfunctions.F("gpu"),
				InstanceType:          nvidiacloudfunctions.F("instanceType"),
				MaxInstances:          nvidiacloudfunctions.F(int64(0)),
				MinInstances:          nvidiacloudfunctions.F(int64(0)),
				Attributes:            nvidiacloudfunctions.F([]string{"string"}),
				AvailabilityZones:     nvidiacloudfunctions.F([]string{"string", "string", "string"}),
				Backend:               nvidiacloudfunctions.F("backend"),
				Clusters:              nvidiacloudfunctions.F([]string{"string"}),
				Configuration:         nvidiacloudfunctions.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvidiacloudfunctions.F(int64(1)),
				PreferredOrder:        nvidiacloudfunctions.F(int64(1)),
				Regions:               nvidiacloudfunctions.F([]string{"string"}),
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

func TestFunctionDeploymentFunctionVersionGet(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *nvidiacloudfunctions.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunctionDeploymentFunctionVersionUpdate(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvidiacloudfunctions.FunctionDeploymentFunctionVersionUpdateParams{
			DeploymentSpecifications: nvidiacloudfunctions.F([]nvidiacloudfunctions.FunctionDeploymentFunctionVersionUpdateParamsDeploymentSpecification{{
				GPU:                   nvidiacloudfunctions.F("gpu"),
				InstanceType:          nvidiacloudfunctions.F("instanceType"),
				MaxInstances:          nvidiacloudfunctions.F(int64(0)),
				MinInstances:          nvidiacloudfunctions.F(int64(0)),
				Attributes:            nvidiacloudfunctions.F([]string{"string"}),
				AvailabilityZones:     nvidiacloudfunctions.F([]string{"string", "string", "string"}),
				Backend:               nvidiacloudfunctions.F("backend"),
				Clusters:              nvidiacloudfunctions.F([]string{"string"}),
				Configuration:         nvidiacloudfunctions.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvidiacloudfunctions.F(int64(1)),
				PreferredOrder:        nvidiacloudfunctions.F(int64(1)),
				Regions:               nvidiacloudfunctions.F([]string{"string"}),
			}, {
				GPU:                   nvidiacloudfunctions.F("gpu"),
				InstanceType:          nvidiacloudfunctions.F("instanceType"),
				MaxInstances:          nvidiacloudfunctions.F(int64(0)),
				MinInstances:          nvidiacloudfunctions.F(int64(0)),
				Attributes:            nvidiacloudfunctions.F([]string{"string"}),
				AvailabilityZones:     nvidiacloudfunctions.F([]string{"string", "string", "string"}),
				Backend:               nvidiacloudfunctions.F("backend"),
				Clusters:              nvidiacloudfunctions.F([]string{"string"}),
				Configuration:         nvidiacloudfunctions.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvidiacloudfunctions.F(int64(1)),
				PreferredOrder:        nvidiacloudfunctions.F(int64(1)),
				Regions:               nvidiacloudfunctions.F([]string{"string"}),
			}, {
				GPU:                   nvidiacloudfunctions.F("gpu"),
				InstanceType:          nvidiacloudfunctions.F("instanceType"),
				MaxInstances:          nvidiacloudfunctions.F(int64(0)),
				MinInstances:          nvidiacloudfunctions.F(int64(0)),
				Attributes:            nvidiacloudfunctions.F([]string{"string"}),
				AvailabilityZones:     nvidiacloudfunctions.F([]string{"string", "string", "string"}),
				Backend:               nvidiacloudfunctions.F("backend"),
				Clusters:              nvidiacloudfunctions.F([]string{"string"}),
				Configuration:         nvidiacloudfunctions.F[any](map[string]interface{}{}),
				MaxRequestConcurrency: nvidiacloudfunctions.F(int64(1)),
				PreferredOrder:        nvidiacloudfunctions.F(int64(1)),
				Regions:               nvidiacloudfunctions.F([]string{"string"}),
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

func TestFunctionDeploymentFunctionVersionDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.FunctionDeployment.Functions.Versions.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		nvidiacloudfunctions.FunctionDeploymentFunctionVersionDeleteParams{
			Graceful: nvidiacloudfunctions.F(true),
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
