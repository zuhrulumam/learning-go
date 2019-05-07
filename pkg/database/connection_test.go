package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection_PingOK(t *testing.T) {
	dbURL := os.Getenv("LEARNING_GO_DB_URL")
	err := ConfigureTest(dbURL)
	defer Quit()
	assert.Nil(t, err, "they should err nil when try to connect")
}
