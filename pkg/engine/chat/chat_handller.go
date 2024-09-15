package chat

import "github.com/gin-gonic/gin"

func SocketCreateChatRoom() gin.HandlerFunc {
	return createChatRoom
}
