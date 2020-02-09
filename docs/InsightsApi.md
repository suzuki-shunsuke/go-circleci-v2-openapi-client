# github.com/suzuki-shunsuke/go-circleci-v2-openapi-client\InsightsApi

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectWorkflowMetrics**](InsightsApi.md#GetProjectWorkflowMetrics) | **Get** /insights/{project-slug}/workflows | Get summary metrics for a project&#39;s workflows
[**GetProjectWorkflowRuns**](InsightsApi.md#GetProjectWorkflowRuns) | **Get** /insights/{project-slug}/workflows/{workflow-name} | Get recent runs of a workflow



## GetProjectWorkflowMetrics

> InlineResponse200 GetProjectWorkflowMetrics(ctx, projectSlug, optional)

Get summary metrics for a project's workflows

Get summary metrics for a project's workflows. The past 250 workflow runs, going back at most 90 days, are included in the aggregations.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
 **optional** | ***GetProjectWorkflowMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProjectWorkflowMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 
 **branch** | **optional.String**| The name of a vcs branch. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectWorkflowRuns

> InlineResponse2001 GetProjectWorkflowRuns(ctx, projectSlug, workflowName, optional)

Get recent runs of a workflow

Get recent runs of a workflow. The past 250 workflow runs, going back at most 90 days, are returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**workflowName** | **string**| The name of the workflow. | 
 **optional** | ***GetProjectWorkflowRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProjectWorkflowRunsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **optional.String**| The name of a vcs branch. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

