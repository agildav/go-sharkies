package sharks

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

// ----------------------------------------------------------------------

// Init registers the routes
func Init(e *echo.Echo) {
	/*
		!: Add new routes here
	*/

	e.GET("/", getSharks)
	e.GET("/sharks", getSharks)
	e.GET("/sharks/:id", getShark)
	e.POST("/sharks", addShark)
	e.DELETE("/sharks/:id", deleteShark)
	e.DELETE("/sharks", deleteSharks)
	e.PATCH("/sharks/:id", patchShark)
}

// ----------------------------------------------------------------------

/*
	!: Add new handlers here
*/

// getSharks returns all the sharks
func getSharks(c echo.Context) error {
	log.Println(`GET("/sharks", getSharks)`)

	shark := new(Shark)
	sharks, err := shark.findAll()
	if err != nil {
		errMsg := "error obtaining sharks"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, sharks)
}

// getShark returns a shark found by id
func getShark(c echo.Context) error {
	log.Println(`GET("/sharks/:id", getShark)`)
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
	log.Printf("Body -> %+v", *shark)

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

// patchShark edits a shark
func patchShark(c echo.Context) error {
	log.Println(`PATCH("/sharks/:id", patchShark)`)
	log.Println("Params -> ", c.ParamNames(), c.ParamValues())

	sharkID := c.Param("id")
	id, err := strconv.ParseInt(sharkID, 10, 64)
	if err != nil {
		log.Println(err)
		errMsg := "error parsing id"
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	shark := new(Shark)
	if err = c.Bind(shark); err != nil {
		log.Println(err)
		errMsg := "error binding shark"
		return c.JSON(http.StatusBadRequest, errMsg)
	}
	log.Printf("Body -> %+v", *shark)

	res, err := shark.patchShark(id, shark)
	if err != nil {
		errMsg := "error patching shark"
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	return c.JSON(http.StatusOK, res)
}
