package engine

import (
	"Locksyra/pkg/engine/authorize"

	"github.com/gin-gonic/gin"
)

func NewEngine(r *gin.Engine) *gin.Engine {
	auth := r.Group("/auth")
	{
		auth.POST("signup", authorize.SignupPost())
		auth.POST("login", authorize.LoginPost())
	}

	return r
}
