package dispatcher

import (
	/*
		!: Add new components to be dispatched here
	*/
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"sharkies/components/sharks"
	"sharkies/config"
	"sharkies/db"
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
