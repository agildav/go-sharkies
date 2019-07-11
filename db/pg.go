package db

import (
	"github.com/go-pg/pg"

	"log"
	"time"
)

// ----------------------------------------------------------------------

// db is the PostgreSQL connected database, retrieve it using GetDatabase()
var db *pg.DB

// dbLogger is the logger attached to the database in order to print the SQL statements
type dbLogger struct{}

// Init establishes the PostgreSQL connection
func Init(env map[string]string) {
	var (
		dbUser     = env["DB_USER"]
		dbPassword = env["DB_PASSWORD"]
		dbHost     = env["DB_HOST"]
		dbPort     = env["DB_PORT"]
		dbName     = env["DB_NAME"]
	)

	Setup(dbUser, dbPassword, dbHost, dbPort, dbName)
}

// ----------------------------------------------------------------------

// GetDatabase returns the database connected to PostgreSQL
func GetDatabase() *pg.DB {
	return db
}

// setDatabase assings the database pointer
func setDatabase(database *pg.DB) {
	db = database
}

// Setup connects to PostgreSQL and sets a database
func Setup(dbUser, dbPassword, dbHost, dbPort, dbName string) {
	log.Println(":: PostgreSQL, starting")
	init := time.Now()

	pgOptions := &pg.Options{
		User:     dbUser,
		Password: dbPassword,
		Addr:     dbHost + ":" + dbPort,
		Database: dbName,
	}

	db := pg.Connect(pgOptions)

	err := testConnection(db)
	if err != nil {
		log.Fatal("error connecting to database -> ", err)
	}

	// Logger
	db.AddQueryHook(dbLogger{})

	setDatabase(db)

	log.Println(":: PostgreSQL, ready in", time.Since(init))
}

// testConnection checks if a client is connected to the database
func testConnection(db *pg.DB) error {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")

	return err
}

// BeforeQuery executes before a query
func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

// AfterQuery executes after a query
func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	log.Println(q.FormattedQuery())
}
