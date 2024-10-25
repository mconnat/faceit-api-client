/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

type Organizer struct {
	Avatar string `json:"avatar,omitempty"`
	Cover string `json:"cover,omitempty"`
	Description string `json:"description,omitempty"`
	Facebook string `json:"facebook,omitempty"`
	FaceitUrl string `json:"faceit_url,omitempty"`
	FollowersCount int64 `json:"followers_count,omitempty"`
	Name string `json:"name,omitempty"`
	OrganizerId string `json:"organizer_id,omitempty"`
	Twitch string `json:"twitch,omitempty"`
	Twitter string `json:"twitter,omitempty"`
	Type_ string `json:"type,omitempty"`
	Vk string `json:"vk,omitempty"`
	Website string `json:"website,omitempty"`
	Youtube string `json:"youtube,omitempty"`
}
