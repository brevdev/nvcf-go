// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/stainless-sdks/nvidia-cloud-functions-go/internal/apijson"
)

// Parties authorized to invoke function
type AuthorizedPartiesResponse struct {
	// Data Transfer Object(DTO) representing a function with authorized accounts
	Function AuthorizedPartiesResponseFunction `json:"function,required"`
	JSON     authorizedPartiesResponseJSON     `json:"-"`
}

// authorizedPartiesResponseJSON contains the JSON metadata for the struct
// [AuthorizedPartiesResponse]
type authorizedPartiesResponseJSON struct {
	Function    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizedPartiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizedPartiesResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a function with authorized accounts
type AuthorizedPartiesResponseFunction struct {
	// Function id
	ID string `json:"id,required" format:"uuid"`
	// NVIDIA Cloud Account Id
	NcaID string `json:"ncaId,required"`
	// Authorized parties allowed to invoke the function
	AuthorizedParties []AuthorizedPartiesResponseFunctionAuthorizedParty `json:"authorizedParties"`
	// Function version id
	VersionID string                                `json:"versionId" format:"uuid"`
	JSON      authorizedPartiesResponseFunctionJSON `json:"-"`
}

// authorizedPartiesResponseFunctionJSON contains the JSON metadata for the struct
// [AuthorizedPartiesResponseFunction]
type authorizedPartiesResponseFunctionJSON struct {
	ID                apijson.Field
	NcaID             apijson.Field
	AuthorizedParties apijson.Field
	VersionID         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AuthorizedPartiesResponseFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizedPartiesResponseFunctionJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an authorized party.
type AuthorizedPartiesResponseFunctionAuthorizedParty struct {
	// NVIDIA Cloud Account authorized to invoke the function
	NcaID string `json:"ncaId,required"`
	// Client Id -- 'sub' claim in the JWT. This field should not be specified anymore.
	ClientID string                                               `json:"clientId"`
	JSON     authorizedPartiesResponseFunctionAuthorizedPartyJSON `json:"-"`
}

// authorizedPartiesResponseFunctionAuthorizedPartyJSON contains the JSON metadata
// for the struct [AuthorizedPartiesResponseFunctionAuthorizedParty]
type authorizedPartiesResponseFunctionAuthorizedPartyJSON struct {
	NcaID       apijson.Field
	ClientID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthorizedPartiesResponseFunctionAuthorizedParty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authorizedPartiesResponseFunctionAuthorizedPartyJSON) RawJSON() string {
	return r.raw
}

// Response body for create function request.
type CreateFunctionResponse struct {
	// Data Transfer Object (DTO) representing a function
	Function CreateFunctionResponseFunction `json:"function,required"`
	JSON     createFunctionResponseJSON     `json:"-"`
}

// createFunctionResponseJSON contains the JSON metadata for the struct
// [CreateFunctionResponse]
type createFunctionResponseJSON struct {
	Function    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateFunctionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createFunctionResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object (DTO) representing a function
type CreateFunctionResponseFunction struct {
	// Unique function id
	ID string `json:"id,required" format:"uuid"`
	// Function creation timestamp
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Used to indicate a STREAMING function. Defaults to DEFAULT.
	FunctionType CreateFunctionResponseFunctionFunctionType `json:"functionType,required"`
	// Health endpoint for the container or helmChart
	HealthUri string `json:"healthUri,required" format:"uri"`
	// Function name
	Name string `json:"name,required"`
	// NVIDIA Cloud Account Id
	NcaID string `json:"ncaId,required"`
	// Function status
	Status CreateFunctionResponseFunctionStatus `json:"status,required"`
	// Unique function version id
	VersionID string `json:"versionId,required" format:"uuid"`
	// List of active instances for this function.
	ActiveInstances []CreateFunctionResponseFunctionActiveInstance `json:"activeInstances"`
	// Invocation request body format
	APIBodyFormat CreateFunctionResponseFunctionAPIBodyFormat `json:"apiBodyFormat"`
	// Args used to launch the container
	ContainerArgs string `json:"containerArgs"`
	// Environment settings used to launch the container
	ContainerEnvironment []CreateFunctionResponseFunctionContainerEnvironment `json:"containerEnvironment"`
	// Optional custom container
	ContainerImage string `json:"containerImage" format:"uri"`
	// Function/version description
	Description string `json:"description"`
	// Data Transfer Object(DTO) representing a function ne
	Health CreateFunctionResponseFunctionHealth `json:"health"`
	// Optional Helm Chart
	HelmChart string `json:"helmChart" format:"uri"`
	// Helm Chart Service Name specified only when helmChart property is specified
	HelmChartServiceName string `json:"helmChartServiceName"`
	// Optional port number where the inference listener is running - defaults to 8000
	// for Triton
	InferencePort int64 `json:"inferencePort"`
	// Entrypoint for invoking the container to process requests
	InferenceURL string `json:"inferenceUrl" format:"uri"`
	// Optional set of models
	Models []CreateFunctionResponseFunctionModel `json:"models"`
	// Indicates whether the function is owned by another account. If the account that
	// is being used to lookup functions happens to be authorized to invoke/list this
	// function which is owned by a different account, then this field is set to true
	// and ncaId will contain the id of the account that owns the function. Otherwise,
	// this field is not set as it defaults to false.
	OwnedByDifferentAccount bool `json:"ownedByDifferentAccount"`
	// Optional set of resources.
	Resources []CreateFunctionResponseFunctionResource `json:"resources"`
	// Optional secret names
	Secrets []string `json:"secrets"`
	// Optional set of tags. Maximum allowed number of tags per function is 64. Maximum
	// length of each tag is 128 chars.
	Tags []string                           `json:"tags"`
	JSON createFunctionResponseFunctionJSON `json:"-"`
}

// createFunctionResponseFunctionJSON contains the JSON metadata for the struct
// [CreateFunctionResponseFunction]
type createFunctionResponseFunctionJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	FunctionType            apijson.Field
	HealthUri               apijson.Field
	Name                    apijson.Field
	NcaID                   apijson.Field
	Status                  apijson.Field
	VersionID               apijson.Field
	ActiveInstances         apijson.Field
	APIBodyFormat           apijson.Field
	ContainerArgs           apijson.Field
	ContainerEnvironment    apijson.Field
	ContainerImage          apijson.Field
	Description             apijson.Field
	Health                  apijson.Field
	HelmChart               apijson.Field
	HelmChartServiceName    apijson.Field
	InferencePort           apijson.Field
	InferenceURL            apijson.Field
	Models                  apijson.Field
	OwnedByDifferentAccount apijson.Field
	Resources               apijson.Field
	Secrets                 apijson.Field
	Tags                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CreateFunctionResponseFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createFunctionResponseFunctionJSON) RawJSON() string {
	return r.raw
}

// Used to indicate a STREAMING function. Defaults to DEFAULT.
type CreateFunctionResponseFunctionFunctionType string

const (
	CreateFunctionResponseFunctionFunctionTypeDefault   CreateFunctionResponseFunctionFunctionType = "DEFAULT"
	CreateFunctionResponseFunctionFunctionTypeStreaming CreateFunctionResponseFunctionFunctionType = "STREAMING"
)

func (r CreateFunctionResponseFunctionFunctionType) IsKnown() bool {
	switch r {
	case CreateFunctionResponseFunctionFunctionTypeDefault, CreateFunctionResponseFunctionFunctionTypeStreaming:
		return true
	}
	return false
}

// Function status
type CreateFunctionResponseFunctionStatus string

const (
	CreateFunctionResponseFunctionStatusActive    CreateFunctionResponseFunctionStatus = "ACTIVE"
	CreateFunctionResponseFunctionStatusDeploying CreateFunctionResponseFunctionStatus = "DEPLOYING"
	CreateFunctionResponseFunctionStatusError     CreateFunctionResponseFunctionStatus = "ERROR"
	CreateFunctionResponseFunctionStatusInactive  CreateFunctionResponseFunctionStatus = "INACTIVE"
	CreateFunctionResponseFunctionStatusDeleted   CreateFunctionResponseFunctionStatus = "DELETED"
)

func (r CreateFunctionResponseFunctionStatus) IsKnown() bool {
	switch r {
	case CreateFunctionResponseFunctionStatusActive, CreateFunctionResponseFunctionStatusDeploying, CreateFunctionResponseFunctionStatusError, CreateFunctionResponseFunctionStatusInactive, CreateFunctionResponseFunctionStatusDeleted:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a spot instance
type CreateFunctionResponseFunctionActiveInstance struct {
	// Backend where the instance is running
	Backend string `json:"backend"`
	// Function executing on the instance
	FunctionID string `json:"functionId" format:"uuid"`
	// Function version executing on the instance
	FunctionVersionID string `json:"functionVersionId" format:"uuid"`
	// GPU name powering the instance
	GPU string `json:"gpu"`
	// Instance creation timestamp
	InstanceCreatedAt time.Time `json:"instanceCreatedAt" format:"date-time"`
	// Unique id of the instance
	InstanceID string `json:"instanceId"`
	// Instance status
	InstanceStatus CreateFunctionResponseFunctionActiveInstancesInstanceStatus `json:"instanceStatus"`
	// GPU instance-type powering the instance
	InstanceType string `json:"instanceType"`
	// Instance's last updated timestamp
	InstanceUpdatedAt time.Time `json:"instanceUpdatedAt" format:"date-time"`
	// Location such as zone name or region where the instance is running
	Location string `json:"location"`
	// NVIDIA Cloud Account Id that owns the function running on the instance
	NcaID string `json:"ncaId"`
	// SIS request-id used to launch this instance
	SisRequestID string                                           `json:"sisRequestId" format:"uuid"`
	JSON         createFunctionResponseFunctionActiveInstanceJSON `json:"-"`
}

// createFunctionResponseFunctionActiveInstanceJSON contains the JSON metadata for
// the struct [CreateFunctionResponseFunctionActiveInstance]
type createFunctionResponseFunctionActiveInstanceJSON struct {
	Backend           apijson.Field
	FunctionID        apijson.Field
	FunctionVersionID apijson.Field
	GPU               apijson.Field
	InstanceCreatedAt apijson.Field
	InstanceID        apijson.Field
	InstanceStatus    apijson.Field
	InstanceType      apijson.Field
	InstanceUpdatedAt apijson.Field
	Location          apijson.Field
	NcaID             apijson.Field
	SisRequestID      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CreateFunctionResponseFunctionActiveInstance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createFunctionResponseFunctionActiveInstanceJSON) RawJSON() string {
	return r.raw
}

// Instance status
type CreateFunctionResponseFunctionActiveInstancesInstanceStatus string

const (
	CreateFunctionResponseFunctionActiveInstancesInstanceStatusActive    CreateFunctionResponseFunctionActiveInstancesInstanceStatus = "ACTIVE"
	CreateFunctionResponseFunctionActiveInstancesInstanceStatusErrored   CreateFunctionResponseFunctionActiveInstancesInstanceStatus = "ERRORED"
	CreateFunctionResponseFunctionActiveInstancesInstanceStatusPreempted CreateFunctionResponseFunctionActiveInstancesInstanceStatus = "PREEMPTED"
	CreateFunctionResponseFunctionActiveInstancesInstanceStatusDeleted   CreateFunctionResponseFunctionActiveInstancesInstanceStatus = "DELETED"
)

func (r CreateFunctionResponseFunctionActiveInstancesInstanceStatus) IsKnown() bool {
	switch r {
	case CreateFunctionResponseFunctionActiveInstancesInstanceStatusActive, CreateFunctionResponseFunctionActiveInstancesInstanceStatusErrored, CreateFunctionResponseFunctionActiveInstancesInstanceStatusPreempted, CreateFunctionResponseFunctionActiveInstancesInstanceStatusDeleted:
		return true
	}
	return false
}

// Invocation request body format
type CreateFunctionResponseFunctionAPIBodyFormat string

const (
	CreateFunctionResponseFunctionAPIBodyFormatPredictV2 CreateFunctionResponseFunctionAPIBodyFormat = "PREDICT_V2"
	CreateFunctionResponseFunctionAPIBodyFormatCustom    CreateFunctionResponseFunctionAPIBodyFormat = "CUSTOM"
)

func (r CreateFunctionResponseFunctionAPIBodyFormat) IsKnown() bool {
	switch r {
	case CreateFunctionResponseFunctionAPIBodyFormatPredictV2, CreateFunctionResponseFunctionAPIBodyFormatCustom:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a container environment entry
type CreateFunctionResponseFunctionContainerEnvironment struct {
	// Container environment key
	Key string `json:"key,required"`
	// Container environment value
	Value string                                                 `json:"value,required"`
	JSON  createFunctionResponseFunctionContainerEnvironmentJSON `json:"-"`
}

// createFunctionResponseFunctionContainerEnvironmentJSON contains the JSON
// metadata for the struct [CreateFunctionResponseFunctionContainerEnvironment]
type createFunctionResponseFunctionContainerEnvironmentJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateFunctionResponseFunctionContainerEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createFunctionResponseFunctionContainerEnvironmentJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a function ne
type CreateFunctionResponseFunctionHealth struct {
	// Expected return status code considered as successful.
	ExpectedStatusCode int64 `json:"expectedStatusCode,required"`
	// Port number where the health listener is running
	Port int64 `json:"port,required"`
	// HTTP/gPRC protocol type for health endpoint
	Protocol CreateFunctionResponseFunctionHealthProtocol `json:"protocol,required"`
	// ISO 8601 duration string in PnDTnHnMn.nS format
	Timeout string `json:"timeout,required" format:"PnDTnHnMn.nS"`
	// Health endpoint for the container or the helmChart
	Uri  string                                   `json:"uri,required" format:"uri"`
	JSON createFunctionResponseFunctionHealthJSON `json:"-"`
}

// createFunctionResponseFunctionHealthJSON contains the JSON metadata for the
// struct [CreateFunctionResponseFunctionHealth]
type createFunctionResponseFunctionHealthJSON struct {
	ExpectedStatusCode apijson.Field
	Port               apijson.Field
	Protocol           apijson.Field
	Timeout            apijson.Field
	Uri                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreateFunctionResponseFunctionHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createFunctionResponseFunctionHealthJSON) RawJSON() string {
	return r.raw
}

// HTTP/gPRC protocol type for health endpoint
type CreateFunctionResponseFunctionHealthProtocol string

const (
	CreateFunctionResponseFunctionHealthProtocolHTTP CreateFunctionResponseFunctionHealthProtocol = "HTTP"
	CreateFunctionResponseFunctionHealthProtocolGRpc CreateFunctionResponseFunctionHealthProtocol = "gRPC"
)

func (r CreateFunctionResponseFunctionHealthProtocol) IsKnown() bool {
	switch r {
	case CreateFunctionResponseFunctionHealthProtocolHTTP, CreateFunctionResponseFunctionHealthProtocolGRpc:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing an artifact
type CreateFunctionResponseFunctionModel struct {
	// Artifact name
	Name string `json:"name,required"`
	// Artifact URI
	Uri string `json:"uri,required" format:"uri"`
	// Artifact version
	Version string                                  `json:"version,required"`
	JSON    createFunctionResponseFunctionModelJSON `json:"-"`
}

// createFunctionResponseFunctionModelJSON contains the JSON metadata for the
// struct [CreateFunctionResponseFunctionModel]
type createFunctionResponseFunctionModelJSON struct {
	Name        apijson.Field
	Uri         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateFunctionResponseFunctionModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createFunctionResponseFunctionModelJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an artifact
type CreateFunctionResponseFunctionResource struct {
	// Artifact name
	Name string `json:"name,required"`
	// Artifact URI
	Uri string `json:"uri,required" format:"uri"`
	// Artifact version
	Version string                                     `json:"version,required"`
	JSON    createFunctionResponseFunctionResourceJSON `json:"-"`
}

// createFunctionResponseFunctionResourceJSON contains the JSON metadata for the
// struct [CreateFunctionResponseFunctionResource]
type createFunctionResponseFunctionResourceJSON struct {
	Name        apijson.Field
	Uri         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateFunctionResponseFunctionResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createFunctionResponseFunctionResourceJSON) RawJSON() string {
	return r.raw
}

// Response body with function details
type FunctionResponse struct {
	// Data Transfer Object (DTO) representing a function
	Function FunctionResponseFunction `json:"function,required"`
	JSON     functionResponseJSON     `json:"-"`
}

// functionResponseJSON contains the JSON metadata for the struct
// [FunctionResponse]
type functionResponseJSON struct {
	Function    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object (DTO) representing a function
type FunctionResponseFunction struct {
	// Unique function id
	ID string `json:"id,required" format:"uuid"`
	// Function creation timestamp
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Used to indicate a STREAMING function. Defaults to DEFAULT.
	FunctionType FunctionResponseFunctionFunctionType `json:"functionType,required"`
	// Health endpoint for the container or helmChart
	HealthUri string `json:"healthUri,required" format:"uri"`
	// Function name
	Name string `json:"name,required"`
	// NVIDIA Cloud Account Id
	NcaID string `json:"ncaId,required"`
	// Function status
	Status FunctionResponseFunctionStatus `json:"status,required"`
	// Unique function version id
	VersionID string `json:"versionId,required" format:"uuid"`
	// List of active instances for this function.
	ActiveInstances []FunctionResponseFunctionActiveInstance `json:"activeInstances"`
	// Invocation request body format
	APIBodyFormat FunctionResponseFunctionAPIBodyFormat `json:"apiBodyFormat"`
	// Args used to launch the container
	ContainerArgs string `json:"containerArgs"`
	// Environment settings used to launch the container
	ContainerEnvironment []FunctionResponseFunctionContainerEnvironment `json:"containerEnvironment"`
	// Optional custom container
	ContainerImage string `json:"containerImage" format:"uri"`
	// Function/version description
	Description string `json:"description"`
	// Data Transfer Object(DTO) representing a function ne
	Health FunctionResponseFunctionHealth `json:"health"`
	// Optional Helm Chart
	HelmChart string `json:"helmChart" format:"uri"`
	// Helm Chart Service Name specified only when helmChart property is specified
	HelmChartServiceName string `json:"helmChartServiceName"`
	// Optional port number where the inference listener is running - defaults to 8000
	// for Triton
	InferencePort int64 `json:"inferencePort"`
	// Entrypoint for invoking the container to process requests
	InferenceURL string `json:"inferenceUrl" format:"uri"`
	// Optional set of models
	Models []FunctionResponseFunctionModel `json:"models"`
	// Indicates whether the function is owned by another account. If the account that
	// is being used to lookup functions happens to be authorized to invoke/list this
	// function which is owned by a different account, then this field is set to true
	// and ncaId will contain the id of the account that owns the function. Otherwise,
	// this field is not set as it defaults to false.
	OwnedByDifferentAccount bool `json:"ownedByDifferentAccount"`
	// Optional set of resources.
	Resources []FunctionResponseFunctionResource `json:"resources"`
	// Optional secret names
	Secrets []string `json:"secrets"`
	// Optional set of tags. Maximum allowed number of tags per function is 64. Maximum
	// length of each tag is 128 chars.
	Tags []string                     `json:"tags"`
	JSON functionResponseFunctionJSON `json:"-"`
}

// functionResponseFunctionJSON contains the JSON metadata for the struct
// [FunctionResponseFunction]
type functionResponseFunctionJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	FunctionType            apijson.Field
	HealthUri               apijson.Field
	Name                    apijson.Field
	NcaID                   apijson.Field
	Status                  apijson.Field
	VersionID               apijson.Field
	ActiveInstances         apijson.Field
	APIBodyFormat           apijson.Field
	ContainerArgs           apijson.Field
	ContainerEnvironment    apijson.Field
	ContainerImage          apijson.Field
	Description             apijson.Field
	Health                  apijson.Field
	HelmChart               apijson.Field
	HelmChartServiceName    apijson.Field
	InferencePort           apijson.Field
	InferenceURL            apijson.Field
	Models                  apijson.Field
	OwnedByDifferentAccount apijson.Field
	Resources               apijson.Field
	Secrets                 apijson.Field
	Tags                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *FunctionResponseFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionResponseFunctionJSON) RawJSON() string {
	return r.raw
}

// Used to indicate a STREAMING function. Defaults to DEFAULT.
type FunctionResponseFunctionFunctionType string

const (
	FunctionResponseFunctionFunctionTypeDefault   FunctionResponseFunctionFunctionType = "DEFAULT"
	FunctionResponseFunctionFunctionTypeStreaming FunctionResponseFunctionFunctionType = "STREAMING"
)

func (r FunctionResponseFunctionFunctionType) IsKnown() bool {
	switch r {
	case FunctionResponseFunctionFunctionTypeDefault, FunctionResponseFunctionFunctionTypeStreaming:
		return true
	}
	return false
}

// Function status
type FunctionResponseFunctionStatus string

const (
	FunctionResponseFunctionStatusActive    FunctionResponseFunctionStatus = "ACTIVE"
	FunctionResponseFunctionStatusDeploying FunctionResponseFunctionStatus = "DEPLOYING"
	FunctionResponseFunctionStatusError     FunctionResponseFunctionStatus = "ERROR"
	FunctionResponseFunctionStatusInactive  FunctionResponseFunctionStatus = "INACTIVE"
	FunctionResponseFunctionStatusDeleted   FunctionResponseFunctionStatus = "DELETED"
)

func (r FunctionResponseFunctionStatus) IsKnown() bool {
	switch r {
	case FunctionResponseFunctionStatusActive, FunctionResponseFunctionStatusDeploying, FunctionResponseFunctionStatusError, FunctionResponseFunctionStatusInactive, FunctionResponseFunctionStatusDeleted:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a spot instance
type FunctionResponseFunctionActiveInstance struct {
	// Backend where the instance is running
	Backend string `json:"backend"`
	// Function executing on the instance
	FunctionID string `json:"functionId" format:"uuid"`
	// Function version executing on the instance
	FunctionVersionID string `json:"functionVersionId" format:"uuid"`
	// GPU name powering the instance
	GPU string `json:"gpu"`
	// Instance creation timestamp
	InstanceCreatedAt time.Time `json:"instanceCreatedAt" format:"date-time"`
	// Unique id of the instance
	InstanceID string `json:"instanceId"`
	// Instance status
	InstanceStatus FunctionResponseFunctionActiveInstancesInstanceStatus `json:"instanceStatus"`
	// GPU instance-type powering the instance
	InstanceType string `json:"instanceType"`
	// Instance's last updated timestamp
	InstanceUpdatedAt time.Time `json:"instanceUpdatedAt" format:"date-time"`
	// Location such as zone name or region where the instance is running
	Location string `json:"location"`
	// NVIDIA Cloud Account Id that owns the function running on the instance
	NcaID string `json:"ncaId"`
	// SIS request-id used to launch this instance
	SisRequestID string                                     `json:"sisRequestId" format:"uuid"`
	JSON         functionResponseFunctionActiveInstanceJSON `json:"-"`
}

// functionResponseFunctionActiveInstanceJSON contains the JSON metadata for the
// struct [FunctionResponseFunctionActiveInstance]
type functionResponseFunctionActiveInstanceJSON struct {
	Backend           apijson.Field
	FunctionID        apijson.Field
	FunctionVersionID apijson.Field
	GPU               apijson.Field
	InstanceCreatedAt apijson.Field
	InstanceID        apijson.Field
	InstanceStatus    apijson.Field
	InstanceType      apijson.Field
	InstanceUpdatedAt apijson.Field
	Location          apijson.Field
	NcaID             apijson.Field
	SisRequestID      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *FunctionResponseFunctionActiveInstance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionResponseFunctionActiveInstanceJSON) RawJSON() string {
	return r.raw
}

// Instance status
type FunctionResponseFunctionActiveInstancesInstanceStatus string

const (
	FunctionResponseFunctionActiveInstancesInstanceStatusActive    FunctionResponseFunctionActiveInstancesInstanceStatus = "ACTIVE"
	FunctionResponseFunctionActiveInstancesInstanceStatusErrored   FunctionResponseFunctionActiveInstancesInstanceStatus = "ERRORED"
	FunctionResponseFunctionActiveInstancesInstanceStatusPreempted FunctionResponseFunctionActiveInstancesInstanceStatus = "PREEMPTED"
	FunctionResponseFunctionActiveInstancesInstanceStatusDeleted   FunctionResponseFunctionActiveInstancesInstanceStatus = "DELETED"
)

func (r FunctionResponseFunctionActiveInstancesInstanceStatus) IsKnown() bool {
	switch r {
	case FunctionResponseFunctionActiveInstancesInstanceStatusActive, FunctionResponseFunctionActiveInstancesInstanceStatusErrored, FunctionResponseFunctionActiveInstancesInstanceStatusPreempted, FunctionResponseFunctionActiveInstancesInstanceStatusDeleted:
		return true
	}
	return false
}

// Invocation request body format
type FunctionResponseFunctionAPIBodyFormat string

const (
	FunctionResponseFunctionAPIBodyFormatPredictV2 FunctionResponseFunctionAPIBodyFormat = "PREDICT_V2"
	FunctionResponseFunctionAPIBodyFormatCustom    FunctionResponseFunctionAPIBodyFormat = "CUSTOM"
)

func (r FunctionResponseFunctionAPIBodyFormat) IsKnown() bool {
	switch r {
	case FunctionResponseFunctionAPIBodyFormatPredictV2, FunctionResponseFunctionAPIBodyFormatCustom:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a container environment entry
type FunctionResponseFunctionContainerEnvironment struct {
	// Container environment key
	Key string `json:"key,required"`
	// Container environment value
	Value string                                           `json:"value,required"`
	JSON  functionResponseFunctionContainerEnvironmentJSON `json:"-"`
}

// functionResponseFunctionContainerEnvironmentJSON contains the JSON metadata for
// the struct [FunctionResponseFunctionContainerEnvironment]
type functionResponseFunctionContainerEnvironmentJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionResponseFunctionContainerEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionResponseFunctionContainerEnvironmentJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a function ne
type FunctionResponseFunctionHealth struct {
	// Expected return status code considered as successful.
	ExpectedStatusCode int64 `json:"expectedStatusCode,required"`
	// Port number where the health listener is running
	Port int64 `json:"port,required"`
	// HTTP/gPRC protocol type for health endpoint
	Protocol FunctionResponseFunctionHealthProtocol `json:"protocol,required"`
	// ISO 8601 duration string in PnDTnHnMn.nS format
	Timeout string `json:"timeout,required" format:"PnDTnHnMn.nS"`
	// Health endpoint for the container or the helmChart
	Uri  string                             `json:"uri,required" format:"uri"`
	JSON functionResponseFunctionHealthJSON `json:"-"`
}

// functionResponseFunctionHealthJSON contains the JSON metadata for the struct
// [FunctionResponseFunctionHealth]
type functionResponseFunctionHealthJSON struct {
	ExpectedStatusCode apijson.Field
	Port               apijson.Field
	Protocol           apijson.Field
	Timeout            apijson.Field
	Uri                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *FunctionResponseFunctionHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionResponseFunctionHealthJSON) RawJSON() string {
	return r.raw
}

// HTTP/gPRC protocol type for health endpoint
type FunctionResponseFunctionHealthProtocol string

const (
	FunctionResponseFunctionHealthProtocolHTTP FunctionResponseFunctionHealthProtocol = "HTTP"
	FunctionResponseFunctionHealthProtocolGRpc FunctionResponseFunctionHealthProtocol = "gRPC"
)

func (r FunctionResponseFunctionHealthProtocol) IsKnown() bool {
	switch r {
	case FunctionResponseFunctionHealthProtocolHTTP, FunctionResponseFunctionHealthProtocolGRpc:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing an artifact
type FunctionResponseFunctionModel struct {
	// Artifact name
	Name string `json:"name,required"`
	// Artifact URI
	Uri string `json:"uri,required" format:"uri"`
	// Artifact version
	Version string                            `json:"version,required"`
	JSON    functionResponseFunctionModelJSON `json:"-"`
}

// functionResponseFunctionModelJSON contains the JSON metadata for the struct
// [FunctionResponseFunctionModel]
type functionResponseFunctionModelJSON struct {
	Name        apijson.Field
	Uri         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionResponseFunctionModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionResponseFunctionModelJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an artifact
type FunctionResponseFunctionResource struct {
	// Artifact name
	Name string `json:"name,required"`
	// Artifact URI
	Uri string `json:"uri,required" format:"uri"`
	// Artifact version
	Version string                               `json:"version,required"`
	JSON    functionResponseFunctionResourceJSON `json:"-"`
}

// functionResponseFunctionResourceJSON contains the JSON metadata for the struct
// [FunctionResponseFunctionResource]
type functionResponseFunctionResourceJSON struct {
	Name        apijson.Field
	Uri         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionResponseFunctionResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionResponseFunctionResourceJSON) RawJSON() string {
	return r.raw
}

// Request queue details of all the functions with same id in an account
type GetQueuesResponse struct {
	// Function id
	FunctionID string `json:"functionId,required" format:"uuid"`
	// Details of all the queues associated with same named functions
	Queues []GetQueuesResponseQueue `json:"queues,required"`
	JSON   getQueuesResponseJSON    `json:"-"`
}

// getQueuesResponseJSON contains the JSON metadata for the struct
// [GetQueuesResponse]
type getQueuesResponseJSON struct {
	FunctionID  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetQueuesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getQueuesResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a request queue for function version
type GetQueuesResponseQueue struct {
	// Function name
	FunctionName string `json:"functionName,required"`
	// Function status
	FunctionStatus GetQueuesResponseQueuesFunctionStatus `json:"functionStatus,required"`
	// Function version id
	FunctionVersionID string `json:"functionVersionId,required" format:"uuid"`
	// Approximate number of messages in the request queue
	QueueDepth int64                      `json:"queueDepth"`
	JSON       getQueuesResponseQueueJSON `json:"-"`
}

// getQueuesResponseQueueJSON contains the JSON metadata for the struct
// [GetQueuesResponseQueue]
type getQueuesResponseQueueJSON struct {
	FunctionName      apijson.Field
	FunctionStatus    apijson.Field
	FunctionVersionID apijson.Field
	QueueDepth        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *GetQueuesResponseQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getQueuesResponseQueueJSON) RawJSON() string {
	return r.raw
}

// Function status
type GetQueuesResponseQueuesFunctionStatus string

const (
	GetQueuesResponseQueuesFunctionStatusActive    GetQueuesResponseQueuesFunctionStatus = "ACTIVE"
	GetQueuesResponseQueuesFunctionStatusDeploying GetQueuesResponseQueuesFunctionStatus = "DEPLOYING"
	GetQueuesResponseQueuesFunctionStatusError     GetQueuesResponseQueuesFunctionStatus = "ERROR"
	GetQueuesResponseQueuesFunctionStatusInactive  GetQueuesResponseQueuesFunctionStatus = "INACTIVE"
	GetQueuesResponseQueuesFunctionStatusDeleted   GetQueuesResponseQueuesFunctionStatus = "DELETED"
)

func (r GetQueuesResponseQueuesFunctionStatus) IsKnown() bool {
	switch r {
	case GetQueuesResponseQueuesFunctionStatusActive, GetQueuesResponseQueuesFunctionStatusDeploying, GetQueuesResponseQueuesFunctionStatusError, GetQueuesResponseQueuesFunctionStatusInactive, GetQueuesResponseQueuesFunctionStatusDeleted:
		return true
	}
	return false
}

// Response body with result from a request for executing a job/task as a cloud
// function using a GPU powered spot/on-demand instance.
type InvokeFunctionResponse struct {
	// Error code from the container while executing cloud function.
	ErrorCode int64 `json:"errorCode"`
	// Progress indicator for the task/job executing cloud function.
	PercentComplete int64 `json:"percentComplete"`
	// Request id
	ReqID string `json:"reqId" format:"uuid"`
	// Response/result of size < 5MB size for the task/job executing cloud function.
	Response string `json:"response"`
	// For large results, responseReference will be a pre-signeddownload URL.
	ResponseReference string `json:"responseReference" format:"url"`
	// Status of the task/job executing cloud function.
	Status InvokeFunctionResponseStatus `json:"status"`
	JSON   invokeFunctionResponseJSON   `json:"-"`
}

// invokeFunctionResponseJSON contains the JSON metadata for the struct
// [InvokeFunctionResponse]
type invokeFunctionResponseJSON struct {
	ErrorCode         apijson.Field
	PercentComplete   apijson.Field
	ReqID             apijson.Field
	Response          apijson.Field
	ResponseReference apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvokeFunctionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invokeFunctionResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the task/job executing cloud function.
type InvokeFunctionResponseStatus string

const (
	InvokeFunctionResponseStatusErrored           InvokeFunctionResponseStatus = "errored"
	InvokeFunctionResponseStatusInProgress        InvokeFunctionResponseStatus = "in-progress"
	InvokeFunctionResponseStatusFulfilled         InvokeFunctionResponseStatus = "fulfilled"
	InvokeFunctionResponseStatusPendingEvaluation InvokeFunctionResponseStatus = "pending-evaluation"
	InvokeFunctionResponseStatusRejected          InvokeFunctionResponseStatus = "rejected"
)

func (r InvokeFunctionResponseStatus) IsKnown() bool {
	switch r {
	case InvokeFunctionResponseStatusErrored, InvokeFunctionResponseStatusInProgress, InvokeFunctionResponseStatusFulfilled, InvokeFunctionResponseStatusPendingEvaluation, InvokeFunctionResponseStatusRejected:
		return true
	}
	return false
}

// Response body containing list of functions
type ListFunctionsResponse struct {
	// List of functions
	Functions []ListFunctionsResponseFunction `json:"functions,required"`
	JSON      listFunctionsResponseJSON       `json:"-"`
}

// listFunctionsResponseJSON contains the JSON metadata for the struct
// [ListFunctionsResponse]
type listFunctionsResponseJSON struct {
	Functions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListFunctionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listFunctionsResponseJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object (DTO) representing a function
type ListFunctionsResponseFunction struct {
	// Unique function id
	ID string `json:"id,required" format:"uuid"`
	// Function creation timestamp
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Used to indicate a STREAMING function. Defaults to DEFAULT.
	FunctionType ListFunctionsResponseFunctionsFunctionType `json:"functionType,required"`
	// Health endpoint for the container or helmChart
	HealthUri string `json:"healthUri,required" format:"uri"`
	// Function name
	Name string `json:"name,required"`
	// NVIDIA Cloud Account Id
	NcaID string `json:"ncaId,required"`
	// Function status
	Status ListFunctionsResponseFunctionsStatus `json:"status,required"`
	// Unique function version id
	VersionID string `json:"versionId,required" format:"uuid"`
	// List of active instances for this function.
	ActiveInstances []ListFunctionsResponseFunctionsActiveInstance `json:"activeInstances"`
	// Invocation request body format
	APIBodyFormat ListFunctionsResponseFunctionsAPIBodyFormat `json:"apiBodyFormat"`
	// Args used to launch the container
	ContainerArgs string `json:"containerArgs"`
	// Environment settings used to launch the container
	ContainerEnvironment []ListFunctionsResponseFunctionsContainerEnvironment `json:"containerEnvironment"`
	// Optional custom container
	ContainerImage string `json:"containerImage" format:"uri"`
	// Function/version description
	Description string `json:"description"`
	// Data Transfer Object(DTO) representing a function ne
	Health ListFunctionsResponseFunctionsHealth `json:"health"`
	// Optional Helm Chart
	HelmChart string `json:"helmChart" format:"uri"`
	// Helm Chart Service Name specified only when helmChart property is specified
	HelmChartServiceName string `json:"helmChartServiceName"`
	// Optional port number where the inference listener is running - defaults to 8000
	// for Triton
	InferencePort int64 `json:"inferencePort"`
	// Entrypoint for invoking the container to process requests
	InferenceURL string `json:"inferenceUrl" format:"uri"`
	// Optional set of models
	Models []ListFunctionsResponseFunctionsModel `json:"models"`
	// Indicates whether the function is owned by another account. If the account that
	// is being used to lookup functions happens to be authorized to invoke/list this
	// function which is owned by a different account, then this field is set to true
	// and ncaId will contain the id of the account that owns the function. Otherwise,
	// this field is not set as it defaults to false.
	OwnedByDifferentAccount bool `json:"ownedByDifferentAccount"`
	// Optional set of resources.
	Resources []ListFunctionsResponseFunctionsResource `json:"resources"`
	// Optional secret names
	Secrets []string `json:"secrets"`
	// Optional set of tags. Maximum allowed number of tags per function is 64. Maximum
	// length of each tag is 128 chars.
	Tags []string                          `json:"tags"`
	JSON listFunctionsResponseFunctionJSON `json:"-"`
}

// listFunctionsResponseFunctionJSON contains the JSON metadata for the struct
// [ListFunctionsResponseFunction]
type listFunctionsResponseFunctionJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	FunctionType            apijson.Field
	HealthUri               apijson.Field
	Name                    apijson.Field
	NcaID                   apijson.Field
	Status                  apijson.Field
	VersionID               apijson.Field
	ActiveInstances         apijson.Field
	APIBodyFormat           apijson.Field
	ContainerArgs           apijson.Field
	ContainerEnvironment    apijson.Field
	ContainerImage          apijson.Field
	Description             apijson.Field
	Health                  apijson.Field
	HelmChart               apijson.Field
	HelmChartServiceName    apijson.Field
	InferencePort           apijson.Field
	InferenceURL            apijson.Field
	Models                  apijson.Field
	OwnedByDifferentAccount apijson.Field
	Resources               apijson.Field
	Secrets                 apijson.Field
	Tags                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ListFunctionsResponseFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listFunctionsResponseFunctionJSON) RawJSON() string {
	return r.raw
}

// Used to indicate a STREAMING function. Defaults to DEFAULT.
type ListFunctionsResponseFunctionsFunctionType string

const (
	ListFunctionsResponseFunctionsFunctionTypeDefault   ListFunctionsResponseFunctionsFunctionType = "DEFAULT"
	ListFunctionsResponseFunctionsFunctionTypeStreaming ListFunctionsResponseFunctionsFunctionType = "STREAMING"
)

func (r ListFunctionsResponseFunctionsFunctionType) IsKnown() bool {
	switch r {
	case ListFunctionsResponseFunctionsFunctionTypeDefault, ListFunctionsResponseFunctionsFunctionTypeStreaming:
		return true
	}
	return false
}

// Function status
type ListFunctionsResponseFunctionsStatus string

const (
	ListFunctionsResponseFunctionsStatusActive    ListFunctionsResponseFunctionsStatus = "ACTIVE"
	ListFunctionsResponseFunctionsStatusDeploying ListFunctionsResponseFunctionsStatus = "DEPLOYING"
	ListFunctionsResponseFunctionsStatusError     ListFunctionsResponseFunctionsStatus = "ERROR"
	ListFunctionsResponseFunctionsStatusInactive  ListFunctionsResponseFunctionsStatus = "INACTIVE"
	ListFunctionsResponseFunctionsStatusDeleted   ListFunctionsResponseFunctionsStatus = "DELETED"
)

func (r ListFunctionsResponseFunctionsStatus) IsKnown() bool {
	switch r {
	case ListFunctionsResponseFunctionsStatusActive, ListFunctionsResponseFunctionsStatusDeploying, ListFunctionsResponseFunctionsStatusError, ListFunctionsResponseFunctionsStatusInactive, ListFunctionsResponseFunctionsStatusDeleted:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a spot instance
type ListFunctionsResponseFunctionsActiveInstance struct {
	// Backend where the instance is running
	Backend string `json:"backend"`
	// Function executing on the instance
	FunctionID string `json:"functionId" format:"uuid"`
	// Function version executing on the instance
	FunctionVersionID string `json:"functionVersionId" format:"uuid"`
	// GPU name powering the instance
	GPU string `json:"gpu"`
	// Instance creation timestamp
	InstanceCreatedAt time.Time `json:"instanceCreatedAt" format:"date-time"`
	// Unique id of the instance
	InstanceID string `json:"instanceId"`
	// Instance status
	InstanceStatus ListFunctionsResponseFunctionsActiveInstancesInstanceStatus `json:"instanceStatus"`
	// GPU instance-type powering the instance
	InstanceType string `json:"instanceType"`
	// Instance's last updated timestamp
	InstanceUpdatedAt time.Time `json:"instanceUpdatedAt" format:"date-time"`
	// Location such as zone name or region where the instance is running
	Location string `json:"location"`
	// NVIDIA Cloud Account Id that owns the function running on the instance
	NcaID string `json:"ncaId"`
	// SIS request-id used to launch this instance
	SisRequestID string                                           `json:"sisRequestId" format:"uuid"`
	JSON         listFunctionsResponseFunctionsActiveInstanceJSON `json:"-"`
}

// listFunctionsResponseFunctionsActiveInstanceJSON contains the JSON metadata for
// the struct [ListFunctionsResponseFunctionsActiveInstance]
type listFunctionsResponseFunctionsActiveInstanceJSON struct {
	Backend           apijson.Field
	FunctionID        apijson.Field
	FunctionVersionID apijson.Field
	GPU               apijson.Field
	InstanceCreatedAt apijson.Field
	InstanceID        apijson.Field
	InstanceStatus    apijson.Field
	InstanceType      apijson.Field
	InstanceUpdatedAt apijson.Field
	Location          apijson.Field
	NcaID             apijson.Field
	SisRequestID      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ListFunctionsResponseFunctionsActiveInstance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listFunctionsResponseFunctionsActiveInstanceJSON) RawJSON() string {
	return r.raw
}

// Instance status
type ListFunctionsResponseFunctionsActiveInstancesInstanceStatus string

const (
	ListFunctionsResponseFunctionsActiveInstancesInstanceStatusActive    ListFunctionsResponseFunctionsActiveInstancesInstanceStatus = "ACTIVE"
	ListFunctionsResponseFunctionsActiveInstancesInstanceStatusErrored   ListFunctionsResponseFunctionsActiveInstancesInstanceStatus = "ERRORED"
	ListFunctionsResponseFunctionsActiveInstancesInstanceStatusPreempted ListFunctionsResponseFunctionsActiveInstancesInstanceStatus = "PREEMPTED"
	ListFunctionsResponseFunctionsActiveInstancesInstanceStatusDeleted   ListFunctionsResponseFunctionsActiveInstancesInstanceStatus = "DELETED"
)

func (r ListFunctionsResponseFunctionsActiveInstancesInstanceStatus) IsKnown() bool {
	switch r {
	case ListFunctionsResponseFunctionsActiveInstancesInstanceStatusActive, ListFunctionsResponseFunctionsActiveInstancesInstanceStatusErrored, ListFunctionsResponseFunctionsActiveInstancesInstanceStatusPreempted, ListFunctionsResponseFunctionsActiveInstancesInstanceStatusDeleted:
		return true
	}
	return false
}

// Invocation request body format
type ListFunctionsResponseFunctionsAPIBodyFormat string

const (
	ListFunctionsResponseFunctionsAPIBodyFormatPredictV2 ListFunctionsResponseFunctionsAPIBodyFormat = "PREDICT_V2"
	ListFunctionsResponseFunctionsAPIBodyFormatCustom    ListFunctionsResponseFunctionsAPIBodyFormat = "CUSTOM"
)

func (r ListFunctionsResponseFunctionsAPIBodyFormat) IsKnown() bool {
	switch r {
	case ListFunctionsResponseFunctionsAPIBodyFormatPredictV2, ListFunctionsResponseFunctionsAPIBodyFormatCustom:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing a container environment entry
type ListFunctionsResponseFunctionsContainerEnvironment struct {
	// Container environment key
	Key string `json:"key,required"`
	// Container environment value
	Value string                                                 `json:"value,required"`
	JSON  listFunctionsResponseFunctionsContainerEnvironmentJSON `json:"-"`
}

// listFunctionsResponseFunctionsContainerEnvironmentJSON contains the JSON
// metadata for the struct [ListFunctionsResponseFunctionsContainerEnvironment]
type listFunctionsResponseFunctionsContainerEnvironmentJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListFunctionsResponseFunctionsContainerEnvironment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listFunctionsResponseFunctionsContainerEnvironmentJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing a function ne
type ListFunctionsResponseFunctionsHealth struct {
	// Expected return status code considered as successful.
	ExpectedStatusCode int64 `json:"expectedStatusCode,required"`
	// Port number where the health listener is running
	Port int64 `json:"port,required"`
	// HTTP/gPRC protocol type for health endpoint
	Protocol ListFunctionsResponseFunctionsHealthProtocol `json:"protocol,required"`
	// ISO 8601 duration string in PnDTnHnMn.nS format
	Timeout string `json:"timeout,required" format:"PnDTnHnMn.nS"`
	// Health endpoint for the container or the helmChart
	Uri  string                                   `json:"uri,required" format:"uri"`
	JSON listFunctionsResponseFunctionsHealthJSON `json:"-"`
}

// listFunctionsResponseFunctionsHealthJSON contains the JSON metadata for the
// struct [ListFunctionsResponseFunctionsHealth]
type listFunctionsResponseFunctionsHealthJSON struct {
	ExpectedStatusCode apijson.Field
	Port               apijson.Field
	Protocol           apijson.Field
	Timeout            apijson.Field
	Uri                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ListFunctionsResponseFunctionsHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listFunctionsResponseFunctionsHealthJSON) RawJSON() string {
	return r.raw
}

// HTTP/gPRC protocol type for health endpoint
type ListFunctionsResponseFunctionsHealthProtocol string

const (
	ListFunctionsResponseFunctionsHealthProtocolHTTP ListFunctionsResponseFunctionsHealthProtocol = "HTTP"
	ListFunctionsResponseFunctionsHealthProtocolGRpc ListFunctionsResponseFunctionsHealthProtocol = "gRPC"
)

func (r ListFunctionsResponseFunctionsHealthProtocol) IsKnown() bool {
	switch r {
	case ListFunctionsResponseFunctionsHealthProtocolHTTP, ListFunctionsResponseFunctionsHealthProtocolGRpc:
		return true
	}
	return false
}

// Data Transfer Object(DTO) representing an artifact
type ListFunctionsResponseFunctionsModel struct {
	// Artifact name
	Name string `json:"name,required"`
	// Artifact URI
	Uri string `json:"uri,required" format:"uri"`
	// Artifact version
	Version string                                  `json:"version,required"`
	JSON    listFunctionsResponseFunctionsModelJSON `json:"-"`
}

// listFunctionsResponseFunctionsModelJSON contains the JSON metadata for the
// struct [ListFunctionsResponseFunctionsModel]
type listFunctionsResponseFunctionsModelJSON struct {
	Name        apijson.Field
	Uri         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListFunctionsResponseFunctionsModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listFunctionsResponseFunctionsModelJSON) RawJSON() string {
	return r.raw
}

// Data Transfer Object(DTO) representing an artifact
type ListFunctionsResponseFunctionsResource struct {
	// Artifact name
	Name string `json:"name,required"`
	// Artifact URI
	Uri string `json:"uri,required" format:"uri"`
	// Artifact version
	Version string                                     `json:"version,required"`
	JSON    listFunctionsResponseFunctionsResourceJSON `json:"-"`
}

// listFunctionsResponseFunctionsResourceJSON contains the JSON metadata for the
// struct [ListFunctionsResponseFunctionsResource]
type listFunctionsResponseFunctionsResourceJSON struct {
	Name        apijson.Field
	Uri         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListFunctionsResponseFunctionsResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listFunctionsResponseFunctionsResourceJSON) RawJSON() string {
	return r.raw
}
