package configs

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

const (
	Version = "version"
	Addr    = "address"
)

func Init(configPath string) {
	dir := filepath.Dir(configPath)
	fileName := filepath.Base(configPath)
	viperConfigName := strings.TrimRight(fileName, ".yaml")
	viper.AddConfigPath(dir)
	viper.SetConfigName(viperConfigName)
	viper.SetConfigType("yaml")
}
