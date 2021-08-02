package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/wshaman/migrate"

	"github.com/wshaman/course-db/src/db/db"
	_ "github.com/wshaman/course-db/src/db/migrations"
	"github.com/wshaman/course-db/src/utils"
)

func help() {
	c := os.Args[0]
	fmt.Printf(`Usage:
	%s command [params]
commands:
	help show this screen and exit
	up run all migrations 
	down rollback 1 last migration
	create creates a new migration file template 
Eg:
%s create add_table_users
`, c, c)
}

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

	if len(os.Args) < 2 {
		help()
		os.Exit(0)
	}
	dbObj, err := db.New(dbsFromEnv())
	if err != nil {
		log.Fatal(err)
		//	fmt.Println(err.Error())
	}
	command := strings.ToLower(os.Args[1])
	switch command {
	case "up":
		err = migrate.Up(dbObj)
	case "down":
		err = migrate.Down(dbObj)
	case "sync":
		err = migrate.Sync(dbObj)
	case "create":
		if len(os.Args) < 3 {
			help()
			os.Exit(1)
		}
		err = migrate.CreateFile(os.Args[2], "main", "./", true)
	default:
		help()
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
