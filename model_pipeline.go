/*
 * CircleCI API
 *
 * This describes the resources that make up the CircleCI API v2. API v2 is currently in Preview. Additional documentation for this API can be found in the [API Preview Docs](https://github.com/CircleCI-Public/api-preview-docs/tree/master/docs). Breaking changes to the API will be announced in the [Breaking Changes log](https://github.com/CircleCI-Public/api-preview-docs/blob/master/docs/breaking.md).
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package circleci
import (
	"time"
)
// Pipeline NOTE: The definition of pipeline is subject to change.
type Pipeline struct {
	// The unique ID of the pipeline.
	Id string `json:"id"`
	// A sequence of errors that have occurred within the pipeline.
	Errors []PipelineErrors `json:"errors"`
	// The project-slug for the pipeline.
	ProjectSlug string `json:"project_slug"`
	// The date and time the pipeline was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The number of the pipeline.
	Number int64 `json:"number"`
	// The current state of the pipeline.
	State string `json:"state"`
	// The date and time the pipeline was created.
	CreatedAt time.Time `json:"created_at"`
	Trigger PipelineTrigger `json:"trigger"`
	Vcs PipelineVcs `json:"vcs,omitempty"`
}
