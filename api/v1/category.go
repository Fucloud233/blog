package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)

	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// GetCategories 查询用户列表
func GetCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetCategories(pageSize, pageNum)
	code := errmsg.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditCategory 编辑用户
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.CheckupCategory(id, data.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteCategory 删除用户
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategory(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
