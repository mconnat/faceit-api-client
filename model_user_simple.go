/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

// The UserSimple holds information about a user.
type UserSimple struct {
	// The Avatar of a user
	Avatar string `json:"avatar,omitempty"`
	// The Country of a user
	Country string `json:"country,omitempty"`
	// The FaceitUrl of a user
	FaceitUrl string `json:"faceit_url,omitempty"`
	// Deprecated: use memberships instead
	MembershipType string `json:"membership_type,omitempty"`
	// The Memberships of a user
	Memberships []string `json:"memberships,omitempty"`
	// The Nickname of a user
	Nickname string `json:"nickname,omitempty"`
	// The SkillLevel of a user
	SkillLevel int64 `json:"skill_level,omitempty"`
	// The ID of a user
	UserId string `json:"user_id,omitempty"`
}
