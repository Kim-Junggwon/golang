package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	Id      int
	Title   string
	Name    string
	Date    time.Time
	Content string
}

func checkError(err error) {
	if err != nil {
		fmt.Println("[Fatal error]", err.Error())
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
		Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		Title TEXT,
		Name TEXT,
		Date DATETIME,
		Content TEXT
	);
	`

	_, err := db.Exec(sql_table)
	checkError(err)
}

func InsertItem(db *sql.DB, item Item) {
	sql_addItem := "INSERT OR REPLACE INTO list(Title, Name, Date, Content) VALUES (?, ?, ?, ?);"

	stmt, err := db.Prepare(sql_addItem)
	checkError(err)
	defer stmt.Close()

	_, err2 := stmt.Exec(item.Title, item.Name, time.Now(), item.Content)
	checkError(err2)
}

func ReadItems(db *sql.DB) []Item {
	sql_readall := "SELECT * FROM list ORDER BY Id DESC;"

	rows, err := db.Query(sql_readall)
	checkError(err)

	result := []Item{}

	for rows.Next() {
		item := Item{}
		err2 := rows.Scan(&item.Id, &item.Title, &item.Name, &item.Date, &item.Content)
		checkError(err2)

		result = append(result, item)
	}

	return result
}

func ReadPage(db *sql.DB, id string) Item {
	page := Item{}

	sql_read := "SELECT * FROM list WHERE Id=" + id

	err := db.QueryRow(sql_read).Scan(&page.Id, &page.Title, &page.Name, &page.Date, &page.Content)
	checkError(err)

	return page
}

func UpdateContent(db *sql.DB, id string, content string) {
	sql_update := "UPDATE list SET Content=?, Date=? where Id=?"

	stmt, err := db.Prepare(sql_update)
	checkError(err)
	defer stmt.Close()

	_, err2 := stmt.Exec(content, time.Now(), id)
	checkError(err2)
}
