/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit
import (
	"time"
)

type ClanSearch struct {
	// The clan's avatar url
	Avatar string `json:"avatar,omitempty"`
	// The game of the clan
	Game string `json:"game,omitempty"`
	// The id of the clan
	Id string `json:"id,omitempty"`
	// The clan's join type
	Join string `json:"join,omitempty"`
	// The time the clan's last match finished
	LastMatchFinished time.Time `json:"last_match_finished,omitempty"`
	// The clan's matches count in the last 24 hours
	MatchesCount24h int64 `json:"matches_count_24h,omitempty"`
	// The clan's maximum skill level
	MaxSkillLevel int64 `json:"max_skill_level,omitempty"`
	// The clan's members count
	MembersCount int64 `json:"members_count,omitempty"`
	// The clan's members count in the last 24 hours
	MembersCount24h int64 `json:"members_count_24h,omitempty"`
	// The clan's minimum skill level
	MinSkillLevel int64 `json:"min_skill_level,omitempty"`
	// The name of the clan
	Name string `json:"name,omitempty"`
	// The clan's organizer id
	OrganizerId string `json:"organizer_id,omitempty"`
	// The region of the clan
	Region string `json:"region,omitempty"`
	// The type of the clan
	Type_ string `json:"type,omitempty"`
}
