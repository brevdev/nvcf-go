// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nvcf

import (
	"github.com/brevdev/nvcf-go/internal/apierror"
	"github.com/brevdev/nvcf-go/shared"
)

type Error = apierror.Error

// Parties authorized to invoke function
//
// This is an alias to an internal type.
type AuthorizedPartiesResponse = shared.AuthorizedPartiesResponse

// Data Transfer Object(DTO) representing a function with authorized accounts
//
// This is an alias to an internal type.
type AuthorizedPartiesResponseFunction = shared.AuthorizedPartiesResponseFunction

// Data Transfer Object(DTO) representing an authorized party.
//
// This is an alias to an internal type.
type AuthorizedPartiesResponseFunctionAuthorizedParty = shared.AuthorizedPartiesResponseFunctionAuthorizedParty

// Response body for create function request.
//
// This is an alias to an internal type.
type CreateFunctionResponse = shared.CreateFunctionResponse

// Data Transfer Object (DTO) representing a function
//
// This is an alias to an internal type.
type CreateFunctionResponseFunction = shared.CreateFunctionResponseFunction

// Used to indicate a STREAMING function. Defaults to DEFAULT.
//
// This is an alias to an internal type.
type CreateFunctionResponseFunctionFunctionType = shared.CreateFunctionResponseFunctionFunctionType

// This is an alias to an internal value.
const CreateFunctionResponseFunctionFunctionTypeDefault = shared.CreateFunctionResponseFunctionFunctionTypeDefault

// This is an alias to an internal value.
const CreateFunctionResponseFunctionFunctionTypeStreaming = shared.CreateFunctionResponseFunctionFunctionTypeStreaming

// Function status
//
// This is an alias to an internal type.
type CreateFunctionResponseFunctionStatus = shared.CreateFunctionResponseFunctionStatus

// This is an alias to an internal value.
const CreateFunctionResponseFunctionStatusActive = shared.CreateFunctionResponseFunctionStatusActive

// This is an alias to an internal value.
const CreateFunctionResponseFunctionStatusDeploying = shared.CreateFunctionResponseFunctionStatusDeploying

// This is an alias to an internal value.
const CreateFunctionResponseFunctionStatusError = shared.CreateFunctionResponseFunctionStatusError

// This is an alias to an internal value.
const CreateFunctionResponseFunctionStatusInactive = shared.CreateFunctionResponseFunctionStatusInactive

// This is an alias to an internal value.
const CreateFunctionResponseFunctionStatusDeleted = shared.CreateFunctionResponseFunctionStatusDeleted

// Data Transfer Object(DTO) representing a spot instance
//
// This is an alias to an internal type.
type CreateFunctionResponseFunctionActiveInstance = shared.CreateFunctionResponseFunctionActiveInstance

// Instance status
//
// This is an alias to an internal type.
type CreateFunctionResponseFunctionActiveInstancesInstanceStatus = shared.CreateFunctionResponseFunctionActiveInstancesInstanceStatus

// This is an alias to an internal value.
const CreateFunctionResponseFunctionActiveInstancesInstanceStatusActive = shared.CreateFunctionResponseFunctionActiveInstancesInstanceStatusActive

// This is an alias to an internal value.
const CreateFunctionResponseFunctionActiveInstancesInstanceStatusErrored = shared.CreateFunctionResponseFunctionActiveInstancesInstanceStatusErrored

// This is an alias to an internal value.
const CreateFunctionResponseFunctionActiveInstancesInstanceStatusPreempted = shared.CreateFunctionResponseFunctionActiveInstancesInstanceStatusPreempted

// This is an alias to an internal value.
const CreateFunctionResponseFunctionActiveInstancesInstanceStatusDeleted = shared.CreateFunctionResponseFunctionActiveInstancesInstanceStatusDeleted

// Invocation request body format
//
// This is an alias to an internal type.
type CreateFunctionResponseFunctionAPIBodyFormat = shared.CreateFunctionResponseFunctionAPIBodyFormat

// This is an alias to an internal value.
const CreateFunctionResponseFunctionAPIBodyFormatPredictV2 = shared.CreateFunctionResponseFunctionAPIBodyFormatPredictV2

// This is an alias to an internal value.
const CreateFunctionResponseFunctionAPIBodyFormatCustom = shared.CreateFunctionResponseFunctionAPIBodyFormatCustom

// Data Transfer Object(DTO) representing a container environment entry
//
// This is an alias to an internal type.
type CreateFunctionResponseFunctionContainerEnvironment = shared.CreateFunctionResponseFunctionContainerEnvironment

// Data Transfer Object(DTO) representing a function ne
//
// This is an alias to an internal type.
type CreateFunctionResponseFunctionHealth = shared.CreateFunctionResponseFunctionHealth

// HTTP/gPRC protocol type for health endpoint
//
// This is an alias to an internal type.
type CreateFunctionResponseFunctionHealthProtocol = shared.CreateFunctionResponseFunctionHealthProtocol

// This is an alias to an internal value.
const CreateFunctionResponseFunctionHealthProtocolHTTP = shared.CreateFunctionResponseFunctionHealthProtocolHTTP

// This is an alias to an internal value.
const CreateFunctionResponseFunctionHealthProtocolGRpc = shared.CreateFunctionResponseFunctionHealthProtocolGRpc

// Data Transfer Object(DTO) representing an artifact
//
// This is an alias to an internal type.
type CreateFunctionResponseFunctionModel = shared.CreateFunctionResponseFunctionModel

// Data Transfer Object(DTO) representing an artifact
//
// This is an alias to an internal type.
type CreateFunctionResponseFunctionResource = shared.CreateFunctionResponseFunctionResource

// Response body with function details
//
// This is an alias to an internal type.
type FunctionResponse = shared.FunctionResponse

// Data Transfer Object (DTO) representing a function
//
// This is an alias to an internal type.
type FunctionResponseFunction = shared.FunctionResponseFunction

// Used to indicate a STREAMING function. Defaults to DEFAULT.
//
// This is an alias to an internal type.
type FunctionResponseFunctionFunctionType = shared.FunctionResponseFunctionFunctionType

// This is an alias to an internal value.
const FunctionResponseFunctionFunctionTypeDefault = shared.FunctionResponseFunctionFunctionTypeDefault

// This is an alias to an internal value.
const FunctionResponseFunctionFunctionTypeStreaming = shared.FunctionResponseFunctionFunctionTypeStreaming

// Function status
//
// This is an alias to an internal type.
type FunctionResponseFunctionStatus = shared.FunctionResponseFunctionStatus

// This is an alias to an internal value.
const FunctionResponseFunctionStatusActive = shared.FunctionResponseFunctionStatusActive

// This is an alias to an internal value.
const FunctionResponseFunctionStatusDeploying = shared.FunctionResponseFunctionStatusDeploying

// This is an alias to an internal value.
const FunctionResponseFunctionStatusError = shared.FunctionResponseFunctionStatusError

// This is an alias to an internal value.
const FunctionResponseFunctionStatusInactive = shared.FunctionResponseFunctionStatusInactive

// This is an alias to an internal value.
const FunctionResponseFunctionStatusDeleted = shared.FunctionResponseFunctionStatusDeleted

// Data Transfer Object(DTO) representing a spot instance
//
// This is an alias to an internal type.
type FunctionResponseFunctionActiveInstance = shared.FunctionResponseFunctionActiveInstance

// Instance status
//
// This is an alias to an internal type.
type FunctionResponseFunctionActiveInstancesInstanceStatus = shared.FunctionResponseFunctionActiveInstancesInstanceStatus

// This is an alias to an internal value.
const FunctionResponseFunctionActiveInstancesInstanceStatusActive = shared.FunctionResponseFunctionActiveInstancesInstanceStatusActive

// This is an alias to an internal value.
const FunctionResponseFunctionActiveInstancesInstanceStatusErrored = shared.FunctionResponseFunctionActiveInstancesInstanceStatusErrored

// This is an alias to an internal value.
const FunctionResponseFunctionActiveInstancesInstanceStatusPreempted = shared.FunctionResponseFunctionActiveInstancesInstanceStatusPreempted

// This is an alias to an internal value.
const FunctionResponseFunctionActiveInstancesInstanceStatusDeleted = shared.FunctionResponseFunctionActiveInstancesInstanceStatusDeleted

// Invocation request body format
//
// This is an alias to an internal type.
type FunctionResponseFunctionAPIBodyFormat = shared.FunctionResponseFunctionAPIBodyFormat

// This is an alias to an internal value.
const FunctionResponseFunctionAPIBodyFormatPredictV2 = shared.FunctionResponseFunctionAPIBodyFormatPredictV2

// This is an alias to an internal value.
const FunctionResponseFunctionAPIBodyFormatCustom = shared.FunctionResponseFunctionAPIBodyFormatCustom

// Data Transfer Object(DTO) representing a container environment entry
//
// This is an alias to an internal type.
type FunctionResponseFunctionContainerEnvironment = shared.FunctionResponseFunctionContainerEnvironment

// Data Transfer Object(DTO) representing a function ne
//
// This is an alias to an internal type.
type FunctionResponseFunctionHealth = shared.FunctionResponseFunctionHealth

// HTTP/gPRC protocol type for health endpoint
//
// This is an alias to an internal type.
type FunctionResponseFunctionHealthProtocol = shared.FunctionResponseFunctionHealthProtocol

// This is an alias to an internal value.
const FunctionResponseFunctionHealthProtocolHTTP = shared.FunctionResponseFunctionHealthProtocolHTTP

// This is an alias to an internal value.
const FunctionResponseFunctionHealthProtocolGRpc = shared.FunctionResponseFunctionHealthProtocolGRpc

// Data Transfer Object(DTO) representing an artifact
//
// This is an alias to an internal type.
type FunctionResponseFunctionModel = shared.FunctionResponseFunctionModel

// Data Transfer Object(DTO) representing an artifact
//
// This is an alias to an internal type.
type FunctionResponseFunctionResource = shared.FunctionResponseFunctionResource

// Request queue details of all the functions with same id in an account
//
// This is an alias to an internal type.
type GetQueuesResponse = shared.GetQueuesResponse

// Data Transfer Object(DTO) representing a request queue for function version
//
// This is an alias to an internal type.
type GetQueuesResponseQueue = shared.GetQueuesResponseQueue

// Function status
//
// This is an alias to an internal type.
type GetQueuesResponseQueuesFunctionStatus = shared.GetQueuesResponseQueuesFunctionStatus

// This is an alias to an internal value.
const GetQueuesResponseQueuesFunctionStatusActive = shared.GetQueuesResponseQueuesFunctionStatusActive

// This is an alias to an internal value.
const GetQueuesResponseQueuesFunctionStatusDeploying = shared.GetQueuesResponseQueuesFunctionStatusDeploying

// This is an alias to an internal value.
const GetQueuesResponseQueuesFunctionStatusError = shared.GetQueuesResponseQueuesFunctionStatusError

// This is an alias to an internal value.
const GetQueuesResponseQueuesFunctionStatusInactive = shared.GetQueuesResponseQueuesFunctionStatusInactive

// This is an alias to an internal value.
const GetQueuesResponseQueuesFunctionStatusDeleted = shared.GetQueuesResponseQueuesFunctionStatusDeleted

// Response body with result from a request for executing a job/task as a cloud
// function using a GPU powered spot/on-demand instance.
//
// This is an alias to an internal type.
type InvokeFunctionResponse = shared.InvokeFunctionResponse

// Status of the task/job executing cloud function.
//
// This is an alias to an internal type.
type InvokeFunctionResponseStatus = shared.InvokeFunctionResponseStatus

// This is an alias to an internal value.
const InvokeFunctionResponseStatusErrored = shared.InvokeFunctionResponseStatusErrored

// This is an alias to an internal value.
const InvokeFunctionResponseStatusInProgress = shared.InvokeFunctionResponseStatusInProgress

// This is an alias to an internal value.
const InvokeFunctionResponseStatusFulfilled = shared.InvokeFunctionResponseStatusFulfilled

// This is an alias to an internal value.
const InvokeFunctionResponseStatusPendingEvaluation = shared.InvokeFunctionResponseStatusPendingEvaluation

// This is an alias to an internal value.
const InvokeFunctionResponseStatusRejected = shared.InvokeFunctionResponseStatusRejected

// Response body containing list of functions
//
// This is an alias to an internal type.
type ListFunctionsResponse = shared.ListFunctionsResponse

// Data Transfer Object (DTO) representing a function
//
// This is an alias to an internal type.
type ListFunctionsResponseFunction = shared.ListFunctionsResponseFunction

// Used to indicate a STREAMING function. Defaults to DEFAULT.
//
// This is an alias to an internal type.
type ListFunctionsResponseFunctionsFunctionType = shared.ListFunctionsResponseFunctionsFunctionType

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsFunctionTypeDefault = shared.ListFunctionsResponseFunctionsFunctionTypeDefault

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsFunctionTypeStreaming = shared.ListFunctionsResponseFunctionsFunctionTypeStreaming

// Function status
//
// This is an alias to an internal type.
type ListFunctionsResponseFunctionsStatus = shared.ListFunctionsResponseFunctionsStatus

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsStatusActive = shared.ListFunctionsResponseFunctionsStatusActive

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsStatusDeploying = shared.ListFunctionsResponseFunctionsStatusDeploying

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsStatusError = shared.ListFunctionsResponseFunctionsStatusError

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsStatusInactive = shared.ListFunctionsResponseFunctionsStatusInactive

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsStatusDeleted = shared.ListFunctionsResponseFunctionsStatusDeleted

// Data Transfer Object(DTO) representing a spot instance
//
// This is an alias to an internal type.
type ListFunctionsResponseFunctionsActiveInstance = shared.ListFunctionsResponseFunctionsActiveInstance

// Instance status
//
// This is an alias to an internal type.
type ListFunctionsResponseFunctionsActiveInstancesInstanceStatus = shared.ListFunctionsResponseFunctionsActiveInstancesInstanceStatus

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsActiveInstancesInstanceStatusActive = shared.ListFunctionsResponseFunctionsActiveInstancesInstanceStatusActive

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsActiveInstancesInstanceStatusErrored = shared.ListFunctionsResponseFunctionsActiveInstancesInstanceStatusErrored

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsActiveInstancesInstanceStatusPreempted = shared.ListFunctionsResponseFunctionsActiveInstancesInstanceStatusPreempted

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsActiveInstancesInstanceStatusDeleted = shared.ListFunctionsResponseFunctionsActiveInstancesInstanceStatusDeleted

// Invocation request body format
//
// This is an alias to an internal type.
type ListFunctionsResponseFunctionsAPIBodyFormat = shared.ListFunctionsResponseFunctionsAPIBodyFormat

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsAPIBodyFormatPredictV2 = shared.ListFunctionsResponseFunctionsAPIBodyFormatPredictV2

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsAPIBodyFormatCustom = shared.ListFunctionsResponseFunctionsAPIBodyFormatCustom

// Data Transfer Object(DTO) representing a container environment entry
//
// This is an alias to an internal type.
type ListFunctionsResponseFunctionsContainerEnvironment = shared.ListFunctionsResponseFunctionsContainerEnvironment

// Data Transfer Object(DTO) representing a function ne
//
// This is an alias to an internal type.
type ListFunctionsResponseFunctionsHealth = shared.ListFunctionsResponseFunctionsHealth

// HTTP/gPRC protocol type for health endpoint
//
// This is an alias to an internal type.
type ListFunctionsResponseFunctionsHealthProtocol = shared.ListFunctionsResponseFunctionsHealthProtocol

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsHealthProtocolHTTP = shared.ListFunctionsResponseFunctionsHealthProtocolHTTP

// This is an alias to an internal value.
const ListFunctionsResponseFunctionsHealthProtocolGRpc = shared.ListFunctionsResponseFunctionsHealthProtocolGRpc

// Data Transfer Object(DTO) representing an artifact
//
// This is an alias to an internal type.
type ListFunctionsResponseFunctionsModel = shared.ListFunctionsResponseFunctionsModel

// Data Transfer Object(DTO) representing an artifact
//
// This is an alias to an internal type.
type ListFunctionsResponseFunctionsResource = shared.ListFunctionsResponseFunctionsResource
