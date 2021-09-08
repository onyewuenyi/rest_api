package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic("env var not found")
}

func InitConfig() (*Configuration, error) {
	var config_dir = getEnv("PWD") + "/configuration"
	viper.SetConfigName("base")
	viper.AddConfigPath(config_dir)
	viper.ReadInConfig()

	switch app_env := getEnv("APP_ENVIRONMENT"); app_env {
	case "production":
		viper.SetConfigName("production")
		viper.AddConfigPath(config_dir + "/production.yaml")
		viper.MergeInConfig()
	case "staging":
		viper.SetConfigName("staging")
		viper.AddConfigPath(config_dir + "/staging.yaml")
		viper.MergeInConfig()
	case "development":
		viper.SetConfigName("development")
		viper.AddConfigPath(config_dir + "/development.yaml")
		viper.MergeInConfig()
	default:
		panic("invalid environment")
	}
	var cfg = new(Configuration)

	err := viper.Unmarshal(cfg)
	// if err is null fatal
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, err
}

type Configuration struct {
	Server   *Server   `mapstructure:"server,omitempty"`
	Database *Database `mapstructure:"database,omitempty"`
}

type Server struct {
	Port int `mapstructure:"port,omitempty"`
	Host int `mapstructure:"host,omitempty"`
}

type Database struct {
	uname         string `mapstructure:"uname,omitempty"`
	password      string `mapstructure:"password,omitempty"`
	port          int    `mapstructure:"port,omitempty"`
	host          string `mapstructure:"host,omitempty"`
	database_name string `mapstructure:"database_name,omitempty"`
}

func (d *Database) conn_str() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", d.uname, d.password, d.host, d.port, d.database_name)
}

func (d *Database) default_conn_str() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", d.uname, d.password, d.host, d.port)
}
