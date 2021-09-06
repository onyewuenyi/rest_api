package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/onyewuenyi/rest_api/tree/main/pkg/util/config"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	os.Setenv("APP_ENVIRONMENT", "production")
	cfg, err := config.InitConfig()

	fmt.Println(cfg)
	assert.Equal(1, cfg.Server)
}
