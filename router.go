package secretservice

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/get", GetHandler)
	r.HandleFunc("/set", SetHandler)
	r.HandleFunc("/delete", DeleteHandler)
	return r
}
