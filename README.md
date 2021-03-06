# Go API client for CircleCI

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/suzuki-shunsuke/go-circleci-v2-openapi-client)
[![Build Status](https://cloud.drone.io/api/badges/suzuki-shunsuke/go-circleci-v2-openapi-client/status.svg)](https://cloud.drone.io/suzuki-shunsuke/go-circleci-v2-openapi-client)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-circleci-v2-openapi-client.svg)](https://github.com/suzuki-shunsuke/go-circleci-v2-openapi-client)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-circleci-v2-openapi-client/master/LICENSE)

This describes the resources that make up the CircleCI API v2. API v2 is currently in Preview. Additional documentation for this API can be found in the [API Preview Docs](https://github.com/CircleCI-Public/api-preview-docs/tree/master/docs). Breaking changes to the API will be announced in the [Breaking Changes log](https://github.com/CircleCI-Public/api-preview-docs/blob/master/docs/breaking.md).

## How to generate

https://github.com/suzuki-shunsuke/go-circleci-v2-openapi-client/pull/1

```
$ docker run -v ${PWD}:/local \
  openapitools/openapi-generator-cli:v4.2.3 \
  generate \
  --invoker-package github.com/suzuki-shunsuke/go-circleci-v2-openapi-client \
  --package-name circleci \
  -i https://circleci.com/api/v2/openapi.yml \
  -g go \
  --git-user-id suzuki-shunsuke \
  --git-repo-id go-circleci-v2-openapi-client \
  -o /local
```

Unfortunately, it is failed to compile generated code, so we fix the bug manually.

https://github.com/suzuki-shunsuke/go-circleci-v2-openapi-client/commit/f8d41c1c386c23228886a48e6560be64e0b789bb

```
$ go vet ./...
vet: ./api_preview.go:1409:6: GetProjectWorkflowMetricsOpts redeclared in this block

$ go vet ./...
vet: ./api_preview.go:1530:6: GetProjectWorkflowRunsOpts redeclared in this block

$ go vet ./...
vet: ./api_project.go:30:6: CreateCheckoutKeyOpts redeclared in this block

$ go vet ./...
vet: ./api_project.go:153:6: CreateEnvVarOpts redeclared in this block
```

Please let us know if you find the cause of this bug.

## Overview

This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Example

Please see [examples](examples).

## Documentation for API Endpoints

All URIs are relative to *https://circleci.com/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*InsightsApi* | [**GetProjectWorkflowMetrics**](docs/InsightsApi.md#getprojectworkflowmetrics) | **Get** /insights/{project-slug}/workflows | Get summary metrics for a project&#39;s workflows
*InsightsApi* | [**GetProjectWorkflowRuns**](docs/InsightsApi.md#getprojectworkflowruns) | **Get** /insights/{project-slug}/workflows/{workflow-name} | Get recent runs of a workflow
*JobApi* | [**CancelJob**](docs/JobApi.md#canceljob) | **Post** /project/{project-slug}/job/{job-number}/cancel | Cancel job
*JobApi* | [**GetJobArtifacts**](docs/JobApi.md#getjobartifacts) | **Get** /project/{project-slug}/{job-number}/artifacts | Get a job&#39;s artifacts
*JobApi* | [**GetJobDetails**](docs/JobApi.md#getjobdetails) | **Get** /project/{project-slug}/job/{job-number} | Get job details
*JobApi* | [**GetTests**](docs/JobApi.md#gettests) | **Get** /project/{project-slug}/{job-number}/tests | Get test metadata
*PipelineApi* | [**GetPipelineById**](docs/PipelineApi.md#getpipelinebyid) | **Get** /pipeline/{pipeline-id} | Get a pipeline
*PipelineApi* | [**GetPipelineByNumber**](docs/PipelineApi.md#getpipelinebynumber) | **Get** /project/{project-slug}/pipeline/{pipeline-number} | Get a pipeline
*PipelineApi* | [**GetPipelineConfigById**](docs/PipelineApi.md#getpipelineconfigbyid) | **Get** /pipeline/{pipeline-id}/config | Get a pipeline&#39;s configuration
*PipelineApi* | [**ListMyPipelines**](docs/PipelineApi.md#listmypipelines) | **Get** /project/{project-slug}/pipeline/mine | Get your pipelines
*PipelineApi* | [**ListPipelinesForProject**](docs/PipelineApi.md#listpipelinesforproject) | **Get** /project/{project-slug}/pipeline | Get all pipelines
*PipelineApi* | [**ListWorkflowsByPipelineId**](docs/PipelineApi.md#listworkflowsbypipelineid) | **Get** /pipeline/{pipeline-id}/workflow | Get a pipeline&#39;s workflows
*PipelineApi* | [**TriggerPipeline**](docs/PipelineApi.md#triggerpipeline) | **Post** /project/{project-slug}/pipeline | Trigger a new pipeline
*PreviewApi* | [**CancelJob**](docs/PreviewApi.md#canceljob) | **Post** /project/{project-slug}/job/{job-number}/cancel | Cancel job
*PreviewApi* | [**CreateCheckoutKey**](docs/PreviewApi.md#createcheckoutkey) | **Post** /project/{project-slug}/checkout-key | Create a new checkout key
*PreviewApi* | [**CreateEnvVar**](docs/PreviewApi.md#createenvvar) | **Post** /project/{project-slug}/envvar | Create an environment variable
*PreviewApi* | [**DeleteCheckoutKey**](docs/PreviewApi.md#deletecheckoutkey) | **Delete** /project/{project-slug}/checkout-key/{fingerprint} | Delete a checkout key
*PreviewApi* | [**DeleteEnvVar**](docs/PreviewApi.md#deleteenvvar) | **Delete** /project/{project-slug}/envvar/{name} | Delete an environment variable
*PreviewApi* | [**GetCheckoutKey**](docs/PreviewApi.md#getcheckoutkey) | **Get** /project/{project-slug}/checkout-key/{fingerprint} | Get a checkout key
*PreviewApi* | [**GetCollaborations**](docs/PreviewApi.md#getcollaborations) | **Get** /me/collaborations | Collaborations
*PreviewApi* | [**GetCurrentUser**](docs/PreviewApi.md#getcurrentuser) | **Get** /me | User Information
*PreviewApi* | [**GetEnvVar**](docs/PreviewApi.md#getenvvar) | **Get** /project/{project-slug}/envvar/{name} | Get a masked environment variable
*PreviewApi* | [**GetJobArtifacts**](docs/PreviewApi.md#getjobartifacts) | **Get** /project/{project-slug}/{job-number}/artifacts | Get a job&#39;s artifacts
*PreviewApi* | [**GetJobDetails**](docs/PreviewApi.md#getjobdetails) | **Get** /project/{project-slug}/job/{job-number} | Get job details
*PreviewApi* | [**GetProjectBySlug**](docs/PreviewApi.md#getprojectbyslug) | **Get** /project/{project-slug} | Get a project
*PreviewApi* | [**GetProjectWorkflowMetrics**](docs/PreviewApi.md#getprojectworkflowmetrics) | **Get** /insights/{project-slug}/workflows | Get summary metrics for a project&#39;s workflows
*PreviewApi* | [**GetProjectWorkflowRuns**](docs/PreviewApi.md#getprojectworkflowruns) | **Get** /insights/{project-slug}/workflows/{workflow-name} | Get recent runs of a workflow
*PreviewApi* | [**GetTests**](docs/PreviewApi.md#gettests) | **Get** /project/{project-slug}/{job-number}/tests | Get test metadata
*PreviewApi* | [**GetUser**](docs/PreviewApi.md#getuser) | **Get** /user/{id} | User Information
*PreviewApi* | [**ListCheckoutKeys**](docs/PreviewApi.md#listcheckoutkeys) | **Get** /project/{project-slug}/checkout-key | Get all checkout keys
*PreviewApi* | [**ListEnvVars**](docs/PreviewApi.md#listenvvars) | **Get** /project/{project-slug}/envvar | List all environment variables
*ProjectApi* | [**CreateCheckoutKey**](docs/ProjectApi.md#createcheckoutkey) | **Post** /project/{project-slug}/checkout-key | Create a new checkout key
*ProjectApi* | [**CreateEnvVar**](docs/ProjectApi.md#createenvvar) | **Post** /project/{project-slug}/envvar | Create an environment variable
*ProjectApi* | [**DeleteCheckoutKey**](docs/ProjectApi.md#deletecheckoutkey) | **Delete** /project/{project-slug}/checkout-key/{fingerprint} | Delete a checkout key
*ProjectApi* | [**DeleteEnvVar**](docs/ProjectApi.md#deleteenvvar) | **Delete** /project/{project-slug}/envvar/{name} | Delete an environment variable
*ProjectApi* | [**GetCheckoutKey**](docs/ProjectApi.md#getcheckoutkey) | **Get** /project/{project-slug}/checkout-key/{fingerprint} | Get a checkout key
*ProjectApi* | [**GetEnvVar**](docs/ProjectApi.md#getenvvar) | **Get** /project/{project-slug}/envvar/{name} | Get a masked environment variable
*ProjectApi* | [**GetProjectBySlug**](docs/ProjectApi.md#getprojectbyslug) | **Get** /project/{project-slug} | Get a project
*ProjectApi* | [**ListCheckoutKeys**](docs/ProjectApi.md#listcheckoutkeys) | **Get** /project/{project-slug}/checkout-key | Get all checkout keys
*ProjectApi* | [**ListEnvVars**](docs/ProjectApi.md#listenvvars) | **Get** /project/{project-slug}/envvar | List all environment variables
*UserApi* | [**GetCollaborations**](docs/UserApi.md#getcollaborations) | **Get** /me/collaborations | Collaborations
*UserApi* | [**GetCurrentUser**](docs/UserApi.md#getcurrentuser) | **Get** /me | User Information
*UserApi* | [**GetUser**](docs/UserApi.md#getuser) | **Get** /user/{id} | User Information
*WorkflowApi* | [**CancelWorkflow**](docs/WorkflowApi.md#cancelworkflow) | **Post** /workflow/{id}/cancel | Cancel a workflow
*WorkflowApi* | [**GetWorkflowById**](docs/WorkflowApi.md#getworkflowbyid) | **Get** /workflow/{id} | Get a workflow
*WorkflowApi* | [**ListWorkflowJobs**](docs/WorkflowApi.md#listworkflowjobs) | **Get** /workflow/{id}/job | Get a workflow&#39;s jobs
*WorkflowApi* | [**RerunWorkflow**](docs/WorkflowApi.md#rerunworkflow) | **Post** /workflow/{id}/rerun | Rerun a workflow


## Documentation For Models

 - [Artifact](docs/Artifact.md)
 - [ArtifactListResponse](docs/ArtifactListResponse.md)
 - [CheckoutKey](docs/CheckoutKey.md)
 - [CheckoutKeyInput](docs/CheckoutKeyInput.md)
 - [CheckoutKeyListResponse](docs/CheckoutKeyListResponse.md)
 - [Collaboration](docs/Collaboration.md)
 - [EnvironmentVariableListResponse](docs/EnvironmentVariableListResponse.md)
 - [EnvironmentVariablePair](docs/EnvironmentVariablePair.md)
 - [EnvironmentVariablePair1](docs/EnvironmentVariablePair1.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2001Items](docs/InlineResponse2001Items.md)
 - [InlineResponse200Items](docs/InlineResponse200Items.md)
 - [InlineResponse200Metrics](docs/InlineResponse200Metrics.md)
 - [InlineResponse200MetricsDurationMetrics](docs/InlineResponse200MetricsDurationMetrics.md)
 - [Job](docs/Job.md)
 - [JobDetails](docs/JobDetails.md)
 - [JobDetailsContexts](docs/JobDetailsContexts.md)
 - [JobDetailsExecutor](docs/JobDetailsExecutor.md)
 - [JobDetailsLatestWorkflow](docs/JobDetailsLatestWorkflow.md)
 - [JobDetailsMessages](docs/JobDetailsMessages.md)
 - [JobDetailsOrganization](docs/JobDetailsOrganization.md)
 - [JobDetailsParallelRuns](docs/JobDetailsParallelRuns.md)
 - [JobDetailsPipeline](docs/JobDetailsPipeline.md)
 - [JobDetailsProject](docs/JobDetailsProject.md)
 - [MessageResponse](docs/MessageResponse.md)
 - [Pipeline](docs/Pipeline.md)
 - [PipelineConfig](docs/PipelineConfig.md)
 - [PipelineErrors](docs/PipelineErrors.md)
 - [PipelineLight](docs/PipelineLight.md)
 - [PipelineListResponse](docs/PipelineListResponse.md)
 - [PipelineTrigger](docs/PipelineTrigger.md)
 - [PipelineTriggerActor](docs/PipelineTriggerActor.md)
 - [PipelineVcs](docs/PipelineVcs.md)
 - [PipelineVcsCommit](docs/PipelineVcsCommit.md)
 - [Project](docs/Project.md)
 - [ProjectVcsInfo](docs/ProjectVcsInfo.md)
 - [RerunWorkflowParameters](docs/RerunWorkflowParameters.md)
 - [TestsResponse](docs/TestsResponse.md)
 - [TestsResponseItems](docs/TestsResponseItems.md)
 - [TriggerPipelineParameters](docs/TriggerPipelineParameters.md)
 - [User](docs/User.md)
 - [Workflow](docs/Workflow.md)
 - [WorkflowJobListResponse](docs/WorkflowJobListResponse.md)
 - [WorkflowListResponse](docs/WorkflowListResponse.md)

## License

[MIT](LICENSE)

About the license of code generated by OpenAPI Generator, please see [the document of OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator/tree/2aa8a6d033699121d2692a3b03993748b3eeb3ed#34---license-information-on-generated-code).
