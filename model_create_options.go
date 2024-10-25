/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

// Volume configuration
type CreateOptions struct {
	ClusterVolumeSpec *ClusterVolumeSpec `json:"ClusterVolumeSpec,omitempty"`
	// Name of the volume driver to use.
	Driver string `json:"Driver,omitempty"`
	// A mapping of driver options and values. These options are passed directly to the driver and are driver specific.
	DriverOpts map[string]string `json:"DriverOpts,omitempty"`
	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`
	// The new volume's name. If not specified, Docker generates a name.
	Name string `json:"Name,omitempty"`
}
