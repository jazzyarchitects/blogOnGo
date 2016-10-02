package main

import (
	_ "blogOnGo/src/models"
	"net/http"
	_ "blogOnGo/src/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesList = Routes{
	Route{
		"GetAllUsers",
		"GET",
		"/user",
		"controllers.UserController.GetUsers",
	},
	Route{
		"GetUserById",
		"GET",
		"/user/{userId}",
		"controllers.UserController.GetUserById",
	},
}