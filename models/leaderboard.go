package models

var leaderboardKey = "leaderboard"

type Leaderboard struct {
	Count int `json:"count"`
	Users []*User
}

func GetLeaderboard() (*Leaderboard, error) {
	scores := redisDB.ZRangeWithScores(Ctx, leaderboardKey, 0, -1)
	if scores == nil {
		return nil, ErrNil
	}
	count := len(scores.Val())
	users := make([]*User, count)
	for idx, member := range scores.Val() {
		users[idx] = &User{
			Username: member.Member.(string),
			Points:   int(member.Score),
			Rank:     idx,
		}
	}
	leaderboard := &Leaderboard{
		Count: count,
		Users: users,
	}
	return leaderboard, nil
}

//To get the complete leaderboard, Redis provides the ZRange command, used to retrieve the members of a sorted set in ascending order of their scores. ZRange also accepts an optional WITHSCORES argument that tells it to return the score of each member as well. Go-redis on the other hand, splits the command into two, providing ZRange and ZRangeWithScores separately.

//The leaderboardKey represents the key used to identify the set in our Redis database. Since we are only running a single command now (ZRangeWithScores), there is no longer a need to batch the commands with transaction pipelines anymore so we store the result directly in the scores variable. The value stored in scores contains a slice of Go maps, whose length is the number of members stored in the set.
