package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	db, err:= sql.Open("postgres",
		"postgres://postgres:pwd123@localhost:15432/course_db?sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	r.Use(func(c *gin.Context){
		c.Set("DB",  db)
	})
	r.GET("/ping", handler)
	r.Run()
}


func handler(c *gin.Context) {
	s := "hi!"
	dbO, ok := c.Get("DB")
	if !ok {
		log.Fatal("---")
	}
	db, ok := dbO.(*sql.DB)
	if !ok {
		log.Fatal("---")
	}
	db.QueryRow("select now();").Scan(&s)
	c.JSON(200, gin.H{
		"message": s,
	})
}