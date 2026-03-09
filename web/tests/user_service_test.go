package tests

import (
	"testing"

	test "github.com/awesome-goose/goose/testing"
	"github.com/awesome-goose/web/app/user"
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

func (s *UserServiceSuite) TestService_NotNil() {
	s.T.Expect(s.service).Not().ToBeNil()
}

func (s *UserServiceSuite) TestService_IsPointer() {
	s.T.Expect(s.service).Not().ToBeNil()
}
