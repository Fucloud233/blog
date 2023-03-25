package routes

import (
	v1 "blog/api/v1"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

// 大写字母表示函数公有 小字母表示函数私有

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	// 测试函数
	//router := r.Group("api/v1")
	//{
	//	router.GET("hello", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{
	//			"msg": "ok",
	//		})
	//	})
	//}

	v1Router := r.Group("api/v1")
	{
		// 用户模块路由接口
		v1Router.POST("user/add", v1.AddUser)
		v1Router.GET("users", v1.GetUsers)
		v1Router.PUT("user/:id", v1.EditUser)
		v1Router.DELETE("user/:id", v1.DeleteUser)
		// 分类模块路由接口

		// 文章模块路由接口
	}

	r.Run(utils.HttpPort)
}
