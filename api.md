# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#CreateFunctionResponse">CreateFunctionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#FunctionResponse">FunctionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#GetQueuesResponse">GetQueuesResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#ListFunctionsResponse">ListFunctionsResponse</a>

# UserSecretManagement

## Functions

### Versions

Methods:

- <code title="put /v2/nvcf/secrets/functions/{functionId}/versions/{versionId}">client.UserSecretManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#UserSecretManagementFunctionVersionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#UserSecretManagementFunctionVersionUpdateParams">UserSecretManagementFunctionVersionUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# FunctionManagement

## Functions

### Versions

Methods:

- <code title="get /v2/nvcf/functions/{functionId}/versions/{functionVersionId}">client.FunctionManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionManagementFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v2/nvcf/metadata/functions/{functionId}/versions/{functionVersionId}">client.FunctionManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionManagementFunctionVersionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionManagementFunctionVersionUpdateParams">FunctionManagementFunctionVersionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/functions/{functionId}/versions/{functionVersionId}">client.FunctionManagement.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionManagementFunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### IDs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ListFunctionIDsResponse">ListFunctionIDsResponse</a>

Methods:

- <code title="get /v2/nvcf/functions/ids">client.FunctionManagement.Functions.IDs.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionManagementFunctionIDService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionManagementFunctionIDListParams">FunctionManagementFunctionIDListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ListFunctionIDsResponse">ListFunctionIDsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FunctionDeployment

## Functions

### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#DeploymentResponse">DeploymentResponse</a>

Methods:

- <code title="post /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionDeploymentFunctionVersionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionDeploymentFunctionVersionNewParams">FunctionDeploymentFunctionVersionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionDeploymentFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionDeploymentFunctionVersionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionDeploymentFunctionVersionUpdateParams">FunctionDeploymentFunctionVersionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#DeploymentResponse">DeploymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/deployments/functions/{functionId}/versions/{functionVersionId}">client.FunctionDeployment.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionDeploymentFunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionDeploymentFunctionVersionDeleteParams">FunctionDeploymentFunctionVersionDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#FunctionResponse">FunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FunctionInvocation

## Functions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionInvocationFunctionInvokeResponse">FunctionInvocationFunctionInvokeResponse</a>

Methods:

- <code title="post /v2/nvcf/pexec/functions/{functionId}">client.FunctionInvocation.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionInvocationFunctionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionInvocationFunctionInvokeParams">FunctionInvocationFunctionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionInvocationFunctionInvokeResponse">FunctionInvocationFunctionInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionInvocationFunctionVersionInvokeResponse">FunctionInvocationFunctionVersionInvokeResponse</a>

Methods:

- <code title="post /v2/nvcf/pexec/functions/{functionId}/versions/{versionId}">client.FunctionInvocation.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionInvocationFunctionVersionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionInvocationFunctionVersionInvokeParams">FunctionInvocationFunctionVersionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#FunctionInvocationFunctionVersionInvokeResponse">FunctionInvocationFunctionVersionInvokeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# EnvelopeFunctionInvocation

## Functions

Methods:

- <code title="post /v2/nvcf/exec/functions/{functionId}">client.EnvelopeFunctionInvocation.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#EnvelopeFunctionInvocationFunctionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#EnvelopeFunctionInvocationFunctionInvokeParams">EnvelopeFunctionInvocationFunctionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="post /v2/nvcf/exec/functions/{functionId}/versions/{versionId}">client.EnvelopeFunctionInvocation.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#EnvelopeFunctionInvocationFunctionVersionService.Invoke">Invoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#EnvelopeFunctionInvocationFunctionVersionInvokeParams">EnvelopeFunctionInvocationFunctionVersionInvokeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Exec

### Status

Methods:

- <code title="get /v2/nvcf/exec/status/{requestId}">client.EnvelopeFunctionInvocation.Exec.Status.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#EnvelopeFunctionInvocationExecStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#InvokeFunctionResponse">InvokeFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Nvcf

## Functions

Methods:

- <code title="post /v2/nvcf/functions">client.Nvcf.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfFunctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfFunctionNewParams">NvcfFunctionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#CreateFunctionResponse">CreateFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/functions">client.Nvcf.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfFunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfFunctionListParams">NvcfFunctionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#ListFunctionsResponse">ListFunctionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="post /v2/nvcf/functions/{functionId}/versions">client.Nvcf.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfFunctionVersionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfFunctionVersionNewParams">NvcfFunctionVersionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#CreateFunctionResponse">CreateFunctionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/functions/{functionId}/versions">client.Nvcf.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfFunctionVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#ListFunctionsResponse">ListFunctionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Authorizations

### Functions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ListAuthorizedPartiesResponse">ListAuthorizedPartiesResponse</a>

Methods:

- <code title="get /v2/nvcf/authorizations/functions/{functionId}">client.Nvcf.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfAuthorizationFunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ListAuthorizedPartiesResponse">ListAuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/authorizations/functions/{functionId}">client.Nvcf.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfAuthorizationFunctionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/nvcf/authorizations/functions/{functionId}">client.Nvcf.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfAuthorizationFunctionService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfAuthorizationFunctionAuthorizeParams">NvcfAuthorizationFunctionAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Versions

Methods:

- <code title="get /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.Nvcf.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfAuthorizationFunctionVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.Nvcf.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfAuthorizationFunctionVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}">client.Nvcf.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfAuthorizationFunctionVersionService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#NvcfAuthorizationFunctionVersionAuthorizeParams">NvcfAuthorizationFunctionVersionAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#CreateAssetResponse">CreateAssetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ListAssetsResponse">ListAssetsResponse</a>

Methods:

- <code title="post /v2/nvcf/assets">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AssetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AssetNewParams">AssetNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#CreateAssetResponse">CreateAssetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/nvcf/assets">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AssetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ListAssetsResponse">ListAssetsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Authorizations

## Functions

Methods:

- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/add">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AuthorizationFunctionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AuthorizationFunctionAddParams">AuthorizationFunctionAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/remove">client.Authorizations.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AuthorizationFunctionService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AuthorizationFunctionRemoveParams">AuthorizationFunctionRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}/add">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AuthorizationFunctionVersionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AuthorizationFunctionVersionAddParams">AuthorizationFunctionVersionAddParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/nvcf/authorizations/functions/{functionId}/versions/{functionVersionId}/remove">client.Authorizations.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AuthorizationFunctionVersionService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, functionVersionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AuthorizationFunctionVersionRemoveParams">AuthorizationFunctionVersionRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#AuthorizedPartiesResponse">AuthorizedPartiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Queues

## Functions

Methods:

- <code title="get /v2/nvcf/queues/functions/{functionId}">client.Queues.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#QueueFunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#GetQueuesResponse">GetQueuesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Versions

Methods:

- <code title="get /v2/nvcf/queues/functions/{functionId}/versions/{versionId}">client.Queues.Functions.Versions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#QueueFunctionVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionID <a href="https://pkg.go.dev/builtin#string">string</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go/shared#GetQueuesResponse">GetQueuesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Position

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#GetPositionInQueueResponse">GetPositionInQueueResponse</a>

Methods:

- <code title="get /v2/nvcf/queues/{requestId}/position">client.Queues.Position.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#QueuePositionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#GetPositionInQueueResponse">GetPositionInQueueResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Pexec

## Status

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#PexecStatusGetResponse">PexecStatusGetResponse</a>

Methods:

- <code title="get /v2/nvcf/pexec/status/{requestId}">client.Pexec.Status.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#PexecStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#PexecStatusGetParams">PexecStatusGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#PexecStatusGetResponse">PexecStatusGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ClusterGroupsAndGPUs

## ClusterGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ClusterGroupsResponse">ClusterGroupsResponse</a>

Methods:

- <code title="get /v2/nvcf/clusterGroups">client.ClusterGroupsAndGPUs.ClusterGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ClusterGroupsAndGPUClusterGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ClusterGroupsResponse">ClusterGroupsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ClientManagementForNvidiaSuperAdmins

## Clients

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ClientManagementForNvidiaSuperAdminClientGetResponse">ClientManagementForNvidiaSuperAdminClientGetResponse</a>

Methods:

- <code title="get /v2/nvcf/clients/{clientId}">client.ClientManagementForNvidiaSuperAdmins.Clients.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ClientManagementForNvidiaSuperAdminClientService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#ClientManagementForNvidiaSuperAdminClientGetResponse">ClientManagementForNvidiaSuperAdminClientGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AssetManagement

## Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AssetResponse">AssetResponse</a>

Methods:

- <code title="get /v2/nvcf/assets/{assetId}">client.AssetManagement.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AssetManagementAssetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go">nvidiacloudfunctions</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AssetResponse">AssetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/nvcf/assets/{assetId}">client.AssetManagement.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/nvidia-cloud-functions-go#AssetManagementAssetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
