package authorize

import (
	"Locksyra/pkg/db"

	"github.com/gin-gonic/gin"
)

type SignupRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signup(c *gin.Context) {

	// ユーザー情報を取得
	var user SignupRequestModel

	// リクエストボディをパース
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	// パスワードを暗号化
	hash, err := PasswordEncrypt(user.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to encrypt password",
		})
		return
	}

	// ユーザー情報をDBに保存
	db.InsertDocument(map[string]string{
		"name": user.Username,
		"hash": hash,
	})

	c.JSON(200, gin.H{
		"message": "Signup",
	})
}
