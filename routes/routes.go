package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

func RoutesSetup() {
	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	// Set up a http server
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)

	// Run the http server
	if err := router.Run(port); err != nil {
		log.Fatalln("could not run server: ", err.Error())
	} else {
		log.Println("Server listening on port: ", port)
	}
}

func initializeRoutes(router *gin.Engine) {
	// Handle the index route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "Up and running..."})
	})
	// Handle the no route case
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})
	//Group user related routes together
	//userRoutes := router.Group("/users")
	//UserRoutes(userRoutes)
}

// func UserRoutes(routes *gin.RouterGroup) {
// 	// Read users at /users
// 	routes.GET("", handlers.GetUsers)
// 	// Read users at /users/ID
// 	routes.GET("/:id", handlers.GetUser)
// 	// Create user at /users
// 	routes.POST("", handlers.CreateUser)
// 	// Update users at /users
// 	routes.PUT("/:id", handlers.UpdateUser)
// 	// Delete users at /users
// 	routes.DELETE("/:id", handlers.DeleteUser)
// }

// func main2() {
// 	database, err := db.NewDatabase(RedisAddr)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to redis: %s", err.Error())
// 	}

// 	router := initRouter(database)
// 	router.Run(ListenAddr)
// }

// func initRouter(database *db.Database) *gin.Engine {
// 	r := gin.Default()
// 	r.GET("/points/:username", func(c *gin.Context) {
// 		username := c.Param("username")
// 		user, err := database.GetUser(username)
// 		if err != nil {
// 			if err == db.ErrNil {
// 				c.JSON(http.StatusNotFound, gin.H{"error": "No record found for " + username})
// 				return
// 			}
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"user": user})
// 	})

// 	r.POST("/points", func(c *gin.Context) {
// 		var userJson db.User
// 		if err := c.ShouldBindJSON(&userJson); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		err := database.SaveUser(&userJson)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"user": userJson})
// 	})

// 	r.GET("/leaderboard", func(c *gin.Context) {
// 		leaderboard, err := database.GetLeaderboard()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"leaderboard": leaderboard})
// 	})

// 	return r
// }
