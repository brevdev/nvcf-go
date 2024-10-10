// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/brevdev/nvcf-go/internal/apijson"
	"github.com/brevdev/nvcf-go/internal/param"
	"github.com/brevdev/nvcf-go/internal/requestconfig"
	"github.com/brevdev/nvcf-go/option"
	"github.com/brevdev/nvcf-go/shared"
)

// NVCFFunctionVersionService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNVCFFunctionVersionService] method instead.
type NVCFFunctionVersionService struct {
	Options []option.RequestOption
}

// NewNVCFFunctionVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNVCFFunctionVersionService(opts ...option.RequestOption) (r *NVCFFunctionVersionService) {
	r = &NVCFFunctionVersionService{}
	r.Options = opts
	return
}

// Creates a version of the specified function within the authenticated NVIDIA
// Cloud Account. Requires a bearer token with 'register_function' scope in the
// HTTP Authorization header.
func (r *NVCFFunctionVersionService) New(ctx context.Context, functionID string, body NVCFFunctionVersionNewParams, opts ...option.RequestOption) (res *shared.CreateFunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/functions/%s/versions", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists details of all the versions of the specified function in the authenticated
// NVIDIA Cloud Account. Requires either a bearer token or an api-key with
// 'list_functions' or 'list_functions_details' scopes in the HTTP Authorization
// header.
func (r *NVCFFunctionVersionService) List(ctx context.Context, functionID string, opts ...option.RequestOption) (res *shared.ListFunctionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/functions/%s/versions", functionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NVCFFunctionVersionNewParams struct {
	// Entrypoint for invoking the container to process a request
	InferenceURL param.Field[string] `json:"inferenceUrl,required" format:"uri"`
	// Function name must start with lowercase/uppercase/digit and can only contain
	// lowercase, uppercase, digit, hyphen, and underscore characters
	Name param.Field[string] `json:"name,required"`
	// Invocation request body format
	APIBodyFormat param.Field[NVCFFunctionVersionNewParamsAPIBodyFormat] `json:"apiBodyFormat"`
	// Args to be passed when launching the container
	ContainerArgs param.Field[string] `json:"containerArgs"`
	// Environment settings for launching the container
	ContainerEnvironment param.Field[[]NVCFFunctionVersionNewParamsContainerEnvironment] `json:"containerEnvironment"`
	// Optional custom container image
	ContainerImage param.Field[string] `json:"containerImage" format:"uri"`
	// Optional function/version description
	Description param.Field[string] `json:"description"`
	// Optional function type, used to indicate a STREAMING function. Defaults to
	// DEFAULT.
	FunctionType param.Field[NVCFFunctionVersionNewParamsFunctionType] `json:"functionType"`
	// Data Transfer Object(DTO) representing a function ne
	Health param.Field[NVCFFunctionVersionNewParamsHealth] `json:"health"`
	// Health endpoint for the container or the helmChart
	HealthUri param.Field[string] `json:"healthUri" format:"uri"`
	// Optional Helm Chart
	HelmChart param.Field[string] `json:"helmChart" format:"uri"`
	// Helm Chart Service Name is required when helmChart property is specified
	HelmChartServiceName param.Field[string] `json:"helmChartServiceName"`
	// Optional port number where the inference listener is running. Defaults to 8000
	// for Triton.
	InferencePort param.Field[int64] `json:"inferencePort"`
	// Optional set of models
	Models param.Field[[]NVCFFunctionVersionNewParamsModel] `json:"models"`
	// Optional set of resources
	Resources param.Field[[]NVCFFunctionVersionNewParamsResource] `json:"resources"`
	// Optional secrets
	Secrets param.Field[[]NVCFFunctionVersionNewParamsSecret] `json:"secrets"`
	// Optional set of tags - could be empty. Provided by user
	Tags param.Field[[]string] `json:"tags"`
}

func (r NVCFFunctionVersionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Invocation request body format
type NVCFFunctionVersionNewParamsAPIBodyFormat string

const (
	NVCFFunctionVersionNewParamsAPIBodyFormatPredictV2 NVCFFunctionVersionNewParamsAPIBodyFormat = "PREDICT_V2"
	NVCFFunctionVersionNewParamsAPIBodyFormatCustom    NVCFFunctionVersionNewParamsAPIBodyFormat = "CUSTOM"
)

func (r NVCFFunctionVersionNewParamsAPIBodyFormat) IsKnown() bool {
	switch r {
	case NVCFFunctionVersionNewParamsAPIBodyFormatPredictV2, NVCFFunctionVersionNewParamsAPIBodyFormatCustom:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a container environment entry
type NVCFFunctionVersionNewParamsContainerEnvironment struct {
	// Container environment key
	Key param.Field[string] `json:"key,required"`
	// Container environment value
	Value param.Field[string] `json:"value,required"`
}

func (r NVCFFunctionVersionNewParamsContainerEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional function type, used to indicate a STREAMING function. Defaults to
// DEFAULT.
type NVCFFunctionVersionNewParamsFunctionType string

const (
	NVCFFunctionVersionNewParamsFunctionTypeDefault   NVCFFunctionVersionNewParamsFunctionType = "DEFAULT"
	NVCFFunctionVersionNewParamsFunctionTypeStreaming NVCFFunctionVersionNewParamsFunctionType = "STREAMING"
)

func (r NVCFFunctionVersionNewParamsFunctionType) IsKnown() bool {
	switch r {
	case NVCFFunctionVersionNewParamsFunctionTypeDefault, NVCFFunctionVersionNewParamsFunctionTypeStreaming:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a function ne
type NVCFFunctionVersionNewParamsHealth struct {
	// Expected return status code considered as successful.
	ExpectedStatusCode param.Field[int64] `json:"expectedStatusCode,required"`
	// Port number where the health listener is running
	Port param.Field[int64] `json:"port,required"`
	// HTTP/gPRC protocol type for health endpoint
	Protocol param.Field[NVCFFunctionVersionNewParamsHealthProtocol] `json:"protocol,required"`
	// ISO 8601 duration string in PnDTnHnMn.nS format
	Timeout param.Field[string] `json:"timeout,required" format:"PnDTnHnMn.nS"`
	// Health endpoint for the container or the helmChart
	Uri param.Field[string] `json:"uri,required" format:"uri"`
}

func (r NVCFFunctionVersionNewParamsHealth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP/gPRC protocol type for health endpoint
type NVCFFunctionVersionNewParamsHealthProtocol string

const (
	NVCFFunctionVersionNewParamsHealthProtocolHTTP NVCFFunctionVersionNewParamsHealthProtocol = "HTTP"
	NVCFFunctionVersionNewParamsHealthProtocolGRpc NVCFFunctionVersionNewParamsHealthProtocol = "gRPC"
)

func (r NVCFFunctionVersionNewParamsHealthProtocol) IsKnown() bool {
	switch r {
	case NVCFFunctionVersionNewParamsHealthProtocolHTTP, NVCFFunctionVersionNewParamsHealthProtocolGRpc:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing an artifact
type NVCFFunctionVersionNewParamsModel struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r NVCFFunctionVersionNewParamsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an artifact
type NVCFFunctionVersionNewParamsResource struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r NVCFFunctionVersionNewParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing secret name/value pair
type NVCFFunctionVersionNewParamsSecret struct {
	// Secret name
	Name param.Field[string] `json:"name,required"`
	// Secret value
	Value param.Field[string] `json:"value,required"`
}

func (r NVCFFunctionVersionNewParamsSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
