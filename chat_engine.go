package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// HandleInputMessage function is the primary handler for the websocket connection which listens for messages from clients
func HandleInputMessage(from *Client, data []byte) {
	var inputJson map[string]string
	json.Unmarshal(data, &inputJson)

	log.Printf("New %v from: %v -> %v", inputJson["action"], from.username, inputJson["message"])

	var newMessageToAll Message

	switch inputJson["action"] {
	case "post_message":
		newMessageToAll = Message{
			SenderID: from.id,
			Username: from.username,
			Message:  inputJson["message"],
		}

	case "initial_connection":
		from.username = inputJson["username"]
		newMessageToAll = Message{
			SenderID: "System",
			Username: "System",
			Message:  fmt.Sprintf("%s joined the chat", from.username),
		}

		chatHistory, _ := json.Marshal(messages)

		newUserIDMessage := Message{
			SenderID: "System",
			Username: "System",
			Message:  fmt.Sprintf("yourId:%v", from.id),
		}

		newUserIDMessage.BroadcastTo(from)

		from.ws.Write(chatHistory)
	}

	if len(newMessageToAll.Message) > 0 {
		newMessageToAll.Post()
	}
}
