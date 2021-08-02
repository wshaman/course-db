package main

import (

	"fmt"
	"github.com/wshaman/course-db/src/db/models"
	"log"

	_ "github.com/lib/pq"

	"github.com/wshaman/course-db/src/db/db"
	"github.com/wshaman/course-db/src/utils"
)

func dbsFromEnv() db.DBSetup {
	dbs := db.DBSetup{
		User:   utils.EnvOrDef("DB_USER", "postgres"),
		Passwd: utils.EnvOrDef("DB_PASSWD", "pwd123"),
		Host:   utils.EnvOrDef("DB_HOST", "localhost"),
		Port:   utils.EnvOrDefInt("DB_PORT", 15432),
		Name:   utils.EnvOrDef("DB_NAME", "course_db"),
		Type:   "postgres",
	}
	return dbs
}

func main() {
	dbObj, err := db.New(dbsFromEnv())
	OnErrPanic(err)
	users, err := models.UserList(dbObj)
	OnErrPanic(err)
	fmt.Println(users)
}


func OnErrPanic(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}