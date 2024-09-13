package protected

import (
	jwt_util "Locksyra/pkg/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
)

// RegisterPersonalRoutes パーソナル関連のエンドポイントを登録
func me(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	err := jwt_util.VerifyJWT(token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	claims, err := jwt_util.ParseJWT(token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Me",
		"claims":  fmt.Sprintf("%v", claims),
	})

}
