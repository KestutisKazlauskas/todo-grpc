package config

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigIsLoaded(t *testing.T) {
	conf, _ := Load("config.yml")

	assert.EqualValues(t, "grpc", conf.Server.Type)
	assert.EqualValues(t, "", conf.Server.Host)
	assert.EqualValues(t, uint(8000), conf.Server.Port)
	assert.EqualValues(t, "mongo", conf.Database.Type)
	assert.EqualValues(t, "", conf.Database.Host)
	assert.EqualValues(t, uint(27017), conf.Database.Port)
}

func TestNoConfigFile(t *testing.T) {
	_, err := Load("not-exists")

	assert.EqualValues(t, errors.New("file does not exists"), err)
}
