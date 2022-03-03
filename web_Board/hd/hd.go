package hd

import (
	"board/db"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var d *sql.DB

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	items := db.ReadItem(d)

	tmpl, err := template.New("Index").ParseFiles("templates/index.html")
	checkError(err)
	err = tmpl.ExecuteTemplate(w, "index.html", items)
	checkError(err)
}

func getListHandler(w http.ResponseWriter, r *http.Request) {
	// template
	tmpl, err := template.New("Write").ParseFiles("templates/write.html")
	checkError(err)
	err = tmpl.ExecuteTemplate(w, "write.html", "test")
	checkError(err)
}

func postListHandler(w http.ResponseWriter, r *http.Request) {
	db.CreateTable(d)

	tmpl, err := template.New("Write").ParseFiles("templates/write.html")
	checkError(err)
	err = tmpl.ExecuteTemplate(w, "write.html", "test")
	checkError(err)

	new := db.Item{
		Title:   r.FormValue("title"),
		Name:    r.FormValue("name"),
		Content: r.FormValue("content"),
	}

	db.InsertItem(d, new)
	// 완료 후 인덱스 페이지로 이동..
	// http.Redirect(w, r, "/", http.StatusOK) / responsewriter 2 번 호출 오류,..
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// row 별 페이지
}

func Handler(dbpath string) http.Handler {
	d = db.InitDB(dbpath)
	db.CreateTable(d)
	// defer d.Close()

	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/{id:[0-9]+}", pageHandler).Methods("GET")

	mux.HandleFunc("/write", getListHandler).Methods("GET")
	mux.HandleFunc("/write", postListHandler).Methods("POST")

	return mux
}
