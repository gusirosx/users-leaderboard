package models

import (
	"context"
	"errors"
	"rediboard/database"

	"github.com/go-redis/redis/v8"
)

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

// Create an unexported global variable to hold the database connection pool.
var redisDB *redis.Client = database.RedisInstance()

//https://blog.logrocket.com/how-to-use-redis-as-a-database-with-go-redis/
