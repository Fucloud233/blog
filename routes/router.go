package routes

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 大写字母表示函数公有 小字母表示函数私有

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)
}
