package engine

import (
	"Locksyra/pkg/engine/authorize"
	"Locksyra/pkg/engine/chat"
	"Locksyra/pkg/engine/protected"

	"github.com/gin-gonic/gin"
)

func NewEngine(r *gin.Engine) *gin.Engine {
	auth := r.Group("/auth")
	{
		auth.POST("signup", authorize.SignupPost())
		auth.POST("login", authorize.LoginPost())
	}
	personal := r.Group("/personal")
	{
		personal.GET("me", protected.MeGet())
	}
	index := r.Group("/")
	{
		index.GET("ws", chat.SocketCreateChatRoom())
	}

	return r
}
