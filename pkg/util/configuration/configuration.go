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

func LoadConfig(config_dir string) (*Configuration, error) {
	if config_dir == "" {
		config_dir = getEnv("PWD")
	}

	viper.SetConfigName("base")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(config_dir)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error reading config file: %w", err)
	}

	switch app_env := getEnv("APP_ENVIRONMENT"); app_env {
	case "production":
		fmt.Println("Loading " + config_dir + "/production.yaml")
		viper.SetConfigName("production")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(config_dir)

		err := viper.MergeInConfig()
		if err != nil {
			return nil, fmt.Errorf("fatal error reading production config file: %w", err)
		}
	case "development":
		fmt.Println("Loading " + config_dir + "/development.yaml")
		viper.SetConfigName("development")
		viper.AddConfigPath(config_dir + "/development.yaml")

		err := viper.MergeInConfig()
		if err != nil {
			return nil, fmt.Errorf("fatal error reading development config file: %w", err)
		}
	default:
		return nil, fmt.Errorf("APP_ENVIRONMENT env var set to invalid val: %w", err)
	}
	var cfg = new(Configuration)

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}
	return cfg, err
}

type Configuration struct {
	Server   *Server   `mapstructure:"server,omitempty"`
	Database *Database `mapstructure:"database,omitempty"`
}

type Server struct {
	Port int    `mapstructure:"port,omitempty"`
	Host string `mapstructure:"host,omitempty"`
}

type Database struct {
	Uname         string `mapstructure:"uname,omitempty"`
	Password      string `mapstructure:"password,omitempty"`
	Port          int    `mapstructure:"port,omitempty"`
	Host          string `mapstructure:"host,omitempty"`
	Database_name string `mapstructure:"database_name,omitempty"`
}

func (d *Database) Conn_str() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", d.Uname, d.Password, d.Host, d.Port, d.Database_name)
}

func (d *Database) Default_conn_str() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", d.Uname, d.Password, d.Host, d.Port)
}
