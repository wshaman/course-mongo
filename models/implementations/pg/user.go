package pg

import (
	"database/sql"
	"github.com/wshaman/course-mongo/utils"

	"github.com/pkg/errors"

	"github.com/wshaman/course-mongo/models"
)

type user struct {
	db *sql.DB
}

func  (u *user) UserList() ([]models.User, error) {
	rows, err := u.db.Query(`select u.id, u.name, email, l.name as location from users u 
left join locations l on u.location_id=u.id`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ListUsers")
	}
	return u.rowsToUsers(rows)
}

func (u *user) UserListEmailLike(eml string) ([]models.User, error) {
	eml = "%" + eml
	rows, err := u.db.Query(`select u.id, u.name, email, l.name as location from users u 
left join locations l on u.location_id=u.id where email LIKE $1`, eml)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ListUsers")
	}
	return u.rowsToUsers(rows)
}

func  (u *user) UserSave(model *models.User) error {
	if model.ID == 0 {
		return u.insertUser(model)
	}
	return u.updateUser(model)
}

//func insertUser(db *sql.DB, u *User) (err error) {
func (u *user) insertUser(model *models.User) error {
	var id int64
	q := "insert into users (name, email) values ($1, $2) returning id"
	if err := u.db.QueryRow(q, model.Name, model.Email).Scan(&id); err != nil {
		return errors.Wrap(err, "failed to insert user")
	}
	model.ID = int(id)
	return nil
}

func (u *user) updateUser(model *models.User) error {
	q := "update users set  name=$1, email=$2 where id=$3;"
	if _, err := u.db.Exec(q, model.Name, model.Email, model.ID); err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (u *user) rowsToUsers(rows *sql.Rows) (users []models.User, err error) {
	users = make([]models.User, 0)
	for rows.Next() {
		u := &models.User{}
		var s *string
		if err = rows.Scan(&u.ID, &u.Name, &u.Email, &s); err != nil {
			return nil, errors.Wrap(err, "failed to scan users (scan)")
		}
		u.Location = utils.StrPtr2Str(s)
		users = append(users, *u)
	}
	return users, nil
}


func NewUser(db *sql.DB) models.UserModel{
	return &user{
		db:db,
	}
}