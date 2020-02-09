# JobDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebUrl** | **string** | URL of the job in CircleCI Web UI. | 
**Project** | [**JobDetailsProject**](Job_Details_project.md) |  | 
**ParallelRuns** | [**[]JobDetailsParallelRuns**](Job_Details_parallel_runs.md) | Info about parallels runs and their status. | 
**StartedAt** | [**time.Time**](time.Time.md) | The date and time the job started. | 
**LatestWorkflow** | [**JobDetailsLatestWorkflow**](Job_Details_latest_workflow.md) |  | 
**Name** | **string** | The name of the job. | 
**Executor** | [**JobDetailsExecutor**](Job_Details_executor.md) |  | 
**Parallelism** | **int64** | A number of parallel runs the job has. | 
**Status** | [**map[string]interface{}**](map[string]interface{}.md) | The current status of the job. | 
**Number** | **int64** | The number of the job. | 
**Pipeline** | [**JobDetailsPipeline**](Job_Details_pipeline.md) |  | 
**Duration** | **int64** | Duration of a job in milliseconds. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The time when the job was created. | 
**Messages** | [**[]JobDetailsMessages**](Job_Details_messages.md) | Messages from CircleCI execution platform. | 
**Contexts** | [**[]JobDetailsContexts**](Job_Details_contexts.md) | List of contexts used by the job. | 
**Organization** | [**JobDetailsOrganization**](Job_Details_organization.md) |  | 
**QueuedAt** | [**time.Time**](time.Time.md) | The time when the job was placed in a queue. | 
**StoppedAt** | [**time.Time**](time.Time.md) | The time when the job stopped. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


