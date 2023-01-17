package tfinfra

type Database struct {
	Amount             int                 `json:"amount" validate:"required"`
	Name               *string             `json:"name" description:"(Required) The name of the database cluster." validate:"required"`
	Engine             *string             `json:"engine" description:"(Required) Database engine used by the cluster (ex. pg for PostreSQL, mysql for MySQL, redis for Redis, or mongodb for MongoDB)." validate:"required"`
	Size               *string             `json:"size" description:"(Required) Database Droplet size associated with the cluster (ex. db-s-1vcpu-1gb). See here for a list of valid size slugs." validate:"required"`
	Region             *string             `json:"region" description:"(Required) DigitalOcean region where the cluster will reside." validate:"required"`
	NodeCount          *int                `json:"node_count" description:"(Required) Number of nodes that will be included in the cluster." validate:"required"`
	Version            *string             `json:"version" description:"(Required) Engine version used by the cluster (ex. 11 for PostgreSQL 11)." validate:"required"`
	Tags               *[]string           `json:"tags" description:"(Optional) A list of tag names to be applied to the database cluster." validate:"omitempty"`
	PrivateNetworkUuid *string             `json:"private_network_uuid" description:"(Optional) The ID of the VPC where the database cluster will be located." validate:"omitempty"`
	EvictionPolicy     *string             `json:"eviction_policy" description:"(Optional) A string specifying the eviction policy for a Redis cluster. Valid values are: noeviction, allkeys_lru, allkeys_random, volatile_lru, volatile_random, or volatile_ttl." validate:"omitempty"`
	SqlMode            *string             `json:"sql_mode" description:"(Optional) A comma separated string specifying the SQL modes for a MySQL cluster." validate:"omitempty"`
	MaintenanceWindow  *maintenance_window `json:"maintenance_window" description:"(Optional) Defines when the automatic maintenance should be performed for the database cluster." validate:"omitempty,dive,required"`
	Timeouts           *timeouts           `json:"timeouts" description:"(Optional) Customize how long certain operations (create, update and delete operations) are allowed to take before being considered to have failed (e.g \"60m\" for 60 minutes, \"10s\" for ten seconds, or \"2h\")" validate:"omitempty,dive,required"`
	Outputs            DBOutputs           `json:"outputs" description:"The attributes exported after creation of the resource."`
}

type maintenance_window struct {
	Day  *string `json:"day" description:"(Required) The day of the week on which to apply maintenance updates." validate:"required"`
	Hour *string `json:"hour" description:"(Required) The hour in UTC at which maintenance updates will be applied in 24 hour format." validate:"required"`
}

type timeouts struct {
	Create *string `json:"create" validate:"omitempty"`
	Delete *string `json:"delete" validate:"omitempty"`
	Update *string `json:"update" validate:"omitempty"`
}

type DBOutputs struct {
	Id           *string `json:"id" description:"The ID of the database cluster."`
	Urn          *string `json:"urn" description:"The uniform resource name of the database cluster."`
	Host         *string `json:"host" description:"Database cluster's hostname."`
	Private_host *string `json:"private_host" description:"Same as host, but only accessible from resources within the account and in the same region."`
	Port         *int    `json:"port" description:"Network port that the database cluster is listening on."`
	Uri          *string `json:"uri" description:"The full URI for connecting to the database cluster."`
	Private_uri  *string `json:"private_uri" description:"Same as uri, but only accessible from resources within the account and in the same region."`
	Database     *string `json:"database" description:"Name of the cluster's default database."`
	User         *string `json:"user" description:"Username for the cluster's default user."`
	Password     *string `json:"password" description:"Password for the cluster's default user."`
}
