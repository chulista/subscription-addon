package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", Ping)

	return router
}

func Ping( w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}


