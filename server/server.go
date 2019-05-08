package server

import (
	"github.com/agildav/go-boilerplate/dispatcher"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

// // // // // // // // // // // // // // // // // // // // // // //

// Init initializes the dispatcher
func Init() (*echo.Echo, map[string]string) {
	init := time.Now()
	log.Println(":: Server init")

	e, env := dispatcher.Init()

	log.Println(":: Server ready, took", time.Since(init))
	return e, env
}

// // // // // // // // // // // // // // // // // // // // // //

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
