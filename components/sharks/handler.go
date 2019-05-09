package sharks

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

// // // // // // // // // // // // // // // // // // // // // //

// Init registers the routes
func Init(e *echo.Echo) {
	/*
		!: Add new routes here
	*/

	e.GET("/", GetSharks)
	e.GET("/sharks", GetSharks)
	e.GET("/sharks/:id", GetShark)
	e.POST("/sharks", addShark)
	e.DELETE("/sharks/:id", deleteShark)
	e.DELETE("/sharks", deleteSharks)
}

// // // // // // // // // // // // // // // // // // // // // //

/*
	!: Add new handlers here
*/

// GetSharks returns all the sharks
func GetSharks(c echo.Context) error {
	log.Println(`GET("/sharks", GetSharks)`)

	shark := new(Shark)
	sharks, err := shark.findAll()
	if err != nil {
		errMsg := "error obtaining sharks"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, sharks)
}

// GetShark returns a shark found by id
func GetShark(c echo.Context) error {
	log.Println(`GET("/sharks/:id", GetShark)`)
	log.Println("Params -> ", c.ParamNames(), c.ParamValues())

	sharkID := c.Param("id")
	id, err := strconv.ParseInt(sharkID, 10, 64)
	if err != nil {
		log.Println(err)
		errMsg := "error parsing id"
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	shark := new(Shark)
	res, err := shark.findByID(id)
	if err != nil {
		errMsg := "error obtaining shark"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, res)
}

// addShark inserts a shark
func addShark(c echo.Context) error {
	log.Println(`POST("/sharks", addShark)`)

	shark := new(Shark)
	if err := c.Bind(shark); err != nil {
		log.Println(err)
		errMsg := "error binding shark"
		return c.JSON(http.StatusBadRequest, errMsg)
	}
	log.Printf("Params -> %+v", *shark)

	res, err := shark.addShark(shark)
	if err != nil {
		errMsg := "error inserting shark"
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	return c.JSON(http.StatusCreated, res)
}

// deleteShark removes a shark by id
func deleteShark(c echo.Context) error {
	log.Println(`DELETE("/sharks", deleteShark)`)
	log.Println("Params -> ", c.ParamNames(), c.ParamValues())

	sharkID := c.Param("id")
	id, err := strconv.ParseInt(sharkID, 10, 64)
	if err != nil {
		log.Println(err)
		errMsg := "error parsing id"
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	shark := new(Shark)
	res, err := shark.deleteShark(id)
	if err != nil {
		errMsg := "error deleting shark"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, res)
}

// deleteSharks removes all the sharks
func deleteSharks(c echo.Context) error {
	log.Println(`DELETE("/sharks", deleteSharks)`)

	shark := new(Shark)
	res, err := shark.deleteAll()
	if err != nil {
		errMsg := "error deleting sharks"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, res)
}
