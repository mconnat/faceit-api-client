/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

// ContainerTopOKBody OK response to ContainerTop operation
type ContainerTopOkBody struct {
	// Each process running in the container, where each is process is an array of values corresponding to the titles.
	Processes [][]string `json:"Processes"`
	// The ps column titles
	Titles []string `json:"Titles"`
}
