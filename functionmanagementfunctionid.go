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
)

// FunctionManagementFunctionIDService contains methods and other services that
// help with interacting with the nvidia-cloud-functions API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionManagementFunctionIDService] method instead.
type FunctionManagementFunctionIDService struct {
	Options []option.RequestOption
}

// NewFunctionManagementFunctionIDService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFunctionManagementFunctionIDService(opts ...option.RequestOption) (r *FunctionManagementFunctionIDService) {
	r = &FunctionManagementFunctionIDService{}
	r.Options = opts
	return
}

// Lists ids of all the functions in the authenticated NVIDIA Cloud Account.
// Requires either a bearer token or an api-key with 'list_functions' or
// 'list_functions_details' scopes in the HTTP Authorization header.
func (r *FunctionManagementFunctionIDService) List(ctx context.Context, query FunctionManagementFunctionIDListParams, opts ...option.RequestOption) (res *ListFunctionIDsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/nvcf/functions/ids"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response body containing list of function ids in an account
type ListFunctionIDsResponse struct {
	// List of function ids
	FunctionIDs []string                    `json:"functionIds,required" format:"uuid"`
	JSON        listFunctionIDsResponseJSON `json:"-"`
}

// listFunctionIDsResponseJSON contains the JSON metadata for the struct
// [ListFunctionIDsResponse]
type listFunctionIDsResponseJSON struct {
	FunctionIDs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListFunctionIDsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listFunctionIDsResponseJSON) RawJSON() string {
	return r.raw
}

type FunctionManagementFunctionIDListParams struct {
	// Query param 'visibility' indicates the kind of functions to be included in the
	// response.
	Visibility param.Field[[]FunctionManagementFunctionIDListParamsVisibility] `query:"visibility"`
}

// URLQuery serializes [FunctionManagementFunctionIDListParams]'s query parameters
// as `url.Values`.
func (r FunctionManagementFunctionIDListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FunctionManagementFunctionIDListParamsVisibility string

const (
	FunctionManagementFunctionIDListParamsVisibilityAuthorized FunctionManagementFunctionIDListParamsVisibility = "authorized"
	FunctionManagementFunctionIDListParamsVisibilityPrivate    FunctionManagementFunctionIDListParamsVisibility = "private"
	FunctionManagementFunctionIDListParamsVisibilityPublic     FunctionManagementFunctionIDListParamsVisibility = "public"
)

func (r FunctionManagementFunctionIDListParamsVisibility) IsKnown() bool {
	switch r {
	case FunctionManagementFunctionIDListParamsVisibilityAuthorized, FunctionManagementFunctionIDListParamsVisibilityPrivate, FunctionManagementFunctionIDListParamsVisibilityPublic:
		return true
	}
	return false
}
