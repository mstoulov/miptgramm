package handler

import (
	"course_fullstack/backend/internal/service"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var addr = flag.String("addr", ":8080", "http service address")

func serveWs(service *service.Service, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error on opening client connection", err)
		return
	}
	client := NewClient(service, conn)

	go client.ListenSocket()
	go client.ListenPing()
	go client.WriteSocket()
}

func NewServer(service *service.Service) *http.Server {
	flag.Parse()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(service, w, r)
	})
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	return server
}
