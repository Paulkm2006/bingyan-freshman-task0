package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Jwt    JwtConfig      `yaml:"jwt"`
	Server ServerConfig   `yaml:"server"`
	DB     PostgresConfig `yaml:"postgres"`
	Admin  AdminConfig    `yaml:"admin"`
}
type JwtConfig struct {
	Secret string `yaml:"secret"`
	Expire int64  `yaml:"expire"`
}
type ServerConfig struct {
	Port string `yaml:"port"`
	Ver  string `yaml:"ver"`
}
type PostgresConfig struct {
	Dsn string `yaml:"dsn"`
}
type AdminConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var Config ConfigStruct

func InitConfig() {
	var configFile []byte
	var err error
	configFile, err = os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err)
	}
}
