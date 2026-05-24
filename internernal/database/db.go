package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

func NewDatabase(host string, port int, user string, pass string, name string) *Database {
	return &Database{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
		Name: name,
	}
}

func (d *Database) Connect() (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Pass, d.Name)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
