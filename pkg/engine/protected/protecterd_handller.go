package protected

import (
	"github.com/gin-gonic/gin"
)

func MeGet() gin.HandlerFunc {
	return me
}
