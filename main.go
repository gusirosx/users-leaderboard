package main

import "rediboard/routes"

var (
	ListenAddr = "localhost:8080"
	RedisAddr  = "localhost:6379"
)

func main() {
	// Setup GinGonic Routes
	routes.RoutesSetup()
}
