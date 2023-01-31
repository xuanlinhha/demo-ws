package config

import (
	"demo-ws/common/log"

	"github.com/spf13/viper"
)

func GetConfig(m map[string]interface{}, configPath string, cfg interface{}) error {
	// set default
	v := viper.New()
	for k, val := range m {
		v.SetDefault(k, val)
	}

	// read from file
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warnf("config file not found: %v", err)
		}
	}

	// read from env var
	v.AutomaticEnv()

	log.Infof("%v", v)

	// unmarshall
	err := v.Unmarshal(cfg)
	if err != nil {
		log.Errorf("unable to decode into struct, %v", err)
		return err
	}
	return nil
}
