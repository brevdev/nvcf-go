// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiacloudfunctions

import (
	"github.com/brevdev/nvcf-go/option"
)

// NvcfService contains methods and other services that help with interacting with
// the nvidia-cloud-functions API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNvcfService] method instead.
type NvcfService struct {
	Options        []option.RequestOption
	Functions      *NvcfFunctionService
	Authorizations *NvcfAuthorizationService
}

// NewNvcfService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNvcfService(opts ...option.RequestOption) (r *NvcfService) {
	r = &NvcfService{}
	r.Options = opts
	r.Functions = NewNvcfFunctionService(opts...)
	r.Authorizations = NewNvcfAuthorizationService(opts...)
	return
}
