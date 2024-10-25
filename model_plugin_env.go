/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

// PluginEnv plugin env
type PluginEnv struct {
	// description
	Description string `json:"Description"`
	// name
	Name string `json:"Name"`
	// settable
	Settable []string `json:"Settable"`
	// value
	Value string `json:"Value"`
}
