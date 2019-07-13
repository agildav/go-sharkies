package sharks

import (
	"fmt"
	"log"
	"sharkies/db"

	"time"
)

// ----------------------------------------------------------------------

/*
	!: Add new methods here
*/

// findAll looks for all the sharks and orders them by id asc
func (s Shark) findAll() ([]Shark, error) {
	init := time.Now()
	pg := db.GetDatabase()

	// Columns to select
	const (
		cID    = "id"
		cName  = "name"
		cBname = "bname"
		cImage = "image"
	)

	var sharks []Shark
	order := fmt.Sprintf("%s", cID+" "+"ASC")
	err := pg.Model(&sharks).Column(cID, cName, cBname, cImage).Order(order).Select()
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))

		sharks = make([]Shark, 0)
		return sharks, err
	}

	rowsReturned := len(sharks)
	if rowsReturned == 0 {
		log.Println("msg -> there are no sharks")
		log.Printf("msg in -> %v", time.Since(init))

		sharks = make([]Shark, 0)
		return sharks, nil
	} else if rowsReturned <= 0 {
		err := fmt.Errorf("error -> could not get sharks")

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))

		sharks = make([]Shark, 0)
		return sharks, err
	}

	log.Printf("rows returned -> %v in %v", rowsReturned, time.Since(init))

	return sharks, nil
}

// findByID looks for a shark with @id
func (s Shark) findByID(id int64) (Shark, error) {
	init := time.Now()
	pg := db.GetDatabase()

	shark := Shark{ID: id}

	err := pg.Select(&shark)
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		shark = Shark{}
		return shark, err
	}

	if shark.ID != id {
		err := fmt.Errorf("error -> id %v not found", id)

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		shark = Shark{}
		return shark, err
	}

	log.Printf("rows returned -> %v in %v", 1, time.Since(init))

	return shark, nil
}

// addShark inserts a new shark
func (s Shark) addShark(shark *Shark) (int64, error) {
	init := time.Now()
	pg := db.GetDatabase()

	query := `INSERT into sharks("name", "bname", "description", "image") VALUES (?name, ?bname, ?description, ?image) RETURNING id`

	var sharkModel Shark
	res, err := pg.Query(&sharkModel, query, *shark)

	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return -1, err
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected <= 0 {
		err := fmt.Errorf("error -> could not add shark")

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return -1, err
	}

	log.Printf("rows affected -> %v in %v", rowsAffected, time.Since(init))
	return sharkModel.ID, nil
}

// deleteShark deletes a shark with @id
func (s Shark) deleteShark(id int64) (string, error) {
	init := time.Now()
	pg := db.GetDatabase()

	var shark Shark

	query := `DELETE FROM sharks WHERE id = ?`

	res, err := pg.Query(&shark, query, id)
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

	return "shark deleted", nil
}

// deleteAll truncates the table
func (s Shark) deleteAll() (string, error) {
	init := time.Now()
	pg := db.GetDatabase()

	beforeCount, err := pg.Model((*Shark)(nil)).Count()
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	var shark Shark

	query := `TRUNCATE TABLE sharks`

	_, err = pg.Query(&shark, query)
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	afterCount, err := pg.Model((*Shark)(nil)).Count()
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	if afterCount != 0 {
		err := fmt.Errorf("could not delete sharks")

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	log.Printf("rows affected -> %v in %v", beforeCount, time.Since(init))

	return "all sharks deleted", nil
}

// patchShark edits an existing shark
func (s Shark) patchShark(id int64, shark *Shark) (string, error) {
	init := time.Now()
	pg := db.GetDatabase()

	res, err := pg.Model(shark).Where("id = ?", id).UpdateNotNull()
	if err != nil {
		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected <= 0 {
		err := fmt.Errorf("error -> could not add shark")

		log.Println(err)
		log.Printf("err in -> %v", time.Since(init))
		return "", err
	}

	log.Printf("rows affected -> %v in %v", rowsAffected, time.Since(init))
	return "shark patched", nil
}
