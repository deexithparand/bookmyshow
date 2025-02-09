package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RouterInit() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("Welcome to bookmyshow\n"))
	})

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		path_url := r.URL.Path
		response_string := "page " + path_url + " is not found\n"
		w.Write([]byte(response_string))
	})

	return router
}
