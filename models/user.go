package models

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Points   int    `json:"points" binding:"required"`
	Rank     int    `json:"rank"`
}

func SetScore(user *User) error {
	member := &redis.Z{
		Score:  float64(user.Points),
		Member: user.Username,
	}
	pipe := redisDB.TxPipeline()
	pipe.ZAdd(Ctx, "leaderboard", member)
	rank := pipe.ZRank(Ctx, leaderboardKey, user.Username)
	_, err := pipe.Exec(Ctx)
	if err != nil {
		return err
	}
	fmt.Println(rank.Val(), err)
	user.Rank = int(rank.Val())
	return nil
}

func GetScore(username string) (*User, error) {
	pipe := redisDB.TxPipeline()
	score := pipe.ZScore(Ctx, leaderboardKey, username)
	rank := pipe.ZRank(Ctx, leaderboardKey, username)
	_, err := pipe.Exec(Ctx)
	if err != nil {
		return nil, err
	}
	if score == nil {
		return nil, ErrNil
	}
	return &User{
		Username: username,
		Points:   int(score.Val()),
		Rank:     int(rank.Val()),
	}, nil
}
