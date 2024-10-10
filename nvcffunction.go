// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"net/http"
	"net/url"

	"github.com/brevdev/nvcf-go/internal/apijson"
	"github.com/brevdev/nvcf-go/internal/apiquery"
	"github.com/brevdev/nvcf-go/internal/param"
	"github.com/brevdev/nvcf-go/internal/requestconfig"
	"github.com/brevdev/nvcf-go/option"
	"github.com/brevdev/nvcf-go/shared"
)

// NVCFFunctionService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNVCFFunctionService] method instead.
type NVCFFunctionService struct {
	Options  []option.RequestOption
	Versions *NVCFFunctionVersionService
}

// NewNVCFFunctionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNVCFFunctionService(opts ...option.RequestOption) (r *NVCFFunctionService) {
	r = &NVCFFunctionService{}
	r.Options = opts
	r.Versions = NewNVCFFunctionVersionService(opts...)
	return
}

// Creates a new function within the authenticated NVIDIA Cloud Account. Requires a
// bearer token with 'register_function' scope in the HTTP Authorization header.
func (r *NVCFFunctionService) New(ctx context.Context, body NVCFFunctionNewParams, opts ...option.RequestOption) (res *shared.CreateFunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all the functions associated with the authenticated NVIDIA Cloud Account.
// Requires either a bearer token or an api-key with 'list_functions' or
// 'list_functions_details' scopes in the HTTP Authorization header.
func (r *NVCFFunctionService) List(ctx context.Context, query NVCFFunctionListParams, opts ...option.RequestOption) (res *shared.ListFunctionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NVCFFunctionNewParams struct {
	// Entrypoint for invoking the container to process a request
	InferenceURL param.Field[string] `json:"inferenceUrl,required" format:"uri"`
	// Function name must start with lowercase/uppercase/digit and can only contain
	// lowercase, uppercase, digit, hyphen, and underscore characters
	Name param.Field[string] `json:"name,required"`
	// Invocation request body format
	APIBodyFormat param.Field[NVCFFunctionNewParamsAPIBodyFormat] `json:"apiBodyFormat"`
	// Args to be passed when launching the container
	ContainerArgs param.Field[string] `json:"containerArgs"`
	// Environment settings for launching the container
	ContainerEnvironment param.Field[[]NVCFFunctionNewParamsContainerEnvironment] `json:"containerEnvironment"`
	// Optional custom container image
	ContainerImage param.Field[string] `json:"containerImage" format:"uri"`
	// Optional function/version description
	Description param.Field[string] `json:"description"`
	// Optional function type, used to indicate a STREAMING function. Defaults to
	// DEFAULT.
	FunctionType param.Field[NVCFFunctionNewParamsFunctionType] `json:"functionType"`
	// Data Transfer Object(DTO) representing a function ne
	Health param.Field[NVCFFunctionNewParamsHealth] `json:"health"`
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
	Models param.Field[[]NVCFFunctionNewParamsModel] `json:"models"`
	// Optional set of resources
	Resources param.Field[[]NVCFFunctionNewParamsResource] `json:"resources"`
	// Optional secrets
	Secrets param.Field[[]NVCFFunctionNewParamsSecret] `json:"secrets"`
	// Optional set of tags - could be empty. Provided by user
	Tags param.Field[[]string] `json:"tags"`
}

func (r NVCFFunctionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Invocation request body format
type NVCFFunctionNewParamsAPIBodyFormat string

const (
	NVCFFunctionNewParamsAPIBodyFormatPredictV2 NVCFFunctionNewParamsAPIBodyFormat = "PREDICT_V2"
	NVCFFunctionNewParamsAPIBodyFormatCustom    NVCFFunctionNewParamsAPIBodyFormat = "CUSTOM"
)

func (r NVCFFunctionNewParamsAPIBodyFormat) IsKnown() bool {
	switch r {
	case NVCFFunctionNewParamsAPIBodyFormatPredictV2, NVCFFunctionNewParamsAPIBodyFormatCustom:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a container environment entry
type NVCFFunctionNewParamsContainerEnvironment struct {
	// Container environment key
	Key param.Field[string] `json:"key,required"`
	// Container environment value
	Value param.Field[string] `json:"value,required"`
}

func (r NVCFFunctionNewParamsContainerEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional function type, used to indicate a STREAMING function. Defaults to
// DEFAULT.
type NVCFFunctionNewParamsFunctionType string

const (
	NVCFFunctionNewParamsFunctionTypeDefault   NVCFFunctionNewParamsFunctionType = "DEFAULT"
	NVCFFunctionNewParamsFunctionTypeStreaming NVCFFunctionNewParamsFunctionType = "STREAMING"
)

func (r NVCFFunctionNewParamsFunctionType) IsKnown() bool {
	switch r {
	case NVCFFunctionNewParamsFunctionTypeDefault, NVCFFunctionNewParamsFunctionTypeStreaming:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a function ne
type NVCFFunctionNewParamsHealth struct {
	// Expected return status code considered as successful.
	ExpectedStatusCode param.Field[int64] `json:"expectedStatusCode,required"`
	// Port number where the health listener is running
	Port param.Field[int64] `json:"port,required"`
	// HTTP/gPRC protocol type for health endpoint
	Protocol param.Field[NVCFFunctionNewParamsHealthProtocol] `json:"protocol,required"`
	// ISO 8601 duration string in PnDTnHnMn.nS format
	Timeout param.Field[string] `json:"timeout,required" format:"PnDTnHnMn.nS"`
	// Health endpoint for the container or the helmChart
	Uri param.Field[string] `json:"uri,required" format:"uri"`
}

func (r NVCFFunctionNewParamsHealth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP/gPRC protocol type for health endpoint
type NVCFFunctionNewParamsHealthProtocol string

const (
	NVCFFunctionNewParamsHealthProtocolHTTP NVCFFunctionNewParamsHealthProtocol = "HTTP"
	NVCFFunctionNewParamsHealthProtocolGRpc NVCFFunctionNewParamsHealthProtocol = "gRPC"
)

func (r NVCFFunctionNewParamsHealthProtocol) IsKnown() bool {
	switch r {
	case NVCFFunctionNewParamsHealthProtocolHTTP, NVCFFunctionNewParamsHealthProtocolGRpc:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing an artifact
type NVCFFunctionNewParamsModel struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r NVCFFunctionNewParamsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an artifact
type NVCFFunctionNewParamsResource struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r NVCFFunctionNewParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing secret name/value pair
type NVCFFunctionNewParamsSecret struct {
	// Secret name
	Name param.Field[string] `json:"name,required"`
	// Secret value
	Value param.Field[string] `json:"value,required"`
}

func (r NVCFFunctionNewParamsSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NVCFFunctionListParams struct {
	// Query param 'visibility' indicates the kind of functions to be included in the
	// response.
	Visibility param.Field[[]NVCFFunctionListParamsVisibility] `query:"visibility"`
}

// URLQuery serializes [NVCFFunctionListParams]'s query parameters as `url.Values`.
func (r NVCFFunctionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NVCFFunctionListParamsVisibility string

const (
	NVCFFunctionListParamsVisibilityAuthorized NVCFFunctionListParamsVisibility = "authorized"
	NVCFFunctionListParamsVisibilityPrivate    NVCFFunctionListParamsVisibility = "private"
	NVCFFunctionListParamsVisibilityPublic     NVCFFunctionListParamsVisibility = "public"
)

func (r NVCFFunctionListParamsVisibility) IsKnown() bool {
	switch r {
	case NVCFFunctionListParamsVisibilityAuthorized, NVCFFunctionListParamsVisibilityPrivate, NVCFFunctionListParamsVisibilityPublic:
		return true
	}
	return false
}
