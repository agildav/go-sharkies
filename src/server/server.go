package server

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
	"sharkies/router"
	"time"
)

// ----------------------------------------------------------------------

// Init initializes the router
func Init() *echo.Echo {

	log.Println(":: Server, starting")

	e := router.Init()

	log.Println(":: App env,", os.Getenv("APP_ENV"))
	return e
}

// ----------------------------------------------------------------------

// Start runs the server
func Start() {
	init := time.Now()
	e := Init()
	var port = os.Getenv("PORT")

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
