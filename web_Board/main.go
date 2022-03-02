package main

import (
	"board/hd"
	"net/http"
)

func main() {
	dbpath := "test.db"

	http.ListenAndServe(":3000", hd.Handler(dbpath))
}
