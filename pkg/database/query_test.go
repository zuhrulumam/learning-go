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

	_, err = Query(`select 1`)
	assert.Nil(t, err, "it should nil on error")
}

func TestQuery_ExecOK(t *testing.T) {
	dbURL := os.Getenv("LEARNING_GO_DB_URL")
	err := ConfigureTest(dbURL)
	defer Quit()
	_, err = Exec(`select 1`) // it should do write operation
	assert.Nil(t, err, "it should nil on error")
}

func TestQuery_QueryGormOK(t *testing.T) {
	dbURL := os.Getenv("LEARNING_GO_DB_URL")
	err := ConfigureGormTest(dbURL)
	defer Quit()

	err = QueryGorm(nil, "select 1")
	assert.Nil(t, err, "it should nil on error")
}

func TestQuery_ExecGormOK(t *testing.T) {
	dbURL := os.Getenv("LEARNING_GO_DB_URL")
	err := ConfigureGormTest(dbURL)
	defer Quit()
	err = ExecGorm(nil, "select 1") // it should do write operation
	assert.Nil(t, err, "it should nil on error")
}
