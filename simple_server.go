package main

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	if vars["name"] != "" {
		io.WriteString(writer, "hello "+vars["name"]+"!")
	} else {
		io.WriteString(writer, "hello!")
	}
}

func NoCacheDecorator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		h.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloHandler)
	router.HandleFunc("/hello/{name}", HelloHandler)
	staticHandler := http.FileServer(http.Dir("."))
	staticHandler = http.StripPrefix("/static/", staticHandler)
	staticHandler = NoCacheDecorator(staticHandler)
	router.PathPrefix("/static/").Handler(staticHandler)
	http.ListenAndServe("localhost:1234", router)
}
