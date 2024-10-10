// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/brevdev/nvcf-go/internal/apijson"
	"github.com/brevdev/nvcf-go/internal/requestconfig"
	"github.com/brevdev/nvcf-go/option"
)

// ClientManagementForNVIDIASuperAdminClientService contains methods and other
// services that help with interacting with the nvcf API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientManagementForNVIDIASuperAdminClientService] method instead.
type ClientManagementForNVIDIASuperAdminClientService struct {
	Options []option.RequestOption
}

// NewClientManagementForNVIDIASuperAdminClientService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewClientManagementForNVIDIASuperAdminClientService(opts ...option.RequestOption) (r *ClientManagementForNVIDIASuperAdminClientService) {
	r = &ClientManagementForNVIDIASuperAdminClientService{}
	r.Options = opts
	return
}

// Gets details of the specified client.Requires a bearer token in the HTTP
// Authorization header with 'account_setup' scope. These endpoints are invoked by
// NVIDIA Super Admins working across accounts.
func (r *ClientManagementForNVIDIASuperAdminClientService) Get(ctx context.Context, clientID string, opts ...option.RequestOption) (res *ClientManagementForNVIDIASuperAdminClientGetResponse, err error) {
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
type ClientManagementForNVIDIASuperAdminClientGetResponse struct {
	// Data Transfer Object(DTO) representing a client
	Client ClientManagementForNVIDIASuperAdminClientGetResponseClient `json:"client,required"`
	JSON   clientManagementForNVIDIASuperAdminClientGetResponseJSON   `json:"-"`
}

// clientManagementForNVIDIASuperAdminClientGetResponseJSON contains the JSON
// metadata for the struct [ClientManagementForNVIDIASuperAdminClientGetResponse]
type clientManagementForNVIDIASuperAdminClientGetResponseJSON struct {
	Client      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientManagementForNVIDIASuperAdminClientGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientManagementForNVIDIASuperAdminClientGetResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a client
type ClientManagementForNVIDIASuperAdminClientGetResponseClient struct {
	// Name of the associated NVIDIA Cloud Account
	Name string `json:"name,required"`
	// Associated NVIDIA Cloud Account id
	NcaID string `json:"ncaId,required"`
	// Client Id
	ClientID string                                                         `json:"clientId"`
	JSON     clientManagementForNVIDIASuperAdminClientGetResponseClientJSON `json:"-"`
}

// clientManagementForNVIDIASuperAdminClientGetResponseClientJSON contains the JSON
// metadata for the struct
// [ClientManagementForNVIDIASuperAdminClientGetResponseClient]
type clientManagementForNVIDIASuperAdminClientGetResponseClientJSON struct {
	Name        apijson.Field
	NcaID       apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientManagementForNVIDIASuperAdminClientGetResponseClient) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientManagementForNVIDIASuperAdminClientGetResponseClientJSON) RawJSON() string {
	return r.raw
}
