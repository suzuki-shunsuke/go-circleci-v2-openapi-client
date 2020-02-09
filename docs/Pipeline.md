# Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the pipeline. | 
**Errors** | [**[]PipelineErrors**](Pipeline_errors.md) | A sequence of errors that have occurred within the pipeline. | 
**ProjectSlug** | **string** | The project-slug for the pipeline. | 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the pipeline was last updated. | [optional] 
**Number** | **int64** | The number of the pipeline. | 
**State** | **string** | The current state of the pipeline. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the pipeline was created. | 
**Trigger** | [**PipelineTrigger**](Pipeline_trigger.md) |  | 
**Vcs** | [**PipelineVcs**](Pipeline_vcs.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


