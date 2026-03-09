package tests

import (
	"testing"

	test "github.com/awesome-goose/goose/testing"
	"github.com/awesome-goose/web/app/user"
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

func (s *UserEntitySuite) TestUserEntity_NotNil() {
	ue := &user.UserEntity{}
	s.T.Expect(ue).Not().ToBeNil()
}
