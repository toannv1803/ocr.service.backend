package config

import "github.com/spf13/viper"

type Config struct {
	*viper.Viper
}

func (q *Config) Refresh() {
	viper.SetDefault("MONGO_HOST", "localhost")
	viper.SetDefault("MONGO_PORT", "27017")
	viper.SetDefault("MONGO_USER", "")
	viper.SetDefault("MONGO_PASSWORD", "")
}

func NewConfig() *Config {
	var q Config
	q.Viper = viper.New()
	q.Refresh()
	return &q
}
