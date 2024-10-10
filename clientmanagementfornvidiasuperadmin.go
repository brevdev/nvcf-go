// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/brevdev/nvcf-go/option"
)

// ClientManagementForNVIDIASuperAdminService contains methods and other services
// that help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientManagementForNVIDIASuperAdminService] method instead.
type ClientManagementForNVIDIASuperAdminService struct {
	Options []option.RequestOption
	Clients *ClientManagementForNVIDIASuperAdminClientService
}

// NewClientManagementForNVIDIASuperAdminService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewClientManagementForNVIDIASuperAdminService(opts ...option.RequestOption) (r *ClientManagementForNVIDIASuperAdminService) {
	r = &ClientManagementForNVIDIASuperAdminService{}
	r.Options = opts
	r.Clients = NewClientManagementForNVIDIASuperAdminClientService(opts...)
	return
}
