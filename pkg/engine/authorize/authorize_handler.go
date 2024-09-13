package authorize

import (
	"github.com/gin-gonic/gin"
)

func SignupPost() gin.HandlerFunc {
	return signup
}

func LoginPost() gin.HandlerFunc {
	return login
}

type LoginRequestSchema struct {
	Uaername string `json:"username"`
	Password string `json:"password"`
}
