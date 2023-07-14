// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package openapi

const (
	ApiKeyScopes = "apiKey.Scopes"
)

// V1HelloRequestSchema defines model for V1HelloRequestSchema.
type V1HelloRequestSchema struct {
	// Name name
	Name string `json:"name"`
}

// V1HelloResponseSchema defines model for V1HelloResponseSchema.
type V1HelloResponseSchema struct {
	// Greet 挨拶
	Greet string `json:"greet"`
}

// V1HelloJSONRequestBody defines body for V1Hello for application/json ContentType.
type V1HelloJSONRequestBody = V1HelloRequestSchema
