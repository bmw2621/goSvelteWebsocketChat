package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func init() {
	http.Handle("/ws", websocket.Handler(HandleNewConnection))
	http.Handle("/", http.FileServer(http.Dir("./chatClient/public")))
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", nil))
}
