package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type keys struct {
	id         int
	key_source string
	key        string
}

func New() *keys {
	return &keys{}
}
func (k keys) AddKeys(db *sql.DB) {

	// saves application from sql injection.
	stmt, err := db.Prepare("INSERT INTO keys (id, key_source, key) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)

	}
	stmt.Exec(1, k.key_source, k.key)
	defer stmt.Close()

	fmt.Println("Added key source & key successfully")
}

func searchForKey(db *sql.DB, searchString string) []keys {

	rows, err := db.Query("SELECT id, key_source, key, email, FROM keys WHERE key_source like '%" + searchString + "%' OR key like '%" + searchString + "%'")

	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	key := make([]keys, 0)

	for rows.Next() {
		ourPerson := keys{}
		err = rows.Scan(&ourPerson.id, &ourPerson.key_source, &ourPerson.key)
		if err != nil {
			log.Fatal(err)
		}

		key = append(key, ourPerson)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return key
}

func deleteKey(db *sql.DB, idToDelete string) int64 {

	stmt, _ := db.Prepare("DELETE FROM people where id = ?")
	defer stmt.Close()

	res, _ := stmt.Exec(idToDelete)

	affected, _ := res.RowsAffected()

	return affected
}
