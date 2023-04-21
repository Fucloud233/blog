package middleware

import (
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

// 通过.use来使用中间件
// https://gin-gonic.com/docs/examples/custom-middleware/

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var secretKey = []byte("")

func GenerateToken(claims *MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析Token
func parseToken(tokenString string) (*MyClaims, int) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if validErr, ok := err.(*jwt.ValidationError); ok {
			if validErr.Errors&jwt.ValidationErrorMalformed != 0 {
				// Token不正确
				return nil, errmsg.ERROR_TOKEN_WRONG
			} else if validErr.Errors&jwt.ValidationErrorExpired != 0 {
				// Token过期
				return nil, errmsg.ERROR_TOKEN_EXPIRED
			} else if validErr.Errors&jwt.ValidationErrorNotValidYet != 0 {
				// Token无效
				return nil, errmsg.ERROR_TOKEN_INVALID
			} else {
				return nil, errmsg.ERROR_TOKEN_INVALID
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, errmsg.SUCCESS
		}
		return nil, errmsg.ERROR_TOKEN_INVALID
	}

	return nil, errmsg.ERROR_TOKEN_INVALID
}

// VerifyJWT 中间件
func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Token")
		var code int
		var claims *MyClaims
		if tokenString != "" {
			// 解析Token
			claims, code = parseToken(tokenString)

			if code == errmsg.SUCCESS {
				c.Set("username", claims)
				c.Next()
				return
			}
		} else {
			// 当token不存在时
			code = errmsg.ERROR_TOKEN_NOT_EXIST
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"data":    nil,
		})
		c.Abort()
	}
}
