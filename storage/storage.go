package storage

import (
	"encoding/json"

	"github.com/blue-jay/core/jsonconfig"
	"github.com/blue-jay/core/storage/driver/mysql"
)

// Info contains the database connection information for the different storage.
type Info struct {
	MySQL mysql.Info `json:"MySQL"`
}

// ParseJSON unmarshals bytes to structs.
func (c *Info) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// LoadConfig reads the configuration file.
func LoadConfig(configFile string) (*Info, error) {
	// Configuration
	config := &Info{}

	// Load the configuration file
	err := jsonconfig.Load(configFile, config)

	// Return the configuration
	return config, err
}
