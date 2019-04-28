package todos

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// // // // // // // // // // // // // // // // // // // // // //

// Init registers the routes
func Init(e *echo.Echo) {
	/*
		!: Add new routes here
	*/

	e.GET("/", GetTodos)
}

// // // // // // // // // // // // // // // // // // // // // //

/*
	!: Add new handlers here
*/

// GetTodos returns all the todos
func GetTodos(c echo.Context) error {
	log.Println(`GET("/", GetTodos)`)

	var todo Todo

	todos, err := todo.searchTodos()
	if err != nil {
		errMsg := "error obtaining todos"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, todos)
}
