package chat

import (
	"github.com/gin-gonic/gin"
)

type CreateDirectMessageChannelRequest struct {
	UserID     string `json:"user_id"`
	SenderName string `json:"sender_name"`
}

func createDirectMessageChannel(c *gin.Context) {

}
