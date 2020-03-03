package modules

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/websocket"
	"log"
)

var Hubs *Hub

type Broadcast struct {
	ID  bson.ObjectId
	Msg string
}

type Hub struct {
	Clients map[bson.ObjectId]*Client
	//Clients    map[*Client]bool
	Broadcast  chan Broadcast
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() {
	Hubs = &Hub{
		Clients: make(map[bson.ObjectId]*Client),
		//Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Broadcast),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
	go Hubs.Run()
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			log.Printf("有客户端连入：%v\t客户端ID：%v\n", client, client.ID.Hex())
			if _, ok := h.Clients[client.ID]; ok {
				log.Println("重复连接")
				h.Clients[client.ID].Conn.Close()
			}
			h.Clients[client.ID] = client
			log.Println("当前客户端连接数：", len(h.Clients))
		case client := <-h.Unregister:
			log.Printf("有客户端退出：%v\t客户端ID：%v\n", client, client.ID.Hex())
			if err := h.Clients[client.ID].Conn.WriteMessage(websocket.PingMessage, nil);err != nil {
				log.Println(err)
				delete(h.Clients, client.ID)
			}
			close(client.Send)
			log.Println("当前客户端连接数：", len(h.Clients))
		case b := <-h.Broadcast:
			message, _ := json.Marshal(b)
			log.Printf("有客户端发出信息：%v\t客户端ID：%v\n", string(message), b.ID.Hex())
			for id, client := range h.Clients {
				if id != b.ID {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(h.Clients, client.ID)
					}
				}
			}
			//for client := range h.Clients {
			//	select {
			//	case client.Send <- message:
			//	default:
			//		close(client.Send)
			//		delete(h.Clients, client)
			//	}
			//}
		}
	}
}
