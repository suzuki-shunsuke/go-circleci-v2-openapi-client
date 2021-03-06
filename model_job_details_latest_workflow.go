/*
 * CircleCI API
 *
 * This describes the resources that make up the CircleCI API v2. API v2 is currently in Preview. Additional documentation for this API can be found in the [API Preview Docs](https://github.com/CircleCI-Public/api-preview-docs/tree/master/docs). Breaking changes to the API will be announced in the [Breaking Changes log](https://github.com/CircleCI-Public/api-preview-docs/blob/master/docs/breaking.md).
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package circleci
// JobDetailsLatestWorkflow Info about the latest workflow the job was a part of.
type JobDetailsLatestWorkflow struct {
	// The unique ID of the workflow.
	Id string `json:"id"`
	// The name of the workflow.
	Name string `json:"name"`
}
