package user

type Repository interface {
	Save(user User) (User, error)
}
