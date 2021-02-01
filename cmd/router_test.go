package cmd

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var app *Application

func TestMain(m *testing.M) {
	log.Println("BEFORE tests")
	app = NewApplication()

	log.Println("RUNNING tests suite")
	exitVal := m.Run()

	log.Println("AFTER tests")
	os.Exit(exitVal)
}

func TestNewRouter(t *testing.T) {
	router := NewRouter()

	assert.NotNil(t,router)
}

func TestPing(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	response := executeRequest(req)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, response.Body.String(), "pong")

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}