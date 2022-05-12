package handlers

import (
	"net/http"
	"rediboard/models"

	"github.com/gin-gonic/gin"
)

func SetScore(ctx *gin.Context) {
	var userJson models.User
	if err := ctx.ShouldBindJSON(&userJson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.SetScore(&userJson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": userJson})
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
