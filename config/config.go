package config

import (
	"bytes"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// Config represents the application configuration
	Config struct {
	}
)

// Validate validates configuration
func (c Config) Validate() error {
	return validator.New().Struct(c)
}

// Init load and initiates configuration
func Init(path string) Config {
	var cfg Config

	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewReader([]byte(Default))); err != nil {
		logrus.Panicf("error loading default configs: %s", err)
	}

	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.SetEnvPrefix(Namespace)
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	err := v.MergeInConfig()
	if err != nil {
		logrus.Warn("no cfg file found. Using defaults and environment variables")
	}

	if err := v.UnmarshalExact(&cfg); err != nil {
		logrus.Fatalf("invalid configuration: %s", err)
	}

	if err := cfg.Validate(); err != nil {
		logrus.Fatalf("invalid configuration: %s", err)
	}

	return cfg
}
