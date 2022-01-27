package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dixonwille/wmenu"
)

func main() {

	db, err := sql.Open("sqlite3", "./keys.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer close
	defer db.Close()

	// _, err = db.Exec("CREATE TABLE `keys` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `key_source` VARCHAR(64) NULL, `key` VARCHAR(255) NOT NULL)")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	menu := wmenu.NewMenu("PLEASE SELECT THE OPTION 0,1,2,3")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })
	menu.Option("Add secret key", 0, true, nil)
	menu.Option("Find the key", 1, false, nil)
	// menu.Option("Update the key information", 2, false, nil)
	menu.Option("Delete a key by ID", 2, false, nil)
	menu.Option("Quit Application", 3, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the key's source")
		keySource, _ := reader.ReadString('\n')
		if keySource != "\n" {
			keySource = strings.TrimSuffix(keySource, "\n")
		}

		fmt.Println("Enter the Key")
		key, _ := reader.ReadString('\n')
		if key != "\n" {
			key = strings.TrimSuffix(key, "\n")
		}
		newKey := New()

		newKey.key_source = keySource
		newKey.key = key
		newKey.AddKeys(db)

	case 1:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the key to search : ")
		searchString, _ := reader.ReadString('\n')
		searchString = strings.TrimSuffix(searchString, "\n")
		key := searchForKey(db, searchString)

		fmt.Printf("Found %v results\n", len(key))

		for _, ourPerson := range key {
			fmt.Printf("Retrieved Information", ourPerson.key_source, ourPerson.key)
		}
		break
	// case 2:
	// 	fmt.Println("Update a Person's information")
	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the ID of the key you want to delete : ")
		searchString, _ := reader.ReadString('\n')
		idToDelete := strings.TrimSuffix(searchString, "\n")
		affected := deleteKey(db, idToDelete)

		if affected == 1 {
			fmt.Println("Deleted key from database")
		}
	case 3:
		fmt.Println("Goodbye!")
		os.Exit(3)
	}
}
