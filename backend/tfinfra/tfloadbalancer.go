package tfinfra

type LoadBalancer struct {
	Amount                       int              `json:"amount" validate:"required"`
	Name                         *string          `json:"name" description:"(Required) The Load Balancer name" validate:"required"`
	Region                       *string          `json:"region" description:"(Required) The region to start in" validate:"required"`
	Size                         *string          `json:"size" description:"(Optional) The size of the Load Balancer. It must be either lb-small, lb-medium, or lb-large. Defaults to lb-small. Only one of size or size_unit may be provided." validate:"omitempty"`
	SizeUnit                     *int             `json:"size_unit" description:"(Optional) The size of the Load Balancer. It must be in the range (1, 100). Defaults to 1. Only one of size or size_unit may be provided." validate:"omitempty"`
	Algorithm                    *string          `json:"algorithm" description:"(Optional) The load balancing algorithm used to determine which backend Droplet will be selected by a client. It must be either round_robin or least_connections. The default value is round_robin." validate:"omitempty"`
	ForwardingRule               *forwarding_rule `json:"forwarding_rule" description:"(Required) A list of forwarding_rule to be assigned to the Load Balancer. The forwarding_rule block is documented below." validate:"required,dive"`
	Healthcheck                  *healthcheck     `json:"healthcheck" description:"(Optional) A healthcheck block to be assigned to the Load Balancer. The healthcheck block is documented below. Only 1 healthcheck is allowed." validate:"omitempty,dive"`
	StickySessions               *sticky_sessions `json:"sticky_sessions" description:"(Optional) A sticky_sessions block to be assigned to the Load Balancer. The sticky_sessions block is documented below. Only 1 sticky_sessions block is allowed." validate:"omitempty,dive"`
	RedirectHttpToHttps          *bool            `json:"redirect_http_to_https" description:"(Optional) A boolean value indicating whether HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443. Default value is false." validate:"omitempty"`
	EnableProxyProtocol          *bool            `json:"enable_proxy_protocol" description:"(Optional) A boolean value indicating whether PROXY Protocol should be used to pass information from connecting client requests to the backend service. Default value is false." validate:"omitempty"`
	EnableBackendKeepalive       *bool            `json:"enable_backend_keepalive" description:"(Optional) A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is false." validate:"omitempty"`
	HttpIdleTimeoutSeconds       *int             `json:"http_idle_timeout_seconds" description:"(Optional) Specifies the idle timeout for HTTPS connections on the load balancer in seconds." validate:"omitempty"`
	DisableLetsEncryptDnsRecords *bool            `json:"disable_lets_encrypt_dns_records" description:"(Optional) A boolean value indicating whether to disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer. Default value is false." validate:"omitempty"`
	ProjectId                    *string          `json:"project_id" description:"(Optional) The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user's default project." validate:"omitempty"`
	VpcUuid                      *string          `json:"vpc_uuid" description:"(Optional) The ID of the VPC where the load balancer will be located." validate:"omitempty"`
	DropletIds                   *[]string        `json:"droplet_ids" description:"(Optional) A list of the IDs of each droplet to be attached to the Load Balancer." validate:"omitempty"`
	DropletTag                   *[]string        `json:"droplet_tag" description:"(Optional) The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer." validate:"omitempty"`
	Outputs                      LBOutputs        `json:"outputs" description:"The attributes exported after creation of the resource."`
}

type forwarding_rule struct {
	EntryProtocol   *string `json:"entry_protocol" description:"(Required) The protocol used for traffic to the Load Balancer. The possible values are: http, https, http2, http3, tcp, or udp." validate:"required"`
	EntryPort       *int    `json:"entry_port" description:"(Required) An integer representing the port on which the Load Balancer instance will listen." validate:"required"`
	TargetProtocol  *string `json:"target_protocol" description:"(Required) The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: http, https, http2, tcp, or udp." validate:"required"`
	TargetPort      *int    `json:"target_port" description:"(Required) An integer representing the port on the backend Droplets to which the Load Balancer will send traffic." validate:"required"`
	CertificateName *string `json:"certificate_name" description:"(Optional) The unique name of the TLS certificate to be used for SSL termination." validate:"omitempty"`
	TlsPassthrough  *bool   `json:"tls_passthrough" description:"(Optional) A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is false." validate:"omitempty"`
}

type healthcheck struct {
	Protocol               *string `json:"protocol" description:"(Required) The protocol used for health checks sent to the backend Droplets. The possible values are http, https or tcp." validate:"required"`
	Port                   *int    `json:"port" description:"(Optional) An integer representing the port on the backend Droplets on which the health check will attempt a connection." validate:"omitempty"`
	Path                   *string `json:"path" description:"(Optional) The path on the backend Droplets to which the Load Balancer instance will send a request." validate:"omitempty"`
	CheckIntervalSeconds   *int    `json:"check_interval_seconds" description:"(Optional) The number of seconds between between two consecutive health checks. If not specified, the default value is 10." validate:"omitempty"`
	ResponseTimeoutSeconds *int    `json:"response_timeout_seconds" description:"(Optional) The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is 5." validate:"omitempty"`
	UnhealthyThreshold     *int    `json:"unhealthy_threshold" description:"(Optional) The number of times a health check must fail for a backend Droplet to be marked \"unhealthy\" and be removed from the pool. If not specified, the default value is 3." validate:"omitempty"`
	HealthyThreshold       *int    `json:"healthy_threshold" description:"(Optional) The number of times a health check must pass for a backend Droplet to be marked \"healthy\" and be re-added to the pool. If not specified, the default value is 5." validate:"omitempty"`
}

type sticky_sessions struct {
	Type             *string `json:"type" description:"(Required) An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are cookies or none. If not specified, the default value is none." validate:"required"`
	CookieName       *string `json:"cookie_name" description:"(Optional) The name to be used for the cookie sent to the client. This attribute is required when using cookies for the sticky sessions type." validate:"omitempty"`
	CookieTtlSeconds *int    `json:"cookie_ttl_seconds" description:"(Optional) The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using cookies for the sticky sessions type." validate:"omitempty"`
}

type LBOutputs struct {
	Id              *string `json:"id" description:"The ID of the Load Balancer"`
	Ip              *string `json:"ip" description:"The ip of the Load Balancer"`
	Urn             *string `json:"urn" description:"The uniform resource name for the Load Balancer"`
	CertificateName *string `json:"certificate_name" description:"The unique name of the TLS certificate that is used for SSL termination."`
}
