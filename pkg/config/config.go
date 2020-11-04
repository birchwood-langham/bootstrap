package config

import (
	"math"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	// VersionKey is the configuration key for retrieving application version number
	VersionKey = "version"
	// ServiceNameKey is the configuration key for retrieving the service name
	ServiceNameKey = "service.name"
	// LogFilePathKey is the configuration key for retrieving the path for the log file generated by the service
	LogFilePathKey = "log.filepath"
	// LogLevelKey is the configuration key for retrieving the logging level
	LogLevelKey = "log.level"
	// LogFileMaxSize is the configuration key for retrieving the log file max size configuration
	LogFileMaxSize = "log.max-size"
	// LogFileMaxBackups is the configuration key for retrieving the log file max backups configuration
	LogFileMaxBackups = "log.max-backups"
	// LogFileMaxAge is the configuration key for retrieving the log file max age configuration
	LogFileMaxAge = "log.max-age"
	// LogFileCompress is the configuration key for retrieving the log file compression configuration
	LogFileCompress = "log.compress"
)

type Config struct {
	path []string
}

func Get(path ...string) *Config {
	return &Config{path: path}
}

func (c *Config) String(d string) string {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetString(k)
	}

	return d
}

func (c *Config) Int(d int) int {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetInt(k)
	}

	return d
}

func (c *Config) Int8(d int8) int8 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		v := viper.GetInt32(k)

		if v >= math.MinInt8 && v <= math.MaxInt8 {
			return int8(v)
		}
	}

	return d
}

func (c *Config) Int16(d int16) int16 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		v := viper.GetInt32(k)

		if v >= math.MinInt16 && v <= math.MaxInt16 {
			return int16(v)
		}
	}

	return d
}

func (c *Config) Int32(d int32) int32 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetInt32(k)
	}

	return d
}

func (c *Config) Int64(d int64) int64 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetInt64(k)
	}

	return d
}

func (c *Config) Value(d interface{}) interface{} {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.Get(k)
	}

	return d
}

func (c *Config) Bool(d bool) bool {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetBool(k)
	}

	return d
}

func (c *Config) Float64(d float64) float64 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetFloat64(k)
	}

	return d
}

func (c *Config) Float32(d float32) float32 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		v := viper.GetFloat64(k)

		if v >= -math.MaxFloat32 && v <= math.MaxFloat32 {
			return float32(v)
		}
	}

	return d
}

func (c *Config) StringMap(d map[string]interface{}) map[string]interface{} {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetStringMap(k)
	}

	return d
}

func (c *Config) StringMapString(d map[string]string) map[string]string {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetStringMapString(k)
	}

	return d
}

func (c *Config) StringSlice(d []string) []string {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetStringSlice(k)
	}

	return d
}

func (c *Config) Time(d time.Time) time.Time {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetTime(k)
	}

	return d
}

func (c *Config) Duration(d time.Duration) time.Duration {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetDuration(k)
	}

	return d
}

func (c *Config) Uint(d uint) uint {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetUint(k)
	}

	return d
}

func (c *Config) Uint8(d uint8) uint8 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		v := viper.GetUint(k)

		if v <= math.MaxUint8 {
			return uint8(v)
		}
	}

	return d
}

func (c *Config) Uint16(d uint16) uint16 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		v := viper.GetUint(k)

		if v <= math.MaxUint16 {
			return uint16(v)
		}
	}

	return d
}

func (c *Config) Uint32(d uint32) uint32 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetUint32(k)
	}

	return d
}

func (c *Config) Uint64(d uint64) uint64 {
	k := mkString(".", c.path...)

	if viper.IsSet(k) {
		return viper.GetUint64(k)
	}

	return d
}

func mkString(sep string, input ...string) string {
	b := strings.Builder{}

	addSep := false

	for _, i := range input {
		if addSep {
			b.WriteString(sep)
		}

		b.WriteString(i)

		addSep = true
	}

	return b.String()
}
