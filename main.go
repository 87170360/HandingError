package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	errors "github.com/pkg"
	"log"
)

func queryUser() error {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
	var name string
	err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("select name from users where id =%v", 1))
	}
	return nil
}

func main() {
	err := queryUser()
	fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
	fmt.Printf("stack track \n%+v\n", err)
}
