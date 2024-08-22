// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiacloudfunctions

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/nvidia-cloud-functions-go/internal/apijson"
	"github.com/stainless-sdks/nvidia-cloud-functions-go/internal/param"
	"github.com/stainless-sdks/nvidia-cloud-functions-go/internal/requestconfig"
	"github.com/stainless-sdks/nvidia-cloud-functions-go/option"
	"github.com/stainless-sdks/nvidia-cloud-functions-go/shared"
)

// FunctionManagementFunctionVersionService contains methods and other services
// that help with interacting with the nvidia-cloud-functions API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionManagementFunctionVersionService] method instead.
type FunctionManagementFunctionVersionService struct {
	Options []option.RequestOption
}

// NewFunctionManagementFunctionVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFunctionManagementFunctionVersionService(opts ...option.RequestOption) (r *FunctionManagementFunctionVersionService) {
	r = &FunctionManagementFunctionVersionService{}
	r.Options = opts
	return
}

// Retrieves detailed information of the specified function version in the
// authenticated NVIDIA Cloud Account. Requires either a bearer token or an api-key
// with 'list_functions' or 'list_functions_details' scopes in the HTTP
// Authorization header.
func (r *FunctionManagementFunctionVersionService) Get(ctx context.Context, functionID string, functionVersionID string, opts ...option.RequestOption) (res *shared.FunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates metadata, such as tags, of the specified function version within the
// authenticated NVIDIA Cloud Account. Values specified in the payload completely
// override the existing values. Requires a bearer token with 'update_function'
// scope in the HTTP Authorization header.
func (r *FunctionManagementFunctionVersionService) Update(ctx context.Context, functionID string, functionVersionID string, body FunctionManagementFunctionVersionUpdateParams, opts ...option.RequestOption) (res *shared.FunctionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/metadata/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes the specified function version in the authenticated NVIDIA Cloud
// Account. Requires a bearer token with 'delete_function' scope in the HTTP
// Authorization header. If the function version is public, then Account Admin
// cannot delete the function.
func (r *FunctionManagementFunctionVersionService) Delete(ctx context.Context, functionID string, functionVersionID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if functionVersionID == "" {
		err = errors.New("missing required functionVersionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/functions/%s/versions/%s", functionID, functionVersionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type FunctionManagementFunctionVersionUpdateParams struct {
	// Set of tags provided by user
	Tags param.Field[[]string] `json:"tags"`
}

func (r FunctionManagementFunctionVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
