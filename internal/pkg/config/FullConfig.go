package config

import (
	"errors"

	"github.com/spf13/viper"

	"gin-code-generator/internal/pkg/util"
)

type FullConfig struct {
	SystemConfig     SystemConfig     `mapstructure:"system" json:"system" yaml:"system"`
	MySqlConfig      MySqlConfig      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PostgreSQLConfig PostgreSQLConfig `mapstructure:"postgreSQL" json:"postgreSQL" yaml:"postgreSQL"`
	OracleConfig     OracleConfig     `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	RedisConfig      RedisConfig      `mapstructure:"redis" json:"redis" yaml:"redis"`
}

func LoadConfig() (FullConfig, error) {
	conifgFile := "./config.yaml"

	if !util.CheckFileIsExist(conifgFile) {
		return FullConfig{}, errors.New("config file not found")
	}

	v := viper.New()
	v.SetConfigFile(conifgFile)
	err := v.ReadInConfig()
	if err != nil {
		return FullConfig{}, err
	}

	var fullConfig FullConfig
	err = v.Unmarshal(&fullConfig)
	if err != nil {
		return FullConfig{}, err
	}

	return fullConfig, nil
}
