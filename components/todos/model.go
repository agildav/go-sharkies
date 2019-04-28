package todos

import (
	"github.com/agildav/go-boilerplate/db"
	"log"
	"time"
)

// // // // // // // // // // // // // // // // // // // // // // // // // //s

// Todo represents a todo table
type Todo struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// // // // // // // // // // // // // // // // // // // // // // // // // //
/*
	!: Add new methods here
*/

// searchTodos looks for all the todos and orders them by id
func (t Todo) searchTodos() ([]Todo, error) {
	init := time.Now()

	pg := db.GetDatabase()

	var todos []Todo

	query := "SELECT * FROM todos ORDER BY id"

	res, err := pg.Query(&todos, query)
	if err != nil {
		log.Println(err)
		return todos, err
	}

	log.Println(res.Model())
	log.Printf("rows returned -> %v in %v", res.RowsReturned(), time.Since(init))

	return todos, nil
}
