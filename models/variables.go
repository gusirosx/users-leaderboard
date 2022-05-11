package models

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

type Database struct {
	Client *redis.Client
}
