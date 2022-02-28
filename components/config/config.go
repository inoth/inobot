package config

import (
	"github.com/spf13/viper"
)

var Cfg *viper.Viper

type ViperConfig struct {
	Path         string
	defaultValue map[string]interface{}
}

func (m ViperConfig) SetDefaultValue(defaultValue map[string]interface{}) *ViperConfig {
	m.defaultValue = make(map[string]interface{})
	for k, v := range defaultValue {
		m.defaultValue[k] = v
	}
	return &m
}

func (m *ViperConfig) loadDefaultValue() {
	for k, v := range m.defaultValue {
		Cfg.SetDefault(k, v)
	}
}

func (m *ViperConfig) Init() error {
	v := viper.New()
	if m.Path == "" {
		m.Path = "config.yaml"
	}
	v.SetConfigFile(m.Path)
	// v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	Cfg = v
	m.loadDefaultValue()
	return nil
}

// func selectConfigName(path []string) string {
// 	if len(path) > 0 {
// 		return path[0]
// 	} else {
// 		e := os.Getenv("GORUNEVN")
// 		if len(e) > 0 {
// 			return "config." + e
// 		}
// 		return "config"
// 	}
// }
