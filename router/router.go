package router

import (
	/*
		!: Add new components here
	*/
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"sharkies/db"
	"sharkies/src/api/sharks"
)

// ----------------------------------------------------------------------

// Init registers all the routes and env variables
func Init() *echo.Echo {
	db.Init()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	// Allow from any origin with these methods
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	/*
		!: Add new controllers here
	*/

	sharks.Init(e)

	return e
}
