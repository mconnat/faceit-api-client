/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

// PluginMount plugin mount
type PluginMount struct {
	// description
	Description string `json:"Description"`
	// destination
	Destination string `json:"Destination"`
	// name
	Name string `json:"Name"`
	// options
	Options []string `json:"Options"`
	// settable
	Settable []string `json:"Settable"`
	// source
	Source string `json:"Source"`
	// type
	Type_ string `json:"Type"`
}
