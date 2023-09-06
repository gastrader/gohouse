package utils

import (
	"os"

	"gopkg.in/ini.v1"
)

func GetIni(section, key, defaultValue string) string {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		Errorf("fail to read INI file: %v", err)
		os.Exit(1)
	}
	if value := cfg.Section(section).Key(key).String(); value != "" {
		return value 
	}
	return defaultValue
}