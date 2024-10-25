/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-faceit

type LeaderboardConfig struct {
	// Max players in the leaderboard.
	MaxPlayers int64 `json:"max_players,omitempty"`
	// User will lose this amount of points if they lose a match
	PointsPerLoss int64 `json:"points_per_loss,omitempty"`
	// User will gain this amount of points if they win a match. When not configured, it's using the global value which is 3
	PointsPerWin int64 `json:"points_per_win,omitempty"`
	Promotion *Promotion `json:"promotion,omitempty"`
	Relegation *Relegation `json:"relegation,omitempty"`
	// Starting points for a player.
	StartingPoints int64 `json:"starting_points,omitempty"`
}
