package hd

import (
	"board/db"
	"database/sql"
	"fmt"
	"net/http"
	"os"
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
	tmpl, err := template.New("Hello").ParseFiles("templates/index.html")
	checkError(err)

	tmpl.ExecuteTemplate(w, "index.html", "Test Msg")
}

func getListHandler(w http.ResponseWriter, r *http.Request) {
	// template
}

func postListHandler(w http.ResponseWriter, r *http.Request) {
	// submit button event

}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// row 별 페이지
}

func Handler(dbpath string) http.Handler {
	d = db.InitDB(dbpath)
	defer d.Close()

	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/{id:[0-9]+}", pageHandler).Methods("GET")

	mux.HandleFunc("/write", getListHandler).Methods("GET")
	mux.HandleFunc("/write", postListHandler).Methods("POST")

	return mux
}
