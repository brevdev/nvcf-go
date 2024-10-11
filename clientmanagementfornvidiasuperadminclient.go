// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvidiacloudfunctions

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/brevdev/nvcf-go/internal/apijson"
	"github.com/brevdev/nvcf-go/internal/requestconfig"
	"github.com/brevdev/nvcf-go/option"
)

// ClientManagementForNvidiaSuperAdminClientService contains methods and other
// services that help with interacting with the nvidia-cloud-functions API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientManagementForNvidiaSuperAdminClientService] method instead.
type ClientManagementForNvidiaSuperAdminClientService struct {
	Options []option.RequestOption
}

// NewClientManagementForNvidiaSuperAdminClientService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewClientManagementForNvidiaSuperAdminClientService(opts ...option.RequestOption) (r *ClientManagementForNvidiaSuperAdminClientService) {
	r = &ClientManagementForNvidiaSuperAdminClientService{}
	r.Options = opts
	return
}

// Gets details of the specified client.Requires a bearer token in the HTTP
// Authorization header with 'account_setup' scope. These endpoints are invoked by
// NVIDIA Super Admins working across accounts.
func (r *ClientManagementForNvidiaSuperAdminClientService) Get(ctx context.Context, clientID string, opts ...option.RequestOption) (res *ClientManagementForNvidiaSuperAdminClientGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clientID == "" {
		err = errors.New("missing required clientId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/clients/%s", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response body for client.
type ClientManagementForNvidiaSuperAdminClientGetResponse struct {
	// Data Transfer Object(DTO) representing a client
	Client ClientManagementForNvidiaSuperAdminClientGetResponseClient `json:"client,required"`
	JSON   clientManagementForNvidiaSuperAdminClientGetResponseJSON   `json:"-"`
}

// clientManagementForNvidiaSuperAdminClientGetResponseJSON contains the JSON
// metadata for the struct [ClientManagementForNvidiaSuperAdminClientGetResponse]
type clientManagementForNvidiaSuperAdminClientGetResponseJSON struct {
	Client      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientManagementForNvidiaSuperAdminClientGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientManagementForNvidiaSuperAdminClientGetResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a client
type ClientManagementForNvidiaSuperAdminClientGetResponseClient struct {
	// Name of the associated NVIDIA Cloud Account
	Name string `json:"name,required"`
	// Associated NVIDIA Cloud Account id
	NcaID string `json:"ncaId,required"`
	// Client Id
	ClientID string                                                         `json:"clientId"`
	JSON     clientManagementForNvidiaSuperAdminClientGetResponseClientJSON `json:"-"`
}

// clientManagementForNvidiaSuperAdminClientGetResponseClientJSON contains the JSON
// metadata for the struct
// [ClientManagementForNvidiaSuperAdminClientGetResponseClient]
type clientManagementForNvidiaSuperAdminClientGetResponseClientJSON struct {
	Name        apijson.Field
	NcaID       apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientManagementForNvidiaSuperAdminClientGetResponseClient) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientManagementForNvidiaSuperAdminClientGetResponseClientJSON) RawJSON() string {
	return r.raw
}
