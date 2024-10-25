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

// ClusterVolume contains options and information specific to, and only present on, Swarm CSI cluster volumes.
type ClusterVolume struct {
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// ID is the Swarm ID of the volume. Because cluster volumes are Swarm objects, they have an ID, unlike non-cluster volumes, which only have a Name. This ID can be used to refer to the cluster volume.
	ID string `json:"ID,omitempty"`
	Info *Info `json:"Info,omitempty"`
	// PublishStatus contains the status of the volume as it pertains to its publishing on Nodes.
	PublishStatus []PublishStatus `json:"PublishStatus,omitempty"`
	Spec *ClusterVolumeSpec `json:"Spec,omitempty"`
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	Version *Version `json:"Version,omitempty"`
}
