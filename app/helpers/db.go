package helpers

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

const UrlTable = "url"

func CreateDB() {

	os.OpenFile("url_shortener.db", os.O_CREATE|os.O_RDWR, 0755)
	database, err := sql.Open("sqlite3", "url_shortener.db")
	defer database.Close()
	if err != nil{
		fmt.Println(err)
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS " + UrlTable + " (id INTEGER PRIMARY KEY AUTOINCREMENT, destination_url TEXT, short_url TEXT)")
	if err != nil{
		fmt.Println(err)
	}
	defer statement.Close()
	statement.Exec()
	statement, err = database.Prepare("CREATE UNIQUE INDEX IF NOT EXISTS short_url ON url(short_url);")
	defer statement.Close()
	if err != nil{
		fmt.Println(err)
	}
	statement.Exec()

}

func ExecSelect(query string, args ...interface{}) (rows *sql.Rows){
	database, err := sql.Open("sqlite3", "url_shortener.db")
	defer database.Close()
	if err != nil{
		fmt.Println(err)
	}
	rows,err = database.Query(query, args...)
	if err != nil{
		fmt.Println(err)
	}
	return rows
}

func ExecDelete(query string, args ...interface{}) (id int64){
	database, err := sql.Open("sqlite3", "url_shortener.db")
	defer database.Close()
	if err != nil{
		fmt.Println(err)
	}
	stmt,err := database.Exec(query, args...)
	if err != nil{
		fmt.Println(err)
	}
	id, err = stmt.RowsAffected()
	if err != nil{
		fmt.Println(err)
	}
	return id
}

func ExecCount(query string, args ...interface{}) (count int){
	database, err := sql.Open("sqlite3", "url_shortener.db")
	defer database.Close()
	if err != nil{
		fmt.Println(err)
	}
	err = database.QueryRow(query, args...).Scan(&count)
	if err != nil{
		fmt.Println(err)
	}
	return count
}


func ExecInsert(query string, args ...interface{}) (id int64){
	database, err := sql.Open("sqlite3", "url_shortener.db")
	defer database.Close()
	if err != nil{
		fmt.Println(err)
	}
	stmt,err := database.Exec(query, args...)
	if err != nil{
		fmt.Println(err)
	}
	id, err = stmt.LastInsertId()
	if err != nil{
		fmt.Println(err)
	}
	return id
}