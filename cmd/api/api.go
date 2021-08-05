package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/wshaman/course-mongo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mongoImpl "github.com/wshaman/course-mongo/models/implementations/mongo"
	"github.com/wshaman/course-mongo/models/implementations/pg"
)

func withPg() models.UserModel {
	uri := "postgres://postgres:pwd123@localhost:15432/course_db?sslmode=disable"
	t, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to open DB"))
	}
	c := pg.NewUser(t)
	return c
}

func withMongo() models.UserModel {
	uri := "mongodb://localhost:27017/corp?retryWrites=true&w=majority\n"
	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = c.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	um, err := mongoImpl.NewUser(c)

	if err != nil {
		log.Fatal(err)
	}
	return um
}

func main() {
	var um models.UserModel
	switch true {
	case os.Getenv("DB") == "mongo":
		um = withMongo()
	case os.Getenv("DB") == "pg":
		um = withPg()
	default:
		log.Fatal("no DB chosen")
	}
	m, err := models.New(
		um,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.UserList())
}
