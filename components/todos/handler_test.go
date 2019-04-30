package todos

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agildav/go-boilerplate/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// // // // // // // // // // // // // // // // // // // // // // // // // //
var (
	env        map[string]string
	err        error
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
	e          *echo.Echo
)

func init() {
	// Config setup
	env, err = godotenv.Read("../../.env")
	if err != nil {
		log.Fatal("error loading .env file -> ", err)
	}

	dbUser = env["TEST_DB_USER"]
	dbPassword = env["TEST_DB_PASSWORD"]
	dbHost = env["TEST_DB_HOST"]
	dbPort = env["TEST_DB_PORT"]
	dbName = env["TEST_DB_NAME"]

	// DB
	db.Setup(dbUser, dbPassword, dbHost, dbPort, dbName)

	// Echo
	e = echo.New()
}

// // // // // // // // // // // // // // // // // // // // // // // // // //

/*
	!: Add new tests here
*/

func Test_GetTodos(t *testing.T) {

	/*
		!: Add new cases here
	*/

	t.Run("returns the list of todos", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/")

		todosJSON := `[{"id":1,"title":"Test Title one","body":"Test Body one"},{"id":2,"title":"Test Title two","body":"Test Body two"}]`
		expectedJSON := string(todosJSON + "\n")

		// Assertions
		if assert.NoError(t, GetTodos(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})
}

func Test_GetTodo(t *testing.T) {

	/*
		!: Add new cases here
	*/

	t.Run("returns an error when id is not an integer", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/todos/:id")
		c.SetParamNames("id")
		c.SetParamValues("a")

		todoJSON := `"error parsing id"`
		expectedJSON := string(todoJSON + "\n")

		// Assertions
		if assert.NoError(t, GetTodo(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns a todo with id = 2", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/todos/:id")
		c.SetParamNames("id")
		c.SetParamValues("2")

		todoJSON := `{"id":2,"title":"Test Title two","body":"Test Body two"}`
		expectedJSON := string(todoJSON + "\n")

		// Assertions
		if assert.NoError(t, GetTodo(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns an error with non-existent id", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/todos/:id")
		c.SetParamNames("id")
		c.SetParamValues("999")

		todoJSON := `"error obtaining todo"`
		expectedJSON := string(todoJSON + "\n")

		// Assertions
		if assert.NoError(t, GetTodo(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})
}
