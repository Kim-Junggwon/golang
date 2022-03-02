package db

import (
	"testing"
	"time"
)

func TestDatabase(t *testing.T) {
	const dbpath = "foo.db"

	db := InitDB(dbpath)
	defer db.Close()
	CreateTable(db)

	items := []Item{
		Item{"0", "T0", "N0", time.Now()},
		Item{"1", "T1", "N1", time.Now()},
	}
	InsertItem(db, items)

	readItems := ReadItem(db)
	t.Log(readItems)

	items2 := []Item{
		Item{"5", "T5", "N5", time.Now()},
		Item{"6", "T6", "N6", time.Now()},
	}
	InsertItem(db, items2)

	readItems2 := ReadItem(db)
	t.Log(readItems2)
}
