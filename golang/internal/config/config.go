package config

import (
	"encoding/json"
	"os"
)

// Config represents the benchmark configuration
type Config struct {
	NumRecords     int        `json:"num_records"`
	NumIterations  int        `json:"num_iterations"`
	DataGeneration DataConfig `json:"data"`
}

// DataConfig represents data generation settings
type DataConfig struct {
	StringLength   int `json:"string_length"`
	BinarySize     int `json:"binary_size"`
	ArrayMinSize   int `json:"array_min_size"`
	ArrayMaxSize   int `json:"array_max_size"`
	MapMinSize     int `json:"map_min_size"`
	MapMaxSize     int `json:"map_max_size"`
	StructMinSize  int `json:"struct_min_size"`
	StructMaxSize  int `json:"struct_max_size"`
}

// LoadConfig reads the benchmark configuration from a JSON file
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// DefaultConfig returns the default configuration if file cannot be loaded
func DefaultConfig() *Config {
	return &Config{
		NumRecords:    1_000_000,
		NumIterations: 10,
		DataGeneration: DataConfig{
			StringLength:  20,
			BinarySize:    32,
			ArrayMinSize:  1,
			ArrayMaxSize:  5,
			MapMinSize:    1,
			MapMaxSize:    3,
			StructMinSize: 1,
			StructMaxSize: 3,
		},
	}
}
