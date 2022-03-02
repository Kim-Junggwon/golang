package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	Id      string
	Title   string
	Name    string
	Date    time.Time
	Content string
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	checkError(err)
	if db == nil {
		panic("db nil")
	}

	return db
}

func CreateTable(db *sql.DB) {
	sql_table := `
	CREATE TABLE IF NOT EXISTS list(
		Id TEXT NOT NULL PRIMARY KEY,
		Title TEXT,
		Name TEXT,
		Date DATETIME,
		Content TEXT
	);
	`

	_, err := db.Exec(sql_table)
	checkError(err)
}

func InsertItem(db *sql.DB, items []Item) {
	sql_addItem := "INSERT OR REPLACE INTO list(Id, Title, Name, Date, Content) VALUES (?, ?, ?, ?, ?);"

	stmt, err := db.Prepare(sql_addItem)
	checkError(err)
	defer stmt.Close()

	for _, item := range items {
		_, err2 := stmt.Exec(item.Id, item.Title, item.Name, item.Date, item.Content)
		checkError(err2)
	}
}

func ReadItem(db *sql.DB) []Item {
	sql_readall := "SELECT * FROM list ORDER BY Date DESC;"

	rows, err := db.Query(sql_readall)
	checkError(err)

	result := []Item{}

	for rows.Next() {
		var item Item // item := Item{}
		err2 := rows.Scan(&item.Id, &item.Title, &item.Name, &item.Date, &item.Content)
		checkError(err2)

		result = append(result, item)
	}

	return result
}
