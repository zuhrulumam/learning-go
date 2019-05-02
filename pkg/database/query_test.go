package database

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestQuery_QueryOK(t *testing.T) {
	dbURL := os.Getenv("LEARNING_GO_DB_URL")
	err := ConfigureTest(dbURL)

	_, err = Query(`select 1`)
	assert.Nil(t, err, "it should nil on error")
}
