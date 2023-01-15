package tfcompute

type Droplets struct {
	Amount           int            `json:"amount"  validate:"required"`
	Image            *string        `json:"image" description:"(Required) The Droplet image ID or slug. This could be either image ID or droplet snapshot ID."  validate:"required"`
	Name             *string        `json:"name" description:"(Required) The Droplet name."  validate:"required"`
	Region           *string        `json:"region" description:"(Required) The region where the Droplet will be created."  validate:"required"`
	Size             *string        `json:"size" description:"(Required) The unique slug that indentifies the type of Droplet. You can find a list of available slugs on DigitalOcean API documentation."  validate:"required"`
	Backups          *bool          `json:"backups" description:"(Optional) Boolean controlling if backups are made. Defaults to false."  validate:"omitempty,boolean"`
	Monitoring       *bool          `json:"monitoring" description:"(Optional) Boolean controlling whether monitoring agent is installed. Defaults to false. If set to true, you can configure monitor alert policies monitor alert resource"  validate:"omitempty,boolean"`
	Ipv6             *bool          `json:"ipv6" description:"(Optional) Boolean controlling if IPv6 is enabled. Defaults to false."  validate:"omitempty,boolean"`
	VpcUuid          *string        `json:"vpc_uuid" description:"(Optional) The ID of the VPC where the Droplet will be located."  validate:"omitempty"`
	SshKeys          *[]string      `json:"ssh_keys" description:"(Optional) A list of SSH key IDs or fingerprints to enable in the format [12345, 123456]. To retrieve this info, use the DigitalOcean API or CLI (doctl compute ssh-key list). Once a Droplet is created keys can not be added or removed via this provider. Modifying this field will prompt you to destroy and recreate the Droplet."  validate:"omitempty"`
	ResizeDisk       *bool          `json:"resize_disk" description:"(Optional) Boolean controlling whether to increase the disk size when resizing a Droplet. It defaults to true. When set to false, only the Droplet's RAM and CPU will be resized. Increasing a Droplet's disk size is a permanent change. Increasing only RAM and CPU is reversible."  validate:"omitempty,boolean"`
	Tags             *[]string      `json:"tags" description:"(Optional) A list of the tags to be applied to this Droplet."  validate:"omitempty"`
	UserData         *string        `json:"user_data" description:"(Optional) A string of the desired User Data for the Droplet."  validate:"omitempty"`
	VolumeIds        *[]string      `json:"volume_ids" description:"(Optional) A list of the IDs of each block storage volume to be attached to the Droplet."  validate:"omitempty"`
	DropletAgent     *bool          `json:"droplet_agent" description:"(Optional) A boolean indicating whether to install the DigitalOcean agent used for providing access to the Droplet web console in the control panel. By default, the agent is installed on new Droplets but installation errors (i.e. OS not supported) are ignored. To prevent it from being installed, set to false. To make installation errors fatal, explicitly set it to true."  validate:"omitempty,boolean"`
	GracefulShutdown *bool          `json:"graceful_shutdown" description:"(Optional) A boolean indicating whether the droplet should be gracefully shut down before it is deleted."  validate:"omitempty,boolean"`
	Outputs          DropletOutputs `json:"outputs" description:"The attributes exported after creation of the resource."`
}

type DropletOutputs struct {
	Id                 string   `json:"output_id" description:"The ID of the Droplet"`
	Urn                string   `json:"output_urn" description:"The uniform resource name of the Droplet"`
	Name               string   `json:"output_name" description:"he name of the Droplet"`
	Region             string   `json:"output_region" description:"The region of the Droplet"`
	Image              string   `json:"output_image" description:"The image of the Droplet"`
	Ipv6               bool     `json:"output_ipv6" description:"Is IPv6 enabled"`
	Ipv6Address        string   `json:"output_ipv6_address" description:"The IPv6 address"`
	Ipv4Address        string   `json:"output_ipv4_address" description:"The IPv4 address"`
	Ipv4AddressPrivate string   `json:"output_ipv4_address_private" description:"The private networking IPv4 address"`
	Locked             bool     `json:"output_locked" description:"Is the Droplet locked"`
	PrivateNetworking  bool     `json:"output_private_networking" description:"Is private networking enabled"`
	PriceHourly        float64  `json:"output_price_hourly" description:"Droplet hourly price"`
	PriceMonthly       float64  `json:"output_price_monthly" description:"Droplet monthly price"`
	Size               string   `json:"output_size" description:"The instance size"`
	Disk               int      `json:"output_disk" description:"The size of the instance's disk in GB"`
	Vcpus              int      `json:"output_vcpus" description:"The number of the instance's virtual CPUs"`
	Status             string   `json:"output_status" description:"The status of the Droplet"`
	Tags               []string `json:"output_tags" description:"The tags associated with the Droplet"`
	VolumeId           []string `json:"output_volume_id" description:"A list of the attached block storage volumes"`
}
