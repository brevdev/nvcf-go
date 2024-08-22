// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiacloudfunctions

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

// NvcfFunctionService contains methods and other services that help with
// interacting with the nvidia-cloud-functions API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNvcfFunctionService] method instead.
type NvcfFunctionService struct {
	Options  []option.RequestOption
	Versions *NvcfFunctionVersionService
}

// NewNvcfFunctionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNvcfFunctionService(opts ...option.RequestOption) (r *NvcfFunctionService) {
	r = &NvcfFunctionService{}
	r.Options = opts
	r.Versions = NewNvcfFunctionVersionService(opts...)
	return
}

// Creates a new function within the authenticated NVIDIA Cloud Account. Requires a
// bearer token with 'register_function' scope in the HTTP Authorization header.
func (r *NvcfFunctionService) New(ctx context.Context, body NvcfFunctionNewParams, opts ...option.RequestOption) (res *shared.CreateFunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all the functions associated with the authenticated NVIDIA Cloud Account.
// Requires either a bearer token or an api-key with 'list_functions' or
// 'list_functions_details' scopes in the HTTP Authorization header.
func (r *NvcfFunctionService) List(ctx context.Context, query NvcfFunctionListParams, opts ...option.RequestOption) (res *shared.ListFunctionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NvcfFunctionNewParams struct {
	// Entrypoint for invoking the container to process a request
	InferenceURL param.Field[string] `json:"inferenceUrl,required" format:"uri"`
	// Function name must start with lowercase/uppercase/digit and can only contain
	// lowercase, uppercase, digit, hyphen, and underscore characters
	Name param.Field[string] `json:"name,required"`
	// Invocation request body format
	APIBodyFormat param.Field[NvcfFunctionNewParamsAPIBodyFormat] `json:"apiBodyFormat"`
	// Args to be passed when launching the container
	ContainerArgs param.Field[string] `json:"containerArgs"`
	// Environment settings for launching the container
	ContainerEnvironment param.Field[[]NvcfFunctionNewParamsContainerEnvironment] `json:"containerEnvironment"`
	// Optional custom container image
	ContainerImage param.Field[string] `json:"containerImage" format:"uri"`
	// Optional function/version description
	Description param.Field[string] `json:"description"`
	// Optional function type, used to indicate a STREAMING function. Defaults to
	// DEFAULT.
	FunctionType param.Field[NvcfFunctionNewParamsFunctionType] `json:"functionType"`
	// Data Transfer Object(DTO) representing a function ne
	Health param.Field[NvcfFunctionNewParamsHealth] `json:"health"`
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
	Models param.Field[[]NvcfFunctionNewParamsModel] `json:"models"`
	// Optional set of resources
	Resources param.Field[[]NvcfFunctionNewParamsResource] `json:"resources"`
	// Optional secrets
	Secrets param.Field[[]NvcfFunctionNewParamsSecret] `json:"secrets"`
	// Optional set of tags - could be empty. Provided by user
	Tags param.Field[[]string] `json:"tags"`
}

func (r NvcfFunctionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Invocation request body format
type NvcfFunctionNewParamsAPIBodyFormat string

const (
	NvcfFunctionNewParamsAPIBodyFormatPredictV2 NvcfFunctionNewParamsAPIBodyFormat = "PREDICT_V2"
	NvcfFunctionNewParamsAPIBodyFormatCustom    NvcfFunctionNewParamsAPIBodyFormat = "CUSTOM"
)

func (r NvcfFunctionNewParamsAPIBodyFormat) IsKnown() bool {
	switch r {
	case NvcfFunctionNewParamsAPIBodyFormatPredictV2, NvcfFunctionNewParamsAPIBodyFormatCustom:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a container environment entry
type NvcfFunctionNewParamsContainerEnvironment struct {
	// Container environment key
	Key param.Field[string] `json:"key,required"`
	// Container environment value
	Value param.Field[string] `json:"value,required"`
}

func (r NvcfFunctionNewParamsContainerEnvironment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional function type, used to indicate a STREAMING function. Defaults to
// DEFAULT.
type NvcfFunctionNewParamsFunctionType string

const (
	NvcfFunctionNewParamsFunctionTypeDefault   NvcfFunctionNewParamsFunctionType = "DEFAULT"
	NvcfFunctionNewParamsFunctionTypeStreaming NvcfFunctionNewParamsFunctionType = "STREAMING"
)

func (r NvcfFunctionNewParamsFunctionType) IsKnown() bool {
	switch r {
	case NvcfFunctionNewParamsFunctionTypeDefault, NvcfFunctionNewParamsFunctionTypeStreaming:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a function ne
type NvcfFunctionNewParamsHealth struct {
	// Expected return status code considered as successful.
	ExpectedStatusCode param.Field[int64] `json:"expectedStatusCode,required"`
	// Port number where the health listener is running
	Port param.Field[int64] `json:"port,required"`
	// HTTP/gPRC protocol type for health endpoint
	Protocol param.Field[NvcfFunctionNewParamsHealthProtocol] `json:"protocol,required"`
	// ISO 8601 duration string in PnDTnHnMn.nS format
	Timeout param.Field[string] `json:"timeout,required" format:"PnDTnHnMn.nS"`
	// Health endpoint for the container or the helmChart
	Uri param.Field[string] `json:"uri,required" format:"uri"`
}

func (r NvcfFunctionNewParamsHealth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP/gPRC protocol type for health endpoint
type NvcfFunctionNewParamsHealthProtocol string

const (
	NvcfFunctionNewParamsHealthProtocolHTTP NvcfFunctionNewParamsHealthProtocol = "HTTP"
	NvcfFunctionNewParamsHealthProtocolGRpc NvcfFunctionNewParamsHealthProtocol = "gRPC"
)

func (r NvcfFunctionNewParamsHealthProtocol) IsKnown() bool {
	switch r {
	case NvcfFunctionNewParamsHealthProtocolHTTP, NvcfFunctionNewParamsHealthProtocolGRpc:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing an artifact
type NvcfFunctionNewParamsModel struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r NvcfFunctionNewParamsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing an artifact
type NvcfFunctionNewParamsResource struct {
	// Artifact name
	Name param.Field[string] `json:"name,required"`
	// Artifact URI
	Uri param.Field[string] `json:"uri,required" format:"uri"`
	// Artifact version
	Version param.Field[string] `json:"version,required"`
}

func (r NvcfFunctionNewParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing secret name/value pair
type NvcfFunctionNewParamsSecret struct {
	// Secret name
	Name param.Field[string] `json:"name,required"`
	// Secret value
	Value param.Field[string] `json:"value,required"`
}

func (r NvcfFunctionNewParamsSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NvcfFunctionListParams struct {
	// Query param 'visibility' indicates the kind of functions to be included in the
	// response.
	Visibility param.Field[[]NvcfFunctionListParamsVisibility] `query:"visibility"`
}

// URLQuery serializes [NvcfFunctionListParams]'s query parameters as `url.Values`.
func (r NvcfFunctionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NvcfFunctionListParamsVisibility string

const (
	NvcfFunctionListParamsVisibilityAuthorized NvcfFunctionListParamsVisibility = "authorized"
	NvcfFunctionListParamsVisibilityPrivate    NvcfFunctionListParamsVisibility = "private"
	NvcfFunctionListParamsVisibilityPublic     NvcfFunctionListParamsVisibility = "public"
)

func (r NvcfFunctionListParamsVisibility) IsKnown() bool {
	switch r {
	case NvcfFunctionListParamsVisibilityAuthorized, NvcfFunctionListParamsVisibilityPrivate, NvcfFunctionListParamsVisibilityPublic:
		return true
	}
	return false
}
