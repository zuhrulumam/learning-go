package models

import (
	"github.com/zuhrulumam/learning-go/pkg/helpers/testhelper"
	"testing"
	"github.com/stretchr/testify/suite"
)

type ModelsTest struct {
	testhelper.Suite
}

func TestModels_Suite(t *testing.T){
	s := &ModelsTest{}
	s.NeedDB = true
	suite.Run(t, s)
}