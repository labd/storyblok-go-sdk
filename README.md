# Go SDK for Storyblok

This repository contains the go sdk for the [Storyblok Management API](https://www.storyblok.com/docs/api/management).

## Installation

This repository depends on redocly to bundle the yaml files and oapi-codegen to generate the Go SDK.

```bash
pnpm install -g @redocly/openapi-cli
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

## Development

The source.yaml is the entrypoint for the code generation.
There we define the REST routes as references to files in the path directory.
These files reference schemas from the schemas directory.

1. edit the source.yaml or any of the referenced files.
2. bundle the yaml files into a new openapi.yaml file
3. generate a new Go SDK with the oapi-codegen

```bash
redocly bundle > openapi.yaml
oapi-codegen -config oapi-config.yaml openapi.yaml
```

## Example

```go
package main

import (
    "context"
    "github.com/deepmap/oapi-codegen/pkg/securityprovider"
    "github.com/labd/storyblok-go-sdk/sbmgmt"
)

const token    = "<your personal api key>"
const spaceId  = "<your space id>"
const endpoint = "https://mapi.storyblok.com"

func ref[T any](s T) *T {
    return &s
}

func main() {
    ctx := context.Background()

    apiKeyProvider, err := securityprovider.NewSecurityProviderApiKey("header", "Authorization", token)
    if err != nil {
        panic(err)
    }

    client, err := sbmgmt.NewClientWithResponses(endpoint, sbmgmt.WithRequestEditorFn(apiKeyProvider.Intercept))
    if err != nil {
        panic(err)
    }

    resp, err := client.GetSpaceWithResponse(ctx, spaceId)
    if err != nil {
        panic(err)
    }
    fmt.Println(resp.JSON200.Space.Name)
}
```

## Progress

- [ ] Stories
- [ ] Collaborators
- [x] Components
- [x] Component Groups
- [ ] Assets
- [x] Asset Folders
- [x] Datasources
- [x] Datasource Entries
- [x] Spaces
- [ ] Space Roles
- [ ] Tasks
- [ ] Approvals
- [ ] Activities
- [ ] Presets
- [ ] Field Types
- [ ] Workflow Stage
- [ ] Workflow Stage Change
- [ ] Releases
- [ ] Branch deployments
