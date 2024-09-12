package db

import "log/slog"

func CreateInitCollection() {
	client := Connect()

	db := client.Database("locksyra")
	db.Collection("user")

	slog.Info("db and collection create successfull.")
}
