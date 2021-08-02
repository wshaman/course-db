package main

import (
	"encoding/json"
	"net/http"

	"github.com/wshaman/server-with-db/datasource"
)

type UsersHandler struct {
	ds *datasource.DS
}

func (u *UsersHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	users, err := u.ds.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	msg, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func main() {
	ds := datasource.NewDB()
	usersHandler := &UsersHandler{
		ds,
	}

	http.Handle("/users", usersHandler)
	http.ListenAndServe("localhost:8080", nil)
}
