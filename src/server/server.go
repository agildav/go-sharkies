package server

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"sharkies/router"
	"time"
)

// ----------------------------------------------------------------------

// Init initializes the router
func Init() (*echo.Echo, map[string]string) {
	init := time.Now()
	log.Println(":: Server init")

	e, env := router.Init()

	log.Println(":: Server ready, took", time.Since(init))
	log.Println(":: App env -> ", env["APP_ENV"])
	return e, env
}

// ----------------------------------------------------------------------

// Start runs the server
func Start() {
	e, env := Init()
	var port = env["PORT"]

	e.HideBanner = true

	s := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  45 * time.Second,
		WriteTimeout: 45 * time.Second,
	}
	log.Println(":: Server listening on", port)
	e.Logger.Fatal(e.StartServer(s))
}
