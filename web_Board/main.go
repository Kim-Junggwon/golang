package main

import (
	"database/sql" // db
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3" // db
)

// 기능 별 코드 분리 예정
func connectDB() {
	// db, err := sql.Open("sqlite3", "file:board.db")
	// db, err := sql.Open("sqlite3", "file::memory:?cache=shared&mode=memory")
	db, err := sql.Open("sqlite3", "./board.db")

	checkError(err)
	defer db.Close()

	// 테이블 생성(존재하지 않을 경우)
	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS board (
			id	INTEGER PRIMARY KEY AUTOINCREMENT,
			title	TEXT,
			name	TEXT,
			date	DATETIME,
			content	TEXT
		)`)
	checkError(err)
	stmt.Exec()

	rows, _ := db.Query("SELECT * FROM board")
	checkError(err)
	fmt.Println(rows)

	// 테스트 값 삽입
	stmt, err = db.Prepare("INSERT INTO board (title, name, date, content) VALUES (?, ?, datetime('now'), ?")
	checkError(err)

	res, err := stmt.Exec("title1", "name1", "content1")
	checkError(err)
	id, _ := res.LastInsertId()

	var tList Board
	tList.ID = int(id)
	tList.Title = "title11"
	tList.Name = "name11"
	tList.CreatedAt = time.Now()
	tList.Content = "content11"

	// 조회
	rows, _ := db.Query("SELECT id, title, name, date, content FROM board")
	checkError(err)

	list := []*Board{}

	defer rows.Close()
	for rows.Next() {
		var b Board
		rows.Scan(&b.ID, &b.Title, &b.Name, &b.CreatedAt, &b.Content)
		list = append(list, &b)
	}

	fmt.Println(list)
}

type Board struct {
	ID        int       `json:"number"`     // 글번호
	Title     string    `json:"title"`      // 제목
	Name      string    `json:"name"`       // 작성자
	CreatedAt time.Time `json:"created_at"` // 작성 시간
	Content   string    `json:"content"`    // 내용
}

type BoardList struct {
	Boards []*Board `json:"boards"` // DB에 저장된 게시물 목록
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
		ID:        0,
		Title:     r.FormValue("title"),
		Name:      r.FormValue("name"),
		Content:   r.FormValue("content"),
		CreatedAt: time.Now(),
	}

	log.Println("Title: ", board.Title)
	log.Println("Name: ", board.Name)
	log.Println("Content: ", board.Content)
	log.Println("Number: ", board.ID)
	log.Println("Time: ", board.CreatedAt)

	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	connectDB() // DB 연동 및 기능 테스트 함수

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/upload", uploadHandler)

	http.ListenAndServe(":3000", nil)
}
