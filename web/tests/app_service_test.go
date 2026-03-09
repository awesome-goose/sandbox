package tests

import (
	"testing"

	test "github.com/awesome-goose/goose/testing"
	"github.com/awesome-goose/web/app"
)

func TestAppService(t *testing.T) {
	test.NewSuiteRunner(t, &AppServiceSuite{}).Run()
}

type AppServiceSuite struct {
	test.Suite
	service *app.AppService
}

func (s *AppServiceSuite) SetupTest() {
	s.service = &app.AppService{}
}

func (s *AppServiceSuite) TestHealth_ReturnsHealthyMessage() {
	result := s.service.Health()
	s.T.Expect(result).ToEqual("Goose is healthy")
}

func (s *AppServiceSuite) TestHealth_NotEmpty() {
	result := s.service.Health()
	s.T.Expect(result).Not().ToBeEmpty()
}

func (s *AppServiceSuite) TestHealth_ContainsGoose() {
	result := s.service.Health()
	s.T.Expect(result).ToContainString("Goose")
}

func (s *AppServiceSuite) TestHealth_ContainsHealthy() {
	result := s.service.Health()
	s.T.Expect(result).ToContainString("healthy")
}
