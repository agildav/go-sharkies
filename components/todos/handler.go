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
	e.GET("/todos", GetTodos)
	e.GET("/todos/:id", GetTodo)
	e.POST("/todos", AddTodo)
	e.DELETE("/todos/:id", DeleteTodo)
	e.DELETE("/todos", DeleteTodos)
}

// // // // // // // // // // // // // // // // // // // // // //

/*
	!: Add new handlers here
*/

// GetTodos returns all the todos
func GetTodos(c echo.Context) error {
	log.Println(`GET("/todos", GetTodos)`)

	todo := new(Todo)
	todos, err := todo.findAll()
	if err != nil {
		errMsg := "error obtaining todos"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, todos)
}

// GetTodo returns a todo found by id
func GetTodo(c echo.Context) error {
	log.Println(`GET("/todos/:id", GetTodo)`)
	log.Println("Params -> ", c.ParamNames(), c.ParamValues())

	todoID := c.Param("id")
	id, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil {
		log.Println(err)
		errMsg := "error parsing id"
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	todo := new(Todo)
	res, err := todo.findByID(id)
	if err != nil {
		errMsg := "error obtaining todo"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, res)
}

// AddTodo inserts a todo
func AddTodo(c echo.Context) error {
	log.Println(`POST("/todos", AddTodo)`)

	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		log.Println(err)
		errMsg := "error binding todo"
		return c.JSON(http.StatusBadRequest, errMsg)
	}
	log.Printf("Params -> %+v", *todo)

	res, err := todo.addTodo(todo)
	if err != nil {
		errMsg := "error inserting todo"
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	return c.JSON(http.StatusCreated, res)
}

// DeleteTodo removes a todo by id
func DeleteTodo(c echo.Context) error {
	log.Println(`DELETE("/todos", DeleteTodo)`)
	log.Println("Params -> ", c.ParamNames(), c.ParamValues())

	todoID := c.Param("id")
	id, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil {
		log.Println(err)
		errMsg := "error parsing id"
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	todo := new(Todo)
	res, err := todo.deleteTodo(id)
	if err != nil {
		errMsg := "error deleting todo"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, res)
}

// DeleteTodos removes all the todos
func DeleteTodos(c echo.Context) error {
	log.Println(`DELETE("/todos", DeleteTodos)`)

	todo := new(Todo)
	res, err := todo.deleteAll()
	if err != nil {
		errMsg := "error deleting todos"
		return c.JSON(http.StatusNotFound, errMsg)
	}

	return c.JSON(http.StatusOK, res)
}
