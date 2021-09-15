package configuration_test

import (
	"testing"

	"os"

	"github.com/onyewuenyi/rest_api/pkg/util/configuration"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	cases := []struct {
		name     string
		path     string
		wantData *configuration.Configuration
		wantErr  bool
	}{
		{
			name:    "Fail on non-existing file",
			wantErr: true,
		},
		{
			name:    "invalid",
			wantErr: true,
		},
		{
			name: "development",
			wantData: &configuration.Configuration{
				Server: &configuration.Server{
					Port: 8000,
					Host: "127.0.0.1",
				},
				Database: &configuration.Database{
					Host:          "localhost",
					Port:          5432,
					Uname:         "charlesonyewuenyi",
					Password:      "password",
					Database_name: "newsletter",
				},
			},
		},
		{
			name: "production",
			wantData: &configuration.Configuration{
				Server: &configuration.Server{
					Port: 8000,
					Host: "0.0.0.0",
				},
				Database: &configuration.Database{
					Host:          "localhost",
					Port:          5432,
					Uname:         "charlesonyewuenyi",
					Password:      "password",
					Database_name: "newsletter",
				},
			},
		},
	}

	for _, tt := range cases {
		// set env var
		err := os.Setenv("APP_ENVIRONMENT", tt.name)
		assert.Nil(t, err)

		t.Run(tt.name, func(t *testing.T) {
			cfg, _ := configuration.LoadConfig("test_data")
			assert.Equal(t, tt.wantData, cfg)
			// assert.Equal(t, tt.wantErr, err != nil)
		})
	}

}
