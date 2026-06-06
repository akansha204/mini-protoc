package user

type UserRequest struct {
	Name string
	Age int32
}

type UserResponse struct {
	Name string
	Age int32
}

type UserService interface {
	GetUser(req UserRequest) (UserResponse, error)
	CreateUser(req UserRequest) (UserResponse, error)
}

