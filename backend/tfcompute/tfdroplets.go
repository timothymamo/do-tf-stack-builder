package tfcompute

type Droplets struct {
	Amount           int       `json:"amount" validate:"required"`
	Name             *string   `json:"name" description:"Info" validate:"required"`
	Image            *string   `json:"image" description:"Info2" validate:"required"`
	Region           *string   `json:"region" validate:"required"`
	Size             *string   `json:"size" validate:"required"`
	Backups          *bool     `json:"backups" validate:"omitempty,boolean"`
	Monitoring       *bool     `json:"monitoring" validate:"omitempty,boolean"`
	Ipv6             *bool     `json:"ipv6" validate:"omitempty,boolean"`
	VpcUuid          *string   `json:"vpc_uuid" validate:"omitempty"`
	SshKeys          *[]string `json:"ssh_keys" validate:"omitempty"`
	ResizeDisk       *bool     `json:"resize_disk" validate:"omitempty,boolean"`
	Tags             *[]string `json:"tags" validate:"omitempty"`
	UserData         *string   `json:"user_data" validate:"omitempty"`
	VolumeIds        *[]string `json:"volume_ids" validate:"omitempty"`
	DropletAgent     *bool     `json:"droplet_agent" validate:"omitempty,boolean"`
	GracefulShutdown *bool     `json:"graceful_shutdown" validate:"omitempty,boolean"`
}
