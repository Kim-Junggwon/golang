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
		fmt.Println("Fatal error : ", err.Error())
		os.Exit(1)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	items := db.ReadItems(d)
	db.CreateTable(d) // table이 없을 경우 생성

	tmpl, err := template.New("Index").ParseFiles("templates/index.html")
	checkError(err)
	err = tmpl.ExecuteTemplate(w, "index.html", items)
	checkError(err)
}

func writeGetHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("Write").ParseFiles("templates/write.html")
	checkError(err)
	err = tmpl.ExecuteTemplate(w, "write.html", "")
	checkError(err)
}

func writePostHandler(w http.ResponseWriter, r *http.Request) {
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

func modifyGetHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimLeft(r.URL.Path, "/modify")
	item := db.ReadPage(d, id)

	tmpl, err := template.New("Modify").ParseFiles("templates/modify.html")
	checkError(err)
	err = tmpl.ExecuteTemplate(w, "modify.html", item)
	checkError(err)
}

func modifyUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimLeft(r.URL.Path, "/modify")
	content := r.FormValue("content")

	db.UpdateContent(d, id, content)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func pageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimLeft(r.URL.Path, "/del")

	db.DeletePage(d, id)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func Handler(dbpath string) http.Handler {
	d = db.InitDB(dbpath)
	db.CreateTable(d)
	// defer d.Close()

	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/page/{id:[0-9]+}", pageGetHandler).Methods("GET")

	mux.HandleFunc("/modify/{id:[0-9]+}", modifyGetHandler).Methods("GET")
	mux.HandleFunc("/modify/{id:[0-9]+}", modifyUpdateHandler).Methods("POST") // 게시물 수정
	mux.HandleFunc("/del/{id:[0-9]+}", pageDeleteHandler).Methods("GET")       // 게시물 삭제

	mux.HandleFunc("/write", writeGetHandler).Methods("GET")
	mux.HandleFunc("/write", writePostHandler).Methods("POST")

	return mux
}
