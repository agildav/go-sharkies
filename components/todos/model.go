package todos

import (
	"fmt"
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

// findAll looks for all the todos and orders them by id
func (t Todo) findAll() ([]Todo, error) {
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

// findByID looks for a todo with @id
func (t Todo) findByID(id int64) (Todo, error) {
	init := time.Now()

	pg := db.GetDatabase()

	var todo Todo

	query := `SELECT * FROM todos WHERE id = ?`

	res, err := pg.Query(&todo, query, id)
	if err != nil {
		todo = Todo{}

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return todo, err
	}

	rowsReturned := res.RowsReturned()

	if rowsReturned <= 0 {
		err := fmt.Errorf("error -> id %v not found", id)
		todo = Todo{}

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return todo, err
	}

	log.Println(res.Model())
	log.Printf("rows returned -> %v in %v", rowsReturned, time.Since(init))

	return todo, nil
}
