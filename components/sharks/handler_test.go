package sharks

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/agildav/sharkies/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// // // // // // // // // // // // // // // // // // // // // // // // // //
var (
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
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load("../../.env")
		if err != nil {
			log.Fatal("error loading .env file -> ", err)
		}
	}

	// DB Config
	dbUser = os.Getenv("TEST_DB_USER")
	dbPassword = os.Getenv("TEST_DB_PASSWORD")
	dbHost = os.Getenv("TEST_DB_HOST")
	dbPort = os.Getenv("TEST_DB_PORT")
	dbName = os.Getenv("TEST_DB_NAME")

	// DB
	db.Setup(dbUser, dbPassword, dbHost, dbPort, dbName)

	// Echo
	e = echo.New()
}

// // // // // // // // // // // // // // // // // // // // // // // // // //

/*
	!: Add new tests here
*/

func Test_GetSharks(t *testing.T) {

	/*
		!: Add new cases here
	*/

	t.Run("returns the list of sharks at index", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/")

		sharksJSON := `[{"id":1,"name":"Basking Shark","bname":"Cetorhinus maximus"},{"id":2,"name":"Zebra Bullhead Shark","bname":"Heterodontus zebra"}]`
		expectedJSON := string(sharksJSON + "\n")

		// Assertions
		if assert.NoError(t, GetSharks(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns the list of sharks at /sharks", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/sharks")

		sharksJSON := `[{"id":1,"name":"Basking Shark","bname":"Cetorhinus maximus"},{"id":2,"name":"Zebra Bullhead Shark","bname":"Heterodontus zebra"}]`
		expectedJSON := string(sharksJSON + "\n")

		// Assertions
		if assert.NoError(t, GetSharks(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns an empty slice when non-existent sharks", func(t *testing.T) {
		shark := new(Shark)
		shark.deleteAll()

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/sharks")

		sharksJSON := `[]`
		expectedJSON := string(sharksJSON + "\n")

		// Assertions
		if assert.NoError(t, GetSharks(c)) {
			if assert.Equal(t, http.StatusOK, rec.Code) {
				assert.Equal(t, expectedJSON, rec.Body.String())

				// adds the shark and go back to the previous state
				newshark1 := &Shark{ID: 1, Name: "Basking Shark", Bname: "Cetorhinus maximus", Description: "Description of basking shark", Image: "Image of basking shark"}
				newshark2 := &Shark{ID: 2, Name: "Zebra Bullhead Shark", Bname: "Heterodontus zebra", Description: "Description of zebra shark", Image: "Image of zebra shark"}
				shark.addShark(newshark1)
				shark.addShark(newshark2)
			}
		}
	})
}

func Test_Getshark(t *testing.T) {

	/*
		!: Add new cases here
	*/

	t.Run("returns a shark with id = 2", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/sharks/:id")
		c.SetParamNames("id")
		c.SetParamValues("2")

		sharkJSON := `{"id":2,"name":"Zebra Bullhead Shark","bname":"Heterodontus zebra","description":"Description of zebra shark","image":"Image of zebra shark"}`
		expectedJSON := string(sharkJSON + "\n")

		// Assertions
		if assert.NoError(t, GetShark(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns an error when invalid id", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/sharks/:id")
		c.SetParamNames("id")
		c.SetParamValues("a")

		sharkJSON := `"error parsing id"`
		expectedJSON := string(sharkJSON + "\n")

		// Assertions
		if assert.NoError(t, GetShark(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns an error when non-existent id", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/sharks/:id")
		c.SetParamNames("id")
		c.SetParamValues("999")

		sharkJSON := `"error obtaining shark"`
		expectedJSON := string(sharkJSON + "\n")

		// Assertions
		if assert.NoError(t, GetShark(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})
}

func Test_addShark(t *testing.T) {

	/*
		!: Add new cases here
	*/

	t.Run("returns shark inserted", func(t *testing.T) {
		json := `{"id":3,"name":"Test Name three","bname":"Test Bname three","description":"Test Description three","image":"Test Image three"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/sharks")

		sharkJSON := `"shark inserted"`
		expectedJSON := string(sharkJSON + "\n")

		// Assertions
		if assert.NoError(t, addShark(c)) {
			if assert.Equal(t, http.StatusCreated, rec.Code) {
				assert.Equal(t, expectedJSON, rec.Body.String())

				// delete the shark and go back to the previous state
				shark := new(Shark)
				var id int64 = 3
				shark.deleteShark(id)
			}
		}
	})

	t.Run("returns an error when invalid id", func(t *testing.T) {
		json := `{"id":a,"name":"Test Name invalid id","bname":"Test Bname invalid id","description":"Test Description invalid id","image":"Test Image invalid id"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/sharks")

		sharkJSON := `"error binding shark"`
		expectedJSON := string(sharkJSON + "\n")

		// Assertions
		if assert.NoError(t, addShark(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns an error when existing id", func(t *testing.T) {
		json := `{"id":2,"name":"Test Name existing id","bname":"Test Bname existing id","description":"Test Description existing id","image":"Test Image existing id"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/sharks")

		sharkJSON := `"error inserting shark"`
		expectedJSON := string(sharkJSON + "\n")

		// Assertions
		if assert.NoError(t, addShark(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})
}

func Test_deleteShark(t *testing.T) {
	/*
		!: Add new cases here
	*/

	t.Run("returns shark deleted", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/sharks/:id")
		c.SetParamNames("id")
		c.SetParamValues("2")

		sharkJSON := `"shark deleted"`
		expectedJSON := string(sharkJSON + "\n")

		// Assertions
		if assert.NoError(t, deleteShark(c)) {
			if assert.Equal(t, http.StatusOK, rec.Code) {
				assert.Equal(t, expectedJSON, rec.Body.String())

				// adds the shark and go back to the previous state
				newshark := &Shark{ID: 2, Name: "Zebra Bullhead Shark", Bname: "Heterodontus zebra", Description: "Description of zebra shark", Image: "Image of zebra shark"}
				shark := new(Shark)
				shark.addShark(newshark)
			}
		}
	})

	t.Run("returns an error when invalid id", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/sharks/:id")
		c.SetParamNames("id")
		c.SetParamValues("a")

		sharkJSON := `"error parsing id"`
		expectedJSON := string(sharkJSON + "\n")

		// Assertions
		if assert.NoError(t, deleteShark(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})

	t.Run("returns an error when non-existent id", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/sharks/:id")
		c.SetParamNames("id")
		c.SetParamValues("999")

		sharkJSON := `"error deleting shark"`
		expectedJSON := string(sharkJSON + "\n")

		// Assertions
		if assert.NoError(t, deleteShark(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, expectedJSON, rec.Body.String())
		}
	})
}

func Test_deleteSharks(t *testing.T) {
	/*
		!: Add new cases here
	*/

	t.Run("returns all sharks deleted", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		c := e.NewContext(req, rec)
		c.SetPath("/sharks")

		sharksJSON := `"all sharks deleted"`
		expectedJSON := string(sharksJSON + "\n")

		// Assertions
		if assert.NoError(t, deleteSharks(c)) {
			if assert.Equal(t, http.StatusOK, rec.Code) {
				assert.Equal(t, expectedJSON, rec.Body.String())

				// adds the shark and go back to the previous state
				newshark1 := &Shark{ID: 1, Name: "Basking Shark", Bname: "Cetorhinus maximus", Description: "Description of basking shark", Image: "Image of basking shark"}
				newshark2 := &Shark{ID: 2, Name: "Zebra Bullhead Shark", Bname: "Heterodontus zebra", Description: "Description of zebra shark", Image: "Image of zebra shark"}
				shark := new(Shark)
				shark.addShark(newshark1)
				shark.addShark(newshark2)
			}
		}
	})
}
