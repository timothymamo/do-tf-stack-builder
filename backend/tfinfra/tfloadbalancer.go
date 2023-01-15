package tfinfra

type LoadBalancer struct {
	Amount                       int             `json:"amount" validate:"required"`
	Name                         *string         `json:"name" validate:"required"`
	Region                       *string         `json:"region" validate:"required"`
	ForwardingRule               *forwardingRule `json:"forwarding_rule" validate:"required,dive"`
	Size                         *string         `json:"size" validate:"omitempty"`
	SizeUnit                     *int            `json:"size_unit" validate:"omitempty"`
	Algorithm                    *string         `json:"algorithm" validate:"omitempty"`
	RedirectHttpToHttps          *bool           `json:"redirect_http_to_https" validate:"omitempty"`
	EnableProxyProtocol          *bool           `json:"enable_proxy_protocol" validate:"omitempty"`
	EnableBackendKeepalive       *bool           `json:"enable_backend_keepalive" validate:"omitempty"`
	HttpIdleTimeoutSeconds       *int            `json:"http_idle_timeout_seconds" validate:"omitempty"`
	DisableLetsEncryptDnsRecords *bool           `json:"disable_lets_encrypt_dns_records" validate:"omitempty"`
	ProjectId                    *string         `json:"project_id" validate:"omitempty"`
	VpcUuid                      *string         `json:"vpc_uuid" validate:"omitempty"`
	DropletIds                   *[]string       `json:"droplet_ids" validate:"omitempty"`
	DropletTag                   *[]string       `json:"droplet_tag" validate:"omitempty"`
	Healthcheck                  *healthcheck    `json:"healthcheck" validate:"omitempty,dive"`
	StickySessions               *stickySessions `json:"sticky_sessions" validate:"omitempty,dive"`
}

type forwardingRule struct {
	EntryProtocol   *string `json:"entry_protocol" validate:"required"`
	EntryPort       *int    `json:"entry_port" validate:"required"`
	TargetProtocol  *string `json:"target_protocol" validate:"required"`
	TargetPort      *int    `json:"target_port" validate:"required"`
	CertificateName *string `json:"certificate_name" validate:"omitempty"`
	TlsPassthrough  *bool   `json:"tls_passthrough" validate:"omitempty"`
}

type healthcheck struct {
	Protocol               *string `json:"protocol" validate:"required"`
	Port                   *int    `json:"port" validate:"omitempty"`
	Path                   *string `json:"path" validate:"omitempty"`
	CheckIntervalSeconds   *int    `json:"check_interval_seconds" validate:"omitempty"`
	ResponseTimeoutSeconds *int    `json:"response_timeout_seconds" validate:"omitempty"`
	UnhealthyThreshold     *int    `json:"unhealthy_threshold" validate:"omitempty"`
	HealthyThreshold       *int    `json:"healthy_threshold" validate:"omitempty"`
}

type stickySessions struct {
	Type             *string `json:"type" validate:"required"`
	CookieName       *string `json:"cookie_name" validate:"omitempty"`
	CookieTtlSeconds *int    `json:"cookie_ttl_seconds" validate:"omitempty"`
}
