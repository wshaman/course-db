package server

import (
	"github.com/wshaman/server-with-db/datasource"
	"github.com/wshaman/server-with-db/server/handlers"
	"net/http"
)

func Run(addr string, ds datasource.DS) {
	usersHandler := handlers.NewUserHandler(ds)
	http.Handle("/users", usersHandler)
	http.ListenAndServe(addr, nil)
}
