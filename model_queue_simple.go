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

type QueueSimple struct {
	EntityId string `json:"entityId,omitempty"`
	EntityType string `json:"entityType,omitempty"`
	Game string `json:"game,omitempty"`
	Id string `json:"id,omitempty"`
	LastModified time.Time `json:"lastModified,omitempty"`
	Open bool `json:"open,omitempty"`
	OrganizerId string `json:"organizerId,omitempty"`
	QueueName string `json:"queueName,omitempty"`
	Region string `json:"region,omitempty"`
	State string `json:"state,omitempty"`
}
