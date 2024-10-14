package service

type Repository interface {
	GetName() string
}

type UserService struct {
	r Repository
}

func NewUserService(r Repository) *UserService {
	return &UserService{r: r}
}

func (u *UserService) GetName() string {
	return u.r.GetName()
}
