package db

import (
	"context"
	"log/slog"
)

func InsertDocument(data any) error {
	cl := Connect()

	coll := cl.Database("locksyra").Collection("user")

	insertRes, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	slog.Info("挿入されたドキュメントのID: ", insertRes.InsertedID)

	return nil

}
