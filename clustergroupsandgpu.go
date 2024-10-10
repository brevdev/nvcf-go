// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/brevdev/nvcf-go/option"
)

// ClusterGroupsAndGPUService contains methods and other services that help with
// interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClusterGroupsAndGPUService] method instead.
type ClusterGroupsAndGPUService struct {
	Options       []option.RequestOption
	ClusterGroups *ClusterGroupsAndGPUClusterGroupService
}

// NewClusterGroupsAndGPUService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClusterGroupsAndGPUService(opts ...option.RequestOption) (r *ClusterGroupsAndGPUService) {
	r = &ClusterGroupsAndGPUService{}
	r.Options = opts
	r.ClusterGroups = NewClusterGroupsAndGPUClusterGroupService(opts...)
	return
}
