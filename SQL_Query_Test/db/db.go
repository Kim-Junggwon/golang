package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	Id    string
	Title string
	Name  string
	Date  time.Time
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
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
		Date DATETIME DEFAULT (DATETIME('now', 'localtime'))
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func InsertItem(db *sql.DB, items []Item) {
	sql_addItem := `
	INSERT OR REPLACE INTO list(Id, Title, Name, Date) VALUES (?, ?, ?, ?);
	`

	stmt, err := db.Prepare(sql_addItem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range items {
		_, err2 := stmt.Exec(item.Id, item.Title, item.Name, item.Date)
		if err2 != nil {
			panic(err2)
		}
	}
}

func ReadItem(db *sql.DB) []Item {
	sql_readall := `
	SELECT * FROM list
	ORDER BY Date DESC;
	`

	rows, err := db.Query(sql_readall)
	if err != nil {
		panic(err)
	}

	var result []Item
	for rows.Next() {
		item := Item{}
		err2 := rows.Scan(&item.Id, &item.Title, &item.Name, &item.Date)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}

	return result
}
