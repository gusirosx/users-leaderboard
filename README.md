# users-leaderboard

Redis golang user leaderboard

rediboard


In this article, we will explore go-redis and use its pipeline feature to build a leaderboard API. The API will use Gin and Redis’ sorted sets under the hood. It will expose the following endpoints:

GET /points/:username — to get a user’s score and their rank in the overall leaderboard
POST /points — to add or update a user and their score. This endpoint will also return the new rank of the user
GET /leaderboard — returns the current leaderboard, with users sorted in ascending order of their ranks


Prerequisites
To follow along with this post, you will need:

A Go installation with modules support
Redis installed on your local computer (alternatively, you can use the Docker image if you have Docker installed)
Experience writing Go

Save users to the sorted sets

The code above creates a User struct to serve as a wrapper around users in the leaderboard. The struct includes how we want the fields to be represented when transformed to JSON as well as when they are translated from HTTP requests using Gin’s binding. It then leverages pipelines to add the new member to the sorted set and gets the member’s new rank. Because the user parameter is a pointer, the Rank value is updated across the board when we mutate it from the SaveUser() function.