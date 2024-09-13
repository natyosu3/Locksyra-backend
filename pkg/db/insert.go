package db

import (
	"Locksyra/pkg/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertDocument(user model.User) error {
	cl := Connect()

	coll := cl.Database("locksyra").Collection("user")

	_, err := coll.InsertOne(context.TODO(), bson.M{
		"username":        user.Username,
		"hashed_password": user.HashedPassword,
	})
	if err != nil {
		return err
	}

	return nil

}
