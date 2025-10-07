package config

import (
	"fmt"
	"os"
)

func GetEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}

type ServiceConfig struct {
	Port         string
	ServiceToken string
}

func LoadConfig() *ServiceConfig {
	return &ServiceConfig{
		Port:         GetEnv("CORE_PORT", "9000"),
		ServiceToken: GetEnv("SERVICE_TOKEN", "servicesecret"),
	}
}

func (c *ServiceConfig) Addr() string {
	return fmt.Sprintf(":%s", c.Port)
}
