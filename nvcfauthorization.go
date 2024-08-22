// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiacloudfunctions

import (
	"github.com/stainless-sdks/nvidia-cloud-functions-go/option"
)

// NvcfAuthorizationService contains methods and other services that help with
// interacting with the nvidia-cloud-functions API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNvcfAuthorizationService] method instead.
type NvcfAuthorizationService struct {
	Options   []option.RequestOption
	Functions *NvcfAuthorizationFunctionService
}

// NewNvcfAuthorizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNvcfAuthorizationService(opts ...option.RequestOption) (r *NvcfAuthorizationService) {
	r = &NvcfAuthorizationService{}
	r.Options = opts
	r.Functions = NewNvcfAuthorizationFunctionService(opts...)
	return
}
