package testhelper

import (
	"os"

	"github.com/stretchr/testify/suite"
	"github.com/zuhrulumam/learning-go/pkg/database"
)

// Suite model
type Suite struct {
	suite.Suite
	suite.SetupAllSuite
	suite.TearDownAllSuite

	NeedDB bool
}

// SetupSuite func
func (s *Suite) SetupSuite() {
	if s.NeedDB {
		s.setupDB()
	}
}

func (s *Suite) setupDB() {
	dbURL := os.Getenv("LEARNING_GO_DB_URL")
	err := database.ConfigureTest(dbURL)

	s.Nil(err)
}

// TearDownSuite tear down suite
func (s *Suite) TearDownSuite() {
	// if s.NeedDB {
	// 	defer database.Quit()
	// }
}
