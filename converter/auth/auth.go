package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func AuthTokenValid(c *gin.Context) bool {
	tokenValue := GetAuthToken(c)
	if len(tokenValue) == 0 {
		return false
	}

	_, err := GetSession(tokenValue)
	if nil != err {
		return false
	}

	return true
}

func GetAuthToken(c *gin.Context) string {
	return c.Request.Header.Get(os.Getenv("AUTH_HEADER"))
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
