package models

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/wshaman/course-db/src/utils"
)

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"name"`
	Location string `db:"location"`
}


func  UserList(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`select u.id, u.name, email, l.name as location from users u 
left join locations l on u.location_id=u.id`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ListUsers")
	}
	return rowsToUsers(rows)
}

func UserListEmailLike(db *sql.DB, eml string) ([]User, error) {
	eml = "%" + eml
	rows, err := db.Query(`select u.id, u.name, email, l.name as location from users u 
left join locations l on u.location_id=u.id where email LIKE $1`, eml)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ListUsers")
	}
	return rowsToUsers(rows)
}

func  UserSave(db *sql.DB, u *User) error {
	if u.ID == 0 {
		return insertUser(db, u)
	}
	return updateUser(db, u)
}

func insertUser(db *sql.DB, u *User) (err error) {
	var id int64
	q := "insert into users (name, email) values ($1, $2) returning id"
	if err = db.QueryRow(q, u.Name, u.Email).Scan(&id); err != nil {
		return errors.Wrap(err, "failed to insert user")
	}
	u.ID = int(id)
	return nil
}

func updateUser(db *sql.DB, u *User) error {
	q := "update users set  name=$1, email=$2 where id=$3;"
	if _, err := db.Exec(q, u.Name, u.Email, u.ID); err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func rowsToUsers(rows *sql.Rows) (users []User, err error) {
	users = make([]User, 0)
	for rows.Next() {
		u := &User{}
		var s *string
		if err = rows.Scan(&u.ID, &u.Name, &u.Email, &s); err != nil {
			return nil, errors.Wrap(err, "failed to scan users (scan)")
		}
		u.Location = utils.StrPtr2Str(s)
		users = append(users, *u)
	}
	return users, nil
}