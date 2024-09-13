package authorize

import (
	"Locksyra/pkg/db"
	"Locksyra/pkg/util"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var loginReq LoginRequestSchema

	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	user, err := db.ReadUser(loginReq.Uaername)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid username or password",
			"error":   err.Error(),
		})
		return
	}

	if err := util.CompareHashAndPassword(user.HashedPassword, loginReq.Password); err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid username or password",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"message": "Login success"})
}
