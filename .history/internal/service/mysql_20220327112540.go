package service

import (
	_ "github.com/go-sql-driver/mysql"
	"time"

	"database/sql"
	"fmt"
)

type MySQLServer struct {
	Connected bool
	Dsn       string
	// DialTimeout time.Duration
	MaxRetries int
	Conn       *sql.DB
}

func (s *MySQLServer) Name() string {
	return "SQL Server"
}

func (s *MySQLServer) AttemptConnection() error {
	currAttempts := 0

	for currAttempts < s.MaxRetries && s.Connected == false {
		conn, err := s.openDB(s.Dsn)
		fmt.Println("Attempting to connect to mysql with dns: " + s.Dsn)

		if err != nil {
			println(err.Error())
		} else {
			s.Conn = conn
			s.Connected = true
		}

		currAttempts++
		time.Sleep(2 * time.Second)

	}

	if s.Connected == false {
		return fmt.Errorf("failed to connect")
	}
	return nil

}

func (s *MySQLServer) openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		// print(err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
