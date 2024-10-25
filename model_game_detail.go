/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

type GameDetail struct {
	FaceitElo int64 `json:"faceit_elo,omitempty"`
	GamePlayerId string `json:"game_player_id,omitempty"`
	GamePlayerName string `json:"game_player_name,omitempty"`
	// Deprecated: no more in use
	GameProfileId string `json:"game_profile_id,omitempty"`
	Region string `json:"region,omitempty"`
	// Deprecated: no more in use
	Regions *interface{} `json:"regions,omitempty"`
	SkillLevel int64 `json:"skill_level,omitempty"`
	// Deprecated: use SkillLevel instead
	SkillLevelLabel string `json:"skill_level_label,omitempty"`
}
