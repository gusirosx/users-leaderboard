package handlers

import (
	"net/http"
	"rediboard/models"

	"github.com/gin-gonic/gin"
)

func SetScore(ctx *gin.Context) {
	// create an empty user of type models.User
	var user models.User

	// decode the json request to user
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.SetScore(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func GetScore(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := models.GetScore(username)
	if err != nil {
		if err == models.ErrNil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No record found for " + username})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func GetLeaderboard(ctx *gin.Context) {
	leaderboard, err := models.GetLeaderboard()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"leaderboard": leaderboard})
}

//err = redis.Set(strconv.Itoa(v.Block_id)+":"+strconv.Itoa(v.Block_parent), json, 0).Err()

// ctx := context.TODO()
// client.Set(ctx, "language", "Go", 0)
// language := client.Get(ctx, "language")
// year := client.Get(ctx, "year")

// fmt.Println(language.Val()) // "Go"
// fmt.Println(year.Val()) // ""

// func GetSize(c *gin.Context) {
// 	redis := dbRedis.ConnectRedis()
// 	defer redis.Close()

// 	dbsize, err := redis.DbSize().Result()
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	c.JSON(200, gin.H{
// 		"size": dbsize,
// 	})

// }

// func GetValues(c *gin.Context) {
// 	redis := dbRedis.ConnectRedis()
// 	defer redis.Close()

// 	keys, _ := redis.Keys("*").Result()

// 	var values []models.BQData
// 	var blocks models.BQData
// 	for _, v := range keys {
// 		value, err := redis.Get(v).Result()
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		json.Unmarshal([]byte(value), &blocks)
// 		values = append(values, blocks)
// 	}

// 	c.IndentedJSON(http.StatusOK, values)
// }

// package main

// import (
// 	"context"
// 	"github.com/go-redis/redis/v8"
// )

// func main() {
// 	ctx := context.Background()

// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:	  "localhost:6379",
// 		Password: "", // no password set
// 		DB:		  0,  // use default DB
// 	})

// 	err := rdb.Set(ctx, "key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := rdb.Get(ctx, "key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := rdb.Get(ctx, "key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// 	// Output: key value
// 	// key2 does not exist
// }

//=========================================

// Using redisClient's Scan function, we can retrieve all key records containing the entered key. We create a string array and after each iteration, we append keys to this string array.

// package main

// import (
// 	"context"
// 	"fmt"
// 	"github.com/go-redis/redis"
// )

// var redisClient *redis.Client

// func main() {
// 	ctx := context.TODO()
// 	connectRedis(ctx)

// 	setToRedis(ctx, "name", "redis-test")
// 	setToRedis(ctx, "name2", "redis-test-2")
// 	val := getFromRedis(ctx,"name")

// 	fmt.Printf("First value with name key : %s \n", val)

// 	values := getAllKeys(ctx, "name*")

// 	fmt.Printf("All values : %v \n", values)

// }

// func connectRedis(ctx context.Context) {
// 	client := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 		Password: "",
// 		DB: 0,
// 	})

// 	pong, err := client.Ping(ctx).Result()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(pong)

// 	redisClient = client
// }

// // Send our command across the connection. The first parameter to
// // Do() is always the name of the Redis command (in this example
// // HMSET), optionally followed by any necessary arguments (in this
// // example the key, followed by the various hash fields and values).
// _, err = conn.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
// if err != nil {
// 	log.Fatal(err)
// }
