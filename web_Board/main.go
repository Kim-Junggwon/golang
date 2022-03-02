package main

import (
	"board/hd"
	"net/http"
)

func main() {
	/* 	const dbpath = "test.db"
	   	d := db.InitDB(dbpath) // db 연동 인스턴스 생성
	   	defer d.Close()

	   	db.CreateTable(d) // table이 존재하지 않는다면 생성

	   	items := []db.Item{
	   		db.Item{"0", "T0", "N0", time.Now()},
	   		db.Item{"1", "T1", "N1", time.Now()},
	   	}
	   	db.InsertItem(d, items) // db에 데이터 삽입

	   	read := db.ReadItem(d) // db 데이터 읽기
	   	for _, v := range read {
	   		fmt.Println(v.Id, v.Name, v.Title, v.Date.Format("2006-01-02 15:04"))
	   	} */
	dbpath := "test.db"

	http.ListenAndServe(":3000", hd.Handler(dbpath))
}
