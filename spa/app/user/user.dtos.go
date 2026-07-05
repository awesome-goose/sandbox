package user

type ListDto struct{}

type GetDto struct {
	ID string `param:"id"`
}

type CreateDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
