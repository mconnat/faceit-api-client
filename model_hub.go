/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

type Hub struct {
	Avatar string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	ChatRoomId string `json:"chat_room_id,omitempty"`
	CoverImage string `json:"cover_image,omitempty"`
	Description string `json:"description,omitempty"`
	FaceitUrl string `json:"faceit_url,omitempty"`
	GameData *Game `json:"game_data,omitempty"`
	GameId string `json:"game_id,omitempty"`
	HubId string `json:"hub_id,omitempty"`
	JoinPermission string `json:"join_permission,omitempty"`
	MaxSkillLevel int64 `json:"max_skill_level,omitempty"`
	MinSkillLevel int64 `json:"min_skill_level,omitempty"`
	Name string `json:"name,omitempty"`
	OrganizerData *Organizer `json:"organizer_data,omitempty"`
	OrganizerId string `json:"organizer_id,omitempty"`
	PlayersJoined int64 `json:"players_joined,omitempty"`
	Region string `json:"region,omitempty"`
	RuleId string `json:"rule_id,omitempty"`
}
