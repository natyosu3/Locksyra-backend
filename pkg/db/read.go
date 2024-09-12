package db

import (
	"context"
	"log/slog"
	"time"
)

func ReadUser() {
	cl := Connect()
	coll := cl.Database("locksyra").Collection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	slog.Info(coll.Find())
}
