package config

import "github.com/owncloud/ocis/v2/ocis-pkg/tracing"

// Tracing sets the tracing parameters for the frontend service.
type Tracing struct {
	Enabled   bool   `yaml:"enabled" env:"OCIS_TRACING_ENABLED;FRONTEND_TRACING_ENABLED" desc:"Activates tracing." introductionVersion:"pre5.0"`
	Type      string `yaml:"type" env:"OCIS_TRACING_TYPE;FRONTEND_TRACING_TYPE" desc:"The type of tracing. Defaults to '', which is the same as 'jaeger'. Allowed tracing types are 'jaeger', 'otlp' and '' as of now." introductionVersion:"pre5.0"`
	Endpoint  string `yaml:"endpoint" env:"OCIS_TRACING_ENDPOINT;FRONTEND_TRACING_ENDPOINT" desc:"The endpoint of the tracing agent." introductionVersion:"pre5.0"`
	Collector string `yaml:"collector" env:"OCIS_TRACING_COLLECTOR;FRONTEND_TRACING_COLLECTOR" desc:"The HTTP endpoint for sending spans directly to a collector, i.e. http://jaeger-collector:14268/api/traces. Only used if the tracing endpoint is unset." introductionVersion:"pre5.0"`
}

// Convert Tracing to the tracing package's Config struct.
func (t Tracing) Convert() tracing.Config {
	return tracing.Config{
		Enabled:   t.Enabled,
		Type:      t.Type,
		Endpoint:  t.Endpoint,
		Collector: t.Collector,
	}
}
