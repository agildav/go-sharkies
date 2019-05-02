package db

import (
	"github.com/go-pg/pg"

	"log"
	"time"
)

// // // // // // // // // // // // // // // // // // // // // //

// DB is the PostgreSQL connected database, retrieve it using GetDatabase()
var DB *pg.DB

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

// // // // // // // // // // // // // // // // // // // // // //

// GetDatabase returns the database connected to PostgreSQL
func GetDatabase() *pg.DB {
	return DB
}

// setDatabase assings the database pointer
func setDatabase(db *pg.DB) {
	DB = db
}

// Setup connects to PostgreSQL and sets a database
func Setup(dbUser, dbPassword, dbHost, dbPort, dbName string) {
	log.Println(":: PostgreSQL init")
	init := time.Now()

	pgOptions := &pg.Options{
		User:     dbUser,
		Password: dbPassword,
		Addr:     dbHost + dbPort,
		Database: dbName,
	}

	db := pg.Connect(pgOptions)

	// Logger
	db.AddQueryHook(dbLogger{})

	err := testConnection(db)
	if err != nil {
		log.Fatal("error connecting to database -> ", err)
	}

	setDatabase(db)

	log.Println(":: PostgreSQL ready, took", time.Since(init))
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
