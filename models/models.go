package models

type Models struct {
	UserModel
}

func New(um UserModel) (Models, error) {
	return Models{
		um,
	}, nil
}
