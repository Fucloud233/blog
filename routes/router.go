package routes

import (
	v1 "blog/api/v1"
	"blog/middleware"
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

	auth := r.Group("api/v1")
	auth.Use(middleware.VerifyJWT())
	{
		auth.GET("users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
	}

	v1Router := r.Group("api/v1")
	{
		// 用户模块路由接口
		v1Router.POST("user/add", v1.AddUser)

		// 登陆模块路由接口
		v1Router.POST("login", v1.Login)

		// 分类模块路由接口
		v1Router.GET("categories", v1.GetCategories)

		// 文章模块路由接口
		v1Router.GET("articles", v1.GetArticles)
		v1Router.GET("article/info/:id", v1.GetArticle)
		v1Router.GET("articles/category/:cid", v1.GetArticleFromCategory)

	}

	r.Run(utils.HttpPort)
}
