package datasource

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "pass"
	dbname   = "db"
)

type DS struct {
	db *sql.DB
}

func (ds *DS) GetUsers() ([]string, error) {
	rows, err := ds.db.Query(`select name from users;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []string{}
	for rows.Next() {
		name := ""
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		users = append(users, name)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func NewDB() *DS {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &DS{
		db,
	}
}
