package main

import (
	"github.com/emersonlugo/users-api/internal/handler"
	"github.com/emersonlugo/users-api/internal/server"
	"github.com/emersonlugo/users-api/internal/users"
)

func main() {
	// Create User Repository from In Memory Repository
	userRepository := users.NewInMemoryUserRepository()
	// Create User Service and Inject User Repository Dependency
	userService := users.NewUserService(userRepository)
	// Create HTTP Handler
	httpRequestHandler := handler.NewHandler(userService)
	// Start HTTP Server
	server.StartHTTPServer("8080", httpRequestHandler)
}