/*
 * CircleCI API
 *
 * This describes the resources that make up the CircleCI API v2. API v2 is currently in Preview. Additional documentation for this API can be found in the [API Preview Docs](https://github.com/CircleCI-Public/api-preview-docs/tree/master/docs). Breaking changes to the API will be announced in the [Breaking Changes log](https://github.com/CircleCI-Public/api-preview-docs/blob/master/docs/breaking.md).
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package circleci
// PipelineVcs VCS information for the pipeline.
type PipelineVcs struct {
	// Name of the VCS provider (e.g. GitHub, Bitbucket).
	ProviderName string `json:"provider_name"`
	// URL for the repository where the trigger originated. For fork-PR pipelines, this is the URL to the fork. For other pipelines the `origin_` and `target_repository_url`s will be the same.
	OriginRepositoryUrl string `json:"origin_repository_url"`
	// URL for the repository the trigger targets (i.e. the repository where the PR will be merged). For fork-PR pipelines, this is the URL to the parent repo. For other pipelines, the `origin_` and `target_repository_url`s will be the same.
	TargetRepositoryUrl string `json:"target_repository_url"`
	// The code revision the pipeline ran.
	Revision string `json:"revision"`
	// The branch where the pipeline ran. The HEAD commit on this branch was used for the pipeline. Note that `branch` and `tag` are mutually exclusive.
	Branch string `json:"branch,omitempty"`
	// The tag used by the pipeline. The commit that this tag points to was used for the pipeline. Note that `branch` and `tag` are mutually exclusive.
	Tag string `json:"tag,omitempty"`
	Commit PipelineVcsCommit `json:"commit,omitempty"`
}
