package chat

import (
	"log"
	"log/slog"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	Type        SocketType `json:"type"`
	ConectionID string     `json:"conection_id"`
	Data        string     `json:"data"`
	Sender      User       `json:"sender"`
	Receiver    User       `json:"receiver"`
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

	// 一意のIDを生成
	connID := uuid.New().String()

	// 接続を管理するマップに追加
	connections.Lock()
	connections.m[connID] = conn
	connections.Unlock()

	defer func() {
		// 接続を管理するマップから削除
		connections.Lock()
		delete(connections.m, connID)
		connections.Unlock()
	}()

	for {
		var msg SocketFormat
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error: ", err)
			break
		}

		if !allowedTypes[msg.Type] {
			log.Println("Invalid message type: ", msg.Type)
			err = conn.WriteJSON(SocketFormat{
				Type: ErrorType,
				Data: "Invalid message type",
			})
			if err != nil {
				log.Println("Write error: ", err)
				break
			}
			continue
		}

		switch msg.Type {
		case FirstConnectionType:
			// メッセージ送信者にメッセージを送信
			err = conn.WriteJSON(SocketFormat{
				Type:        MessageType,
				Data:        "Message received by server",
				ConectionID: connID,
				Sender: User{
					ID:   msg.Sender.ID,
					Name: msg.Sender.Name,
					UUID: connID,
				},
				Receiver: User{
					ID:   "admin",
					Name: "admin",
					UUID: "admin",
				},
			})
			if err != nil {
				log.Println("Write error: ", err)
			}
		case MessageType:
			log.Printf("Message from %s: %s", msg.Sender, msg.Data)
		case InviteType:
			log.Printf("Invite from %s to %s: %s", msg.Sender, msg.Receiver, msg.Data)
			// 受信者に通知を送信
			connections.RLock()
			receiverConn, ok := connections.m[msg.Receiver.ID]
			connections.RUnlock()
			if ok {
				err = receiverConn.WriteJSON(SocketFormat{
					Type:   InviteType,
					Data:   msg.Data,
					Sender: msg.Sender,
				})
				if err != nil {
					log.Println("Write error: ", err)
				}
			} else {
				slog.Info("Receiver not found: " + msg.Receiver.ID)
			}
		case ErrorType:
			log.Printf("Error from %s: %s", msg.Sender, msg.Data)
		}
	}
}
