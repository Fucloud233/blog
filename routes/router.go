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
		v1Router.POST("category/add", v1.AddCategory)
		v1Router.GET("categories", v1.GetCategories)
		v1Router.PUT("category/:id", v1.EditCategory)
		v1Router.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块路由接口
		v1Router.POST("article/add", v1.AddArticle)
		v1Router.GET("articles", v1.GetArticles)
		v1Router.PUT("article/:id", v1.EditArticle)
		v1Router.DELETE("article/:id", v1.DeleteArticle)
	}

	r.Run(utils.HttpPort)
}
