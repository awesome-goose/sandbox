package tests

import (
	"testing"

	"github.com/awesome-goose/cli/app/user"
	test "github.com/awesome-goose/goose/testing"
)

func TestUserEntity(t *testing.T) {
	test.NewSuiteRunner(t, &UserEntitySuite{}).Run()
}

type UserEntitySuite struct {
	test.Suite
}

func (s *UserEntitySuite) TestUser_Initialization() {
	u := &user.User{Name: "Test User"}
	s.T.Expect(u.Name).ToEqual("Test User")
}

func (s *UserEntitySuite) TestUser_EmptyName() {
	u := &user.User{}
	s.T.Expect(u.Name).ToEqual("")
}

func (s *UserEntitySuite) TestUser_NotNil() {
	u := &user.User{}
	s.T.Expect(u).Not().ToBeNil()
}

func (s *UserEntitySuite) TestUser_NameNotEmpty() {
	u := &user.User{Name: "John Doe"}
	s.T.Expect(u.Name).Not().ToBeEmpty()
}
