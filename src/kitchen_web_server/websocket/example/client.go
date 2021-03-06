package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan Message
	id   string
}

type Message struct {
	To      string      `json:"receiverID"`
	From    string      `json:"senderID"`
	Payload OrderStruct `json:"payload"`
}

var upgrader = websocket.Upgrader{}

func serveClient(hub *Hub, w http.ResponseWriter, r *http.Request) {
	var conn, err = upgrader.Upgrade(w, r, nil)

	if err != nil {
		println("Error in upgrading http to websocket. Check log for more info")
		log.Println(err)
		return
	}

	var initialsMsg Message
	err = conn.ReadJSON(&initialsMsg)
	if err != nil {
		println("Error! Something wrong")
		conn.Close()
		return
	}

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan Message),
		id:   initialsMsg.From,
	}

	client.hub.join <- client
	fmt.Println("client joined", client)

	go client.readMsg()
	go client.writeMsg()
}

func (client *Client) readMsg() {
	for {
		var msg Message
		err := client.conn.ReadJSON(&msg)
		if err != nil {
			println("Client left")
			client.conn.Close()
			break
		}

		// println("Reading from user", string(msg.From), "type: ", string(msg.Type))
		// switch msg.Type {
		// case "broadcast":
		// client.hub.broadcast <- msg
		// case "private":
		client.hub.private <- msg
		// }
	}
}

func (client *Client) writeMsg() {
	for {
		select {
		case msg := <-client.send:
			// println("Receving message: ", msg.Payload)
			client.conn.WriteJSON(Message{
				To:      msg.To,
				Payload: msg.Payload,
			})
		}
	}
}
