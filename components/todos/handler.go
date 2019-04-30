package todos

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

	e.GET("/", GetTodos)
	e.GET("/todos/:id", GetTodo)
}

// // // // // // // // // // // // // // // // // // // // // //

/*
	!: Add new handlers here
*/

// GetTodos returns all the todos
func GetTodos(c echo.Context) error {
	log.Println(`GET("/", GetTodos)`)

	var todo Todo

	todos, err := todo.findAll()
	if err != nil {
		errMsg := "error obtaining todos"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, todos)
}

// GetTodo returns a todo found by id
func GetTodo(c echo.Context) error {
	log.Println(`GET("/todos/:id", GetTodos)`)
	log.Println("Params -> ", c.ParamNames(), c.ParamValues())
	todoID := c.Param("id")

	id, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil {
		log.Println(err)
		errMsg := "error parsing id"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	var todo Todo

	res, err := todo.findByID(id)
	if err != nil {
		errMsg := "error obtaining todo"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, res)
}
