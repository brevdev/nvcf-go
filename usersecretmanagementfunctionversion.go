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
)

// UserSecretManagementFunctionVersionService contains methods and other services
// that help with interacting with the nvidia-cloud-functions API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSecretManagementFunctionVersionService] method instead.
type UserSecretManagementFunctionVersionService struct {
	Options []option.RequestOption
}

// NewUserSecretManagementFunctionVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewUserSecretManagementFunctionVersionService(opts ...option.RequestOption) (r *UserSecretManagementFunctionVersionService) {
	r = &UserSecretManagementFunctionVersionService{}
	r.Options = opts
	return
}

// Updates secrets for the specified function version. This endpoint requires
// either a bearer token or an api-key with 'update_secrets' scope in the HTTP
// Authorization header.
func (r *UserSecretManagementFunctionVersionService) Update(ctx context.Context, functionID string, versionID string, body UserSecretManagementFunctionVersionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required versionId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/secrets/functions/%s/versions/%s", functionID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type UserSecretManagementFunctionVersionUpdateParams struct {
	// Secrets
	Secrets param.Field[[]UserSecretManagementFunctionVersionUpdateParamsSecret] `json:"secrets,required"`
}

func (r UserSecretManagementFunctionVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data Transfer Object(DTO) representing secret name/value pair
type UserSecretManagementFunctionVersionUpdateParamsSecret struct {
	// Secret name
	Name param.Field[string] `json:"name,required"`
	// Secret value
	Value param.Field[string] `json:"value,required"`
}

func (r UserSecretManagementFunctionVersionUpdateParamsSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
