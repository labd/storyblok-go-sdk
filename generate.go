package main

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config oapi-config.yaml ./openapi.yaml
