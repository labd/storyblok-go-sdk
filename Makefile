generate:
	npx @redocly/cli bundle --output openapi.yaml
	oapi-codegen -config oapi-config.yaml openapi.yaml
