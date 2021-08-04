package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var clients []*Client

// Client type holds websocket connections
type Client struct {
	id       string
	username string
	ws       *websocket.Conn
}

// StartListening method of the Client type sets a listener for incoming messages and directs them to the appropriate handler
func (c *Client) StartListening() {

	// Set buffer byte array at 512 bytes
	buf := make([]byte, 512)

	// Loop listening for messages from the client and direct to handler
	for {
		n, err := c.ws.Read(buf[0:])

		// Error only exists if the connection is broken
		if err != nil {
			ReleaseConnection(c)
			errorMessage := Message{
				SenderID: "System",
				Username: "System",
				Message:  fmt.Sprintf("%s disconnected", c.username),
			}
			errorMessage.Post()
			break
		} else {
			HandleInputMessage(c, buf[0:n])
		}
	}
}

// HandleNewConnection is called when a new connection to the websocket server is created
func HandleNewConnection(ws *websocket.Conn) {
	log.Println("New Connection!")

	// Create a new Connection object and add it to the clients array
	connection := Client{generateId(), "", ws}
	clients = append(clients, &connection)

	// Start Listening
	connection.StartListening()
}

// ReleaseConnection function removes a client from the clients array
func ReleaseConnection(c *Client) {
	log.Println("ReleaseConnection:", c.username, c.id)

	index := -1
	for k, v := range clients {
		if c == v {
			index = k
			break
		}
	}

	if index >= 0 {
		clients = append(clients[:index], clients[index+1:]...)
	}

	log.Println("Connection count:", len(clients))
}
