package handlers

import (
	"net/http"
	"rediboard/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

//https://blog.logrocket.com/how-to-use-redis-as-a-database-with-go-redis/

var database *redis.Client = database.RedisInstance()

func GetScore(c *gin.Context) {
	var userJson models.User
	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.SaveUser(&userJson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userJson})
}

func SetScore(c *gin.Context) {
	username := c.Param("username")
	user, err := database.GetUser(username)
	if err != nil {
		if err == models.ErrNil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No record found for " + username})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetLeaderboardfunc(c *gin.Context) {
	leaderboard, err := database.GetLeaderboard()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"leaderboard": leaderboard})
}
