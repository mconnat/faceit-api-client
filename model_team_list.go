/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

// The TeamList holds teams information.
type TeamList struct {
	End int64 `json:"end,omitempty"`
	// The teams list.
	Items []Team `json:"items"`
	Start int64 `json:"start,omitempty"`
}
