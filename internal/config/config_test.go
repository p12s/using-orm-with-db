package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/p12s/using-orm-with-db/internal/config"
	"github.com/stretchr/testify/assert"
)

const (
	OK_ENV_PATH    = "fixtures/.env.ok.example"
	EMPTY_ENV_PATH = "fixtures/.env.empty"
)

func TestNew(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.Equal(t, nil, err)

	// empty .env file should return err
	err = godotenv.Load(fmt.Sprintf("%s/%s", currentDir, EMPTY_ENV_PATH))
	assert.Equal(t, nil, err)

	_, err = config.New()
	assert.NotNil(t, err)

	// not empty .env file should not
	err = godotenv.Load(fmt.Sprintf("%s/%s", currentDir, OK_ENV_PATH))
	assert.Equal(t, nil, err)

	_, err = config.New()
	assert.Equal(t, nil, err)
}
