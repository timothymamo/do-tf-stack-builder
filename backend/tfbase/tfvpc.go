package tfbase

type Vpc struct {
	Amount      int     `json:"amount" validate:"required"`
	Name        *string `json:"name" variable:"vpc_name" validate:"required"`
	Region      *string `json:"region" validate:"required"`
	Description *string `json:"description"`
	Iprange     *string `json:"ip_range" validate:"omitempty,cidrv4"`
}
