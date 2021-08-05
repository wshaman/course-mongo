package mongo

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/wshaman/course-mongo/models"
)

type user struct {
	db *mongo.Database
}

func  (u *user) UserList() ([]models.User, error) {
	ctx := context.Background()
	c, err := u.db.Collection("employee").Find(ctx, bson.M{"email": bson.M{"$regex": ".*party.com"}})
	if err != nil {
		return nil, errors.Wrap(err, "failed to UserList")
	}
	var users []bson.M
	if err = c.All(ctx, &users); err != nil {
		return nil, errors.Wrap(err, "failed to UserList(unwrap)")
	}
	fmt.Println(users)
	res := make([]models.User, 0, len(users))
	for _, v := range users {
		usr := models.User{
			Name: v["userName"].(string),
			Email: v["email"].(string),
		}
		res = append(res, usr)
	}
	return res, nil
}

func (u *user) UserListEmailLike(eml string) ([]models.User, error) {
	return nil, nil
}

func  (u *user) UserSave(model *models.User) error {
	if model.ID == 0 {
		return u.insertUser(model)
	}
	return u.updateUser(model)
}

func (u *user) insertUser(model *models.User) error {

	return nil
}

func (u *user) updateUser(model *models.User) error {

	return nil
}


func NewUser(c *mongo.Client) (models.UserModel, error) {
	return &user{
		db: c.Database("corp"),
	}, nil
}