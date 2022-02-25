package main

import (
	"board/db"
	"fmt"
	"time"
)

func main() {
	const dbpath = "test2.db"

	d := db.InitDB(dbpath)
	defer d.Close()

	db.CreateTable(d)

	items := []db.Item{
		db.Item{"0", "T0", "N0", time.Now()},
		db.Item{"1", "T1", "N1", time.Now()},
	}
	db.InsertItem(d, items)

	read := db.ReadItem(d)

	for _, v := range read {
		fmt.Println(v.Id, v.Name, v.Title, v.Date)
	}

}
