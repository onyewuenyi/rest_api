package config_test

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Setenv("APP_ENVIRONMENT", "production")
}
