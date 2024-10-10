// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/brevdev/nvcf-go/option"
)

// NVCFAuthorizationService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNVCFAuthorizationService] method instead.
type NVCFAuthorizationService struct {
	Options   []option.RequestOption
	Functions *NVCFAuthorizationFunctionService
}

// NewNVCFAuthorizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNVCFAuthorizationService(opts ...option.RequestOption) (r *NVCFAuthorizationService) {
	r = &NVCFAuthorizationService{}
	r.Options = opts
	r.Functions = NewNVCFAuthorizationFunctionService(opts...)
	return
}
