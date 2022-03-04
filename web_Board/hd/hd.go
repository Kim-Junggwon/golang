package hd

import (
	"board/db"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

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
	tmpl, err := template.New("Write").ParseFiles("templates/write.html")
	checkError(err)
	err = tmpl.ExecuteTemplate(w, "write.html", "")
	checkError(err)
}

func postListHandler(w http.ResponseWriter, r *http.Request) {
	db.CreateTable(d) // table이 없을 경우 생성

	new := db.Item{
		Title:   r.FormValue("title"),
		Name:    r.FormValue("name"),
		Content: r.FormValue("content"),
	}

	db.InsertItem(d, new)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func pageGetHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimLeft(r.URL.Path, "/page")
	item := db.ReadPage(d, id)

	item.Content = strings.Replace(item.Content, "\n", "<br>", -1) // -1: 모든 문자

	tmpl, err := template.New("Page").ParseFiles("templates/page.html")
	checkError(err)
	err = tmpl.ExecuteTemplate(w, "page.html", item)
	checkError(err)
}

func pageUpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func pageDeleteHandler(w http.ResponseWriter, r *http.Request) {

}

func Handler(dbpath string) http.Handler {
	d = db.InitDB(dbpath)
	db.CreateTable(d)
	// defer d.Close()

	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/page/{id:[0-9]+}", pageGetHandler).Methods("GET")
	mux.HandleFunc("/page/{id:[0-9]+}", pageUpdateHandler).Methods("UPDATE")
	mux.HandleFunc("/page/{id:[0-9]+}", pageDeleteHandler).Methods("DELETE")

	mux.HandleFunc("/write", getListHandler).Methods("GET")
	mux.HandleFunc("/write", postListHandler).Methods("POST")

	return mux
}
