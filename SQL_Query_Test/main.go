package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3" // db
)

type Person struct {
	Name string
	Age  int
}

type Db *sql.DB

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

func main() {
	// db connect
	Db, _ := sql.Open("sqlite3", "./test.db")
	defer Db.Close()

	// create table
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`
	_, err := Db.Exec(cmd)
	checkError(err)

	// insert data
	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err = Db.Exec(cmd, "Kim", 25)
	checkError(err)

	// select data
	cmd = "SELECT * FROM person"
	rows, _ := Db.Query(cmd)
	defer rows.Close()

	var pp []Person
	for rows.Next() {
		var p Person
		err = rows.Scan(&p.Name, &p.Age)
		checkError(err)
		pp = append(pp, p)
	}

	err = rows.Err()
	checkError(err)

	// Print data
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
