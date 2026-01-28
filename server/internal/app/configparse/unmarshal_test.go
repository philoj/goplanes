package configparse_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/philoj/goplanes/server/internal/app/configparse"
	"github.com/stretchr/testify/assert"
)

type config struct {
	Name              string `mapstructure:"name"`
	Port              int    `mapstructure:"port"`
	Debug             bool   `mapstructure:"debug"`
	MonitoringService svcCfg `mapstructure:"monitoring_service"`
}

type svcCfg struct {
	HostURL string        `mapstructure:"host_url"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func TestFileOverrideDefaults(t *testing.T) {
	defaults := map[string]any{
		"name":  "default",
		"port":  8080,
		"debug": true,
		"monitoring_service": map[string]interface{}{
			"host_url": "http://localhost",
			"timeout":  5 * time.Second,
		},
	}

	expected := config{
		Name:  "test",
		Port:  9090,
		Debug: true,
		MonitoringService: svcCfg{
			HostURL: "http://localhost",
			Timeout: 10 * time.Second,
		},
	}
	fileContents := fmt.Sprintf(`
name: %s
port: %d
monitoring_service:
  timeout: 10s
`,
		expected.Name, expected.Port)
	testFile := "test_config.yaml"
	err := os.WriteFile(testFile, []byte(fileContents), 0600)
	defer func() {
		err = os.Remove(testFile)
		assert.NoError(t, err, "Failed to rollback temporary config file")
	}()
	assert.NoError(t, err, "Failed to create temporary config file")

	var cfg config
	err = configparse.Unmarshal(&cfg, configparse.FromMap(defaults), configparse.FromPath("test_config.yaml"))
	assert.NoError(t, err, "Failed to unmarshal")

	assert.Equal(t, expected, cfg)
}
