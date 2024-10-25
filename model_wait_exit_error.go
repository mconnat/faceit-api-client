/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

// WaitExitError container waiting error, if any
type WaitExitError struct {
	// Details of an error
	Message string `json:"Message,omitempty"`
}
