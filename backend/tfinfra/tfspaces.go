package tfinfra

type Spaces struct {
	Amount        int             `json:"amount" validate:"required"`
	Name          *string         `json:"name" description:"(Required) The name of the bucket" validate:"required"`
	Region        *string         `json:"region" description:"The region where the bucket resides (Defaults to nyc3)" validate:"required"`
	Acl           *string         `json:"acl" description:"Canned ACL applied on bucket creation (private or public-read)" validate:"required"`
	CorsRule      *[]cors_rule    `json:"cors_rule" description:"(Optional) A rule of Cross-Origin Resource Sharing (documented below)." validate:"omitempty,dive"`
	LifecycleRule *lifecycle_rule `json:"lifecycle_rule" description:"(Optional) A configuration of object lifecycle management (documented below)." validate:"omitempty,dive"`
	Versioning    *versioning     `json:"versioning" description:"(Optional) A state of versioning (documented below)" validate:"omitempty" validate:"omitempty,dive"`
	ForceDestroy  *bool           `json:"force_destroy" description:"Unless true, the bucket will only be destroyed if empty (Defaults to false)" validate:"omitempty"`
	Outputs       SpacesOutputs   `json:"outputs" description:"The attributes exported after creation of the resource."`
}

type cors_rule struct {
	Allowed_methods *[]string `json:"allowed_methods" description:"(Optional) A list of headers that will be included in the CORS preflight request's Access-Control-Request-Headers. A header may contain one wildcard (e.g. x-amz-*)." validate:"required"`
	Allowed_origins *[]string `json:"allowed_origins" description:"(Required) A list of HTTP methods (e.g. GET) which are allowed from the specified origin." validate:"required"`
	Allowed_headers *[]string `json:"allowed_headers" description:"(Required) A list of hosts from which requests using the specified methods are allowed. A host may contain one wildcard (e.g. http://*.example.com)." validate:"omitempty"`
	Max_age_seconds *int      `json:"max_age_seconds" description:"(Optional) The time in seconds that browser can cache the response for a preflight request." validate:"omitempty"`
}

type lifecycle_rule struct {
	Enabled                            *bool                          `json:"enabled" description:"(Optional) Unique identifier for the rule." validate:"required,boolean"`
	Id                                 *string                        `json:"id" description:"(Optional) Object key prefix identifying one or more objects to which the rule applies." validate:"omitempty"`
	Prefix                             *string                        `json:"prefix" description:"(Required) Specifies lifecycle rule status." validate:"omitempty"`
	AbortIncompleteMultipartUploadDays *int                           `json:"abort_incomplete_multipart_upload_days" description:"(Optional) Specifies the number of days after initiating a multipart upload when the multipart upload must be completed or else Spaces will abort the upload." validate:"omitempty"`
	Expiration                         *expiration                    `json:"expiration" description:"(Optional) Specifies a time period after which applicable objects expire (documented below)." validate:"omitempty,dive,required"`
	NoncurrentVersionExpiration        *noncurrent_version_expiration `json:"noncurrent_version_expiration" description:"(Optional) Specifies when non-current object versions expire (documented below)." validate:"omitempty,dive,required"`
}

// // At least one of expiration or noncurrent_version_expiration must be specified.

type expiration struct {
	Date                      *string `json:"date" description:"(Optional) Specifies the date/time after which you want applicable objects to expire. The argument uses RFC3339 format, e.g. \"2020-03-22T15:03:55Z\" or parts thereof e.g. \"2019-02-28\"." validate:"omitempty"`
	Days                      *int    `json:"days" description:"(Optional) Specifies the number of days after object creation when the applicable objects will expire." validate:"omitempty"`
	ExpiredObjectDeleteMarker *bool   `json:"expired_object_delete_marker" description:"(Optional) On a versioned bucket (versioning-enabled or versioning-suspended bucket), setting this to true directs Spaces to delete expired object delete markers." validate:"omitempty"`
}

type noncurrent_version_expiration struct {
	Days *int `json:"days" description:"(Required) Specifies the number of days after which an object's non-current versions expire." validate:"required"`
}

type versioning struct {
	Enabled *bool `json:"enabled" description:"(Optional) Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket." validate:"omitempty"`
}

type SpacesOutputs struct {
	Name               *string `json:"name" description:"The name of the bucket"`
	Urn                *string `json:"urn" description:"The uniform resource name for the bucket"`
	Region             *string `json:"region" description:"The name of the region"`
	Bucket_domain_name *string `json:"bucket_domain_name" description:"The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)"`
	Endpoint           *string `json:"endpoint" description:"The FQDN of the bucket without the bucket name (e.g. nyc3.digitaloceanspaces.com)"`
}
