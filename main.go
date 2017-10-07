package main
import (
	//"fmt"
	"net/http"
	"strings"
	"log"
	"html/template"
	"io"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		if t, err := template.ParseFiles("login.html"); err != nil {
			log.Fatal("template.ParseFiles", err)
		} else {
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {
		for k, v := range r.Form {
			io.WriteString(w, "key: " + k)
			io.WriteString(w, " val: " + strings.Join(v, " ") + "\n")
		}
		// connect db
		if db, err := sql.Open("mysql", "oprecruit:oprecruit@/oprecruit?charset=utf8"); err != nil {
			log.Fatal("DB connect", err)
		} else {
			defer db.Close()
			if stmt, err := db.Prepare("INSERT user_info SET netid=?, number=?, nickname=?, pwd=?, grade=?"); err != nil {
				log.Fatal("sql syntax", err)
			} else {
				if _, err := stmt.Exec(r.Form["netid"][0], r.Form["number"][0], r.Form["nickname"][0], r.Form["pwd"][0], r.Form["grade"][0]); err != nil {
					log.Fatal("sql execute", err)
				}
			}
		}
	}
}
func main() {
	http.HandleFunc("/", login)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
