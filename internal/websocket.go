package internal

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

type Server struct {
	clients map[*websocket.Conn]bool
}

func InitServer() *Server {
	return &Server{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleConnection(ws *websocket.Conn) {
	fmt.Println("New connection from :", ws.RemoteAddr())
	s.clients[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := ws.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client: ", err.Error())
			continue
		}
		s.broadcast(buffer[:n])
	}
}

func (s *Server) broadcast(message []byte) {
	for client := range s.clients {
		_, err := client.Write(message)
		if err != nil {
			fmt.Println("Error writing to client: ", err.Error())
			client.Close()
			delete(s.clients, client)
		}
	}
}
