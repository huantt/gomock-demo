package user

//go:generate go run go.uber.org/mock/mockgen@latest -source=interface.go -destination=mock.go -package=user

type IUserRepo interface {
	GetUserByID(id int) (*User, error)
	Insert(user User) error
	Update(id int, user User) error
}
