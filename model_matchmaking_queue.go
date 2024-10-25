/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

type MatchmakingQueue struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Open bool `json:"open,omitempty"`
	OrganizerId string `json:"organizer_id,omitempty"`
	Paused bool `json:"paused,omitempty"`
}
