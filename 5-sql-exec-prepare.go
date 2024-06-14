package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3307)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO tb_student (id, name, age, grade) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt.Exec("G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println("Error executing statement:", err.Error())
		return
	}
	fmt.Println("insert success!")

	stmt, err = db.Prepare("UPDATE tb_student SET name=?, age=?, grade=? WHERE id=?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt.Exec("Galahad Updated", 30, 3, "G001")
	if err != nil {
		fmt.Println("Error executing statement:", err.Error())
		return
	}
	fmt.Println("update success!")

	stmt, err = db.Prepare("DELETE FROM tb_student WHERE id=?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = stmt.Exec("G001")
	if err != nil {
		fmt.Println("Error executing statement:", err.Error())
		return
	}
	fmt.Println("delete success!")
}

func main() {
	sqlExec()
}
