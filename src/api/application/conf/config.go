// Package conf contains project configurations, for example, the url of a database.
package conf

import (
	"time"

	"github.com/spf13/viper"
)

// FuryConfiguration interface.
type FuryConfiguration interface {
	GetBool(string, bool) bool
	GetString(string, string) string
	GetInt(string, int) int
	GetFloat64(string, float64) float64
	GetUint(string, uint) uint
	GetDuration(string, time.Duration) time.Duration
}

// ViperConfiguration interface.
type ViperConfiguration interface {
	GetBool(string) bool
	GetString(string) string
	GetStringSlice(string) []string
	GetInt(string) int
	GetFloat64(string) float64
	GetDuration(string) time.Duration
}

type configViper struct{}

func getYMLNewConfig() ViperConfiguration {
	return &configViper{}
}

func (c *configViper) GetString(key string) string {
	return viper.GetString(key)
}

func (c *configViper) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func (c *configViper) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (c *configViper) GetInt(key string) int {
	return viper.GetInt(key)
}

func (c *configViper) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func (c *configViper) GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}
