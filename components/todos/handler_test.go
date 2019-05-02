package todos

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
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

	t.Run("returns the list of todos at index", func(t *testing.T) {
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

	t.Run("returns the list of todos at /todos", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/todos")

		todosJSON := `[{"id":1,"title":"Test Title one","body":"Test Body one"},{"id":2,"title":"Test Title two","body":"Test Body two"}]`
		expectedJSON := string(todosJSON + "\n")

		// Assertions
		if assert.NoError(t, GetTodos(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns an empty slice when non-existent todos", func(t *testing.T) {
		todo := new(Todo)
		todo.deleteAll()

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/todos")

		todosJSON := `[]`
		expectedJSON := string(todosJSON + "\n")

		// Assertions
		if assert.NoError(t, GetTodos(c)) {
			if assert.Equal(t, http.StatusOK, rec.Code) {
				assert.Equal(t, expectedJSON, rec.Body.String())

				// adds the todo and go back to the previous state
				newTodo1 := &Todo{ID: 1, Title: "Test Title one", Body: "Test Body one"}
				newTodo2 := &Todo{ID: 2, Title: "Test Title two", Body: "Test Body two"}
				todo := new(Todo)
				todo.addTodo(newTodo1)
				todo.addTodo(newTodo2)
			}
		}
	})
}

func Test_GetTodo(t *testing.T) {

	/*
		!: Add new cases here
	*/

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

	t.Run("returns an error when invalid id", func(t *testing.T) {
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
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns an error when non-existent id", func(t *testing.T) {
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

func Test_AddTodo(t *testing.T) {

	/*
		!: Add new cases here
	*/

	t.Run("returns todo inserted", func(t *testing.T) {
		json := `{"id":3,"title":"Test Title three","body":"Test Body three"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todos")

		todoJSON := `"todo inserted"`
		expectedJSON := string(todoJSON + "\n")

		// Assertions
		if assert.NoError(t, AddTodo(c)) {
			if assert.Equal(t, http.StatusCreated, rec.Code) {
				assert.Equal(t, expectedJSON, rec.Body.String())

				// delete the todo and go back to the previous state
				todo := new(Todo)
				var id int64 = 3
				todo.deleteTodo(id)
			}
		}
	})

	t.Run("returns an error when invalid id", func(t *testing.T) {
		json := `{"id":"a","title":"Test Title invalid id","body":"Test Body invalid id"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todos")

		todoJSON := `"error binding todo"`
		expectedJSON := string(todoJSON + "\n")

		// Assertions
		if assert.NoError(t, AddTodo(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns an error when existing id", func(t *testing.T) {
		json := `{"id":2,"title":"Test Title existing id","body":"Test Body existing id"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todos")

		todoJSON := `"error inserting todo"`
		expectedJSON := string(todoJSON + "\n")

		// Assertions
		if assert.NoError(t, AddTodo(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})
}

func Test_DeleteTodo(t *testing.T) {
	/*
		!: Add new cases here
	*/

	t.Run("returns todo deleted", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/todos/:id")
		c.SetParamNames("id")
		c.SetParamValues("2")

		todoJSON := `"todo deleted"`
		expectedJSON := string(todoJSON + "\n")

		// Assertions
		if assert.NoError(t, DeleteTodo(c)) {
			if assert.Equal(t, http.StatusOK, rec.Code) {
				assert.Equal(t, expectedJSON, rec.Body.String())

				// adds the todo and go back to the previous state
				newTodo := &Todo{ID: 2, Title: "Test Title two", Body: "Test Body two"}
				todo := new(Todo)
				todo.addTodo(newTodo)
			}
		}
	})

	t.Run("returns an error when invalid id", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/todos/:id")
		c.SetParamNames("id")
		c.SetParamValues("a")

		todoJSON := `"error parsing id"`
		expectedJSON := string(todoJSON + "\n")

		// Assertions
		if assert.NoError(t, DeleteTodo(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns an error when non-existent id", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/todos/:id")
		c.SetParamNames("id")
		c.SetParamValues("999")

		todoJSON := `"error deleting todo"`
		expectedJSON := string(todoJSON + "\n")

		// Assertions
		if assert.NoError(t, DeleteTodo(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})
}

func Test_DeleteTodos(t *testing.T) {
	/*
		!: Add new cases here
	*/

	t.Run("returns all todos deleted", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/todos")

		todosJSON := `"all todos deleted"`
		expectedJSON := string(todosJSON + "\n")

		// Assertions
		if assert.NoError(t, DeleteTodos(c)) {
			if assert.Equal(t, http.StatusOK, rec.Code) {
				assert.Equal(t, expectedJSON, rec.Body.String())

				// adds the todo and go back to the previous state
				newTodo1 := &Todo{ID: 1, Title: "Test Title one", Body: "Test Body one"}
				newTodo2 := &Todo{ID: 2, Title: "Test Title two", Body: "Test Body two"}
				todo := new(Todo)
				todo.addTodo(newTodo1)
				todo.addTodo(newTodo2)
			}
		}
	})
}
