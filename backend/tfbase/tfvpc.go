package tfbase

type VPC struct {
	Amount      int        `json:"amount" validate:"required"`
	Name        *string    `json:"name" description:"(Required) A name for the VPC. Must be unique and contain alphanumeric characters, dashes, and periods only." variable:"vpc_name" validate:"required"`
	Region      *string    `json:"region" description:"(Required) The DigitalOcean region slug for the VPC's location." validate:"required"`
	Description *string    `json:"description" description:"(Optional) A free-form text field up to a limit of 255 characters to describe the VPC."`
	Iprange     *string    `json:"ip_range" description:"(Optional) The range of IP addresses for the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be larger than /16 or smaller than /24." validate:"omitempty,cidrv4"`
	Outputs     VPCOutputs `json:"outputs" description:"The attributes exported after creation of the resource."`
}

type VPCOutputs struct {
	Iprange    string `json:"output_ip_range" description:"The range of IP addresses for the VPC in CIDR notation."`
	Id         string `json:"output_id" description:"The unique identifier for the VPC."`
	Urn        string `json:"output_urn" description:"The uniform resource name (URN) for the VPC."`
	Default    bool   `json:"output_default" description:"A boolean indicating whether or not the VPC is the default one for the region."`
	Created_at string `json:"output_created_at" description:"The date and time of when the VPC was created."`
}
