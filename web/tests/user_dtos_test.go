package tests

import (
	"testing"

	test "github.com/awesome-goose/goose/testing"
	"github.com/awesome-goose/web/app/user"
)

func TestUserDtos(t *testing.T) {
	test.NewSuiteRunner(t, &UserDtosSuite{}).Run()
}

type UserDtosSuite struct {
	test.Suite
}

func (s *UserDtosSuite) TestGetDto_Initialization() {
	dto := &user.GetDto{Type: "detailed"}
	s.T.Expect(dto.Type).ToEqual("detailed")
}

func (s *UserDtosSuite) TestGetDto_EmptyValue() {
	dto := &user.GetDto{}
	s.T.Expect(dto.Type).ToEqual("")
}

func (s *UserDtosSuite) TestGetDto_NotNil() {
	dto := &user.GetDto{}
	s.T.Expect(dto).Not().ToBeNil()
}

func (s *UserDtosSuite) TestListDto_Initialization() {
	dto := &user.ListDto{Type: "paginated"}
	s.T.Expect(dto.Type).ToEqual("paginated")
}

func (s *UserDtosSuite) TestListDto_EmptyValue() {
	dto := &user.ListDto{}
	s.T.Expect(dto.Type).ToEqual("")
}

func (s *UserDtosSuite) TestListDto_NotNil() {
	dto := &user.ListDto{}
	s.T.Expect(dto).Not().ToBeNil()
}
