package tfinfra

type Database struct {
	Amount             int          `json:"amount" validate:"required"`
	Name               *string      `json:"name" validate:"required"`
	Engine             *string      `json:"engine" validate:"required"`
	Size               *string      `json:"size" validate:"required"`
	Region             *string      `json:"region" validate:"required"`
	NodeCount          *int         `json:"node_count" validate:"required"`
	Version            *string      `json:"version" validate:"required"`
	Tags               *[]string    `json:"tags" validate:"omitempty"`
	PrivateNetworkUuid *string      `json:"private_network_uuid" validate:"omitempty"`
	EvictionPolicy     *string      `json:"eviction_policy" validate:"omitempty"`
	SqlMode            *string      `json:"sql_mode" validate:"omitempty"`
	MaintenanceWindow  *maintenance `json:"maintenance_window" validate:"omitempty,dive,required"`
	Timeouts           *timeOuts    `json:"timeouts" validate:"omitempty,dive,required"`
}

type maintenance struct {
	Day  *string `json:"day" validate:"required"`
	Hour *string `json:"hour" validate:"required"`
}

type timeOuts struct {
	Create *string `json:"create" validate:"omitempty"`
	Delete *string `json:"delete" validate:"omitempty"`
	Update *string `json:"update" validate:"omitempty"`
}
