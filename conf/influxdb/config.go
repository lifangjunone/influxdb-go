package influxdb_conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

const (
	Name string = "influxdb"
)

var (
	config *Config = &Config{}
)

func GetConfig() *Config {
	return config
}

func GetAddr() string {
	addr := fmt.Sprintf("http://%s:%s", config.Host, config.Port)
	return addr
}

func LoadConfigFromToml(filePath string) (*Config, error) {
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return nil, fmt.Errorf("load config from file %s is error: %s", filePath, err.Error())
	}
	return config, nil
}

type Config struct {
	Token  string `toml:"token" env:"INFLUXDB_TOKEN"`
	Bucket string `toml:"bucket" env:"INFLUXDB_BUCKET"`
	Org    string `toml:"org" env:"INFLUXDB_ORG"`
	Host   string `toml:"host" env:"INFLUXDB_HOST"`
	Port   string `toml:"port" env:"INFLUXDB_PORT"`
}
