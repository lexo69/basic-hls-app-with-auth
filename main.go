package main

import (
	"hls/app"
)

// Create a test user

func init() {
	app.InitHTTP()
	app.InitVideoRoutes()
}

func main() {
	var TestUser = app.User{
		Name:     "Test User",
		Login:    "test",
		Password: "12345",
	}

	app.AppUsers.Users = append(app.AppUsers.Users, TestUser)

	run_err := app.HTTP.Listen(app.HOST)
	if run_err != nil {
		panic("error starting up HTTP server")
	}

}
