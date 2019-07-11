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

	log.Println(":: Server, starting")

	e, env := router.Init()

	log.Println(":: App env,", env["APP_ENV"])
	return e, env
}

// ----------------------------------------------------------------------

// Start runs the server
func Start() {
	init := time.Now()
	e, env := Init()
	var port = env["PORT"]

	e.HideBanner = true

	s := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  45 * time.Second,
		WriteTimeout: 45 * time.Second,
	}

	log.Println(":: Server, ready in", time.Since(init))
	err := e.StartServer(s)
	if err != nil {
		log.Println(":: Server, FAIL in", time.Since(init))
		e.Logger.Fatal(err)
	}

}
