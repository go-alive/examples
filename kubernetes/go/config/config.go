// Package config implements go-config with env and k8s configmap
package config

import (
	"github.com/go-alive/go-micro/config"
	"github.com/go-alive/go-micro/config/source/env"
	"github.com/go-alive/go-plugins/config/source/configmap/v2"
)

// NewConfig returns config with env and k8s configmap setup
func NewConfig(opts ...config.Option) config.Config {
	cfg, _ := config.NewConfig()
	cfg.Load(
		env.NewSource(),
		configmap.NewSource(),
	)
	return cfg
}
