package tests

import (
	"testing"

	test "github.com/awesome-goose/goose/testing"
	apiapp "github.com/awesome-goose/multi/app/api"
	cliapp "github.com/awesome-goose/multi/app/cli"
)

func TestControllerInputs(t *testing.T) {
	test.NewSuiteRunner(t, &ControllerInputsSuite{}).Run()
}

type ControllerInputsSuite struct {
	test.Suite
}

func (s *ControllerInputsSuite) TestApiStatusInput_NotNil() {
	input := &apiapp.ApiStatusInput{}
	s.T.Expect(input).Not().ToBeNil()
}

func (s *ControllerInputsSuite) TestApiHealthInput_NotNil() {
	input := &apiapp.ApiHealthInput{}
	s.T.Expect(input).Not().ToBeNil()
}

func (s *ControllerInputsSuite) TestCliInfoInput_NotNil() {
	input := &cliapp.InfoInput{}
	s.T.Expect(input).Not().ToBeNil()
}

func (s *ControllerInputsSuite) TestCliHelpInput_NotNil() {
	input := &cliapp.HelpInput{}
	s.T.Expect(input).Not().ToBeNil()
}

func (s *ControllerInputsSuite) TestCliVersionInput_NotNil() {
	input := &cliapp.VersionInput{}
	s.T.Expect(input).Not().ToBeNil()
}
