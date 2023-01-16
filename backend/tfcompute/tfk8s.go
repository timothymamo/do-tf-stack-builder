package tfcompute

type K8s struct {
	Amount            int                 `json:"amount" validate:"required"`
	Name              *string             `json:"name" description:"(Required) A name for the Kubernetes cluster." validate:"required"`
	Region            *string             `json:"region" description:"(Required) The slug identifier for the region where the Kubernetes cluster will be created." validate:"required"`
	Version           *string             `json:"version" description:"(Required) The slug identifier for the version of Kubernetes used for the cluster. Use doctl to find the available versions doctl kubernetes options versions. (Note: A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)" validate:"required,oneof=1.25.4-do.0 1.24.8-do.0 1.23.14-do.0"`
	Vpcuuid           *string             `json:"vpc_uuid" description:"(Optional) The ID of the VPC where the Kubernetes cluster will be located." validate:"omitempty"`
	AutoUpgrade       *bool               `json:"auto_upgrade" description:"(Optional) A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window." validate:"omitempty"`
	SurgeUpgrade      *bool               `json:"surge_upgrade" description:"(Optional) Enable/disable surge upgrades for a cluster. Default: false" validate:"omitempty"`
	Ha                *bool               `json:"ha" description:"(Optional) Enable/disable the high availability control plane for a cluster. High availability can only be set when creating a cluster. Any update will create a new cluster. Default: false" validate:"omitempty"`
	NodePool          *node_pool          `json:"node_pool" description:"(Required) A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the digitalocean_kubernetes_node_pool resource. The following arguments may be specified:" validate:"required"`
	Tags              *[]string           `json:"tags" description:"(Optional) A list of tag names to be applied to the Kubernetes cluster." validate:"omitempty"`
	MaintenancePolicy *maintenance_policy `json:"maintenance_policy" description:"(Optional) A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. auto_upgrade must be set to true for this to have an effect." validate:"omitempty,dive,required"`
	Outputs           K8sOutputs          `json:"outputs" description:"The attributes exported after creation of the resource."`
}

type node_pool struct {
	Name      *string            `json:"name" description:"(Required) A name for the node pool." validate:"required"`
	Size      *string            `json:"size" description:"(Required) The slug identifier for the type of Droplet to be used as workers in the node pool." validate:"required"`
	NodeCount *int               `json:"node_count" description:"(Optional) The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value." validate:"omitempty"`
	AutoScale *bool              `json:"auto_scale" description:"(Optional) Enable auto-scaling of the number of nodes in the node pool within the given min/max range." validate:"omitempty"`
	MinNodes  *int               `json:"min_nodes" description:"(Optional) If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to." validate:"omitempty"`
	MaxNodes  *int               `json:"max_nodes" description:"(Optional) If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to." validate:"omitempty"`
	Tags      *[]string          `json:"tags" description:"(Optional) A list of tag names applied to the node pool." validate:"omitempty"`
	Labels    *map[string]string `json:"labels" description:"(Optional) A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding Node resources." validate:"omitempty"`
}

type maintenance_policy struct {
	Day       *string `json:"day" description:"(Required) The day of the maintenance window policy. May be one of \"monday\" through \"sunday\", or \"any\" to indicate an arbitrary week day." validate:"required"`
	StartTime *string `json:"start_time" description:"(Required) The start time in UTC of the maintenance window policy in 24-hour clock format / HH:MM notation (e.g., 15:00)." validate:"required"`
}

type K8sOutputs struct {
	Id                string `json:"output_id" description:"A unique ID that can be used to identify and reference a Kubernetes cluster."`
	Cluster_subnet    string `json:"output_cluster_subnet" description:"The range of IP addresses in the overlay network of the Kubernetes cluster."`
	Service_subnet    string `json:"output_service_subnet" description:"The range of assignable IP addresses for services running in the Kubernetes cluster."`
	Ipv4_address      string `json:"output_ipv4_address" description:"The public IPv4 address of the Kubernetes master node. This will not be set if high availability is configured on the cluster (v1.21+)"`
	Endpoint          string `json:"output_endpoint" description:"The base URL of the API server on the Kubernetes master node."`
	Status            string `json:"output_status" description:"A string indicating the current status of the cluster. Potential values include running, provisioning, and errored."`
	Vpcuuid           string `json:"output_vpc_uuid" description:"The ID of the VPC where the Kubernetes cluster is located."`
	CreatedAt         string `json:"output_created_at" description:"The date and time when the Kubernetes cluster was created."`
	UpdatedAt         string `json:"output_updated_at" description:"The date and time when the Kubernetes cluster was last updated."`
	AutoUpgrade       bool   `json:"output_auto_upgrade" description:"A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window."`
	KubeConfig        string `json:"output_kube_config.0" description:"A representation of the Kubernetes cluster's kubeconfig with the following attributes:"`
	NodePool          string `json:"output_node_pool" description:"In addition to the arguments provided, these additional attributes about the cluster's default node pool are exported:"`
	Urn               string `json:"output_urn" description:"The uniform resource name (URN) for the Kubernetes cluster."`
	MaintenancePolicy string `json:"output_maintenance_policy" description:"A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen."`
}
