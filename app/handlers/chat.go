package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"vibly/app/models"
	"vibly/app/store"
	"vibly/pkg/utils"

	"github.com/gorilla/websocket"
)

var (
	// ChatStore instance for persistence
	PersistChatStore = &store.ChatStore{BaseDir: "data"}

	// upgrader for WebSockets
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	// channelClients maintains active WebSocket connections per channel
	channelClients = make(map[string]map[*websocket.Conn]bool)
	clientsMu      sync.Mutex
)

// ChatWebSocketHandler handles real-time messaging per channel
func ChatWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	channelID := r.URL.Query().Get("channelID")
	if channelID == "" {
		http.Error(w, "channelID is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	// Register client
	clientsMu.Lock()
	if channelClients[channelID] == nil {
		channelClients[channelID] = make(map[*websocket.Conn]bool)
	}
	channelClients[channelID][conn] = true
	clientsMu.Unlock()

	defer func() {
		clientsMu.Lock()
		delete(channelClients[channelID], conn)
		clientsMu.Unlock()
	}()

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		// Set metadata
		msg.ID = utils.GenerateUUID()
		msg.ChannelID = channelID
		msg.CreatedAt = time.Now()

		// Persist message
		_ = PersistChatStore.SaveMessage(channelID, msg)

		// Broadcast to all clients in this channel
		broadcastToChannel(channelID, msg)
	}
}

// ChatHistoryHandler retrieves persistent messages for a channel
func ChatHistoryHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.AllowMethod(w, r, http.MethodGet) {
		return
	}

	channelID := r.URL.Query().Get("channelID")
	if channelID == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 50
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	messages, err := PersistChatStore.GetMessages(channelID, limit)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, true, "Chat history retrieved", messages)
}

func broadcastToChannel(channelID string, msg models.Message) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	payload, _ := json.Marshal(msg)
	for client := range channelClients[channelID] {
		_ = client.WriteMessage(websocket.TextMessage, payload)
	}
}
