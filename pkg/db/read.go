package db

import (
	"Locksyra/pkg/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// ユーザ名を指定してユーザ情報を取得
func ReadUser(username string) (model.User, error) {
	cl := Connect()
	coll := cl.Database("locksyra").Collection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user model.User
	if err := coll.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		log.Printf("Failed to read user: %v", err)
		return model.User{}, err
	}

	log.Println("ReadUser", "username", user)
	log.Println("ReadUser", "user.HashedPassword", user.HashedPassword)

	return user, nil
}

func ReadAllUsers() ([]model.User, error) {
	cl := Connect()
	coll := cl.Database("locksyra").Collection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Failed to find users: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []model.User
	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			log.Printf("Failed to decode user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return users, nil
}
