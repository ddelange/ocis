package config

// Debug defines the available debug configuration. Not used at the moment
type Debug struct {
	Addr   string `yaml:"addr" env:"COLLABORATION_DEBUG_ADDR" desc:"Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed." introductionVersion:"5.1"`
	Token  string `yaml:"token" env:"COLLABORATION_DEBUG_TOKEN" desc:"Token to secure the metrics endpoint." introductionVersion:"5.1"`
	Pprof  bool   `yaml:"pprof" env:"COLLABORATION_DEBUG_PPROF" desc:"Enables pprof, which can be used for profiling." introductionVersion:"5.1"`
	Zpages bool   `yaml:"zpages" env:"COLLABORATION_DEBUG_ZPAGES" desc:"Enables zpages, which can be used for collecting and viewing in-memory traces." introductionVersion:"5.1"`
}