/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

// Secret represents a Swarm Secret value that must be passed to the CSI storage plugin when operating on this Volume. It represents one key-value pair of possibly many.
type Secret struct {
	// Key is the name of the key of the key-value pair passed to the plugin.
	Key string `json:"Key,omitempty"`
	// Secret is the swarm Secret object from which to read data. This can be a Secret name or ID. The Secret data is retrieved by Swarm and used as the value of the key-value pair passed to the plugin.
	Secret string `json:"Secret,omitempty"`
}
