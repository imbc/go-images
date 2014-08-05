package main

import (
	"code.google.com/p/go.net/websocket"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func WebsocketHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	router := mux.NewRouter()
	router.Handle("/ws", websocket.Handler(WebsocketHandler))
	staticHandler := http.FileServer(http.Dir("."))
	router.PathPrefix("/").Handler(staticHandler)
	http.ListenAndServe("localhost:1234", router)
}
