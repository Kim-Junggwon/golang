package db

import (
	"testing"
	"time"
)

func TestDatabase(t *testing.T) {
	const dbpath = "../test.db"

	db := InitDB(dbpath)
	defer db.Close()
	CreateTable(db)

	item := Item{2, "T2", "N2", time.Now(), "Cont2"}
	InsertItem(db, item)

	readItems := ReadItem(db)
	t.Log(readItems)

}
