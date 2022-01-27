package handlers

import (
	"encoding/json"
	"github.com/wshaman/server-with-db/datasource"
	"net/http"
)

type UsersHandler struct {
	ds datasource.DS
}

func NewUserHandler(ds datasource.DS) UsersHandler {
	return UsersHandler{ds: ds}
}

func (u UsersHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
