package tests

import (
	"testing"

	test "github.com/awesome-goose/goose/testing"
	"github.com/awesome-goose/spa/app"
	"github.com/awesome-goose/spa/app/user"
)

func TestAppDtos(t *testing.T) {
	test.NewSuiteRunner(t, &AppDtosSuite{}).Run()
}

type AppDtosSuite struct {
	test.Suite
}

func (s *AppDtosSuite) TestHealthDto_Initialization() {
	dto := &app.HealthDto{Type: "full"}
	s.T.Expect(dto.Type).ToEqual("full")
}

func (s *AppDtosSuite) TestVersionDto_NotNil() {
	dto := &app.VersionDto{}
	s.T.Expect(dto).Not().ToBeNil()
}

func (s *AppDtosSuite) TestGetDto_Initialization() {
	dto := &user.GetDto{ID: "42"}
	s.T.Expect(dto.ID).ToEqual("42")
}

func (s *AppDtosSuite) TestCreateDto_Initialization() {
	dto := &user.CreateDto{Name: "Grace", Email: "grace@example.com"}
	s.T.Expect(dto.Name).ToEqual("Grace")
	s.T.Expect(dto.Email).ToEqual("grace@example.com")
}
