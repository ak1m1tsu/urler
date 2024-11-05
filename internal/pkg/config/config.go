// Package config contains configuration for application
package config

import "time"

// Config configuration for application
type Config struct {
	HTTP HTTPConfig `yaml:"http"`
}

// HTTPConfig configuration for HTTP connection for application
type HTTPConfig struct {
	// Port HTTP port for application
	Port int `yaml:"port"`
	// Timeout timeouts for HTTP server
	Timeout HTTPTimeoutConfig
}

// HTTPTimeoutConfig configuration for HTTP timeouts
type HTTPTimeoutConfig struct {
	// Read timeout for HTTP request
	Read time.Duration `yaml:"read"`
	// Write timeout for HTTP response
	Write time.Duration `yaml:"write"`
	// Idle timeout for idle connection with client
	Idle time.Duration `yaml:"idle"`
}
