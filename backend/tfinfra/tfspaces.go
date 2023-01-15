package tfinfra

type Spaces struct {
	Amount        int            `json:"amount" validate:"required"`
	Name          *string        `json:"name" validate:"required"`
	Region        *string        `json:"region" validate:"required"`
	Acl           *string        `json:"acl" validate:"required"`
	CorsRule      *[]corsRule    `json:"cors_rule" validate:"omitempty,dive"`
	LifecycleRule *lifecycleRule `json:"lifecycle_rule" validate:"omitempty,dive"`
	Versioning    *versioning    `json:"versioning" validate:"omitempty" validate:"omitempty,dive"`
	ForceDestroy  *bool          `json:"force_destroy" validate:"omitempty"`
}

type corsRule struct {
	Allowed_methods *[]string `json:"allowed_methods" validate:"required"`
	Allowed_origins *[]string `json:"allowed_origins" validate:"required"`
	Allowed_headers *[]string `json:"allowed_headers" validate:"omitempty"`
	Max_age_seconds *int      `json:"max_age_seconds" validate:"omitempty"`
}

type lifecycleRule struct {
	Enabled                            *bool                        `json:"enabled" validate:"required,boolean"`
	Id                                 *string                      `json:"id" validate:"omitempty"`
	Prefix                             *string                      `json:"prefix" validate:"omitempty"`
	AbortIncompleteMultipartUploadDays *int                         `json:"abort_incomplete_multipart_upload_days" validate:"omitempty"`
	Expiration                         *expiration                  `json:"expiration" validate:"omitempty,dive,required"`
	NoncurrentVersionExpiration        *noncurrentVersionExpiration `json:"noncurrent_version_expiration" validate:"omitempty,dive,required"`
}

// // At least one of expiration or noncurrent_version_expiration must be specified.

type expiration struct {
	Date                      *string `json:"date" validate:"omitempty"`
	Days                      *int    `json:"days" validate:"omitempty"`
	ExpiredObjectDeleteMarker *bool   `json:"expired_object_delete_marker" validate:"omitempty"`
}

type noncurrentVersionExpiration struct {
	Days *int `json:"days" validate:"required"`
}

type versioning struct {
	Enabled *bool `json:"enabled" validate:"omitempty"`
}
