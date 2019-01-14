package config

import (
	serviceconfig "github.com/IMQS/serviceconfigsgo"
)

const configVersion = 1
const serviceConfigFileName = "gateway-store.json"
const serviceName = "Gateway Store"

// DB hosts the configuration specific for the database package
type DB struct {
	Dbname   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

// Server contains the configuration for the HTTP server
type Server struct {
	Port string `json:"port"`
}

// Config hosts the general configuration for the entire service
type Config struct {
	Server    *Server `json:"server"`
	LogFile   string  `json:"logfile"`
	LogLevel  string  `json:"loglevel"`
	PublicDir string  `json:"publicdir"`
	Db        *DB     `json:"db"`
}

// NewConfig gets the service configuration either from file or the configuration service
func NewConfig(filename string) (*Config, error) {
	c := &Config{}

	if err := serviceconfig.GetConfig(filename, serviceName, configVersion, serviceConfigFileName, c); err != nil {
		return nil, err
	}
	return c, nil
}
