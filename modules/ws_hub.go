package modules

import "log"

var Hubs *Hub

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() {
	Hubs = &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
	go Hubs.Run()
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			log.Println("有客户端连入：", client)
			h.Clients[client] = true
			log.Println("当前客户端连接数：", len(h.Clients))
		case client := <-h.Unregister:
			log.Println("有客户端退出：", client)
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			log.Println("有客户端发出信息：", message)
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
