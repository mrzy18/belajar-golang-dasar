package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type students struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=mrzy18 dbname=sandbox sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("SELECT id, name, grade FROM students WHERE age=$1", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []students

	for rows.Next() {
		student := students{}
		err := rows.Scan(&student.id, &student.name, &student.grade)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, student)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func main() {
	sqlQuery()
}
