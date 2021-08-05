package models

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"name"`
	Location string `db:"location"`
}

type UserModel interface {
	UserSave(*User) error
	UserList() ([]User, error)
	UserListEmailLike(eml string) ([]User, error)
}


