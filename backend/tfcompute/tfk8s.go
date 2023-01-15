package tfcompute

type K8s struct {
	Amount            int          `json:"amount" validate:"required"`
	Name              *string      `json:"name" validate:"required"`
	Region            *string      `json:"region" validate:"required"`
	Version           *string      `json:"version" validate:"required,oneof=1.25.4-do.0 1.24.8-do.0 1.23.14-do.0"`
	NodePool          *nodePool    `json:"node_pool" validate:"required"`
	Vpcuuid           *string      `json:"vpc_uuid" validate:"omitempty"`
	AutoUpgrade       *bool        `json:"auto_upgrade" validate:"omitempty"`
	SurgeUpgrade      *bool        `json:"surge_upgrade" validate:"omitempty"`
	Ha                *bool        `json:"ha" validate:"omitempty"`
	Tags              *[]string    `json:"tags" validate:"omitempty"`
	MaintenancePolicy *maintenance `json:"maintenance_policy" validate:"omitempty,dive,required"`
}

type nodePool struct {
	Name      *string            `json:"name" validate:"required"`
	Size      *string            `json:"size" validate:"required"`
	NodeCount *int               `json:"node_count" validate:"omitempty"`
	AutoScale *bool              `json:"auto_scale" validate:"omitempty"`
	MinNodes  *int               `json:"min_nodes" validate:"omitempty"`
	MaxNodes  *int               `json:"max_nodes" validate:"omitempty"`
	Tags      *[]string          `json:"tags" validate:"omitempty"`
	Labels    *map[string]string `json:"labels" validate:"omitempty"`
}

type maintenance struct {
	Day       *string `json:"day" validate:"required"`
	StartTime *string `json:"start_time" validate:"required"`
}
