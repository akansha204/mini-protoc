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

type UserServiceClient struct {
}

func NewUserServiceClient() *UserServiceClient {
	return &UserServiceClient{}
}

func (c *UserServiceClient) GetUser(req UserRequest) (UserResponse, error) {
	panic("not implemented")
}

func (c *UserServiceClient) CreateUser(req UserRequest) (UserResponse, error) {
	panic("not implemented")
}

func RegisterUserService(service UserService) {
	panic("not implemented")
}

