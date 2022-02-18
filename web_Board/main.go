package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Board struct {
	Title     string    `json:"title"`      // 제목
	Name      string    `json:"name"`       // 작성자
	Content   string    `json:"content"`    // 내용
	Number    int       `json:"number"`     // 글번호
	CreatedAt time.Time `json:"created_at"` // 작성 시간
}

type BoardList struct {
	Boards []*Board `json:"boards"`
}

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

func writeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFiles("templates/write.html")
	checkError(err)

	tmpl.ExecuteTemplate(w, "write.html", "")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	board := Board{
		Title:     r.FormValue("title"),
		Name:      r.FormValue("name"),
		Content:   r.FormValue("content"),
		Number:    0, // DB에 저장된 List Max 값 +1?
		CreatedAt: time.Now(),
	}

	log.Println("Title: ", board.Title)
	log.Println("Name: ", board.Name)
	log.Println("Content: ", board.Content)
	log.Println("Number: ", board.Number)
	log.Println("Time: ", board.CreatedAt)

	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/upload", uploadHandler)

	http.ListenAndServe(":3000", nil)
}
