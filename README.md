# Go SDK for Storyblok


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
