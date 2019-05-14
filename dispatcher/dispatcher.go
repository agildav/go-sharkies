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
	"net/http"
)

// // // // // // // // // // // // // // // // // // // // // //

// Init registers all the routes and env variables
func Init() (*echo.Echo, map[string]string) {
	env := config.Init()
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// Allow from any origin
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	db.Init(env)

	/*
		!: Add new components initializations here
	*/
	sharks.Init(e)

	return e, env
}
