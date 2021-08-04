package main

import "encoding/json"

var messages []*Message

// Message type is a message...
type Message struct {
	SenderID string `json:"senderId"`
	Username string `json:"username,omitempty"`
	Message  string `json:"message"`
}

// Post method of Message type broadcast message to all Clients
func (m *Message) Post() {
	m.Broadcast()
	messages = append(messages, m)
}

// BroadcastTo method sends the message to a given client
func (m *Message) BroadcastTo(to *Client) {
	byteMessage, _ := json.Marshal(m)
	to.ws.Write(byteMessage)
}

// Broadcast method sends the message to all clients
func (m *Message) Broadcast() {
	for _, connection := range clients {
		m.BroadcastTo(connection)
	}
}
