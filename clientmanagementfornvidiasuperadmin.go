// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiacloudfunctions

import (
	"github.com/brevdev/nvcf-go/option"
)

// ClientManagementForNvidiaSuperAdminService contains methods and other services
// that help with interacting with the nvidia-cloud-functions API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientManagementForNvidiaSuperAdminService] method instead.
type ClientManagementForNvidiaSuperAdminService struct {
	Options []option.RequestOption
	Clients *ClientManagementForNvidiaSuperAdminClientService
}

// NewClientManagementForNvidiaSuperAdminService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewClientManagementForNvidiaSuperAdminService(opts ...option.RequestOption) (r *ClientManagementForNvidiaSuperAdminService) {
	r = &ClientManagementForNvidiaSuperAdminService{}
	r.Options = opts
	r.Clients = NewClientManagementForNvidiaSuperAdminClientService(opts...)
	return
}
