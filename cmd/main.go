package main

import (
	"fmt"
	"go-websocket/internal"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	server := internal.InitServer()
	http.Handle("/ws", websocket.Handler(server.HandleConnection))
	fmt.Println("Starting websocket server on localhost:3000/ws")
	http.ListenAndServe("localhost:3000", nil)
}
