package request

type UpdateUserRequest struct {
	ID    string `validate:"required,min=1,max=200" json:"id"`
	Name  string `validate:"required,min=1,max=200" json:"name"`
	Email string `validate:"required,min=1,max=200" json:"email"`
}
