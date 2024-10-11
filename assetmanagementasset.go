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

// AssetManagementAssetService contains methods and other services that help with
// interacting with the nvidia-cloud-functions API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetManagementAssetService] method instead.
type AssetManagementAssetService struct {
	Options []option.RequestOption
}

// NewAssetManagementAssetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssetManagementAssetService(opts ...option.RequestOption) (r *AssetManagementAssetService) {
	r = &AssetManagementAssetService{}
	r.Options = opts
	return
}

// Returns details for the specified asset-id belonging to the current NVIDIA Cloud
// Account. Requires either a bearer token or an api-key with 'invoke_function'
// scope in the HTTP Authorization header.
func (r *AssetManagementAssetService) Get(ctx context.Context, assetID string, opts ...option.RequestOption) (res *AssetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes asset belonging to the current NVIDIA Cloud Account. Requires either a
// bearer token or an api-key with 'invoke_function' scope in the HTTP
// Authorization header.
func (r *AssetManagementAssetService) Delete(ctx context.Context, assetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return
	}
	path := fmt.Sprintf("v2/nvcf/assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AssetResponse struct {
	// Data Transfer Object(DTO) representing an asset
	Asset AssetResponseAsset `json:"asset"`
	JSON  assetResponseJSON  `json:"-"`
}

// assetResponseJSON contains the JSON metadata for the struct [AssetResponse]
type assetResponseJSON struct {
	Asset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an asset
type AssetResponseAsset struct {
	// Asset id
	AssetID string `json:"assetId" format:"uuid"`
	// Content-type specified when creating the asset
	ContentType string `json:"contentType"`
	// Description specified when creating the asset
	Description string                 `json:"description"`
	JSON        assetResponseAssetJSON `json:"-"`
}

// assetResponseAssetJSON contains the JSON metadata for the struct
// [AssetResponseAsset]
type assetResponseAssetJSON struct {
	AssetID     apijson.Field
	ContentType apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssetResponseAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assetResponseAssetJSON) RawJSON() string {
	return r.raw
}
