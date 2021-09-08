package configuration_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/onyewuenyi/rest_api/pkg/util/configuration"
)

func TestConfig(t *testing.T) {
	os.Setenv("APP_ENVIRONMENT", "production")
	cfg, err := configuration.InitConfig()
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}

	fmt.Printf("%#v\n", cfg)
	// assert.Equal(t, 1, cfg{"port"})
}
