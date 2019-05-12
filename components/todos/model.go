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

// findAll looks for all the todos and orders them by id asc
func (t Todo) findAll() ([]Todo, error) {
	init := time.Now()

	pg := db.GetDatabase()

	var todos []Todo
	err := pg.Model(&todos).Select()
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))

		todos = make([]Todo, 0)
		return todos, err
	}

	rowsReturned := len(todos)
	if rowsReturned == 0 {
		log.Println("msg -> there are no todos")
		log.Printf("msg in -> %v", time.Since(init))

		todos = make([]Todo, 0)
		return todos, nil
	} else if rowsReturned <= 0 {
		err := fmt.Errorf("error -> could not get todos")

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))

		todos = make([]Todo, 0)
		return todos, err
	}

	log.Printf("rows returned -> %v in %v", rowsReturned, time.Since(init))

	return todos, nil
}

// findByID looks for a todo with @id
func (t Todo) findByID(id int64) (Todo, error) {
	init := time.Now()

	pg := db.GetDatabase()

	todo := Todo{ID: id}

	err := pg.Select(&todo)
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		todo = Todo{}
		return todo, err
	}

	if todo.ID != id {
		err := fmt.Errorf("error -> id %v not found", id)

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		todo = Todo{}
		return todo, err
	}

	log.Printf("rows returned -> %v in %v", 1, time.Since(init))

	return todo, nil
}

// addTodo inserts a new todo
func (t Todo) addTodo(todo *Todo) (string, error) {
	init := time.Now()

	pg := db.GetDatabase()

	query := `INSERT into todos("id", "title", "body") VALUES (?id, ?title, ?body)`

	var todoModel Todo
	res, err := pg.Query(&todoModel, query, *todo)
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected <= 0 {
		err := fmt.Errorf("error -> could not add todo")

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	log.Printf("rows affected -> %v in %v", rowsAffected, time.Since(init))
	return "todo inserted", nil
}

// deleteTodo deletes a todo with @id
func (t Todo) deleteTodo(id int64) (string, error) {
	init := time.Now()

	pg := db.GetDatabase()

	var todo Todo

	query := `DELETE FROM todos WHERE id = ?`

	res, err := pg.Query(&todo, query, id)
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected <= 0 {
		err := fmt.Errorf("error -> id %v not found", id)

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	log.Printf("rows affected -> %v in %v", rowsAffected, time.Since(init))

	return "todo deleted", nil
}

// deleteAll truncates the table
func (t Todo) deleteAll() (string, error) {
	init := time.Now()

	pg := db.GetDatabase()

	beforeCount, err := pg.Model((*Todo)(nil)).Count()
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	var todo Todo

	query := `TRUNCATE TABLE todos`

	_, err = pg.Query(&todo, query)
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	afterCount, err := pg.Model((*Todo)(nil)).Count()
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	if afterCount != 0 {
		err := fmt.Errorf("could not delete todos")

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	log.Printf("rows affected -> %v in %v", beforeCount, time.Since(init))

	return "all todos deleted", nil
}
