package chat

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type SocketType string

const (
	FirstConnectionType SocketType = "first_connection"
	MessageType         SocketType = "message"
	InviteType          SocketType = "invite"
	ErrorType           SocketType = "error"
)

var allowedTypes = map[SocketType]bool{
	MessageType:         true,
	InviteType:          true,
	ErrorType:           true,
	FirstConnectionType: true,
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type SocketFormat struct {
	Type         SocketType `json:"type"`
	ConnectionID string     `json:"connection_id"`
	Data         string     `json:"data"`
	Sender       User       `json:"sender"`
	Receiver     User       `json:"receiver"`
}

var connections = struct {
	sync.RWMutex
	m map[string]*websocket.Conn
}{m: make(map[string]*websocket.Conn)}

func createChatRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: ", err)
		return
	}
	defer conn.Close()

	// 最初のメッセージでユーザIDを受信
	var initialMsg SocketFormat
	err = conn.ReadJSON(&initialMsg)
	if err != nil {
		log.Println("Failed to read initial message: ", err)
		return
	}

	// ユーザIDをコネクションIDとして使用
	userID := initialMsg.Sender.ID

	// 接続を管理するマップに追加
	connections.Lock()
	connections.m[userID] = conn
	connections.Unlock()

	defer func() {
		// 接続を管理するマップから削除
		connections.Lock()
		delete(connections.m, userID)
		connections.Unlock()
	}()

	// 最初のメッセージを処理
	if initialMsg.Type == FirstConnectionType {
		sendMessage(conn, SocketFormat{
			Type:         MessageType,
			Data:         "Message received by server",
			ConnectionID: userID,
			Sender: User{
				ID:   initialMsg.Sender.ID,
				Name: initialMsg.Sender.Name,
				UUID: userID,
			},
			Receiver: User{
				ID:   "admin",
				Name: "admin",
				UUID: "admin",
			},
		})
	}

	for {
		var msg SocketFormat
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error: ", err)
			break
		}

		if !allowedTypes[msg.Type] {
			sendError(conn, "Invalid message type")
			continue
		}

		switch msg.Type {
		case MessageType:
			log.Printf("Message from %s: %s", msg.Sender.Name, msg.Data)
		case InviteType:
			log.Printf("Invite from %s to %s: %s", msg.Sender.Name, msg.Receiver.Name, msg.Data)
			notifyReceiver(msg)
		case ErrorType:
			log.Printf("Error from %s: %s", msg.Sender.Name, msg.Data)
		}
	}
}

func sendError(conn *websocket.Conn, errorMsg string) {
	err := conn.WriteJSON(SocketFormat{
		Type: ErrorType,
		Data: errorMsg,
	})
	if err != nil {
		log.Println("Write error: ", err)
	}
}

func sendMessage(conn *websocket.Conn, msg SocketFormat) {
	err := conn.WriteJSON(msg)
	if err != nil {
		log.Println("Write error: ", err)
	}
}

func notifyReceiver(msg SocketFormat) {
	connections.RLock()
	receiverConn, ok := connections.m[msg.Receiver.ID]
	connections.RUnlock()
	if ok {
		sendMessage(receiverConn, SocketFormat{
			Type:     InviteType,
			Data:     msg.Data,
			Sender:   msg.Sender,
			Receiver: msg.Receiver,
		})
	} else {
		log.Printf("Receiver not found: %s", msg.Receiver.ID)
	}
}
