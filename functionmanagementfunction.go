// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiacloudfunctions

import (
	"github.com/brevdev/nvcf-go/option"
)

// FunctionManagementFunctionService contains methods and other services that help
// with interacting with the nvidia-cloud-functions API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionManagementFunctionService] method instead.
type FunctionManagementFunctionService struct {
	Options  []option.RequestOption
	Versions *FunctionManagementFunctionVersionService
	IDs      *FunctionManagementFunctionIDService
}

// NewFunctionManagementFunctionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFunctionManagementFunctionService(opts ...option.RequestOption) (r *FunctionManagementFunctionService) {
	r = &FunctionManagementFunctionService{}
	r.Options = opts
	r.Versions = NewFunctionManagementFunctionVersionService(opts...)
	r.IDs = NewFunctionManagementFunctionIDService(opts...)
	return
}
