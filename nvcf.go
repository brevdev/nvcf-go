// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/brevdev/nvcf-go/option"
)

// NVCFService contains methods and other services that help with interacting with
// the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNVCFService] method instead.
type NVCFService struct {
	Options        []option.RequestOption
	Functions      *NVCFFunctionService
	Authorizations *NVCFAuthorizationService
}

// NewNVCFService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNVCFService(opts ...option.RequestOption) (r *NVCFService) {
	r = &NVCFService{}
	r.Options = opts
	r.Functions = NewNVCFFunctionService(opts...)
	r.Authorizations = NewNVCFAuthorizationService(opts...)
	return
}
