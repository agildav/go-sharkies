package dispatcher

import (
	/*
		!: Add new components to be dispatched here
	*/
	"github.com/agildav/go-boilerplate/components/todos"
	"github.com/agildav/go-boilerplate/config"
	"github.com/agildav/go-boilerplate/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// // // // // // // // // // // // // // // // // // // // // //

// Init registers all the routes and env variables
func Init() (*echo.Echo, map[string]string) {
	env := config.Init()
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	db.Init(env)

	/*
		!: Add new components initializations here
	*/
	todos.Init(e)

	return e, env
}
