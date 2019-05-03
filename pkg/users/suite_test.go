package users

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zuhrulumam/learning-go/pkg/helpers/testhelper"
)

type UsersTest struct {
	testhelper.Suite
}

func TestUsers_Suite(t *testing.T) {
	s := &UsersTest{}
	s.NeedDB = true
	suite.Run(t, s)
}
