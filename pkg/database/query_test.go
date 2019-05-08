package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_QueryOK(t *testing.T) {
	dbURL := os.Getenv("LEARNING_GO_DB_URL")
	err := ConfigureTest(dbURL)
	defer Quit()

	err = Query(nil, "select 1")
	assert.NotNil(t, err, "it should not nil on error")
}

func TestQuery_ExecOK(t *testing.T) {
	dbURL := os.Getenv("LEARNING_GO_DB_URL")
	err := ConfigureTest(dbURL)
	defer Quit()
	err = Exec("select 1") // it should do write operation
	assert.Nil(t, err, "it should nil on error")
}
