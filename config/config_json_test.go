// Tests for gocron
package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Settings struct {
	Name string `json:"name"`
}

func TestInitJSON(*testing.T) {
	settings := Settings{}
	settings.Name = "testname"
	InitJSON("config_test.json", false, settings)
}
func TestNewJSON(*testing.T) {
	settings := Settings{}
	settings.Name = "testname"
	NewJSON("config_test.json", false, settings)
}

func TestLoadJSON(t *testing.T) {
	settings := Settings{}
	LoadJSON("config_test.json", &settings)
	assert.Equal(t, "testname", settings.Name)
}
