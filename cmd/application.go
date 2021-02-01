package cmd

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Application struct {
	Router *mux.Router
}

func NewApplication() *Application {
	app := &Application {
		Router:NewRouter(),
	}
	return app
}

func (app *Application) Run( port string ) {
	server := http.ListenAndServe(port, app.Router)
	log.Fatal(server)
}
