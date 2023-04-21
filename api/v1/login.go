package v1

import (
	"blog/middleware"
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)

	var token string
	var code int

	// 使用数据库中的登陆验证程序
	formData, code = model.CheckLogin(formData.Username, formData.Password)

	if code == errmsg.SUCCESS {
		// 如果验证成功 则返回token
		setToken(c, formData)
	} else {
		// 如果验证失败 则返回空
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	}
}

func setToken(c *gin.Context, user model.User) {
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "Fucloud",
		},
	}

	// 使用中间件生成token
	token, err := middleware.GenerateToken(&claims)
	var code int
	if err != nil {
		code = errmsg.ERROR
	} else {
		code = errmsg.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"username": user.Username,
		"id":       user.ID,
		"message":  errmsg.GetErrMsg(code),
		"token":    token,
	})
}
