package main

import (
	"github.com/wshaman/server-with-db/datasource"
	"github.com/wshaman/server-with-db/server"
)

func main() {
	ds := datasource.NewDB()
	server.Run("localhost:8080", ds)
}
