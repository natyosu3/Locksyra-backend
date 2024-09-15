package model

type User struct {
	ID             string `json:"id"`
	Username       string `json:"username" bson:"name"`
	HashedPassword string `json:"hashed_password" bson:"hashed_password"`
}
