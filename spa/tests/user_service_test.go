package tests

import (
	"testing"

	test "github.com/awesome-goose/goose/testing"
	"github.com/awesome-goose/spa/app/user"
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

func (s *UserServiceSuite) TestList_ReturnsSeededUsers() {
	users := s.service.List()
	s.T.Expect(len(users)).ToEqual(3)
}

func (s *UserServiceSuite) TestGet_ExistingUser() {
	found, ok := s.service.Get(1)
	s.T.Expect(ok).ToBeTrue()
	s.T.Expect(found.Name).ToEqual("Grace Hopper")
}

func (s *UserServiceSuite) TestGet_MissingUser() {
	_, ok := s.service.Get(999)
	s.T.Expect(ok).ToBeFalse()
}

func (s *UserServiceSuite) TestCreate_AppendsUser() {
	created := s.service.Create("Katherine Johnson", "katherine@example.com")

	s.T.Expect(created.ID).ToEqual(4)
	s.T.Expect(created.Name).ToEqual("Katherine Johnson")

	users := s.service.List()
	s.T.Expect(len(users)).ToEqual(4)
}

func (s *UserServiceSuite) TestCreate_IncrementsIDs() {
	first := s.service.Create("A", "a@example.com")
	second := s.service.Create("B", "b@example.com")

	s.T.Expect(second.ID).ToEqual(first.ID + 1)
}
