package dispatcher

import (
	/*
		!: Add new components to be dispatched here
	*/
	"github.com/agildav/sharkies/components/sharks"
	"github.com/agildav/sharkies/config"
	"github.com/agildav/sharkies/db"
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
	sharks.Init(e)

	return e, env
}
