package tests

import (
	"testing"

	"github.com/awesome-goose/cli/app/user"
	test "github.com/awesome-goose/goose/testing"
)

func TestUserService(t *testing.T) {
	test.NewSuiteRunner(t, &UserServiceSuite{}).Run()
}

type UserServiceSuite struct {
	test.Suite
	service *user.UserService
}

func (s *UserServiceSuite) SetupTest() {
	s.service = &user.UserService{}
}

func (s *UserServiceSuite) TestCreate_ReturnsUser() {
	result := s.service.Create()
	s.T.Expect(result).Not().ToBeNil()
}

func (s *UserServiceSuite) TestCreate_ReturnsAwesomeGoose() {
	result := s.service.Create()
	s.T.Expect(result.Name).ToEqual("Awesome Goose")
}

func (s *UserServiceSuite) TestList_ReturnsUser() {
	result := s.service.List()
	s.T.Expect(result).Not().ToBeNil()
}

func (s *UserServiceSuite) TestList_ReturnsAwesomeGoose() {
	result := s.service.List()
	s.T.Expect(result.Name).ToEqual("Awesome Goose")
}

func (s *UserServiceSuite) TestGet_ReturnsUser() {
	result := s.service.Get()
	s.T.Expect(result).Not().ToBeNil()
}

func (s *UserServiceSuite) TestGet_ReturnsAwesomeGoose() {
	result := s.service.Get()
	s.T.Expect(result.Name).ToEqual("Awesome Goose")
}

func (s *UserServiceSuite) TestUpdate_ReturnsUser() {
	result := s.service.Update()
	s.T.Expect(result).Not().ToBeNil()
}

func (s *UserServiceSuite) TestUpdate_ReturnsAwesomeGoose() {
	result := s.service.Update()
	s.T.Expect(result.Name).ToEqual("Awesome Goose")
}

func (s *UserServiceSuite) TestDelete_ReturnsUser() {
	result := s.service.Delete()
	s.T.Expect(result).Not().ToBeNil()
}

func (s *UserServiceSuite) TestDelete_ReturnsAwesomeGoose() {
	result := s.service.Delete()
	s.T.Expect(result.Name).ToEqual("Awesome Goose")
}
