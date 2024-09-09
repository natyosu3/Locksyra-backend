package authorize

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 暗号(Hash)化
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 暗号(Hash)と入力された平パスワードの比較
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func Signup(c *gin.Context) {

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
	_, err := PasswordEncrypt(user.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to encrypt password",
		})
		return
	}

	// ユーザー情報をDBに保存
	// ここでは省略

	c.JSON(200, gin.H{
		"message": "Signup",
	})
}
