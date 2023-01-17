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

func SetScore(user User) error {
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

//The code above creates a User struct to serve as a wrapper around users in the leaderboard. The struct includes how we want the fields to be represented when transformed to JSON as well as when they are translated from HTTP requests using Gin’s binding. It then leverages pipelines to add the new member to the sorted set and gets the member’s new rank. Because the user parameter is a pointer, the Rank value is updated across the board when we mutate it from the SaveUser() function.

//Here, we are also leveraging pipelines to get the user’s score and rank, with their username as a key.

//We also signal to the caller if no matching record was found (using ErrNil) so that it is up to the caller to handle such cases separately (for instance, they could choose to display a 404 response).

//Adds all the specified members with the specified scores to the sorted set stored at key. It is possible to specify multiple score / member pairs. If a specified member is already a member of the sorted set, the score is updated and the element reinserted at the right position to ensure the correct ordering.

//If key does not exist, a new sorted set with the specified members as sole members is created, like if the sorted set was empty. If the key exists but does not hold a sorted set, an error is returned.
