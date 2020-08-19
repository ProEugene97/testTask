package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Provider string
	Redis    string
	HttpHost string
	HttpPort string
	GrpcHost string
	GrpcPort string
	LogLevel string
	Sports   []string
	Timeouts []int
}

func GetConfig() (Config, error) {
	v := viper.New()
	c := Config{}

	v.SetConfigName("config")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return c, err
	}

	c.Provider = v.GetString("linesProvider")
	c.Redis = v.GetString("redis")
	c.HttpHost = v.GetString("http.host")
	c.HttpPort = v.GetString("http.port")
	c.GrpcHost = v.GetString("grpc.host")
	c.GrpcPort = v.GetString("grpc.port")
	c.LogLevel = v.GetString("level")
	c.Sports = v.GetStringSlice("sports")
	c.Timeouts = make([]int, len(c.Sports))
	for i, sport := range c.Sports {
		c.Timeouts[i] = v.GetInt("timeouts." + sport)
	}

	return c, nil
}
