package models

import (
	"StudyProjects/ToDoAPI/db"

	_ "github.com/lib/pq"
)

func Insert(todo Todo) (ID int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	sql := `INSERT INTO todos(title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
