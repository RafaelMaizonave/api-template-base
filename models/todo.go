package models

import (
	"database/sql"
	"fmt"

	"github.com/RafaelMaizonave/api-template-base/entities"

	"github.com/RafaelMaizonave/api-template-base/db"
)

func Insert(todo entities.Todo) (id int64, err error) {

	fmt.Println(todo)

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO todos(title,description,done) VALUES ($1,$2,$3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return id, err

}

func Get(id int64) (todo entities.Todo, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT id,title,description,done FROM todos WHERE id=$1`, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	if err != nil && err != sql.ErrNoRows {
		return
	}

	return todo, nil
}

func GetAll() (todos []entities.Todo, err error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT id,title,description,done FROM todos`)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo entities.Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

		if err != nil {
			fmt.Println(err)
			continue
		}

		todos = append(todos, todo)
	}

	if err != nil && err != sql.ErrNoRows {
		return
	}

	return todos, nil
}

func Update(id int64, todo entities.Todo) (int64, error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$1 , description=$2 , done=$3 WHERE id=$4 `, todo.Title, todo.Description, todo.Done, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

func Delete(id int64) (int64, error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1 `, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}
