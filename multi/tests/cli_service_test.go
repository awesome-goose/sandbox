package tests

import (
	"testing"

	test "github.com/awesome-goose/goose/testing"
	cliapp "github.com/awesome-goose/multi/app/cli"
)

func TestCliService(t *testing.T) {
	test.NewSuiteRunner(t, &CliServiceSuite{}).Run()
}

type CliServiceSuite struct {
	test.Suite
	service *cliapp.CliService
}

func (s *CliServiceSuite) SetupTest() {
	s.service = &cliapp.CliService{}
}

func (s *CliServiceSuite) TestService_NotNil() {
	s.T.Expect(s.service).Not().ToBeNil()
}

func (s *CliServiceSuite) TestGetVersion_ReturnsVersion() {
	result := s.service.GetVersion()
	s.T.Expect(result).ToEqual("0.0.0")
}

func (s *CliServiceSuite) TestGetVersion_NotEmpty() {
	result := s.service.GetVersion()
	s.T.Expect(result).Not().ToBeEmpty()
}
