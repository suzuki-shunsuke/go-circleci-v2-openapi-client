package examples

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-circleci-v2-openapi-client"
)

func ExamplePreviewApiService_GetProjectBySlug() {
	token := "xxxxx" // CircleCI API token
	ctx := context.WithValue(context.Background(), circleci.ContextAPIKey, circleci.APIKey{
		Key: token,
	})

	cfg := circleci.NewConfiguration()

	// mocking
	cfg.HTTPClient = newHTTPClientPreviewApiService_GetProjectBySlug(token)

	client := circleci.NewAPIClient(cfg)
	prj, _, err := client.PreviewApi.GetProjectBySlug(ctx, "gh/suzuki-shunsuke/example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("slug: %s, name: %s, org: %s\n", prj.Slug, prj.Name, prj.OrganizationName)
	// Output: slug: gh/suzuki-shunsuke/example, name: example, org: suzuki-shunsuke
}

func newHTTPClientPreviewApiService_GetProjectBySlug(token string) *http.Client {
	return &http.Client{
		Transport: &flute.Transport{
			Services: []flute.Service{
				{
					Endpoint: "https://circleci.com",
					Routes: []flute.Route{
						{
							Name: "get a project by slug",
							Matcher: &flute.Matcher{
								Method: "GET",
								Path:   "/api/v2/project/gh/suzuki-shunsuke/example",
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
									Header: http.Header{
										"Content-Type": []string{"application/json"},
									},
								},
								BodyString: `{
  "slug" : "gh/suzuki-shunsuke/example",
  "organization_name" : "suzuki-shunsuke",
  "name" : "example",
  "vcs_info" : {
    "vcs_url" : "https://github.com/suzuki-shunsuke/example",
    "default_branch" : "master",
    "provider" : "GitHub"
  }
}`,
							},
						},
					},
				},
			},
		},
	}
}
