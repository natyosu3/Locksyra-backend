package authorize

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignupPost() gin.HandlerFunc {
	return signup
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

type LoginRequestSchema struct {
	Uaername string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginReq LoginRequestSchema

	c.ShouldBindJSON(&loginReq)
}
