package db

import "database/sql"

type TodoModel struct {
	DB *sql.Conn
}
